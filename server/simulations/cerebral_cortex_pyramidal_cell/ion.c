#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "param.h"
#include "cell.h"
#include "ion.h"

static double vtrap(const double x, const double y) {
    if (fabs(x / y) < 1e-6) {
        return y * (1.0 - (x / y) / 2.0);
    } else {
        return x / (exp(x / y) - 1.0);
    }
}

static double alpha_m_NaTs(const double v) {
    const double malphaF = 0.182;
    const double mvhalf = -40.0; // (mV)
    const double mk = 6.0; // (mV)
    return malphaF * vtrap(-(v - mvhalf), mk);
}

static double beta_m_NaTs(const double v) {
    const double mbetaF = 0.124;
    const double mvhalf = -40.0; // (mV)
    const double mk = 6.0; // (mV)
    return mbetaF * vtrap((v - mvhalf), mk);
}

static double alpha_h_NaTs(const double v) {
    const double halphaF = 0.015;
    const double hvhalf = -66.0; // (mV)
    const double hk = 6.0; // (mV)
    return halphaF * vtrap((v - hvhalf), hk);
}

static double beta_h_NaTs(const double v) {
    const double hbetaF = 0.015;
    const double hvhalf = -66.0; // (mV)
    const double hk = 6.0; // (mV)
    return hbetaF * vtrap(-(v - hvhalf), hk);
}

static double inf_m_NaTs(const double v) { return alpha_m_NaTs(v) / (alpha_m_NaTs(v) + beta_m_NaTs(v)); }

static double inf_h_NaTs(const double v) { return alpha_h_NaTs(v) / (alpha_h_NaTs(v) + beta_h_NaTs(v)); }

static double tau_m_NaTs(const double v) {
    const double celsius = 34.0;
    const double qt = pow(2.3, (celsius - 23.0) / 10.0);
    return (1.0 / (alpha_m_NaTs(v) + beta_m_NaTs(v))) / qt;
}

static double tau_h_NaTs(const double v) {
    const double celsius = 34.0;
    const double qt = pow(2.3, (celsius - 23.0) / 10.0);
    return (1.0 / (alpha_h_NaTs(v) + beta_h_NaTs(v))) / qt;
}


static double alpha_m_NaTa(const double v) {
    const double malphaF = 0.182;
    const double mvhalf = -48.0;
    const double mk = 6.0;
    return malphaF * vtrap(-(v - mvhalf), mk);
}

static double beta_m_NaTa(const double v) {
    const double mbetaF = 0.124;
    const double mvhalf = -48.0;
    const double mk = 6.0;
    return mbetaF * vtrap((v - mvhalf), mk);
}

static double alpha_h_NaTa(const double v) {
    const double halphaF = 0.015;
    const double hvhalf = -69.0;
    const double hk = 6.0;
    return halphaF * vtrap(v - hvhalf, hk);
}

static double beta_h_NaTa(const double v) {
    const double hbetaF = 0.015;
    const double hvhalf = -69.0;
    const double hk = 6.0;
    return hbetaF * vtrap(-(v - hvhalf), hk);
}

static double inf_m_NaTa(const double v) {
    return alpha_m_NaTa(v) / (alpha_m_NaTa(v) + beta_m_NaTa(v));
}

static double inf_h_NaTa(const double v) {
    return alpha_h_NaTa(v) / (alpha_h_NaTa(v) + beta_h_NaTa(v));
}

static double tau_m_NaTa(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 23.0) / 10.0);
    return (1.0 / (alpha_m_NaTa(v) + beta_m_NaTa(v))) / qt;
}

static double tau_h_NaTa(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 23.0) / 10.0);
    return (1.0 / (alpha_h_NaTa(v) + beta_h_NaTa(v))) / qt;
}


static double alpha_h_Nap(const double v) {
    return 2.88e-6 * vtrap(v + 17.0, 4.63);
    //return 2.88e-3 * vtrap(v + 17.0, 4.63);
}

static double beta_h_Nap(const double v) {
    return 6.94e-6 * vtrap(-(v + 64.4), 2.63);
    //return 6.94e-3 * vtrap(-(v + 64.4), 2.63);
}

extern double inf_m_Nap(const double v) {
    return 1.0 / (1.0 + exp((v - (-52.6)) / -4.6));
}

static double inf_h_Nap(const double v) {
    return 1.0 / (1.0 + exp((v - (-48.8)) / 10.0));
}

static double tau_h_Nap(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (1.0 / (alpha_h_Nap(v) + beta_h_Nap(v))) / qt;
}

static double alpha_m_Kv2(const double v) {
    return 0.12 * vtrap(-(v - 43.0), 11.0);
}

static double beta_m_Kv2(const double v) {
    return 0.02 * exp(-(v + 1.27) / 120.0);
}

static double inf_m_Kv2(const double v) {
    return alpha_m_Kv2(v) / (alpha_m_Kv2(v) + beta_m_Kv2(v));
}

static double inf_h_Kv2(const double v) {
    return 1.0 / (1.0 + exp((v + 58.0) / 11.0));
}

static double tau_m_Kv2(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return 2.5 * (1.0 / (qt * (alpha_m_Kv2(v) + beta_m_Kv2(v))));
}

static double tau_h1_Kv2(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (360 + (1010 + 23.7 * (v + 54)) * exp(-((v + 75) / 48) * ((v + 75) / 48))) / qt;
}

static double tau_h2_Kv2(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (2350 + 1380 * exp(-0.011 * v) - 210 * exp(-0.03 * v)) / qt;
}

static double inf_m_Kv3(const double v) {
    const double vshift = 4.0;
    return 1.0 / (1.0 + exp(((v - (18.700 + vshift)) / (-9.700))));
}

static double tau_m_Kv3(const double v) {
    const double vshift = 4.0;
    return 0.2 * 20.000 / (1 + exp(((v - (-46.560 + vshift)) / (-44.140))));
}


static double inf_m_KP(const double v) {
    const double vshift = -1.1;
    return 1.0 / (1.0 + exp(-(v - (-14.3 + vshift)) / 14.6));
}

static double inf_h_KP(const double v) {
    const double vshift = -1.1;
    return 1.0 / (1.0 + exp(-(v - (-54.0 + vshift)) / -11.0));
}

static double tau_m_KP(const double v) {

    const double vshift = -1.1;
    const double tauF = 1.0;
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    if (v < -50.0 + vshift) {
        return tauF * (1.25 + 175.03 * exp(-(v - vshift) * -0.026)) / qt;
    } else {
        return tauF * (1.25 + 13 * exp(-(v - vshift) * 0.026)) / qt;
    }
}

static double tau_h_KP(const double v) {
    const double vshift = -1.1;
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (360.0 + (1010.0 + 24.0 * (v - (-55.0 + vshift))) *
                    exp(-((v - (-75.0 + vshift)) / 48) * ((v - (-75.0 + vshift)) / 48.0))) / qt;
}

static double inf_m_KT(const double v) {
    const double vshift = 0.0;
    return 1.0 / (1.0 + exp(-(v - (-47 + vshift)) / 29.0));
}

static double inf_h_KT(const double v) {
    const double vshift = 0.0;
    return 1.0 / (1.0 + exp(-(v + 66.0 - vshift) / -10.0));
}

static double tau_m_KT(const double v) {
    double mTauF = 1.0;
    double vshift = 0.0;
    double celcius = 34.0;
    double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (0.34 + mTauF * 0.92 * exp(-((v + 71.0 - vshift) / 59.0) * ((v + 71.0 - vshift) / 59.0))) / qt;
}

static double tau_h_KT(const double v) {
    double hTauF = 1.0;
    double vshift = 0.0;
    double celcius = 34.0;
    double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (8.0 + hTauF * 49.0 * exp(-((v + 73.0 - vshift) / 23.0) * ((v + 73.0 - vshift) / 23.0))) / qt;
}

static double inf_m_Kd(const double v) {
    return 1.0 - 1.0 / (1.0 + exp((v - (-43.0)) / 8.0));
}

static double inf_h_Kd(const double v) {
    return 1.0 / (1.0 + exp((v - (-67.0)) / 7.3));
}

static double tau_m_Kd(const double v) {
    return 1.0;
}

static double tau_h_Kd(const double v) {
    return 1500.0;
}


static double alpha_m_Im(const double v) {
    return 3.3e-3 * exp(2.5 * 0.04 * (v - (-35.0)));
}

static double beta_m_Im(const double v) {
    return 3.3e-3 * exp(-2.5 * 0.04 * (v - (-35.0)));
}

static double inf_m_Im(const double v) {
    return alpha_m_Im(v) / (alpha_m_Im(v) + beta_m_Im(v));
}

static double tau_m_Im(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (1.0 / (alpha_m_Im(v) + beta_m_Im(v))) / qt;
}


static double alpha_m_Imv2(const double v) {
    return 0.007 * exp((6.0 * 0.4 * (v - (-48.0))) / 26.12);
}

static double beta_m_Imv2(const double v) {
    return 0.007 * exp((-6.0 * (1.0 - 0.4) * (v - (-48.0))) / 26.12);
}

static double inf_m_Imv2(const double v) {
    return alpha_m_Imv2(v) / (alpha_m_Imv2(v) + beta_m_Imv2(v));
}

static double tau_m_Imv2(const double v) {
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 30.0) / 10.0);
    return (15.0 + 1.0 / (alpha_m_Imv2(v) + beta_m_Imv2(v))) / qt;
}

static double alpha_m_Ih(const double v) {
    return 0.001 * 6.43 * vtrap(v + 154.9, 11.9);
    //return 6.43 * vtrap(v + 154.9, 11.9);
}

static double beta_m_Ih(const double v) {
    return 0.001 * 193.0 * exp(v / 33.1);
    //return 193.0 * exp( v /33.1);
}

static double inf_m_Ih(const double v) {
    return alpha_m_Ih(v) / (alpha_m_Ih(v) + beta_m_Ih(v));
}

static double tau_m_Ih(const double v) {
    return 1.0 / (alpha_m_Ih(v) + beta_m_Ih(v));
}


static double inf_z_SK(const double v, double ca) {
    if (ca < 1e-7) {
        ca = ca + 1e-07;
    }
    return 1.0 / (1.0 + pow((0.00043 / ca), 4.8));
}

static double tau_z_SK(const double v) {
    return 1.0;
}


static double alpha_m_CaHVA(const double v) {
    return 0.055 * vtrap(-27.0 - v, 3.8);
}

static double beta_m_CaHVA(const double v) {
    return (0.94 * exp((-75.0 - v) / 17.0));
}

static double inf_m_CaHVA(const double v) {
    return alpha_m_CaHVA(v) / (alpha_m_CaHVA(v) + beta_m_CaHVA(v));
}

static double tau_m_CaHVA(const double v) {
    return 1.0 / (alpha_m_CaHVA(v) + beta_m_CaHVA(v));
}

static double alpha_h_CaHVA(const double v) {
    return (0.000457 * exp((-13.0 - v) / 50.0));
}

static double beta_h_CaHVA(const double v) {
    return (0.0065 / (exp((-v - 15.0) / 28.0) + 1.0));
}

static double inf_h_CaHVA(const double v) {
    return alpha_h_CaHVA(v) / (alpha_h_CaHVA(v) + beta_h_CaHVA(v));
}

static double tau_h_CaHVA(const double v) {
    return 1.0 / (alpha_h_CaHVA(v) + beta_h_CaHVA(v));
}


static double inf_m_CaLVA(const double v) {
    const double v_new = v + 10.0;
    return 1.0 / (1.0 + exp((v_new - (-30.0)) / -6.0));
}

static double inf_h_CaLVA(const double v) {
    const double v_new = v + 10.0;
    return 1.0 / (1.0 + exp((v_new - (-80.0)) / 6.4));
}

static double tau_m_CaLVA(const double v) {
    const double v_new = v + 10.0;
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (5.0 + 20.0 / (1.0 + exp((v_new - (-25.0)) / 5.0))) / qt;
}

static double tau_h_CaLVA(const double v) {
    const double v_new = v + 10.0;
    const double celcius = 34.0;
    const double qt = pow(2.3, (celcius - 21.0) / 10.0);
    return (20.0 + 50.0 / (1.0 + exp((v_new - (-40.0)) / 7.0))) / qt;
}

void ion_update(cell_t *cell) {
    double **elem = cell->elem;
    double **ion = cell->ion;
    double **cond = cell->cond;

    //for ( int id = 0; id < N_COMP; id++ )
    const int id = 0;
    {
        int type = elem[TYPE][id];
        //double I_Ca1 = ( 1e-3 * cell -> cond [ g_Ca ] [ id ] / elem [ AREA ] [ id ] * ion [ ch_Ca ] [ id ] * ion [ ch_Ca ] [ id ] * ion [ ci_Ca ] [ id ] *
        //		     ( elem [ V ] [ id ]  - V_Ca ) ); //I_Ca [mA/cm^2]
        const double celcius = 34.0;
        double v_val = elem[V][id];
        double Ca_val = elem[CA][id];
        double rev_ca = 1000 * ((RVAL * (273.15 + celcius)) / (2 * F)) * log(Ca1OUT / Ca_val);
        double gamma = (type == SOMA) ? 0.0012510775510599999 : 0.05;
        double decay = (type == SOMA) ? 717.91660042899991 : 80.0; // ms

        double i_ca =
                0.0265 * (v_val - rev_ca) * (cond[G_CAHVA][id] * ion[M_CAHVA][id] * ion[M_CAHVA][id] * ion[H_CAHVA][id]
                                             + cond[G_CALVA][id] * ion[M_CALVA][id] * ion[M_CALVA][id] *
                                               ion[H_CALVA][id]);
        double dCa = (-(10000) * (i_ca * gamma / (2 * F * DEPTH)) - (Ca_val - MIN_CA) / decay);

        ion[M_NATS][id] = inf_m_NaTs(v_val) + (ion[M_NATS][id] - inf_m_NaTs(v_val)) * exp(-DT / tau_m_NaTs(v_val));
        ion[H_NATS][id] = inf_h_NaTs(v_val) + (ion[H_NATS][id] - inf_h_NaTs(v_val)) * exp(-DT / tau_h_NaTs(v_val));
        ion[M_NATA][id] = inf_m_NaTa(v_val) + (ion[M_NATA][id] - inf_m_NaTa(v_val)) * exp(-DT / tau_m_NaTa(v_val));
        ion[H_NATA][id] = inf_h_NaTa(v_val) + (ion[H_NATA][id] - inf_h_NaTa(v_val)) * exp(-DT / tau_h_NaTa(v_val));
        ion[H_NAP][id] = inf_h_Nap(v_val) + (ion[H_NAP][id] - inf_h_Nap(v_val)) * exp(-DT / tau_h_Nap(v_val));
        ion[M_KV2][id] = inf_m_Kv2(v_val) + (ion[M_KV2][id] - inf_m_Kv2(v_val)) * exp(-DT / tau_m_Kv2(v_val));
        ion[H1_KV2][id] = inf_h_Kv2(v_val) + (ion[H1_KV2][id] - inf_h_Kv2(v_val)) * exp(-DT / tau_h1_Kv2(v_val));
        ion[H2_KV2][id] = inf_h_Kv2(v_val) + (ion[H2_KV2][id] - inf_h_Kv2(v_val)) * exp(-DT / tau_h2_Kv2(v_val));

        //m = m + (1. - exp(dt*(( ( ( - 1.0 ) ) ) / mTau)))*(- ( ( ( mInf ) ) / mTau ) / ( ( ( ( - 1.0 ) ) ) / mTau ) - m)
        //ion [ M_KV3    ] [ id ] += (1. - exp(DT*(( ( ( - 1.0 ) ) ) / tau_m_Kv3(v_val) )))*(- ( ( ( inf_m_Kv3(v_val) ) ) / tau_m_Kv3(v_val) ) / ( ( ( ( - 1.0 ) ) ) / tau_m_Kv3(v_val) ) - ion [ M_KV3    ] [ id ] );
        //						inf_m_Kv3   ( v_val ) + ( ion [ M_KV3    ] [ id ] - inf_m_Kv3   ( v_val ) ) * exp ( - DT / tau_m_Kv3    ( v_val ) );
        ion[M_KV3][id] = inf_m_Kv3(v_val) + (ion[M_KV3][id] - inf_m_Kv3(v_val)) * exp(-DT / tau_m_Kv3(v_val));
        ion[M_KP][id] = inf_m_KP(v_val) + (ion[M_KP][id] - inf_m_KP(v_val)) * exp(-DT / tau_m_KP(v_val));
        ion[H_KP][id] = inf_h_KP(v_val) + (ion[H_KP][id] - inf_h_KP(v_val)) * exp(-DT / tau_h_KP(v_val));
        ion[M_KT][id] = inf_m_KT(v_val) + (ion[M_KT][id] - inf_m_KT(v_val)) * exp(-DT / tau_m_KT(v_val));
        ion[H_KT][id] = inf_h_KT(v_val) + (ion[H_KT][id] - inf_h_KT(v_val)) * exp(-DT / tau_h_KT(v_val));
        ion[M_KD][id] = inf_m_Kd(v_val) + (ion[M_KD][id] - inf_m_Kd(v_val)) * exp(-DT / tau_m_Kd(v_val));
        ion[H_KD][id] = inf_h_Kd(v_val) + (ion[H_KD][id] - inf_h_Kd(v_val)) * exp(-DT / tau_h_Kd(v_val));
        ion[M_IM][id] = inf_m_Im(v_val) + (ion[M_IM][id] - inf_m_Im(v_val)) * exp(-DT / tau_m_Im(v_val));
        ion[M_IMV2][id] = inf_m_Imv2(v_val) + (ion[M_IMV2][id] - inf_m_Imv2(v_val)) * exp(-DT / tau_m_Imv2(v_val));
        ion[M_IH][id] = inf_m_Ih(v_val) + (ion[M_IH][id] - inf_m_Ih(v_val)) * exp(-DT / tau_m_Ih(v_val));
        ion[Z_SK][id] =
                inf_z_SK(v_val, Ca_val) + (ion[Z_SK][id] - inf_z_SK(v_val, Ca_val)) * exp(-DT / tau_z_SK(v_val));
        ion[M_CAHVA][id] = inf_m_CaHVA(v_val) + (ion[M_CAHVA][id] - inf_m_CaHVA(v_val)) * exp(-DT / tau_m_CaHVA(v_val));
        ion[H_CAHVA][id] = inf_h_CaHVA(v_val) + (ion[H_CAHVA][id] - inf_h_CaHVA(v_val)) * exp(-DT / tau_h_CaHVA(v_val));
        ion[M_CALVA][id] = inf_m_CaLVA(v_val) + (ion[M_CALVA][id] - inf_m_CaLVA(v_val)) * exp(-DT / tau_m_CaLVA(v_val));
        ion[H_CALVA][id] = inf_h_CaLVA(v_val) + (ion[H_CALVA][id] - inf_h_CaLVA(v_val)) * exp(-DT / tau_h_CaLVA(v_val));

        elem[CA][id] += dCa * DT;
    }
}

void ion_initialize(cell_t *cell) {
    double **elem = cell->elem;
    //double **ion  = cell -> ion;
    //double init_v_rand = 0.0;
    for (int i = 0; i < N_COMP; i++) {

        //if ( i % N_COMP == 0 ) {
        //  init_v_rand = ( ( double ) rand ( ) / RAND_MAX ) - 0.5;
        //}

        //elem [ V     ] [ i ] = V_INIT + 5.0 * init_v_rand;
        elem[V][i] = V_INIT;
        elem[CA][i] = Ca1_0;
        elem[I_EXT][i] = 0.0;
        double v_val = elem[V][i];
        double ca_val = elem[CA][i];
        cell->ion[M_NATS][i] = inf_m_NaTs(v_val);
        cell->ion[H_NATS][i] = inf_h_NaTs(v_val);
        cell->ion[M_NATA][i] = inf_m_NaTa(v_val);
        cell->ion[H_NATA][i] = inf_h_NaTa(v_val);
        cell->ion[H_NAP][i] = inf_h_Nap(v_val);
        cell->ion[M_KV2][i] = inf_m_Kv2(v_val);
        cell->ion[H1_KV2][i] = inf_h_Kv2(v_val);
        cell->ion[H2_KV2][i] = inf_h_Kv2(v_val);
        cell->ion[M_KV3][i] = inf_m_Kv3(v_val);
        cell->ion[M_KP][i] = inf_m_KP(v_val);
        cell->ion[H_KP][i] = inf_h_KP(v_val);
        cell->ion[M_KT][i] = inf_m_KT(v_val);
        cell->ion[H_KT][i] = inf_h_KT(v_val);
        cell->ion[M_KD][i] = inf_m_Kd(v_val);
        cell->ion[H_KD][i] = inf_h_Kd(v_val);
        cell->ion[M_IM][i] = inf_m_Im(v_val);
        cell->ion[M_IMV2][i] = inf_m_Imv2(v_val);
        cell->ion[M_IH][i] = inf_m_Ih(v_val);
        cell->ion[Z_SK][i] = inf_z_SK(v_val, ca_val);
        cell->ion[M_CAHVA][i] = inf_m_CaHVA(v_val);
        cell->ion[H_CAHVA][i] = inf_h_CaHVA(v_val);
        cell->ion[M_CALVA][i] = inf_m_CaLVA(v_val);
        cell->ion[H_CALVA][i] = inf_h_CaLVA(v_val);

    }
}
