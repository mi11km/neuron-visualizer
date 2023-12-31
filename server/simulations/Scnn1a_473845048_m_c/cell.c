#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <assert.h>
#include "param.h"
#include "cell.h"
#include "ion.h"

static void read_swc_file(cell_t *cell) {
    double x_cor[N_COMP], y_cor[N_COMP], z_cor[N_COMP];
    {
        FILE *file = fopen(SWC_FILE, "r");
        if (!file) {
            fprintf(stderr, "no such file %s\n", SWC_FILE);
            exit(1);
        }

        char buf[1024];
        int i = 0;
        while (fgets(buf, 1024, file)) {
            if (buf[0] == '#') { continue; }
            int id, type, parent;
            double rad, x, y, z;

            sscanf(buf, "%d %d %lf %lf %lf %lf %d", &id, &type, &x, &y, &z, &rad, &parent);

            cell->rad[id] = rad * 1e-4;
            x_cor[id] = x;
            y_cor[id] = y;
            z_cor[id] = z;

            cell->elem[PARENT][id] = (double) parent;
            cell->elem[TYPE][id] = (double) type;

            assert(id == i);
            i++;
        }
        fclose(file);
        assert(i == N_COMP);
    }

    const double RA_TABLE[] = {-1, RA_SOMA, RA_AXON, RA_DEND, RA_DEND};
    const double CM_TABLE[] = {-1, CM_SOMA, CM_AXON, CM_DEND, CM_DEND};

    for (int i = 0; i < N_COMP; i++) {

        double rad = cell->rad[i];
        double len = 0.0;

        int parent = (int) cell->elem[PARENT][i];
        int type = (int) cell->elem[TYPE][i];

        if (parent == -1) {
            len = 2.0 * rad;
        } else {
            double dx = x_cor[i] - x_cor[parent];
            double dy = y_cor[i] - y_cor[parent];
            double dz = z_cor[i] - z_cor[parent];
            len = sqrt(dx * dx + dy * dy + dz * dz) * 1.0e-4; // [mum -> cm]
        }

        cell->len[i] = len;
        cell->Ra[i] = RA_TABLE[type] * 1e-3;
        cell->elem[AREA][i] = 2.0 * M_PI * rad * len; // [cm^2]
        cell->elem[CM][i] = CM_TABLE[type] * cell->elem[AREA][i]; // [muF]
    }
}

static void read_ion_file(cell_t *cell) {
    FILE *file = fopen(ION_FILE, "r");
    if (!file) {
        fprintf(stderr, "no such file %s\n", ION_FILE);
        exit(1);
    }

    char buf[1024];
    int i = 0;
    while (fgets(buf, 1024, file)) {
        double i_l1, i_NATS, i_NATA, i_NAP, i_KV2, i_KV3, i_KP, i_KT, i_KD, i_IM, i_IMV2, i_IH, i_SK, i_CAHVA, i_CALVA;

        sscanf(buf, "%lf %lf %lf %lf %lf %lf %lf %lf %lf %lf %lf %lf %lf %lf %lf",
               &i_l1, &i_NATS, &i_NATA, &i_NAP, &i_KV2, &i_KV3, &i_KP, &i_KT, &i_KD, &i_IM, &i_IMV2, &i_IH, &i_SK,
               &i_CAHVA, &i_CALVA);

        double area = cell->elem[AREA][i];
        cell->cond[G_LEAK][i] = i_l1 * area;
        cell->cond[G_NATS][i] = i_NATS * area;
        cell->cond[G_NATA][i] = i_NATA * area;
        cell->cond[G_NAP][i] = i_NAP * area;
        cell->cond[G_KV2][i] = i_KV2 * area;
        cell->cond[G_KV3][i] = i_KV3 * area;
        cell->cond[G_KP][i] = i_KP * area;
        cell->cond[G_KT][i] = i_KT * area;
        cell->cond[G_KD][i] = i_KD * area;
        cell->cond[G_IM][i] = i_IM * area;
        cell->cond[G_IMV2][i] = i_IMV2 * area;
        cell->cond[G_IH][i] = i_IH * area;
        cell->cond[G_SK][i] = i_SK * area;
        cell->cond[G_CAHVA][i] = i_CAHVA * area;
        cell->cond[G_CALVA][i] = i_CALVA * area;


        i++;
    }
    fclose(file);
    assert(i == N_COMP);
}


cell_t *cell_initialize(void) {
    cell_t *cell = (cell_t *) malloc(sizeof(cell_t));

    // set structures and arrays
    cell->elem = (double **) malloc(N_ELEM * sizeof(double *));
    cell->cond = (double **) malloc(N_COND * sizeof(double *));
    cell->ion = (double **) malloc(N_ION * sizeof(double *));
    for (int i = 0; i < N_ELEM; i++) { cell->elem[i] = (double *) malloc(N_COMP * sizeof(double)); }
    for (int i = 0; i < N_COND; i++) { cell->cond[i] = (double *) malloc(N_COMP * sizeof(double)); }
    for (int i = 0; i < N_ION; i++) { cell->ion[i] = (double *) malloc(N_COMP * sizeof(double)); }

    cell->rad = (double *) malloc(N_COMP * sizeof(double));
    cell->len = (double *) malloc(N_COMP * sizeof(double));
    cell->Ra = (double *) malloc(N_COMP * sizeof(double));

    read_swc_file(cell);
    read_ion_file(cell);
    ion_initialize(cell);

    return cell;
}

void cell_finalize(cell_t *cell) {
    for (int i = 0; i < N_ELEM; i++) { free(cell->elem[i]); }
    free(cell->elem);
    for (int i = 0; i < N_COND; i++) { free(cell->cond[i]); }
    free(cell->cond);
    for (int i = 0; i < N_ION; i++) { free(cell->ion[i]); }
    free(cell->ion);

    free(cell->rad);
    free(cell->len);
    free(cell->Ra);
    free(cell);
}

void cell_set_current(cell_t *cell, const double t) {
    cell->elem[I_EXT][0] = (10 <= t && t < 410) ? 0.12 * 0.001 : 0.0;
}

void cell_output_file(cell_t *cell, const double t) {
    fprintf(cell->outfile, "%lf,%lf,%lf\n", t, cell->elem[V][0], cell->elem[CA][0]);
}

void cell_update_ion(cell_t *cell) { ion_update(cell); }

void cell_output_stdout(cell_t *cell, const double t) {
    printf("%lf", t);
    for (int i = 0; i < N_COMP; ++i) {
        printf(",%lf", cell->elem[V][i]);
    }
    printf("\n");
}
