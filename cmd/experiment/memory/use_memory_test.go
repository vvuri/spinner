package memory

import (
	"fmt"
	"runtime"
	"testing"
	"time"
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
