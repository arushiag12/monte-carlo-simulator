import matplotlib
matplotlib.use('Agg')

import subprocess
import matplotlib.pyplot as plt

def run_command(command):
    result = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, shell=True)
    return result.stdout

def run_analysis():
    versions = ["small", "medium", "large"]
    num_threads = [2, 4, 6, 8, 12]
    iters = 5
    serial_times = {}
    parallel_times = {}
    workstealing_times = {}
    speedups_p = {}
    speedups_w = {}

    for version in versions:
        times = []
        for _ in range(iters):
            command = f"go run ./driver/driver.go {version} s"
            time = float(run_command(command))
            print(f"version: {version}, serial, time: {time}")
            print("-----------------------------------------------")
            times.append(time)
        serial_times[version] = sum(times) / len(times)

        parallel_times[version] = {}
        for num in num_threads:
            times = []
            for _ in range(iters):
                command = f"go run ./driver/driver.go {version} p {num}"
                time = float(run_command(command))
                print(f"version: {version}, parallel num threads: {num}, time: {time}")
                print("-----------------------------------------------")
                times.append(time)
            parallel_times[version][num] = sum(times) / len(times)
    
        speedups_p[version] = {}
        for num in num_threads:
            speedups_p[version][num] = serial_times[version] / parallel_times[version][num]

        workstealing_times[version] = {}
        for num in num_threads:
            times = []
            for _ in range(iters):
                command = f"go run ./driver/driver.go {version} w {num}"
                time = float(run_command(command))
                print(f"version: {version}, work stealing num threads: {num}, time: {time}")
                print("-----------------------------------------------")
                times.append(time)
            workstealing_times[version][num] = sum(times) / len(times)
    
        speedups_w[version] = {}
        for num in num_threads:
            speedups_w[version][num] = serial_times[version] / workstealing_times[version][num]

    plt.figure()
    plt.plot(list(speedups_p["small"].keys()), list(speedups_p["small"].values()), marker='o', linestyle='-', color='g', label='small')
    plt.plot(list(speedups_p["medium"].keys()), list(speedups_p["medium"].values()), marker='o', linestyle='-', color='r', label='medium')
    plt.plot(list(speedups_p["large"].keys()), list(speedups_p["large"].values()), marker='o', linestyle='-', color='c', label='large')
    plt.xlabel("Number of threads")
    plt.ylabel("Speedup")
    plt.legend()
    plt.title(f"Speedup graph of parallel version")
    plt.savefig(f"./benchmark/parallel_speedup.png")

    plt.figure()
    plt.plot(list(speedups_w["small"].keys()), list(speedups_w["small"].values()), marker='o', linestyle='-', color='g', label='small')
    plt.plot(list(speedups_w["medium"].keys()), list(speedups_w["medium"].values()), marker='o', linestyle='-', color='r', label='medium')
    plt.plot(list(speedups_w["large"].keys()), list(speedups_w["large"].values()), marker='o', linestyle='-', color='c', label='large')
    plt.xlabel("Number of threads")
    plt.ylabel("Speedup")
    plt.legend()
    plt.title(f"Speedup graph of work stealing version")
    plt.savefig(f"./benchmark/work_stealing_speedup.png")

    plt.figure()
    plt.plot(list(speedups_p["small"].keys()), list(speedups_p["small"].values()), marker='o', linestyle='-', color='b', label='parallel')
    plt.plot(list(speedups_w["small"].keys()), list(speedups_w["small"].values()), marker='o', linestyle='-', color='g', label='work-stealing')
    plt.xlabel("Number of threads")
    plt.ylabel("Speedup")
    plt.legend()
    plt.title(f"Speedup comparison graph for small")
    plt.savefig(f"./benchmark/speedup_comparison_small.png")

    plt.figure()
    plt.plot(list(speedups_p["medium"].keys()), list(speedups_p["medium"].values()), marker='o', linestyle='-', color='b', label='parallel')
    plt.plot(list(speedups_w["medium"].keys()), list(speedups_w["medium"].values()), marker='o', linestyle='-', color='g', label='work-stealing')
    plt.xlabel("Number of threads")
    plt.ylabel("Speedup")
    plt.legend()
    plt.title(f"Speedup comparison graph for medium")
    plt.savefig(f"./benchmark/speedup_comparison_medium.png")

    plt.figure()
    plt.plot(list(speedups_p["large"].keys()), list(speedups_p["large"].values()), marker='o', linestyle='-', color='b', label='parallel')
    plt.plot(list(speedups_w["large"].keys()), list(speedups_w["large"].values()), marker='o', linestyle='-', color='g', label='work-stealing')
    plt.xlabel("Number of threads")
    plt.ylabel("Speedup")
    plt.legend()
    plt.title(f"Speedup comparison graph for large")
    plt.savefig(f"./benchmark/speedup_comparison_large.png")

if __name__ == "__main__":
    run_analysis()