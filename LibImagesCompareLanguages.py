import matplotlib.pyplot as plt

# Valores en el eje X Tamaños de entrada
x = [100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000, 50000]

#Promedios de Algoritmo Burbuja en Python
#y1 = [0.000479,0.045480,0.178412,0.417040,0.709100,1.137107,1.598331,2.203492,2.904818,3.584382,4.521556,18.006864,40.850948,74.645574,117.585469]
#Promedio de Algoritmo Burbuja en Golang
#y2 = [0.000000,0.000691,0.002436,0.006084,0.010834,0.017216,0.023767,0.034475,0.042540,0.054674,0.070389,0.459646,1.262913,3.157021,4.303893]
#Promedio de Algoritmo Burbuja en C++
#y3 = [0.000045, 0.003888, 0.020262, 0.040093, 0.073201, 0.106796, 0.157617, 0.219056, 0.281735, 0.360846, 0.446404, 2.061460, 4.518771, 14.913841, 21.809275]

#Promedio de Algoritmo Quick Sort Python
#y4 = [0.000201, 0.002110, 0.004997, 0.006291, 0.009177, 0.011786, 0.013992, 0.016472, 0.019708, 0.023137, 0.024815, 0.052479, 0.069037, 0.085803, 0.107082]
#Promedio de Algoritmo Quick Sort en Golang
#y5 = [0.000000, 0.000125, 0.000412, 0.000525, 0.000406, 0.000432, 0.000600, 0.000695, 0.000965, 0.000312, 0.000105, 0.001116, 0.004019, 0.002109, 0.005453]
#Promedio de Algoritmo Quick Sort en C++
#y6 = [0.000008, 0.000129, 0.000298, 0.000368, 0.000964, 0.001062, 0.001295, 0.001369, 0.001491, 0.001754, 0.002355, 0.004488, 0.007978, 0.008828, 0.010853]

#Promedio de Algoritmo Merge Sort Python
#y7 = [0.000100, 0.002285, 0.004450, 0.006505, 0.009359, 0.012793, 0.015398, 0.016159, 0.018647, 0.022962, 0.025098, 0.057421, 0.080890, 0.112015, 0.137197]
#Promedio de Algoritmo Merge Sort en Golang
#y8 = [0.000108, 0.000249, 0.000449, 0.000777, 0.000987, 0.000545, 0.000917, 0.000422, 0.002008, 0.005437, 0.000900, 0.003971, 0.007809, 0.007563, 0.008518]
#Promedio de Algoritmo Merge Sort en C++
#y9 = [0.000059, 0.000565, 0.001163, 0.002485, 0.004167, 0.005608, 0.005760, 0.006567, 0.007915, 0.008410, 0.012026, 0.019554, 0.026017, 0.033197, 0.043074]

#Promedio de Algoritmo Selection Sort Python
y10 = [0.000199, 0.021535, 0.076218, 0.157918, 0.281186, 0.437503, 0.635146, 0.872537, 1.163518, 1.435196, 1.758675, 7.189240, 16.416556, 31.755306, 51.601579]
#Promedio de Algoritmo Selection Sort en Golang
y11 = [0.000000, 0.000590, 0.003146, 0.006548, 0.011019, 0.017917, 0.026804, 0.033217, 0.046790, 0.053022, 0.069445, 0.271104, 0.591428, 1.014313, 1.612249]
#Promedio de Algoritmo Selection Sort en C++
y12 = [0.000022, 0.001700, 0.007110, 0.017355, 0.031897, 0.045371, 0.065459, 0.082890, 0.110462, 0.134743, 0.166781, 0.626202, 1.412749, 2.476501, 3.926257]

# Crear el gráfico
#plt.plot(x, y1, 'o-', label='Python', color='red', linewidth=2)
#plt.plot(x, y2, 'o-', label='Golang', color="green", linewidth=2)
#plt.plot(x, y3, 'o-', label='C++', color='skyblue', linewidth=4)
#plt.plot(x, y4, 'o-', label='Python', color="red", linewidth=2)
#plt.plot(x, y5, 'o-', label='Golang', color='green', linewidth=2)
#plt.plot(x, y6, 'o-', label='C++', color='purple', linewidth=4)
#plt.plot(x, y7, 'o-', label='Python', color='red', linewidth=2)
#plt.plot(x, y8, 'o-', label='Golang', color="green", linewidth=2)
#plt.plot(x, y9, 'o-', label='C++', color='orange', linewidth=4)
plt.plot(x, y10, 'o-', label='Python', color="red", linewidth=2)
plt.plot(x, y11, 'o-', label='Golang', color="green", linewidth=2)
plt.plot(x, y12, 'o-', label='C++', color='purple', linewidth=4)
plt.xticks(x)
plt.xticks(fontsize='x-small')
# Agregar títulos y etiquetas
#plt.title('Algoritmo Bubble Sort')
#plt.title('Algoritmo Quick Sort')
#plt.title('Algoritmo Merge Sort')
plt.title('Algoritmo Selection Sort')
plt.xlabel('Valores Cantidad de Elementos')
plt.ylabel('Tiempo de Procesamiento (seg)')
plt.legend()

# Mostrar el gráfico
plt.show()