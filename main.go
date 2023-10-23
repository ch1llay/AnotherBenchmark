package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func initBenchmark(cpuCountPercent int) int {

	if cpuCountPercent > 100 {
		cpuCountPercent = 100
	}

	cpuCount := runtime.NumCPU()
	newCpuCount := cpuCount * cpuCountPercent / 100
	runtime.GOMAXPROCS(newCpuCount)

	goroutineAmount := newCpuCount * 2

	fmt.Printf("Исползуем %d логических ядер, и %d горутин\n", newCpuCount, goroutineAmount)

	return newCpuCount * 2

}
func timer(wg *sync.WaitGroup, hardTimeSeconds int) {
	startTime := time.Now()
	curSeconds := 0
	curDiff := hardTimeSeconds
	for curSeconds < hardTimeSeconds {
		curSeconds = int(time.Now().Sub(startTime).Seconds())
		dif := hardTimeSeconds - curSeconds
		if curDiff != curSeconds {
			curDiff = curSeconds
		}

		if curDiff%10 == 0 {
			fmt.Printf("До бешбармаков осталось %d секунд\n", dif)
		}

	}
	fmt.Printf("Прошло %d секунд\n", hardTimeSeconds)
	wg.Done()

}
func startBenchmark(goroutineAmount int, hardTimeSeconds int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go timer(&wg, hardTimeSeconds)
	points := make([]int64, goroutineAmount, goroutineAmount)

	for i := 0; i < goroutineAmount; i++ {
		ii := i
		go func() {
			benchmarkFunc(ii, points)
		}()
	}

	wg.Wait()
	core := max(points)
	all := sum(points)
	fmt.Printf("core - %d, all - %d", core, all)
}

func benchmarkFunc(i int, points []int64) {
	for {
		n := time.Now().UnixNano()
		monsterNum := n%(rand.Int63n(int64(math.Sqrt(float64(n))+1)%100)+1)/(rand.Int63n(1000)+1) + int64(math.Log(float64(n/10000)))
		var megaMonsterNum = monsterNum*int64(math.Sin(float64(n))) + n%1000
		megaMonsterNum++
		points[i]++
	}

}

func max(slice []int64) int64 {
	maxNum := slice[0]

	for _, element := range slice {
		maxNum = int64(math.Max(float64(maxNum), float64(element)))
	}

	return maxNum
}

func sum(arr []int64) int64 {
	var s int64
	for _, valueInt := range arr {
		s += valueInt
	}
	return s
}

func main() {

	var hardTimeSeconds int
	var defaultCpuCountPercent int

	fmt.Println("Подготавливаем процессор к приготовлению бешбармаков")
	fmt.Println("Введите время(в секундах) для подготовки процессора к приготовлению бешбармаков")
	_, err := fmt.Scanf("%d\n", &hardTimeSeconds)

	if err != nil {
		return
	}

	fmt.Println("Введите число в процентах (0-100) на сколько хорошо процессор будет готовиться к приготовлению бешбармаков")
	_, err = fmt.Scanf("%d\n", &defaultCpuCountPercent)
	if err != nil {
		return
	}

	goroutineAmount := initBenchmark(defaultCpuCountPercent)

	startBenchmark(goroutineAmount, hardTimeSeconds)

	var s string
	fmt.Scanf("%s,\n", s)
}
