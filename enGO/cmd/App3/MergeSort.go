package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// BetzyYarinR
// Generar un array de tamaño n con elementos aleatorios
func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100000) // Genera números aleatorios entre 0 y 100000
	}
	return arr
}

func main() {
	nro_pruebas := 5
	arraySizes := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}

	// Iterar sobre los tamaños de los arrays
	for _, n := range arraySizes {
		arraytp := []float64{}
		sum_time := 0.0
		for i := 0; i < nro_pruebas; i++ {
			// Generar un array aleatorio de tamaño 'n'
			values := generateRandomArray(n)
			//fmt.Println("Array desordenado:", values)
			// Copiar el array para realizar el Merge Sort sin modificar el original
			valuesCopy := make([]int, n)
			copy(valuesCopy, values)

			// Medir el tiempo de procesamiento
			startTime := time.Now()
			valuesCopy = mergeSort(valuesCopy)
			elapsedTime := time.Since(startTime).Seconds()
			//fmt.Println("Array ordenado:", valuesCopy)

			num_tp := fmt.Sprintf("%.6f", elapsedTime)
			fmt.Println("N Entrada: ", n, "  TP(seg): ", num_tp)
			tiempo, err := strconv.ParseFloat(num_tp, 64)
			if err != nil {
				// Manejar el error en caso de que la conversión falle
				fmt.Println("Error al convertir el número:", err)
				return
			}
			sum_time += tiempo
			arraytp = append(arraytp, tiempo)
		}
		promedio := calculateAverage(sum_time, nro_pruebas)
		desvEst := calculateStandardDeviation(arraytp, promedio)
		fmt.Printf("Promedio: %.6f \n", promedio)
		fmt.Printf("Desviacion Estandar: %.6f \n\n", desvEst)
	}
}

// Implementación del algoritmo Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// Función auxiliar para combinar dos slices ordenados
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			merged = append(merged, left[l])
			l++
		} else {
			merged = append(merged, right[r])
			r++
		}
	}

	merged = append(merged, left[l:]...)
	merged = append(merged, right[r:]...)
	return merged
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
