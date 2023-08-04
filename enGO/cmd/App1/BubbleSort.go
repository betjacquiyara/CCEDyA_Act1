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

	// Semilla para generar números aleatorios diferentes en cada ejecución
	rand.Seed(time.Now().UnixNano())

	// Tamaños de los arrays
	arraySizes := []int{100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000}

	// Bucle para procesar los arrays con diferentes tamaños
	for _, n := range arraySizes {
		arraytp := []float64{}
		sum_time := 0.0
		for i := 0; i < nro_pruebas; i++ {
			// Generar un nuevo array con valores aleatorios
			values := generateRandomArray(n)

			// Mostrar el array antes de ordenarlo (opcional)
			//fmt.Println("Array desordenado:", values)

			// Copia el array para medir el tiempo de procesamiento sin modificar el original
			valuesCopy := make([]int, n)
			copy(valuesCopy, values)

			// Medir el tiempo de procesamiento del algoritmo de burbuja
			startTime := time.Now()
			bubbleSort(valuesCopy)
			elapsedTime := time.Since(startTime)

			// Mostrar el array ordenado (opcional)
			//fmt.Println("Array ordenado:", valuesCopy)
			// Mostrar el tiempo de procesamiento en segundos para este array
			num_tp := fmt.Sprintf("%.6f", elapsedTime.Seconds())
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
		}
		//fmt.Printf("Promedio manual: %.6f ", sum_time/5)
		//promedio := fmt.Sprintf("%.6f", calculateAverage(sum_time, nro_pruebas))
		//desvEst := fmt.Sprintf("%.6f", calculateStandardDeviation(arraytp, promedio))
		promedio := calculateAverage(sum_time, nro_pruebas)
		desvEst := calculateStandardDeviation(arraytp, promedio)
		fmt.Printf("Promedio: %.6f \n", promedio)
		fmt.Printf("Desviacion Estandar: %.6f \n\n", desvEst)
	}
}

// Función para generar un array de tamaño n con valores aleatorios entre 0 y 100000
func generateRandomArray(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(100000) // Genera números aleatorios entre 0 y 100000
	}
	return array
}

// Implementación del algoritmo de ordenación burbuja
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Intercambiar los elementos si están en el orden incorrecto
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
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