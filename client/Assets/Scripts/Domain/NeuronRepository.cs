using System;
using System.Collections.Generic;
using System.IO;
using System.Net.Http;
using System.Text;
using System.Threading;
using Cysharp.Threading.Tasks;
using UnityEngine;
using UnityEngine.Networking;
using Openapi;

namespace Domain
{
    public class NeuronRepository
    {
        private readonly string _endpoint;
        private readonly HttpClient _httpClient;
        private CancellationTokenSource _cancellationTokenSource;

        private readonly Dictionary<string, Neuron> _neuronCache;

        public NeuronRepository(string endpoint)
        {
            _endpoint = endpoint;
            _httpClient = new HttpClient();
            _neuronCache = new Dictionary<string, Neuron>();
        }

        public List<string> GetNeuronNames()
        {
            using var request = UnityWebRequest.Get(_endpoint + "/api/v1/neurons");
            request.downloadHandler = new DownloadHandlerBuffer();
            request.SetRequestHeader("Accept", "application/json");
            request.SendWebRequest();

            while (request.result == UnityWebRequest.Result.InProgress)
            {
            }

            if (request.result != UnityWebRequest.Result.Success)
            {
                throw new Exception(request.error);
            }

            var names = new List<string>();
            var response = JsonUtility.FromJson<GetNeuronsResponse>(request.downloadHandler.text);
            foreach (var neuron in response.neurons) names.Add(neuron.name);

            return names;
        }

        public async UniTask<Neuron> GetNeuron(string name)
        {
            Neuron neuron;
            if (_neuronCache.TryGetValue(name, out neuron)) return neuron;

            using var request = UnityWebRequest.Get(_endpoint + "/api/v1/neurons/" + name + "/compartments");
            request.downloadHandler = new DownloadHandlerBuffer();
            request.SetRequestHeader("Accept", "application/json");
            await request.SendWebRequest();

            while (request.result == UnityWebRequest.Result.InProgress)
            {
            }

            if (request.result != UnityWebRequest.Result.Success)
            {
                throw new Exception(request.error);
            }

            neuron = new Neuron(name);
            var response = JsonUtility.FromJson<GetNeuronCompartmentsResponse>(request.downloadHandler.text);
            foreach (var compartment in response.compartments)
            {
                neuron.Compartments.Add(compartment.id,
                    new NeuronCompartment(compartment.id, (CompartmentType) compartment.type.id, compartment.positionX,
                        compartment.positionY, compartment.positionZ, compartment.radius, compartment.parentId));
            }

            _neuronCache.Add(name, neuron);
            return neuron;
        }

        public IEnumerable<GetNeuronCompartmentsMembranePotentialResponse> GetMembranePotentials(string name)
        {
            // 前回の処理があればキャンセルする
            _cancellationTokenSource?.Cancel();
            _cancellationTokenSource = new CancellationTokenSource();

            using var request = new HttpRequestMessage(HttpMethod.Get,
                _endpoint + "/api/v1/neurons/" + name + "/compartments/membranePotentials");
            request.Headers.Add("Accept", "text/event-stream");
            using var response = _httpClient.SendAsync(request, HttpCompletionOption.ResponseHeadersRead,
                _cancellationTokenSource.Token);

            if (!response.Result.IsSuccessStatusCode)
            {
                Console.WriteLine(response.Result.StatusCode);
                yield break;
            }

            using var stream = response.Result.Content.ReadAsStreamAsync();
            _cancellationTokenSource.Token.ThrowIfCancellationRequested();
            using var reader = new StreamReader(stream.Result, Encoding.UTF8);

            string line;
            while ((line = reader.ReadLine()) != null)
            {
                _cancellationTokenSource.Token.ThrowIfCancellationRequested();
                if (string.IsNullOrEmpty(line)) continue;
                if (!line.StartsWith("{\"membranePotentials\"")) continue;
                if (line == "0") break;
                yield return JsonUtility.FromJson<GetNeuronCompartmentsMembranePotentialResponse>(line);
            }

            _cancellationTokenSource.Dispose();
            _cancellationTokenSource = null;
        }

        public void CancelGetMembranePotentials()
        {
            _cancellationTokenSource?.Cancel();
            _cancellationTokenSource = null;
        }
    }
}