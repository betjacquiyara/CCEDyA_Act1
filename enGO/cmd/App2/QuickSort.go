package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// BetzyYarinR
func main() {
	nro_pruebas := 5
	arraySizes := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}

	// Iterar sobre los tamaños de los arrays
	for _, n := range arraySizes {

		arraytp := []float64{}
		//fmt.Println("QS Array iniciando", arraytp)
		sum_time := 0.0
		for i := 0; i < nro_pruebas; i++ {
			// Generar un array aleatorio de tamaño 'n'
			values := generateRandomArray(n)
			// Mostrar el array antes de ordenarlo (opcional)
			//fmt.Println("Array desordenado:\n", values)

			// Copiar el array para realizar el QuickSort sin modificar el original
			valuesCopy := make([]int, n)
			copy(valuesCopy, values)
			// Medir el tiempo de procesamiento
			startTime := time.Now()
			quickSort(valuesCopy)
			elapsedTime := time.Since(startTime).Seconds()

			// Mostrar el array ordenado (opcional)
			//fmt.Println("Array ordenado:", valuesCopy)
			// Mostrar el tiempo de procesamiento para este array
			num_tp := fmt.Sprintf("%.6f", elapsedTime)
			fmt.Println("N Entrada: ", n, "  TP(seg): ", num_tp)

			tiempo, err := strconv.ParseFloat(num_tp, 64)
			if err != nil {
				// Manejar el error en caso de que la conversión falle
				fmt.Println("Error al convertir el número:", err)
				return
			}
			sum_time += tiempo
			//fmt.Println("Suma de promedios: ", sum_time)
			// Agregar elementos al slice
			arraytp = append(arraytp, tiempo)
			//fmt.Println("QS Array agregando...\n", arraytp)
		}
		//fmt.Printf("Promedio manual: %.6f ", sum_time/5)
		promedio := calculateAverage(sum_time, nro_pruebas)
		desvEst := calculateStandardDeviation(arraytp, promedio)
		fmt.Printf("Promedio: %.6f \n", promedio)
		fmt.Printf("Desviacion Estandar: %.6f \n\n", desvEst)
	}
}

// Implementación del algoritmo QuickSort
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left := 0
	right := len(arr) - 1

	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quickSort(arr[:right+1])
	quickSort(arr[left:])
}

// Generar un array de tamaño n con elementos aleatorios
func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100000) // Genera números aleatorios entre 0 y 100000
	}
	return arr
}

// Función para calcular el Promedio del tiempo de poocesamiento
func calculateAverage(sum float64, nPruebas int) float64 {
	return sum / float64(nPruebas)
}

// Función para calcular la Desviación Estándar
func calculateStandardDeviation(arr []float64, prom float64) float64 {
	variance := 0.0
	for _, value := range arr {
		variance += math.Pow(value-prom, 2)
	}
	variance /= float64(len(arr))
	return math.Sqrt(variance)
}

// BetzyYarinR
