clear all;

subplot(1,3,1);

x = linspace(0,2*pi);
y = sin(x)

plot(x,y)
xlabel('\tau', 'FontSize', 17)
ylabel('T(\tau)', 'FontSize', 13)

xticks([0 2*pi])
xticklabels({'0','\tau_{final}'})
xlim([0 2*pi])

yticks([-1 1])
yticklabels({'T_1','T_2'})

ax = gca;
ax.FontSize = 14; 

subplot(1,3,2);

x = 0:0.1:1
y = x

plot(x,y)
xlabel('\tau', 'FontSize', 17)
ylabel('T(\tau)', 'FontSize', 13)

xticks([0 1])
xticklabels({'0','\tau_{final}'})

yticks([0 1])
yticklabels({'T_1','T_2'})

ax = gca;
ax.FontSize = 14; 

subplot(1,3,3);

x = 0:0.1:1
y = 1-x

plot(x,y)
xlabel('\tau', 'FontSize', 17)
ylabel('T(\tau)', 'FontSize', 13)

xticks([0 1])
xticklabels({'0','\tau_{final}'})

ax = gca;
ax.FontSize = 14; 

yticks([0 1])
yticklabels({'T_1','T_2'})