clear all;
close all;
clc;

stickingProbabilities = [1:10];
dfs = [];

ensembles = 1000;

for probability = stickingProbabilities
   
    i = 0;

    % Columns 1 is number of particles data
    % Column 2 is cluster radius data
    % Each row is new system in ensemble
    % datas = cell(ensembles, 2)

    Rs = [];
    Ns = [];

    while( i < ensembles )
        fname = ['../results/stick2/ensemble-p', num2str(probability) ,'-#', num2str(i) ,'.csv'];
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

    % Fit a straight line to data
    [p1, errorInP1] = polyfit(logNs(:,1), meanLogRs, 1);
    [pevaled, delta] = polyval(p1, logNs(:, 1), errorInP1);

    % Calculate df
    df = 1/p1(1)
    
    dfs = [dfs df]   
end

figure;
hold on;
plot(stickingProbabilities/10, dfs, 'x');
xlabel('$p_{stick}$', 'Interpreter', 'latex', 'FontSize', 16);
ylabel('$d_f$', 'Interpreter', 'latex', 'FontSize', 16);

hold off;
