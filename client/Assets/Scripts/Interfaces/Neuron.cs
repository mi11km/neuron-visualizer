using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Unity.VisualScripting;
using Random = UnityEngine.Random;
using Domain;


namespace Interfaces
{
    public class NeuronGenerator : MonoBehaviour
    {
        [SerializeField] private GameObject neuronPrefab; // ニューロンのプレハブ
        [SerializeField] private GameObject compartmentPrefab; // 細胞体以外のニューロンのコンパートメントプレハブ
        [SerializeField] private GameObject somaPrefab; // 細胞体プレハブ
        [SerializeField] private Player player;
        [SerializeField] private string endpoint;
        [SerializeField] private string currentNeuronName;

        private NeuronRepository _neuronRepository;
        private Coroutine _neuronFiringCoroutine;
        private Neuron _currentNeuron;
        private List<GameObject> _neurons;

        private void Start()
        {
            _neuronRepository = new NeuronRepository(endpoint);
            _neurons = new List<GameObject>();
            // GenerateMultiNeuron();

            // ニューロンを生成して、プレイヤーを細胞体の前に配置する
            Generate(currentNeuronName, new Vector3(0, 0, 0));
            var soma = _currentNeuron.GetSoma();
            player.RepositionInFrontOf(soma.PositionX, soma.PositionY, soma.PositionZ);
        }

        private void Update()
        {
            if (Input.GetKeyDown(KeyCode.F)) StartFiring(currentNeuronName);
            if (Input.GetKeyUp(KeyCode.R)) StopFiring(currentNeuronName);
        }

        // ニューロンのゲームオブジェクトを取得する
        private GameObject FindGeneratedNeuron(string neuronName)
        {
            foreach (var n in _neurons)
            {
                if (n.name == neuronName)
                {
                    return n;
                }
            }

            return null;
        }

        // ニューロンをゲームオブジェクトとして生成する
        private void Generate(string neuronName, Vector3 position, string neuronNameSuffix = "")
        {
            var neuron = Instantiate(neuronPrefab, position, Quaternion.identity);
            neuron.name = neuronName + neuronNameSuffix;
            _currentNeuron = _neuronRepository.GetNeuron(neuronName);
            foreach (var nc in _currentNeuron.Compartments.Values)
            {
                var compartmentPosition = new Vector3(
                    nc.PositionX + position.x,
                    nc.PositionY + position.y,
                    nc.PositionZ + position.z
                );

                GameObject compartmentObj;
                if (nc.Type == CompartmentType.Soma)
                {
                    compartmentObj = Instantiate(somaPrefab, compartmentPosition, Quaternion.identity);
                    compartmentObj.transform.localScale = new Vector3(nc.Radius, nc.Radius, nc.Radius);
                }
                else
                {
                    // 親コンパートメントの位置を取得
                    NeuronCompartment parent;
                    if (!_currentNeuron.Compartments.TryGetValue(nc.ParentId, out parent)) continue;
                    var parentPosition = new Vector3(
                        parent.PositionX + position.x,
                        parent.PositionY + position.y,
                        parent.PositionZ + position.z
                    );

                    // 親コンパートメントとの中点を計算し、その位置にコンパートメントを生成する
                    var offset = parentPosition - compartmentPosition;
                    var scale = new Vector3(nc.Radius, offset.magnitude / 2.0f, nc.Radius);
                    var p = compartmentPosition + (offset / 2.0f);
                    compartmentObj = Instantiate(compartmentPrefab, p, Quaternion.identity);
                    compartmentObj.transform.localScale = scale;
                    compartmentObj.transform.up = offset;
                }

                compartmentObj.name = $"{nc.ID}";
                compartmentObj.transform.parent = neuron.transform;
                compartmentObj.GetComponent<Renderer>().material.color = new Color(1.0f, 0.3f, 0.25f);
            }

            _neurons.Add(neuron);
        }

        // ニューロンのゲームオブジェクトを削除する
        private void Destroy(string neuronName)
        {
            var neuron = FindGeneratedNeuron(neuronName);
            _neurons.Remove(neuron);
            if (!neuron.IsUnityNull()) Destroy(neuron);
        }

        private void GenerateMultiNeuron()
        {
            const float r = 60.0f;
            const float h = 200.0f;
            for (var i = 0; i < 8; i++)
            {
                var x = Random.Range(-r, r);
                var y = Random.Range(-h, h);
                var z = Random.Range(-r, r);
                Generate("cerebral_cortex_pyramidal_cell", new Vector3(x, y, z), i.ToString());
            }

            for (var i = 0; i < 2; i++)
            {
                var x = Random.Range(-r, r);
                var y = Random.Range(-h, h);
                var z = Random.Range(-r, r);
                Generate("Pvalb_470522102_m_c", new Vector3(x, y, z), i.ToString());
            }
        }

        // 生成されたニューロンのゲームオブジェクトを全て削除する
        private void DestroyAll()
        {
            foreach (var n in _neurons)
            {
                Destroy(n);
            }

            _neurons.Clear();
        }

        // ニューロン発火のシミュレーションを可視化する (ニューロンのシミュレーション結果の膜電位を取得し、それに応じた色を設定する)
        private void StartFiring(string neuronName)
        {
            var neuron = FindGeneratedNeuron(neuronName);
            if (neuron.IsUnityNull())
            {
                Debug.Log("ニューロンが見つかりません");
                return;
            }

            _neuronFiringCoroutine = StartCoroutine(VisualizeFiringSimulation(neuron));
        }

        // ニューロン発火のシミュレーションを停止する
        private void StopFiring(string neuronName)
        {
            _neuronRepository.CancelGetMembranePotentials();
            StopCoroutine(_neuronFiringCoroutine);

            // ニューロンを再生成する
            Destroy(neuronName);
            Generate(neuronName, new Vector3(0, 0, 0));
        }

        private IEnumerator VisualizeFiringSimulation(GameObject neuron)
        {
            const float minMembranePotential = -70.0f;
            const float maxMembranePotential = -30.0f - minMembranePotential;
            const float hsvColorMapMin = 0.5f;
            const float hsvColorMapMax = 1.0f;
            var membranePotentialsIterator = _neuronRepository.GetMembranePotentials(neuron.name);
            foreach (var membranePotentials in membranePotentialsIterator)
            {
                for (int i = 0; i < neuron.transform.childCount; i++)
                {
                    var compartment = neuron.transform.GetChild(i).gameObject;
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

                yield return new WaitForSeconds(Time.deltaTime * 5);
            }
        }
    }
}