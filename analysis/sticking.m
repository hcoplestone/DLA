clear all;
close all;
clc;

stickingProbabilities = [1 1.5 2 3 4 5 6 7 8 9 10];
stickingProbabilityFileIndices = [1 15 2 3 4 5 6 7 8 9 10]

dfs = [];
dfErrors = [];

ensembles = 1000;

for probability = stickingProbabilityFileIndices
   
    i = 0;

    % Columns 1 is number of particles data
    % Column 2 is cluster radius data
    % Each row is new system in ensemble
    % datas = cell(ensembles, 2)

    Rs = [];
    Ns = [];

    while( i < ensembles )
        fname = ['../results/stick4/ensemble-p', num2str(probability) ,'-#', num2str(i) ,'.csv'];
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
    
    [P, gof] = fit(meanLogRs, logNs(:,1), 'poly1')

    % Calculate df
    df = P.p1
    
    dfs = [dfs df]
    dfErrors = [dfErrors gof.rmse]
end

figure;
hold on;
plot(stickingProbabilities/10, dfs, 'x');
xlabel('$p_{stick}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$d_f$', 'Interpreter', 'latex', 'FontSize', 16);

errorbar(stickingProbabilities/10, dfs, dfErrors, 'LineStyle', 'none')

legend_handle = legend('Ensemble average for $d_f$','Standard error in $d_f$');
set(legend_handle,'Interpreter','latex');

f=fit(stickingProbabilities'/10,dfs','poly2')
plot(f)

hold off;
