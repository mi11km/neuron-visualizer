#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include "param.h"
#include "cell.h"
#include "solve.h"

extern void timer_start(void);

extern double timer_elapsed(void);

int main(int argc, char *argv[]) {
    cell_t *cell = cell_initialize();
    solve_t *solve = solve_initialize(cell, BE);

    timer_start();
    for (int t_count = 0; t_count < Tmax - 1; t_count++) {
        double t = t_count * DT;
        if (t_count % 500 == 0) { printf("time: %d \n", (int) t); }
        cell_set_current(cell, t);
        cell_update_ion(cell);
        solve_update_v(cell, solve);
        cell_output_file(cell, t);
    }
    double elapsedTime = timer_elapsed();
    printf("Elapsed time = %f sec.\n", elapsedTime);

    cell_finalize(cell);
    solve_finalize(solve);
}
