using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using Cysharp.Threading.Tasks;
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
        [SerializeField] private string endpoint;

        // 膜電位の色を計算するための定数
        private const float MinMembranePotential = -70.0f;
        private const float MaxMembranePotential = -30.0f;
        private const float HsvColorMapMin = 0.5f;
        private const float HsvColorMapMax = 1.0f;

        private NeuronRepository _neuronRepository;
        private List<GameObject> _generatedNeuronGameObjects;
        private Dictionary<string, Coroutine> _neuronFiringCoroutines; // string はニューロンのゲームオブジェクト名

        private void Awake()
        {
            _neuronRepository = new NeuronRepository(endpoint);
            _generatedNeuronGameObjects = new List<GameObject>();
            _neuronFiringCoroutines = new Dictionary<string, Coroutine>();
        }

        /// <summary>
        /// 生成できる単一ニューロン名一覧を返す
        /// </summary>
        /// <returns>ニューロン名リスト</returns>
        public List<string> GetAvailableNeuronNames()
        {
            return _neuronRepository.GetNeuronNames();
        }

        /// <summary>
        /// 生成されているニューロンの中に指定した名前のニューロンがあるかどうかを返す
        /// </summary>
        /// <param name="neuronName">ニューロン名</param>
        /// <returns>ニューロンのゲームオブジェクト or null</returns>
        public GameObject FindGeneratedNeuron(string neuronName)
        {
            return _generatedNeuronGameObjects.FirstOrDefault(neuron => neuron.name == neuronName);
        }

        /// <summary>
        /// 指定した名前のニューロンのゲームオブジェクトを生成する
        /// </summary>
        /// <param name="neuronName">ニューロン名</param>
        /// <param name="position">ニューロンの位置</param>
        /// <param name="generateNeuronNameSuffix">生成するニューロンのゲームオブジェクト名につけるサフィックス</param>
        public async UniTask<GameObject> GenerateSingleNeuron(string neuronName, Vector3 position,
            string generateNeuronNameSuffix = "")
        {
            var neuronObj = Instantiate(neuronPrefab, position, Quaternion.identity);
            neuronObj.name = neuronName + generateNeuronNameSuffix;
            var neuron = await _neuronRepository.GetNeuron(neuronName);
            foreach (var nc in neuron.Compartments.Values)
            {
                // position を基準に相対的な位置を計算する
                var compartmentPosition = new Vector3(nc.PositionX + position.x, nc.PositionY + position.y,
                    nc.PositionZ + position.z);

                GameObject compartmentObj;
                if (nc.Type == CompartmentType.Soma)
                {
                    compartmentObj = Instantiate(somaPrefab, compartmentPosition, Quaternion.identity);
                    compartmentObj.transform.localScale = new Vector3(nc.Radius, nc.Radius, nc.Radius);
                }
                else
                {
                    // 親コンパートメントの位置を取得
                    if (!neuron.Compartments.TryGetValue(nc.ParentId, out NeuronCompartment parent)) continue;
                    var parentPosition = new Vector3(parent.PositionX + position.x, parent.PositionY + position.y,
                        parent.PositionZ + position.z);

                    // 親コンパートメントとの中点を計算し、その位置にコンパートメントを生成する
                    var offset = parentPosition - compartmentPosition;
                    var scale = new Vector3(nc.Radius, offset.magnitude / 2.0f, nc.Radius);
                    var midpoint = compartmentPosition + (offset / 2.0f);
                    compartmentObj = Instantiate(compartmentPrefab, midpoint, Quaternion.identity);
                    compartmentObj.transform.localScale = scale;
                    compartmentObj.transform.up = offset;
                }

                compartmentObj.name = $"{nc.ID}";
                compartmentObj.transform.parent = neuronObj.transform;
            }

            _generatedNeuronGameObjects.Add(neuronObj);
            return neuronObj;
        }

        /// <summary>
        /// 指定した名前のニューロンのゲームオブジェクトを削除する
        /// </summary>
        /// <param name="neuronObj">削除するニューロンのゲームオブジェクト</param>
        public void DestroySingleNeuron(GameObject neuronObj)
        {
            Destroy(neuronObj);
            _generatedNeuronGameObjects.Remove(neuronObj);
        }

        /// <summary>
        /// 膜電位の値に応じた色を返す
        /// </summary>
        /// <param name="membranePotential">膜電位の値</param>
        /// <returns>色</returns>
        private static Color GetCompartmentColorFromMembranePotential(float membranePotential)
        {
            var colorBase =
                ((membranePotential - MinMembranePotential) / (MaxMembranePotential - MinMembranePotential)) *
                (HsvColorMapMax - HsvColorMapMin) + HsvColorMapMin; // colorBase hsvColorMapMin ~ hsvColorMapMax
            if (membranePotential > MaxMembranePotential) colorBase = HsvColorMapMax;
            if (membranePotential < MinMembranePotential) colorBase = HsvColorMapMin;
            return Color.HSVToRGB(colorBase, 1.0f, 1.0f);
        }

        /// <summary>
        /// 指定したニューロンのゲームオブジェクトに、膜電位のシミュレーション結果を色として可視化する
        /// </summary>
        /// <param name="neuronObj">ニューロンのゲームオブジェクト</param>
        /// <returns>IEnumerator</returns>
        private IEnumerator VisualizeSingleNeuronFiringSimulation(GameObject neuronObj)
        {
            var availableNeuronNames = GetAvailableNeuronNames();
            if (!availableNeuronNames.Contains(neuronObj.name)) yield break;
            var membranePotentialsIterator = _neuronRepository.GetMembranePotentials(neuronObj.name);
            foreach (var membranePotentials in membranePotentialsIterator)
            {
                for (var i = 0; i < neuronObj.transform.childCount; i++)
                {
                    var compartment = neuronObj.transform.GetChild(i).gameObject;
                    if (!Int32.TryParse(compartment.name, out Int32 compartmentId)) continue;
                    var membranePotential = membranePotentials.membranePotentials[compartmentId];
                    // 膜電位の値に応じた色を設定する
                    compartment.GetComponent<Renderer>().material.color =
                        GetCompartmentColorFromMembranePotential(membranePotential);
                }

                yield return new WaitForSeconds(Time.deltaTime * 5);
            }
        }

        /// <summary>
        /// ニューロン発火のシミュレーションを可視化する (ニューロンのシミュレーション結果の膜電位を取得し、それに応じた色を設定する)
        /// </summary>
        /// <param name="neuronObj">ニューロンのゲームオブジェクト</param>
        public void StartSingleNeuronFiring(GameObject neuronObj)
        {
            if (_neuronFiringCoroutines.TryGetValue(neuronObj.name, out Coroutine coroutine)) return;
            var neuronFiringCoroutine = StartCoroutine(VisualizeSingleNeuronFiringSimulation(neuronObj));
            _neuronFiringCoroutines.Add(neuronObj.name, neuronFiringCoroutine);
        }

        /// <summary>
        /// ニューロン発火のシミュレーションを停止する
        /// </summary>
        /// <param name="neuronObj">ニューロンのゲームオブジェクト</param>
        public void StopSingleNeuronFiring(GameObject neuronObj)
        {
            if (!_neuronFiringCoroutines.TryGetValue(neuronObj.name, out Coroutine coroutine)) return;
            _neuronRepository.CancelGetMembranePotentials();
            StopCoroutine(coroutine);
            _neuronFiringCoroutines.Remove(neuronObj.name);
        }

        public async void GenerateMultiNeuron()
        {
            const float r = 60.0f;
            const float h = 200.0f;
            for (var i = 0; i < 8; i++)
            {
                var x = Random.Range(-r, r);
                var y = Random.Range(-h, h);
                var z = Random.Range(-r, r);
                await GenerateSingleNeuron("cerebral_cortex_pyramidal_cell", new Vector3(x, y, z), i.ToString());
            }

            for (var i = 0; i < 2; i++)
            {
                var x = Random.Range(-r, r);
                var y = Random.Range(-h, h);
                var z = Random.Range(-r, r);
                await GenerateSingleNeuron("Pvalb_470522102_m_c", new Vector3(x, y, z), i.ToString());
            }
        }

        /// <summary>
        /// 生成されているニューロンのゲームオブジェクトを全て削除する
        /// </summary>
        public void DestroyAllNeurons()
        {
            foreach (var generatedNeuronGameObject in _generatedNeuronGameObjects) Destroy(generatedNeuronGameObject);
            _generatedNeuronGameObjects.Clear();
        }
    }
}