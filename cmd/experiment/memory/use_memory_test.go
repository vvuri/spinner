package memory

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"
	"unsafe"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Человеко-читаемый формат
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func PrintAsBinary(a any) {
	type iface struct {
		t, v unsafe.Pointer
	}
	p := uintptr((*(*iface)(unsafe.Pointer(&a))).v)

	t := reflect.TypeOf(a)

	for i := 0; i < int(t.Size()); i++ {
		n := *(*byte)(unsafe.Pointer(p))
		fmt.Printf("%08b ", n)
		p += unsafe.Sizeof(n)
	}

	fmt.Print("\n")
}

func TestUseMemorySlice(t *testing.T) {
	// Напечатать использование памяти перед началом работы
	printMemUsage()

	// Ваш код
	for i := 0; i < 10; i++ {
		_ = make([]byte, 10<<20) // Создание 10 MB среза
		time.Sleep(1 * time.Second)
		printMemUsage()
	}

	// Напечатать использование памяти в конце работы
	printMemUsage()
}

func TestUseMemoryRead(t *testing.T) {
	printMemUsage()
	str := "TCP client exiting..."
	fmt.Println(str)
	printMemUsage()
}

func TestByteUse(t *testing.T) {
	var x byte
	x = 254
	PrintAsBinary(x)
}

func TestUint16Use(t *testing.T) {
	var x uint16
	x = 32768
	fmt.Println("little-endian (или «от младшего к старшему»)")
	PrintAsBinary(x)
}

func TestAllocatorWork(t *testing.T) {
	var b1, b2 [257]byte

	e1 := unsafe.Pointer(&b1)
	e2 := unsafe.Pointer(&b2)
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println("Size of b1:", unsafe.Sizeof(b1))
	fmt.Println("Size of b2:", unsafe.Sizeof(b2))
	fmt.Println("Distance b1 to b2:", uintptr(e2)-uintptr(e1))
	// Явно видно, что размер массивов равен объявленному нами — 257.
	// Это на единицу больше, чем доступно в 18-м классе-размере.
	// А вот расстояние в памяти между ними 288 байт,
	// что как раз равно размерности 19-го класса размера.
}

func TestRunFunction(t *testing.T) {
	pointTest := func(t *int) uintptr {
		fmt.Println("pointer of t", &t)
		fmt.Println("pointer in func before", t)
		z := 321
		t = &z
		fmt.Println("pointer in func after", t)
		return uintptr(unsafe.Pointer(&t))
	}

	x := 123

	fmt.Println("pointer in main", &x)
	p := pointTest(&x)
	fmt.Println(x)
	fmt.Println("Distance main to func:", uintptr(unsafe.Pointer(&x))-p)
}

func TestRunStructure(t *testing.T) {
	// можем преобразовать любую структуру в массив байтов
	//  «сортируйте поля в ваших» структурах по размеру
	type struct4 struct {
		b1 byte
		b2 byte
		i  int16
	}

	x := struct4{
		b1: 2,
		b2: 4,
		i:  261,
	}

	fmt.Println(*(*[4]byte)(unsafe.Pointer(&x)))
}

func TestRunInterface(t *testing.T) {
	type iface struct {
		t     unsafe.Pointer
		value unsafe.Pointer
	}

	var x any

	x = 10

	i := *(*iface)(unsafe.Pointer(&x))
	v := *(*int)(i.value)

	fmt.Println(i)
	fmt.Println(v)
}

func TestRunSlice(t *testing.T) {
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := unsafe.Pointer(&x)
	fmt.Println(*(*slice)(s1))
	fmt.Println(*(*[9]int)((*(*slice)(s1)).array))
}

func TestRunSliceTwo(t *testing.T) {
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x2 := x1[3:6]
	x2[2] = 99

	s1 := unsafe.Pointer(&x1)
	fmt.Println("x1 pointer", s1)
	fmt.Println("x1 slice", *(*slice)(s1))
	fmt.Println("x1 array", *(*[9]int)((*(*slice)(s1)).array))

	s2 := unsafe.Pointer(&x2)
	fmt.Println("x2 pointer", s2)
	fmt.Println("x2 slice", *(*slice)(s2))
	fmt.Println("x2 array", *(*[6]int)((*(*slice)(s2)).array))
}

func TestRunMap(t *testing.T) {
	type hmap struct {
		count     int
		flags     uint8
		B         uint8
		noverflow uint16
		hash0     uint32

		buckets    unsafe.Pointer
		oldbuckets unsafe.Pointer
		nevacuate  uintptr
	}

	x := make(map[int]int)
	for i := 0; i < 8; i++ {
		x[50+i] = 10 + i
	}

	s1 := unsafe.Pointer(&x)
	hm1 := *(*hmap)(*(*unsafe.Pointer)(s1))
	fmt.Printf("hash map 1 %+v", hm1)
	fmt.Println(*(*[17]int)(hm1.buckets))
	fmt.Println(*(*[16]int)(unsafe.Pointer(uintptr(hm1.buckets) + uintptr(0x8))))

	x[5] = 2 // тут мапа будет расти

	s2 := unsafe.Pointer(&x)
	hm2 := *(*hmap)(*(*unsafe.Pointer)(s2))
	fmt.Printf("hash map 2 %+v", hm2)
	fmt.Println(*(*[16]int)(unsafe.Pointer(uintptr(hm2.buckets) + uintptr(0x8))))
}

func TestRunChanel(t *testing.T) {
	type hchan struct {
		qcount   uint
		dataqsiz uint
		buf      unsafe.Pointer
		elemsize uint16
		closed   uint32
		elemtype unsafe.Pointer
		sendx    uint
		recvx    uint
	}
	// qcount — количество элементов в канале;
	// buf — ссылка на буфер с элементами;
	// closed — флаг закрытости канала.

	ch := make(chan int, 100)

	ch <- 3
	ch <- 2
	ch <- 1
	<-ch
	ch <- 4

	s1 := unsafe.Pointer(&ch)
	fmt.Println(s1)
	fmt.Printf("chan %+v\n", *(*hchan)(*(*unsafe.Pointer)(s1)))
	p := *(*hchan)(*(*unsafe.Pointer)(s1))
	fmt.Println(*(*[10]int)(p.buf))
}
