import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.axes_grid1.inset_locator import zoomed_inset_axes
from mpl_toolkits.axes_grid1.inset_locator import mark_inset
import matplotlib.patches as patches


fig = plt.figure()
i = 0

control = 'results/thermal/ensemble#' + str(i) + '.dat'
steadyTempIncrease = 'results/thermal/ensemble#' + str(i) + '-steady-temp-increase.dat'
sinTemp = 'results/thermal/ensemble#' + str(i) + '-sin.dat'

# Control
grid = np.loadtxt(control, int)
ax = fig.add_subplot(1, 3, 1)
ax.imshow(grid, 'binary')
ax.set_aspect('equal')
ax.get_yaxis().set_visible(False)
ax.get_xaxis().set_visible(False)

# Steady temp increase
grid = np.loadtxt(steadyTempIncrease, int)
ax = fig.add_subplot(1, 3, 2)
ax.imshow(grid, 'binary')
ax.set_aspect('equal')
ax.get_yaxis().set_visible(False)
ax.get_xaxis().set_visible(False)

# Sin temp 
grid = np.loadtxt(sinTemp, int)
ax = fig.add_subplot(1, 3, 3)
ax.imshow(grid, 'binary')
ax.set_aspect('equal')
ax.get_yaxis().set_visible(False)
ax.get_xaxis().set_visible(False)

plt.show()