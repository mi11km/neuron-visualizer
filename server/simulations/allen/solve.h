#pragma once

#include "csparse.h"

typedef enum {
    BE, CN, N_SOLVER
} solve_solver_t;
typedef struct {
    solve_solver_t type;
    double *val, *b;
    int *dig;
    cs *A;
    double *x;
    css *S;
    csn *N;
    int *xi;
    cs *L, *U;
} solve_t;

solve_t *solve_initialize(cell_t *, solve_solver_t);

void solve_finalize(solve_t *);

void solve_update_v(cell_t *, solve_t *);
