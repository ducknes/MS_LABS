import numpy as np
import matplotlib.pyplot as plt
from scipy.stats import kstest


# Функция для генерации случайных чисел с заданным законом распределения
def generate_random_numbers(n):
    # Генерация равномерно распределенных случайных чисел
    u = np.random.uniform(0, 1, n)
    # Преобразование равномерных чисел в числа с заданным законом распределения
    z = np.exp(1) - np.exp(-u)
    return z


# Функция для оценки математического ожидания и дисперсии
def estimate_mean_and_variance(z):
    # Вычисление математического ожидания
    mean = np.mean(z)
    # Вычисление дисперсии
    variance = np.var(z)
    return mean, variance


# Функция для оценки соответствия закона распределения
def kolmogorov_test(z):
    # Применение критерия Колмогорова
    D, p = kstest(z, 'uniform')
    return D, p


# Функция для построения гистограммы
def plot_histogram(z, bins=30):
    # Построение гистограммы распределения сгенерированных чисел
    plt.hist(z, bins=bins, density=True, alpha=0.6, color='g')
    plt.title('Гистограмма')
    plt.show()


# Основная функция
def main():
    # Размеры выборок для анализа
    sample_sizes = [50, 100, 1000, 100000]
    for n in sample_sizes:
        # Генерация случайных чисел
        z = generate_random_numbers(n)
        # Оценка математического ожидания и дисперсии
        mean, variance = estimate_mean_and_variance(z)
        print(f"Размер выборки: {n}, Среднее: {mean}, Дисперсия: {variance}")
        # Оценка соответствия закона распределения
        D, p = kolmogorov_test(z)
        print(f"Тест Колмогорова: D = {D}, p-значение = {p}")
        # Построение гистограммы
        plot_histogram(z)


if __name__ == "__main__":
    main()
