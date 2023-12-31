using System;
using System.Collections;
using UnityEngine;
using Domain;


public class NeuronGenerator : MonoBehaviour
{
    [SerializeField] private GameObject compartmentPrefab; // 細胞体以外のニューロンのコンパートメントプレハブ
    [SerializeField] private GameObject somaPrefab; // 細胞体プレハブ

    private readonly NeuronRepository _repository = new();

    private readonly string _currentNeuronName = "Scnn1a_473845048_m_c";
    private Domain.Neuron _currentNeuron;
    private Coroutine _neuronFiringCoroutine;

    void Start()
    {
        GenerateNeuron();
    }

    void Update()
    {
        if (Input.GetKeyDown(KeyCode.Space)) DestroyNeuronIfExist();
        if (Input.GetKeyDown(KeyCode.F)) StartFiringNeuron();
        if (Input.GetKeyDown(KeyCode.R)) StopFiringNeuron();
    }

    // ニューロンをゲームオブジェクトとして生成する
    private void GenerateNeuron()
    {
        _currentNeuron = _repository.GetNeuron(_currentNeuronName);
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

    private void StartFiringNeuron()
    {
        _neuronFiringCoroutine = StartCoroutine(VisualizeFiringSimulation());
    }

    private void StopFiringNeuron()
    {
        StopCoroutine(_neuronFiringCoroutine);
    }

    // ニューロン発火のシミュレーションを可視化する (ニューロンのシミュレーション結果の膜電位を取得し、それに応じた色を設定する)
    private IEnumerator VisualizeFiringSimulation()
    {
        var minMembranePotential = -70.0f;
        var maxMembranePotential = -30.0f - minMembranePotential;
        const float hsvColorMapMin = 0.5f;
        const float hsvColorMapMax = 1.0f;
        foreach (var membranePotentials in _repository.GetMembranePotentials(_currentNeuronName))
        {
            for (int i = 0; i < transform.childCount; i++)
            {
                var compartment = transform.GetChild(i).gameObject;
                var membranePotential = membranePotentials.MembranePotentials[Int32.Parse(compartment.name)];

                // 膜電位の値に応じた色を設定する
                var colorBase =
                    ((membranePotential - minMembranePotential) / maxMembranePotential) *
                    (hsvColorMapMax - hsvColorMapMin) +
                    hsvColorMapMin; // colorBase hsvColorMapMin ~ hsvColorMapMax
                if (membranePotential - minMembranePotential > maxMembranePotential) colorBase = hsvColorMapMax;
                if (membranePotential - minMembranePotential < 0) colorBase = hsvColorMapMin;
                compartment.GetComponent<Renderer>().material.color = Color.HSVToRGB(colorBase, 1.0f, 1.0f);
            }

            yield return new WaitForSeconds(Time.deltaTime * 10);
        }
    }
}