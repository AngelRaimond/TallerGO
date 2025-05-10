package main

import (
	"fmt"
	"sync"
)

// factorial calcula el factorial de un número
func factorial(n int) uint64 {
	if n <= 1 {
		return 1
	}
	result := uint64(1)
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	return result
}

// worker procesa números del canal jobs y envía resultados al canal results
func worker(id int, jobs <-chan int, results chan<- struct {
	number int
	result uint64
}, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		result := factorial(n)
		results <- struct {
			number int
			result uint64
		}{n, result}
	}
}

func main() {
	// Números de entrada para calcular el factorial
	numbers := []int{5, 7, 10, 3, 8, 12, 4, 6}

	// Crear canales para trabajos y resultados
	jobs := make(chan int, len(numbers))
	results := make(chan struct {
		number int
		result uint64
	}, len(numbers))

	// Crear pool de trabajadores
	numWorkers := 3
	var wg sync.WaitGroup

	// Iniciar trabajadores
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Enviar trabajos a los workers
	for _, n := range numbers {
		jobs <- n
	}
	close(jobs)

	// Iniciar una goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recolectar e imprimir resultados
	fmt.Println("Resultados de los factoriales:")
	for result := range results {
		fmt.Printf("%d! = %d\n", result.number, result.result)
	}
}
