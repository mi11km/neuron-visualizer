using UnityEngine;
using Domain;


public class NeuronGenerator : MonoBehaviour
{
    [SerializeField] private GameObject compartmentPrefab; // 細胞体以外のニューロンのコンパートメントプレハブ
    [SerializeField] private GameObject somaPrefab; // 細胞体プレハブ

    private readonly NeuronRepository _repository = new();
    private Domain.Neuron _currentNeuron;

    void Start()
    {
        GenerateNeuron();
    }

    void Update()
    {
        if (Input.GetKeyDown(KeyCode.Space)) DestroyNeuronIfExist();
    }

    // ニューロンをゲームオブジェクトとして生成する
    private void GenerateNeuron()
    {
        _currentNeuron = _repository.GetNeuron("Scnn1a_473845048_m_c");
        foreach (var nc in _currentNeuron.Compartments.Values)
        {
            Vector3 position = new Vector3(nc.PositionX, nc.PositionY, nc.PositionZ);

            GameObject compartmentObj;
            if (nc.Type == CompartmentType.Soma)
            {
                compartmentObj = Instantiate(somaPrefab, position, Quaternion.identity);
                compartmentObj.transform.localScale = new Vector3(nc.Radius, nc.Radius, nc.Radius);
            }
            else
            {
                // 親コンパートメントの位置を取得
                NeuronCompartment parent;
                if (!_currentNeuron.Compartments.TryGetValue(nc.ParentId, out parent)) continue;
                Vector3 parentPosition = new Vector3(parent.PositionX, parent.PositionY, parent.PositionZ);

                Vector3 offset = parentPosition - position;
                Vector3 scale = new Vector3(nc.Radius, offset.magnitude / 2.0f, nc.Radius);
                Vector3 p = position + (offset / 2.0f);
                compartmentObj = Instantiate(compartmentPrefab, p, Quaternion.identity);
                compartmentObj.transform.localScale = scale;
                compartmentObj.transform.up = offset;
            }

            compartmentObj.name = $"{nc.ID}";
            compartmentObj.transform.parent = transform; // このスクリプトがアタッチされているゲームオブジェクトの子にする
            compartmentObj.GetComponent<Renderer>().material.color = new Color(1.0f, 0.3f, 0.25f);
        }
    }

    // ニューロンが表示されていたら破棄する
    private void DestroyNeuronIfExist()
    {
        for (int i = 0; i < transform.childCount; i++)
        {
            Destroy(transform.GetChild(i).gameObject);
        }
    }
}