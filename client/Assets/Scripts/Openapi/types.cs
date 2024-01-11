using System;

namespace Openapi
{
    [Serializable]
    public class GetNeuronsResponse
    {
        public Neuron neurons;
    }

    [Serializable]
    public class GetNeuronCompartmentsResponse
    {
        public NeuronCompartment[] compartments;
    }

    [Serializable]
    public class GetNeuronCompartmentsMembranePotentialResponse
    {
        public float timeStep;
        public float[] membranePotentials;
    }

    [Serializable]
    public class Neuron
    {
        public string name;
    }

    [Serializable]
    public class NeuronCompartmentType
    {
        public long id;
        public string name;
    }

    [Serializable]
    public class NeuronCompartment
    {
        public long id;
        public NeuronCompartmentType type;
        public float positionX;
        public float positionY;
        public float positionZ;
        public float radius;
        public long parentId;
    }
}