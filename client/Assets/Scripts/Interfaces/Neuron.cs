using System;
using System.Collections;
using UnityEngine;
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

        private void Start()
        {
            _neuronRepository = new NeuronRepository(endpoint);

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

        // ニューロンをゲームオブジェクトとして生成する
        private void Generate(string neuronName, Vector3 position)
        {
            var neuron = Instantiate(neuronPrefab, position, Quaternion.identity);
            neuron.name = neuronName;
            _currentNeuron = _neuronRepository.GetNeuron(neuronName);
            foreach (var nc in _currentNeuron.Compartments.Values)
            {
                var compartmentPosition = new Vector3(nc.PositionX, nc.PositionY, nc.PositionZ);

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
                    var parentPosition = new Vector3(parent.PositionX, parent.PositionY, parent.PositionZ);

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
        }

        // ニューロンが表示されていたら破棄する
        private void DestroyIfExist(string neuronName)
        {
            var neuron = GameObject.Find(neuronName); // NOTE: ニューロンの名前が一意である必要がある
            Destroy(neuron);
        }

        // ニューロン発火のシミュレーションを可視化する (ニューロンのシミュレーション結果の膜電位を取得し、それに応じた色を設定する)
        private void StartFiring(string neuronName)
        {
            var neuron = GameObject.Find(neuronName); // NOTE: ニューロンの名前が一意である必要がある
            _neuronFiringCoroutine = StartCoroutine(VisualizeFiringSimulation(neuron));
        }

        // ニューロン発火のシミュレーションを停止する
        private void StopFiring(string neuronName)
        {
            _neuronRepository.CancelGetMembranePotentials();
            StopCoroutine(_neuronFiringCoroutine);

            // ニューロンを再生成する
            DestroyIfExist(neuronName);
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