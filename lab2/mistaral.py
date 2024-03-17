import numpy as np
import matplotlib.pyplot as plt
import time

# Параметры системы
k = 2
l = 8
m = 2
n = 10
kt = 100
b = 26000
i1 = 9
i2 = 2
s = 200
V = 800
T = 10
x1 = 0.3
x2 = 0.3
x3 = 0
x4 = 0
x5 = 500
Amax = 0.5

# Функция alfa_star
def alfa_star(alfa):
    if alfa <= Amax:
        return alfa
    else:
        return Amax * np.sign(alfa)

# Функция правых частей системы уравнений
def f(vector, t):
    F = np.zeros(5)
    alfa = vector[1] - vector[0]
    F[0] = k * alfa_star(alfa)
    F[1] = vector[2]
    F[2] = l * vector[0] - l * vector[1] - m * vector[2] + n * vector[3]
    theta = (10000 - vector[4]) / (b - V * t)
    F[3] = -kt * vector[3] - i1 * vector[1] - i2 * vector[2] + s * (theta - vector[0])
    F[4] = V * np.sin(vector[0])
    return F

# Метод Рунге-Кутты 4-го порядка
def rk4(f, vector, t, h):
    k1 = f(vector, t)
    k2 = f(vector + h * k1 / 2, t + h / 2)
    k3 = f(vector + h * k2 / 2, t + h / 2)
    k4 = f(vector + h * k3, t + h)
    return vector + h * (k1 + 2 * k2 + 2 * k3 + k4) / 6

# Решение системы дифференциальных уравнений
def solve(h):
    t = 0
    vector = np.array([x1, x2, x3, x4, x5])
    t_vals = [t]
    vector_vals = [vector]
    while t < T:
        t += h
        vector = rk4(f, vector, t, h)
        t_vals.append(t)
        vector_vals.append(vector)
    return np.array(t_vals), np.array(vector_vals)

# Решение системы с заданным шагом
h = 0.01
t_vals, vector_vals = solve(h)

# Вывод графиков xi(t)
for i in range(5):
    plt.plot(t_vals, vector_vals[:, i])
plt.xlabel('t')
plt.ylabel('x(t)')
plt.show()

# Вывод значений переменных состояния в конце интервала интегрирования и относительной погрешности
vector_T = vector_vals[-1]
delta = np.abs((vector_T - vector_vals[-2]) / vector_T)
print("x(T) = ", vector_T)
print("Relative error = ", delta)

# Анализ зависимости точности и трудоемкости от шага интегрирования
h_vals = np.logspace(-4, -1, 10)
delta_vals = []
time_vals = []
for h in h_vals:
    t0 = time.time()
    t_vals, vector_vals = solve(h)
    t1 = time.time()
    vector_T = vector_vals[-1]
    delta = np.abs((vector_T - vector_vals[-2]) / vector_T)
    delta_vals.append(delta)
    time_vals.append(t1 - t0)

# Вывод графиков зависимостей относительной погрешности и оценки трудоемкости от величины шага h
plt.semilogx(h_vals, delta_vals)
plt.xlabel('h')
plt.ylabel('Relative error')
plt.show()

plt.semilogx(h_vals, time_vals)
plt.xlabel('h')
plt.ylabel('Time')
plt.show()

# Автоматический выбор величины шага интегрирования для достижения относительной погрешности не более 1%
eps = 0.01
h_min = 1e-6
h_max = 1
h = (h_min + h_max) / 2
while True:
    t_vals, vector_vals = solve(h)
    vector_T = vector_vals[-1]
    if np.mean(np.abs(vector_T)) != 0:
        delta = np.mean(np.abs(vector_T - vector_vals[-2])) / np.mean(np.abs(vector_T))
    else:
        delta = np.inf
    if delta <= eps:
        break
    elif delta > eps and h >= h_max:
        h_max = 2 * h
        h = (h_min + h_max) / 2
    elif delta > eps and h < h_max:
        h_min = h
        h = (h_min + h_max) / 2

# Вывод итоговых результатов
t_vals, vector_vals = solve(h)
vector_T = vector_vals[-1]
delta = np.abs((vector_T - vector_vals[-2]) / vector_T)
print("Step size = ", h)
print("x(T) = ", vector_T)
print("Relative error = ", delta)
