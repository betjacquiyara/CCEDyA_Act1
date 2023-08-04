#include <iostream>
#include <vector>
#include <algorithm>
#include <chrono>
#include <ctime>
#include <cstdlib>
#include <cmath>
#include <iomanip>

using namespace std;
// BetzyYarinR
// Implementación del algoritmo Quick Sort
void quickSort(vector<int>& arr, int low, int high) {
    if (low < high) {
        int pivot = arr[high];
        int i = low - 1;

        for (int j = low; j < high; j++) {
            if (arr[j] <= pivot) {
                i++;
                swap(arr[i], arr[j]);
            }
        }
        swap(arr[i+1], arr[high]);

        quickSort(arr, low, i);
        quickSort(arr, i + 2, high);
    }
}

// Generar un vector de tamaño n con elementos aleatorios
vector<int> generateRandomVector(int n) {
    vector<int> arr(n);
    for (int i = 0; i < n; i++) {
        arr[i] = rand() % 100000; // Genera números aleatorios 
    }
    return arr;
}

// Función para calcular el tiempo de procesamiento en segundos
double timeElapsed(const chrono::time_point<chrono::high_resolution_clock>& start) {
    auto end = chrono::high_resolution_clock::now();
    return chrono::duration_cast<chrono::microseconds>(end - start).count() / 1000000.0;
}

double calculateAverage(float sum, int nPruebas)  {
	return sum / nPruebas;
}

double calculateStandardDeviation(vector<double>& arr, double prom, int nItems)  {
	 // Calcular la suma de los cuadrados de las diferencias entre cada elemento y la media
    double sumaCuadradosDiferencias = 0.0;
    for (const double& elemento : arr) {
        double diferencia = elemento - prom;
        sumaCuadradosDiferencias += diferencia * diferencia;
    }
    // Dividir la suma de los cuadrados por el número de elementos en el array menos 1
    double varianza = sumaCuadradosDiferencias / (nItems - 1);
    // Calcular la raíz cuadrada para obtener la Desviación Estándar
    double desviacionEstandar = std::sqrt(varianza);
    return desviacionEstandar;
}

int main() {
    int nro_pruebas = 5;
    double sum_time =0.0, num_tp=0.0, promedio=0.0, desvEst = 0;
    srand(time(0)); // Semilla para números aleatorios

    // Tamaños de los arrays
    vector<int> arraySizes = {100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000};

    // Iterar sobre los tamaños de los arrays
    for (int size : arraySizes) {
        std::vector<double> arraytp;
        for(int i = 0; i < nro_pruebas; i++){
            // Generar un vector aleatorio de tamaño 'size'
            vector<int> arr = generateRandomVector(size);
            // Imprimir el Array en desorden
            /**for (int i = 0; i < size; i++) {
                std::cout << arr[i] << " ";
            }**/ 
            // Medir el tiempo de procesamiento
            chrono::time_point<chrono::high_resolution_clock> startTime = chrono::high_resolution_clock::now();
            quickSort(arr, 0, size - 1);
            double elapsedTime = timeElapsed(startTime);
            // Imprimir el Array en orden
            /**for (int i = 0; i < size; i++) {
                std::cout << arr[i] << " ";
            }**/
            num_tp = std::floor(elapsedTime * 1000000) / 1000000;
            std::cout << "N Entrada: " << size << "  TP(seg): " << std::fixed << std::setprecision(6) << num_tp << std::endl;
            sum_time += num_tp;
            arraytp.push_back(num_tp);
        }
        promedio = calculateAverage(sum_time, nro_pruebas);
        desvEst = calculateStandardDeviation(arraytp, promedio, nro_pruebas);
        std::cout << "Promedio: " << std::fixed << std::setprecision(6) << promedio << std::endl;
        std::cout << "Desviacion Estandar: " << std::fixed << std::setprecision(6) << desvEst << std::endl;
        std::cout << "\n"; 
        sum_time =0.0, num_tp=0.0, promedio=0.0, desvEst = 0.0;
    }
    return 0;
}
// BetzyYarinR