#pragma once

typedef enum {
    M_NATS,
    H_NATS,
    M_NATA,
    H_NATA,
    H_NAP,
    M_KV2,
    H1_KV2,
    H2_KV2,
    M_KV3,
    M_KP,
    H_KP,
    M_KT,
    H_KT,
    M_KD,
    H_KD,
    M_IM,
    M_IMV2,
    M_IH,
    Z_SK,
    M_CAHVA,
    H_CAHVA,
    M_CALVA,
    H_CALVA,
    N_ION
} ion_t;

extern double inf_m_Nap(const double v);

void ion_update(cell_t *);

void ion_initialize(cell_t *);
