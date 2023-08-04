import random
import time
import numpy as np
# BetzyYarinR
def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        # La variable swapped nos ayudará a determinar si el arreglo ya está ordenado
        swapped = False
        for j in range(0, n-i-1):
            # Compara elementos adyacentes y los intercambia si el primero es mayor que el segundo
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = True
        # Si no hubo intercambios en la iteración, el arreglo ya está ordenado, por lo que podemos detener el algoritmo
        if not swapped:
            break

def process(n, total):
    # Generamos un arreglo de ejemplo con números aleatorios
    random_list = [random.randint(inicio_random, fin_random) for _ in range(total)]

    # Medimos el tiempo de procesamiento del algoritmo en segundos
    start_time = time.time() 
    bubble_sort(random_list)
    end_time = time.time() 

    # Calculamos el tiempo de procesamiento en segundos
    processing_time_ms = end_time - start_time
    avg_list.append(float("{:.6f}".format(processing_time_ms))) 
    print(f"Prueba N° [{n+1}] = {processing_time_ms:.6f}")
    start_time = end_time = 0

if __name__ == "__main__":
    nro_pruebas = 5
    inicio_random = 1
    fin_random = 100000    
    array_test = [100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000]
    #array_test = [9000]
    for x in range(len(array_test)):
        print(f"Se ha generado un array de {array_test[x]} elementos aleatorios entre : {inicio_random} y {fin_random}")
        avg_list = []
        for i in range(nro_pruebas):  
            process(i,array_test[x])  
        avg_list = np.array(avg_list)
        promedio = np.mean(avg_list)
        desviacion_estandar = np.std(avg_list)
        print(f"\nPromedio = {promedio:.6f} | Desviacion Estandar = {desviacion_estandar:.6f}")

# BetzyYarinR