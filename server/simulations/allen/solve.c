#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "param.h"
#include "cell.h"
#include "ion.h"
#include "solve.h"

solve_t *solve_initialize(cell_t *cell, solve_solver_t type) {
    solve_t *solver = (solve_t *) malloc(sizeof(solve_t));

    // Matrix
    double *mat = (double *) malloc(N_COMP * N_COMP * sizeof(double));
    {
        double rad[N_COMP], len[N_COMP], Ra[N_COMP];
        for (int i = 0; i < N_COMP; i++) {
            rad[i] = cell->rad[i];
            len[i] = cell->len[i];
            Ra[i] = cell->Ra[i];
        }
        for (int i = 0; i < N_COMP * N_COMP; i++) { mat[i] = 0.0; }
        for (int i = 0; i < N_COMP; i++) {
            int d = cell->elem[PARENT][i];
            if (d >= 0) {
                //double xr = ( d == 0 ) ? 0.0 : Ra [ i ];
                //double yr = ( d == 0 ) ? 0.0 : Ra [ d ];
                double r = (2.0 / ((Ra[i] * len[i]) / (rad[i] * rad[i] * M_PI) +
                                   (Ra[d] * len[d]) / (rad[d] * rad[d] * M_PI))); // -1 * [mS]
                mat[d + N_COMP * i] = r;
                mat[i + N_COMP * d] = mat[d + N_COMP * i]; // i*NGO -> set Rows, +d -> set Columns
            }
        }
        for (int i = 0; i < N_COMP; i++) {
            double r = 0;
            for (int j = 0; j < N_COMP; j++) {
                r += mat[j + N_COMP * i];
            }
            mat[i + N_COMP * i] = 1000 + r; // dirty hack to spacify diagonal elements
        }
    }
    { // generate sparse matrix A for csparse
        cs *T = cs_spalloc(0, 0, 1, 1, 1);
        for (int i = 0; i < N_COMP; i++) {
            for (int j = 0; j < N_COMP; j++) {
                if (mat[j + N_COMP * i] > 0) { cs_entry(T, i, j, ((i == j) ? 1 : -1) * mat[j + N_COMP * i]); }
            }
        }
        solver->A = cs_triplet(T);
    }
    free(mat);

    solver->dig = (int *) malloc(N_COMP * sizeof(int)); // dig index
    {
        int n = 0;
        for (int i = 0; i < solver->A->nzmax; i++) {
            if (solver->A->x[i] > 500) { // if i is a diagonal element
                solver->A->x[i] -= 1000;
                solver->dig[n] = i;
                n++;
            }
        }
    }

    solver->val = (double *) malloc(solver->A->nzmax * sizeof(double));
    for (int i = 0; i < solver->A->nzmax; i++) { solver->val[i] = solver->A->x[i]; }
    solver->b = (double *) malloc(N_COMP * sizeof(double)); // b value

    solver->type = type;

    solver->type = type;

    const int n = solver->A->n;
    solver->S = cs_calloc(1, sizeof(css));        /* allocate symbolic analysis */
    solver->x = cs_malloc(n, sizeof(double));
    solver->xi = cs_malloc(2 * n, sizeof(int));

    return solver;
}

void solve_cs_sqr(const cs *A, css *S) {
    const int n = A->n;
    S->Q = cs_amd(A, 0);            /* fill-reducing ordering */
    S->unz = 4 * (A->p[n]) + n;        /* for LU factorization only, */
    S->lnz = S->unz;            /* guess nnz(L) and nnz(U) */
}

void solve_finalize(solve_t *solver) {
    free(solver->val);
    free(solver->b);
    free(solver->dig);
    cs_spfree(solver->A);
    cs_free(solver->x);
    cs_sfree(solver->S);
    cs_free(solver->xi);
}

csn *solve_cs_lu(solve_t *solver, const cs *A, const css *S, double tol) {
    cs *L, *U;
    csn *N;
    double pivot, *Lx, *Ux, a, t;
    int *Lp, *Li, *Up, *Ui, *Pinv, *Q, n, ipiv, k, top, p, i, col, lnz, unz;
    if (!A || !S) return (NULL);            /* check inputs */
    n = A->n;
    Q = S->Q;
    lnz = S->lnz;
    unz = S->unz;

    //x = cs_malloc (n, sizeof (double)) ;
    //xi = cs_malloc (2*n, sizeof (int)) ;
    double *x = solver->x;
    int *xi = solver->xi;
    N = cs_calloc(1, sizeof(csn));
    //if (!x || !xi || !N) return (cs_ndone (N, NULL, xi, x, 0)) ;
    N->L = L = cs_spalloc(n, n, lnz, 1, 0);        /* initial L and U */
    N->U = U = cs_spalloc(n, n, unz, 1, 0);
    N->Pinv = Pinv = cs_malloc(n, sizeof(int));
    //if (!L || !U || !Pinv) return (cs_ndone (N, NULL, xi, x, 0)) ;
    Lp = L->p;
    Up = U->p;
    for (i = 0; i < n; i++) x[i] = 0;        /* clear workspace */
    for (i = 0; i < n; i++) Pinv[i] = -1;        /* no rows pivotal yet */
    for (k = 0; k <= n; k++) Lp[k] = 0;        /* no cols of L yet */
    lnz = unz = 0;
    for (k = 0; k < n; k++)        /* compute L(:,k) and U(:,k) */
    {
        /* --- Triangular solve --------------------------------------------- */
        Lp[k] = lnz;            /* L(:,k) starts here */
        Up[k] = unz;            /* U(:,k) starts here */
        if ((lnz + n > L->nzmax && !cs_sprealloc(L, 2 * L->nzmax + n)) ||
            (unz + n > U->nzmax && !cs_sprealloc(U, 2 * U->nzmax + n))) {
            //return (cs_ndone (N, NULL, xi, x, 0)) ;
        }
        Li = L->i;
        Lx = L->x;
        Ui = U->i;
        Ux = U->x;
        col = Q ? (Q[k]) : k;
        top = cs_splsolve(L, A, col, xi, x, Pinv); /* x = L\A(:,col) */
        /* --- Find pivot --------------------------------------------------- */
        ipiv = -1;
        a = -1;
        for (p = top; p < n; p++) {
            i = xi[p];        /* x(i) is nonzero */
            if (Pinv[i] < 0)        /* row i is not pivotal */
            {
                if ((t = fabs(x[i])) > a) {
                    a = t;        /* largest pivot candidate so far */
                    ipiv = i;
                }
            } else            /* x(i) is the entry U(Pinv[i],k) */
            {
                Ui[unz] = Pinv[i];
                Ux[unz++] = x[i];
            }
        }
        //if (ipiv == -1 || a <= 0) return (cs_ndone (N, NULL, xi, x, 0)) ;
        if (Pinv[col] < 0 && fabs(x[col]) >= a * tol) ipiv = col;
        /* --- Divide by pivot ---------------------------------------------- */
        pivot = x[ipiv];        /* the chosen pivot */
        Ui[unz] = k;            /* last entry in U(:,k) is U(k,k) */
        Ux[unz++] = pivot;
        Pinv[ipiv] = k;        /* ipiv is the kth pivot row */
        Li[lnz] = ipiv;        /* first entry in L(:,k) is L(k,k) = 1 */
        Lx[lnz++] = 1;
        for (p = top; p < n; p++) /* L(k+1:n,k) = x / pivot */
        {
            i = xi[p];
            if (Pinv[i] < 0)        /* x(i) is an entry in L(:,k) */
            {
                Li[lnz] = i;        /* save unpermuted row in L */
                Lx[lnz++] = x[i] / pivot;    /* scale pivot column */
            }
            x[i] = 0;            /* x [0..n-1] = 0 for next k */
        }
    }
    /* --- Finalize L and U ------------------------------------------------- */
    Lp[n] = lnz;
    Up[n] = unz;
    Li = L->i;                /* fix row indices of L for final Pinv */
    for (p = 0; p < lnz; p++) Li[p] = Pinv[Li[p]];
    cs_sprealloc(L, 0);        /* remove extra space from L and U */
    cs_sprealloc(U, 0);
    //return (cs_ndone (N, NULL, xi, x, 1)) ;	/* success */
    return N;
}

void solve_cs_lusol(solve_t *solver) {
    const int n = solver->A->n;
    const double tol = 0.001;

    solve_cs_sqr(solver->A, solver->S);        /* ordering and symbolic analysis */
    solver->N = solve_cs_lu(solver, solver->A, solver->S, tol);        /* numeric LU factorization */

    cs_ipvec(n, solver->N->Pinv, solver->b, solver->x);    /* x = P*b */
    cs_lsolve(solver->N->L, solver->x);        /* x = L\x */
    cs_usolve(solver->N->U, solver->x);        /* x = U\x */
    cs_ipvec(n, solver->S->Q, solver->x, solver->b);    /* b = Q*x */
    cs_nfree(solver->N);
}

static void update_matrix(cell_t *cell, solve_t *solver) {
    double **elem = cell->elem;
    double **cond = cell->cond;
    double **ion = cell->ion;
    const double Celcius = 34.0;

    double foo[N_COND][N_COMP];
    for (int i = 0; i < N_COMP; i++) {
        foo[G_NATS][i] = cond[G_NATS][i] * ion[M_NATS][i] * ion[M_NATS][i] * ion[M_NATS][i] * ion[H_NATS][i];
        foo[G_NATA][i] = cond[G_NATA][i] * ion[M_NATA][i] * ion[M_NATA][i] * ion[M_NATA][i] * ion[H_NATA][i];
        foo[G_NAP][i] = cond[G_NAP][i] * inf_m_Nap(elem[V][i]) * ion[H_NAP][i];
        foo[G_KV2][i] = cond[G_KV2][i] * ion[M_KV2][i] * ion[M_KV2][i] * (0.5 * ion[H1_KV2][i] + 0.5 * ion[H2_KV2][i]);
        foo[G_KV3][i] = cond[G_KV3][i] * ion[M_KV3][i];
        foo[G_KP][i] = cond[G_KP][i] * ion[M_KP][i] * ion[M_KP][i] * ion[H_KP][i];
        foo[G_KT][i] = cond[G_KT][i] * ion[M_KT][i] * ion[M_KT][i] * ion[M_KT][i] * ion[M_KT][i] * ion[H_KT][i];
        foo[G_KD][i] = cond[G_KD][i] * ion[M_KD][i] * ion[H_KD][i];
        foo[G_IM][i] = cond[G_IM][i] * ion[M_IM][i];
        foo[G_IMV2][i] = cond[G_IMV2][i] * ion[M_IMV2][i];
        foo[G_IH][i] = cond[G_IH][i] * ion[M_IH][i];
        foo[G_SK][i] = cond[G_SK][i] * ion[Z_SK][i];
        foo[G_CAHVA][i] = cond[G_CAHVA][i] * ion[M_CAHVA][i] * ion[M_CAHVA][i] * ion[H_CAHVA][i];
        foo[G_CALVA][i] = cond[G_CALVA][i] * ion[M_CALVA][i] * ion[M_CALVA][i] * ion[H_CALVA][i];
    }

    for (int i = 0; i < solver->A->nzmax; i++) { solver->A->x[i] = solver->val[i]; }

    for (int i = 0; i < N_COMP; i++) {
        const double ca = elem[CA][i];
        const double rev_ca = 1000 * ((RVAL * (273.0 + Celcius)) / (2 * F)) * log(Ca1OUT / ca);
        int d = solver->dig[i];
        solver->A->x[d] += ((elem[CM][i] / DT)
                            + cond[G_LEAK][i]
                            + foo[G_NATS][i]
                            + foo[G_NATA][i]
                            + foo[G_NAP][i]
                            + foo[G_KV2][i]
                            + foo[G_KV3][i]
                            + foo[G_KP][i]
                            + foo[G_KT][i]
                            + foo[G_KD][i]
                            + foo[G_IM][i]
                            + foo[G_IMV2][i]
                            + foo[G_IH][i]
                            + foo[G_SK][i]
                            + foo[G_CAHVA][i]
                            + foo[G_CALVA][i]
        );

        solver->b[i] = ((elem[CM][i] / DT) * elem[V][i]
                        + cond[G_LEAK][i] * V_LEAK + elem[I_EXT][i]
                        + foo[G_NATS][i] * V_NA
                        + foo[G_NATA][i] * V_NA
                        + foo[G_NAP][i] * V_NA
                        + foo[G_KV2][i] * V_K
                        + foo[G_KV3][i] * V_K
                        + foo[G_KP][i] * V_K
                        + foo[G_KT][i] * V_K
                        + foo[G_KD][i] * V_K
                        + foo[G_IM][i] * V_K
                        + foo[G_IMV2][i] * V_K
                        + foo[G_IH][i] * V_HCN
                        + foo[G_SK][i] * V_K
                        + foo[G_CAHVA][i] * rev_ca
                        + foo[G_CALVA][i] * rev_ca
        );
    }
}

void solve_by_backward_euler(cell_t *cell, solve_t *solver) {
    update_matrix(cell, solver);

    //const int order = -1;
    //const double tol = 0.001;
    //cs_lusol ( solver -> A, solver -> b, order, tol );
    solve_cs_lusol(solver);
    for (int i = 0; i < N_COMP; i++) { cell->elem[V][i] = solver->b[i]; }
}

void solve_update_v(cell_t *cell, solve_t *solver) {
    switch (solver->type) {
        default:
            solve_by_backward_euler(cell, solver);
            break;
    }
}
