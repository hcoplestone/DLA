import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation
from scipy.optimize import curve_fit

number_of_subplots = 32
subplot_dimension = np.ceil(np.sqrt(32))

fig = plt.figure()

def fractal_dimension_curve(R, alpha, df, beta):
    return np.power(alpha*R, df) + beta

def determine_fractal_dimension(R, Nc):
    popt, pcov = curve_fit(fractal_dimension_curve, R, Nc)
    return popt, pcov

def analyse_system(i):
    filename = 'results/ensemble{0}.csv'.format(i)
    print('Analysing ' + filename + '\n')

    ax = fig.add_subplot(subplot_dimension,subplot_dimension,1)
    data = np.loadtxt(filename, delimiter=',')

    Nc = data[:, 0]
    R = data[:, 1]

    popt, pcov = determine_fractal_dimension(R, Nc)
    NcFit = fractal_dimension_curve(R, *popt)

    lnNc = np.log(Nc)
    lnR = np.log(R)

    ax.plot(Nc, R, 'o', markerSize=2)
    ax.plot(NcFit, R, 'r-')
    ax.set_xlabel('Nc')
    ax.set_ylabel('R')
    # ax.title.set_text(r'$\alpha = {0:.3f}, d_f = {1:.3f}, \beta = {2:.3f}$'.format(*popt))
    ax.title.set_text(r'$d_f = {1:.3f}$'.format(popt[1]))

def main():
    for i in range(32, 33):
        analyse_system(i)

    plt.show()

if __name__ == '__main__':
    main()