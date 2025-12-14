package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	idle1, total1, err := readCPUinfo()
	if err != nil {
		fmt.Println("Error de lectura 1", err)
		return
	}

	fmt.Println("Midiendo CPU por 1 segundo...")
	time.Sleep(1 * time.Second)

	idle2, total2, err := readCPUinfo()
	if err != nil {
		fmt.Println("Error de lectura 2", err)
		return
	}

	deltaIdle := idle2 - idle1
	deltaTotal := total2 - total1

	usage := (float64(deltaTotal-deltaIdle) / float64(deltaTotal)) * 100

	fmt.Printf("Uso de CPU: %.2f%%\n", usage)
}

func readCPUinfo() (uint64, uint64, error) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	if s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)

		if len(fields) > 0 && fields[0] == "cpu" {
			var total uint64 = 0
			var idle uint64 = 0

			for i, field := range fields[1:] {
				val, err := strconv.ParseUint(field, 10, 64)
				if err != nil {
					return 0, 0, err
				}
				total += val
				if i == 3 {
					idle = val
				}
			}
			return idle, total, nil
		}
	}
	return 0, 0, fmt.Errorf("invalid /proc/stat format")
}
