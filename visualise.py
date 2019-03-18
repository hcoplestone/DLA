import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.axes_grid1.inset_locator import zoomed_inset_axes
from mpl_toolkits.axes_grid1.inset_locator import mark_inset
import matplotlib.patches as patches


fig = plt.figure()
i = 1

filename = 'results/second/ensemble' + str(i) + '.dat'

grid = np.loadtxt(filename, int)
ax = fig.add_subplot(1, 1, 1)

extent = [500, 600,500,600]

# ax.set_title('Ensemble ' + str(i))
# plt.imshow(grid)
plt.imshow(grid, 'binary')
ax.set_aspect('equal')

# axins = zoomed_inset_axes(ax, 3.5, loc=4) # zoom = 6
# axins.imshow(grid, interpolation="nearest",
            #  origin="lower")

# sub region of the original image
x1, x2, y1, y2 = 160, 220, 300, 240 # specify the limits
ax.set_xlim(x1, x2)
ax.set_ylim(y1, y2)

# mark_inset(ax, axins, loc1=4, loc2=1, fc="none", ec="0.5")

# rect = patches.Rectangle((160,240),60,60,linewidth=1,edgecolor='r',facecolor='none')
# ax.add_patch(rect)

# axins = zoomed_inset_axes(ax, 3, loc=4) # zoom-factor: 2.5, location: upper-left
# axins.imshow(grid, 'binary')
# x1, x2, y1, y2 = 500, 600, 500, 600 # specify the limits
# axins.set_xlim(x1, x2) # apply the x-limits
# axins.set_ylim(y1, y2) # apply the y-limits

# mark_inset(ax, axins, loc1=2, loc2=4, fc="none", ec="0.5")

# # plt.yticks(visible=False)
# # plt.xticks(visible=False)

ax.get_yaxis().set_visible(False)
ax.get_xaxis().set_visible(False)
# axins.get_yaxis().set_visible(False)
# axins.get_xaxis().set_visible(False)
# plt.colorbar(orientation='vertical')

plt.show()