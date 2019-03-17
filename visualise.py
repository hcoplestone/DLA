import numpy as np
import matplotlib.pyplot as plt

fig = plt.figure()

for i in range(0, 3):
    filename = 'results/ensemble' + str(i) + '.dat'

    grid = np.loadtxt(filename, int)
    ax = fig.add_subplot(1, 3, i+1)

    ax.set_title('Ensemble ' + str(i))
    plt.imshow(grid)
    ax.set_aspect('equal')

    ax.get_yaxis().set_visible(False)
    ax.get_xaxis().set_visible(False)
    # plt.colorbar(orientation='vertical')

plt.show()