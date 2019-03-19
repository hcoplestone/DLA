clear all;
close all;
clc;

i = 0;
ensembles = 1000;

% Columns 1 is number of particles data
% Column 2 is cluster radius data
% Each row is new system in ensemble
% datas = cell(ensembles, 2)

Rs = [];
Ns = [];

while( i < ensembles )
    fname = ['../results/ensemble', num2str(i) ,'.csv'];
    data = load(fname);
    
    % Number of particles
    N = data(:,1);
    % Cluster radius
    R = data(:,2);
    
    Rs = [Rs R];
    Ns = [Ns N];
    
    i = i+1;
end

logRs = log(Rs)
logNs = log(Ns)

meanLogRs = mean(logRs, 2)

standDevInLogR = std(logRs, 0, 2);

% plot(log(datas{i+1,2}), log(datas{i+1,1}), 'x');
% plot(log(Rs(1,:), Ns(1,:)))

figure;
hold on;

plot(logNs(:,1), meanLogRs, 'x')

errorbar(logNs(:,1), meanLogRs, standDevInLogR, 'LineStyle', 'none')

% Fit a straight line to data
% [p1, errorInP1] = polyfit(logNs(:,1), meanLogRs, 1);
[p1, errorInP1] = polyfit(meanLogRs,logNs(:,1), 1);
[pevaled, delta] = polyval(p1, meanLogRs, errorInP1);
plot(pevaled, meanLogRs, 'r'); % plot x(y) and y
% plot(logNs(:,1),pevaled+5*delta,'m--',logNs(:,1),pevaled-5*delta,'m--');
legend_handle = legend('Ensemble average','Standard deviation in $\ln r_{max}$', 'Linear fit');
set(legend_handle,'Interpreter','latex');

% Calculate df
df = p1(1)
errorInP1

hold off;
xlabel('$\ln N_c$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$\ln r_{max}$', 'Interpreter', 'latex', 'FontSize', 16);
% title('Ensemble average of n=1000 independent clusters')