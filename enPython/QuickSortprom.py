import random
import time
import numpy as np
# BetzyYarinR
def quick_sort(arr):
    if len(arr) <= 1:
        return arr

    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]

    return quick_sort(left) + middle + quick_sort(right)

def process(n, total):
    # Generamos un arreglo de ejemplo con números aleatorios
    random_list = [random.randint(inicio_random, fin_random) for _ in range(total)]

    start_time = time.time()  
    sorted_array = quick_sort(random_list)
    end_time = time.time() 
    processing_time_ms = end_time - start_time

    avg_list.append(float("{:.6f}".format(processing_time_ms))) 
    print(f"Prueba N° [{n+1}] = {processing_time_ms:.6f}")
    start_time = end_time = 0
    #print("Array original:")
    #print(random_list)  
    #print("Array ordenado:")
    #print(sorted_array)

if __name__ == "__main__":
    nro_pruebas = 5
    inicio_random = 1
    fin_random = 100000
    array_test = [100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000]
    
    for x in range(len(array_test)):
        print(f"\nSe ha generado un array de {array_test[x]} elementos aleatorios entre : {inicio_random} y {fin_random}")
        avg_list = []
        for i in range(nro_pruebas):        
            process(i,array_test[x])
        avg_list = np.array(avg_list)
        promedio = np.mean(avg_list)
        desviacion_estandar = np.std(avg_list)
        print(f"\nPromedio = {promedio:.6f} | Desviacion Estandar = {desviacion_estandar:.6f}")
# BetzyYarinR