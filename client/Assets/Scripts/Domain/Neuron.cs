using System.Collections.Generic;


namespace Domain
{
    public class Neuron
    {
        // コンパートメントのディクショナリ（key: NeuronCompartment.ID, value: NeuronCompartment）
        public Dictionary<long, NeuronCompartment> Compartments { get; }

        public string Name { get; }

        public Neuron(string name)
        {
            Name = name;
            Compartments = new Dictionary<long, NeuronCompartment>();
        }

        // 細胞体を取得する
        public NeuronCompartment GetSoma()
        {
            foreach (var compartment in Compartments.Values)
            {
                if (compartment.Type == CompartmentType.Soma)
                {
                    return compartment;
                }
            }

            return null;
        }
    }

    public class NeuronCompartment
    {
        // （同一ニューロンで一意な）コンパートメントの id
        public long ID { get; }

        // 構成要素の種類
        public CompartmentType Type { get; }

        // x座標の位置
        public float PositionX { get; }

        // y座標の位置
        public float PositionY { get; }

        // z座標の位置
        public float PositionZ { get; }

        // コンパートメントの半径
        public float Radius { get; }

        // 親のコンパートメントの ID
        public long ParentId { get; }

        public NeuronCompartment(
            long id, CompartmentType type, float positionX, float positionY, float positionZ, float radius,
            long parentId)
        {
            ID = id;
            Type = type;
            PositionX = positionX;
            PositionY = positionY;
            PositionZ = positionZ;
            Radius = radius;
            ParentId = parentId;
        }
    }

    // ニューロンの構成要素の種類
    public enum CompartmentType
    {
        Soma = 1, // 細胞体
        Axon = 2, // 軸索
        BasalDendrite = 3, // 基底樹状突起
        ApicalDendrite = 4 // 尖端樹状突起
    }
}