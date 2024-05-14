import numpy as np
import matplotlib.pyplot as plt
import scipy.stats as stats


# Функция для генерации выборки с помощью метода обратных функций
# f(z) = 1 / (e - z), 0 <= z <= e - 1
# F(z) = 1 - e^(-z - 1), 0 <= z <= e - 1
# F^(-1)(u) = -ln(1 - u) - 1, u ~ U(0, 1)
def inverse_transform_sampling(n):
    samples = -np.log(1 - np.random.rand(n)) - 1
    return samples


# Функция для вычисления среднего значения и дисперсии выборки
# Среднее значение: E[X] = ∫xf(x)dx
# Дисперсия: D[X] = E[X^2] - E[X]^2 = ∫x^2f(x)dx - E[X]^2
def calculate_mean_variance(samples):
    mean = np.mean(samples)  # Вычисление среднего значения
    variance = np.var(samples)  # Вычисление дисперсии
    return mean, variance


# Функция для вычисления статистики критерия Колмогорова
# D_n = max(|F_n(x) - F(x)|)
def calculate_ks_statistic(samples):
    _, ks_statistic = stats.kstest(samples, 'expon', args=(1,), alternative='two-sided')
    # Применение критерия Колмогорова из библиотеки scipy.stats.kstest
    # Вычисление статистики критерия Колмогорова
    return ks_statistic


# Функция для построения гистограммы выборки
def plot_histogram(samples, bins=50):
    plt.hist(samples, bins, density=True, alpha=0.5, label='Исследуемое распределение')
    x = np.linspace(0, np.e - 1, 100)
    y = 1 / (np.e - x)
    plt.plot(x, y, 'r', label='Теоретическое распределение')
    plt.xlabel('Значение')
    plt.ylabel('Плотность')
    plt.legend()
    plt.show()


# Размеры выборок
sample_sizes = [50, 100, 1000, 10000, 100000]

# Генерация выборок, вычисление среднего значения, дисперсии и статистики критерия Колмогорова
for sample_size in sample_sizes:
    smpls = inverse_transform_sampling(sample_size)
    mean, variance = calculate_mean_variance(smpls)
    ks_statistic = calculate_ks_statistic(smpls)

    print(f'Размер выборки: {sample_size}')
    print(f'Среднее значение: {mean} (ожидаемое: 1 - 1/e)')
    print(f'Дисперсия: {variance} (ожидаемое: 1 - 2/e)')
    print("Критерий Колмогорова: {:.100f}".format(ks_statistic))
    print()

    plot_histogram(smpls)
