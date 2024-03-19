#!/bin/bash
#
#SBATCH --mail-user=arushiag@uchicago.edu
#SBATCH --mail-type=ALL
#SBATCH --job-name=proj3_analysis
#SBATCH --output=./benchmark/slurm/out/analysis_results.stdout
#SBATCH --error=./benchmark/slurm/out/analysis_results.stderr
#SBATCH --chdir=/home/arushiag/Documents/project-3-arushiag12/proj3
#SBATCH --partition=debug
#SBATCH --nodes=1
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=16
#SBATCH --mem-per-cpu=3200
#SBATCH --exclusive
#SBATCH --time=04:00:00

module load golang/1.19
python3 ./benchmark/analysis.py
