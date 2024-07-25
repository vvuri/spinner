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
