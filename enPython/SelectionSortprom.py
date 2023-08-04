import random
import time
import numpy as np
# BetzyYarinR
def selection_sort(arr):
    n = len(arr)
    for i in range(n):
        min_idx = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr
   
def process(n, total):
    
    random_list = [random.randint(inicio_random, fin_random) for _ in range(total)]
    #print(f"Array original: {random_list}")
    start_time = time.time()  
    sorted_array = selection_sort(random_list)
    end_time = time.time() 
    processing_time_ms = end_time - start_time
    avg_list.append(float("{:.6f}".format(processing_time_ms))) 
    print(f"Prueba N° [{n+1}] = {processing_time_ms:.6f}")
    start_time = end_time = 0
    #print(f"Array ordenado: {sorted_array}\n")

if __name__ == "__main__":
    nro_pruebas = 5
    inicio_random = 1
    fin_random = 100000
    array_test = [100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000]
    #array_test = [200,130,1000]
    for x in range(len(array_test)):
        print(f"Se ha generado un array de {array_test[x]} elementos aleatorios entre : {inicio_random} y {fin_random}")
        avg_list = []
        for i in range(nro_pruebas):       
            process(i, array_test[x])
        avg_list = np.array(avg_list)
        promedio = np.mean(avg_list)
        desviacion_estandar = np.std(avg_list)
        print(f"\nPromedio = {promedio:.6f} | Desviacion Estandar = {desviacion_estandar:.6f}")
# BetzyYarinR