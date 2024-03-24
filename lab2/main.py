from matplotlib import pyplot as plt
import numpy as np
import math

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

prev_step = []
prev_lost = []


def alfa_star(alfa):
    if np.abs(alfa) <= Amax:
        return alfa
    return Amax * np.sign(alfa)


def system(vector):
    F = np.zeros(5)
    t = T
    alfa = vector[1] - vector[0]
    F[0] = k * alfa_star(alfa)
    F[1] = vector[2]
    F[2] = l * vector[0] - l * vector[1] - m * vector[2] + n * vector[3]
    theta = (10000 - vector[4]) / (b - V * t)
    F[3] = -kt * vector[3] - i1 * vector[1] - i2 * vector[2] + s * theta - vector[0] * s
    F[4] = V * np.sin(vector[0])
    return F

def Euler(vector, step_size, calculation_end):
    step_count = int(calculation_end / step_size) + 1
    result = [vector]
    for i in range(1, step_count):
        result.append(result[-1] + step_size * system(result[-1]))
    return result

def fixed_step():
    step_size = prev_step[-1]

    R1 = Euler(np.array([x1, x2, x3, x4, x5]), step_size, T)
    R2 = Euler(np.array([x1, x2, x3, x4, x5]), step_size / 2, T)
    steps = np.arange(0, T, step_size)

    res_x1 = [R1[i][0] for i in range(0, len(steps))]
    res_x2 = [R1[i][1] for i in range(0, len(steps))]
    res_x3 = [R1[i][2] for i in range(0, len(steps))]
    res_x4 = [R1[i][3] for i in range(0, len(steps))]
    res_x5 = [R1[i][4] for i in range(0, len(steps))]

    plt.plot(steps, res_x1, label='x1')
    plt.plot(steps, res_x2, label='x2')
    plt.plot(steps, res_x3, label='x3')
    plt.plot(steps, res_x4, label='x4')
    plt.plot(steps, res_x5, label='x5')

    plt.xlabel('Интервал')
    plt.ylabel('Значения')
    plt.legend()
    plt.show()

    print(f'Погрешность переменной Х1 при шаге {round(step_size, 5)} составляет {round(abs((R2[-1][0] - R1[-1][0]) / R2[-1][0]) * 100, 5)}%')
    print(f'Погрешность переменной X2 при шаге {round(step_size, 5)} составляет {round(abs((R2[-1][1] - R1[-1][1]) / R2[-1][1]) * 100, 5)}%')
    print(f'Погрешность переменной X3 при шаге {round(step_size, 5)} составляет {round(abs((R2[-1][2] - R1[-1][2]) / R2[-1][2]) * 100, 5)}%')
    print(f'Погрешность переменной X4 при шаге {round(step_size, 5)} составляет {round(abs((R2[-1][3] - R1[-1][3]) / R2[-1][3]) * 100, 5)}%')
    print(f'Погрешность переменной X5 при шаге {round(step_size, 5)} составляет {round(abs((R2[-1][4] - R1[-1][4]) / R2[-1][4]) * 100, 5)}%')


def dynamic_step():
    curr_step_size = 1
    first_step_size = curr_step_size
    while curr_step_size > 0:
        R1 = Euler(np.array([x1, x2, x3, x4, x5]), curr_step_size, T)
        R2 = Euler(np.array([x1, x2, x3, x4, x5]), curr_step_size / 2, T)
        sigma = abs((R2[-1][4] - R1[-1][4]) / R2[-1][4]) * 100  # Calculate sigma based on x5
        # print(f'Размер шага {round(curr_step_size, 6)} потери {round(sigma, 3)}%')
        prev_step.append(curr_step_size)
        prev_lost.append(sigma)
        if sigma < 1 and abs((R2[-1][0] - R1[-1][0]) / R2[-1][0]) * 100 < 1 and abs(
                (R2[-1][1] - R1[-1][1]) / R2[-1][1]) * 100 < 1 and abs(
                (R2[-1][2] - R1[-1][2]) / R2[-1][2]) * 100 < 1 and abs((R2[-1][3] - R1[-1][3]) / R2[-1][3]) * 100 < 1:
            break
        curr_step_size -= 0.0001

    fig, ax = plt.subplots()
    ax.plot(prev_step, prev_lost)
    plt.xlabel('Размер шага')
    plt.ylabel('Погрешность, %')
    ax.set_xlim(first_step_size, curr_step_size)
    plt.show()
    print(f"Итоговый размер шага: {round(prev_step[-1], 6)}")


if __name__ == '__main__':
    dynamic_step()
    fixed_step()

