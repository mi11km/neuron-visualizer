#pragma once

#define N_COMP ( 3682 )
#define SWC_FILE "./Scnn1a_473845048_m_c.swc"
#define ION_FILE "./Scnn1a_473845048_m_c_ion.txt"

typedef enum {
    _DUMMY_, SOMA, AXON, BASAL, APICAL, N_COMPTYPE
} cell_type_t;

#define RA_SOMA ( 138.27999999999997 )
#define RA_AXON ( 138.27999999999997 )
#define RA_DEND ( 138.27999999999997 )
#define CM_SOMA ( 1.0 )
#define CM_AXON ( 1.0 )
#define CM_DEND ( 2.12 )

#define V_LEAK  ( -92.4991 )
#define V_INIT  ( V_LEAK )
#define V_NA    (  53.0  )
#define V_K     ( -107.0 )
#define V_HCN   ( -45.0  )
#define V_Ca    ( 129.33 )

#define SHELL1_D    ( 0.2e-4 ) // [0.2mum -> 0.2e-4cm]
#define B_Ca1       ( 1.5 ) // [/ms]
#define Ca1_0       ( 1e-4 ) // [mM]
#define Ca1OUT      ( 2.0 )  // [mM]
#define F           ( 9.6485e4 ) // Faraday constant [s*A/mol]
#define RVAL        ( 8.31446261815324 ) // [J/K*mol]
#define MIN_CA      ( 1e-4 ) // [mM]
#define DEPTH       ( 0.1e-4 ) // [ 0.1mu m -> 0.1e-4cm]

typedef enum {
    V, CA, CM, AREA, PARENT, TYPE, I_EXT, N_ELEM
} cell_elem_t;
typedef enum {
    G_LEAK, G_NATS, G_NATA, G_NAP, G_KV2, G_KV3, G_KP, G_KT, G_KD, G_IM, G_IMV2, G_IH, G_SK, G_CAHVA, G_CALVA, N_COND
} cell_cond_t;

typedef struct {
    double *rad, *len, *Ra;
    double **elem, **ion, **cond;
    FILE *outfile;
} cell_t;

cell_t *cell_initialize(void);

void cell_finalize(cell_t *);

void cell_set_current(cell_t *, const double);

void cell_update_ion(cell_t *);

void cell_output_file(cell_t *, const double);
