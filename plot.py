import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation
from scipy.optimize import curve_fit

fig = plt.figure()
ax1 = fig.add_subplot(1,2,1)

def fractal_dimension_curve(R, alpha, df, beta):
    return np.power(alpha*R, df) + beta

def determine_fractal_dimension(R, Nc):
    popt, pcov = curve_fit(fractal_dimension_curve, R, Nc)
    return popt, pcov

def animate(i):
    data = np.loadtxt('data.csv', delimiter=',')

    Nc = data[:, 0]
    R = data[:, 1]

    popt, pcov = determine_fractal_dimension(R, Nc)
    NcFit = fractal_dimension_curve(R, *popt)

    lnNc = np.log(Nc)
    lnR = np.log(R)

    ax1.clear()
    ax1.plot(Nc, R, 'o', markerSize=2)
    ax1.plot(NcFit, R, 'r-')
    ax1.set_xlabel('Nc')
    ax1.set_ylabel('R')
    ax1.title.set_text(r'$\alpha = {0:.3f}, d_f = {1:.3f}, \beta = {2:.3f}$'.format(*popt))

    ax2.clear()
    ax2.plot(lnNc, lnR, 'o', markerSize=2)
    ax2.set_xlabel('ln Nc')
    ax2.set_ylabel('ln R')

ani = animation.FuncAnimation(fig, animate, interval=100)
plt.show()