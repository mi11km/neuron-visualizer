using System;
using System.Collections;
using UnityEngine;
using Domain;


namespace Interfaces
{
    public class Neuron : MonoBehaviour
    {
        [SerializeField] private GameObject compartmentPrefab; // 細胞体以外のニューロンのコンパートメントプレハブ
        [SerializeField] private GameObject somaPrefab; // 細胞体プレハブ
        [SerializeField] private Player player;
        [SerializeField] private string endpoint;

        private readonly string _currentNeuronName = "cerebral_cortex_pyramidal_cell";

        private NeuronRepository _neuronRepository;
        private Domain.Neuron _currentNeuron;
        private Coroutine _neuronFiringCoroutine;

        private void Start()
        {
            _neuronRepository = new NeuronRepository(endpoint);
            Generate();

            // プレイヤーを細胞体の前に配置する
            var soma = _currentNeuron.GetSoma();
            player.RepositionInFrontOf(soma.PositionX, soma.PositionY, soma.PositionZ);
        }

        private void Update()
        {
            if (Input.GetKeyDown(KeyCode.F)) StartFiring();
            if (Input.GetKeyUp(KeyCode.S)) StopFiring();
        }

        // ニューロンをゲームオブジェクトとして生成する
        private void Generate()
        {
            _currentNeuron = _neuronRepository.GetNeuron(_currentNeuronName);
            foreach (var nc in _currentNeuron.Compartments.Values)
            {
                var position = new Vector3(nc.PositionX, nc.PositionY, nc.PositionZ);

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
                    var parentPosition = new Vector3(parent.PositionX, parent.PositionY, parent.PositionZ);

                    // 親コンパートメントとの中点を計算し、その位置にコンパートメントを生成する
                    var offset = parentPosition - position;
                    var scale = new Vector3(nc.Radius, offset.magnitude / 2.0f, nc.Radius);
                    var p = position + (offset / 2.0f);
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
        private void DestroyIfExist()
        {
            for (int i = 0; i < transform.childCount; i++)
            {
                Destroy(transform.GetChild(i).gameObject);
            }
        }

        private void StartFiring()
        {
            _neuronFiringCoroutine = StartCoroutine(VisualizeFiringSimulation());
        }

        private void StopFiring()
        {
            _neuronRepository.CancelGetMembranePotentials();
            StopCoroutine(_neuronFiringCoroutine);

            // ニューロンを再生成する
            DestroyIfExist();
            Generate();
        }

        // ニューロン発火のシミュレーションを可視化する (ニューロンのシミュレーション結果の膜電位を取得し、それに応じた色を設定する)
        private IEnumerator VisualizeFiringSimulation()
        {
            const float minMembranePotential = -70.0f;
            const float maxMembranePotential = -30.0f - minMembranePotential;
            const float hsvColorMapMin = 0.5f;
            const float hsvColorMapMax = 1.0f;
            var membranePotentialsIterator = _neuronRepository.GetMembranePotentials(_currentNeuronName);
            foreach (var membranePotentials in membranePotentialsIterator)
            {
                for (int i = 0; i < transform.childCount; i++)
                {
                    var compartment = transform.GetChild(i).gameObject;
                    var membranePotential = membranePotentials.membranePotentials[Int32.Parse(compartment.name)];

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
}