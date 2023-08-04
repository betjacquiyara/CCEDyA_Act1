package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// BetzyYarinR
// Implementación del algoritmo Selection Sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// Generar un array de tamaño n con elementos aleatorios
func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100000) // Genera números aleatorios
	}
	return arr
}

func main() {
	nro_pruebas := 5
	arraySizes := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}
	//arraySizes := []int{5, 6, 8, 12} array de prueba

	// Iterar sobre los tamaños de los arrays
	for _, n := range arraySizes {
		arraytp := []float64{}
		sum_time := 0.0
		for i := 0; i < nro_pruebas; i++ {
			// Generar un array aleatorio de tamaño 'n'
			values := generateRandomArray(n)
			//fmt.Println("Array desordenado:", values)
			// Copiar el array para realizar el Selection Sort sin modificar el original
			valuesCopy := make([]int, n)
			copy(valuesCopy, values)

			// Medir el tiempo de procesamiento
			startTime := time.Now()
			selectionSort(valuesCopy)
			elapsedTime := time.Since(startTime).Seconds()

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
			// Agregar elementos al slice
			arraytp = append(arraytp, tiempo)
		}
		promedio := calculateAverage(sum_time, nro_pruebas)
		desvEst := calculateStandardDeviation(arraytp, promedio)
		fmt.Printf("Promedio: %.6f \n", promedio)
		fmt.Printf("Desviacion Estandar: %.6f \n\n", desvEst)
	}
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
