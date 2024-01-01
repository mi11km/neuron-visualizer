using Grpc.Core;
using System.Collections.Generic;
using System.Threading;
using Neuron.V1;


namespace Domain
{
    public class NeuronRepository
    {
        private readonly string _endpoint;
        private CancellationTokenSource _cancellationTokenSource;

        public NeuronRepository(string endpoint)
        {
            _endpoint = endpoint;
        }

        public Neuron GetNeuron(string name)
        {
            var channel = new Channel(_endpoint, ChannelCredentials.Insecure);

            var neuronServiceClient = new NeuronService.NeuronServiceClient(channel);
            var response = neuronServiceClient.GetNeuronShape(new GetNeuronShapeRequest { NeuronName = name });

            var neuron = new Neuron();

            foreach (var compartment in response.NeuronCompartments)
            {
                neuron.Compartments.Add(compartment.Id, new NeuronCompartment(
                    compartment.Id,
                    (CompartmentType)compartment.Type,
                    compartment.PositionX,
                    compartment.PositionY,
                    compartment.PositionZ,
                    compartment.Radius,
                    compartment.ParentId
                ));
            }

            return neuron;
        }

        public IEnumerable<GetMembranePotentialsResponse> GetMembranePotentials(string name)
        {
            // 前回の処理があればキャンセルする
            _cancellationTokenSource?.Cancel();
            _cancellationTokenSource = new CancellationTokenSource();

            var channel = new Channel(_endpoint, ChannelCredentials.Insecure);

            var neuronServiceClient = new NeuronService.NeuronServiceClient(channel);
            var response = neuronServiceClient.GetMembranePotentials(
                new GetMembranePotentialsRequest { NeuronName = name });

            while (response.ResponseStream.MoveNext(_cancellationTokenSource.Token).Result)
            {
                yield return response.ResponseStream.Current;
            }
        }

        public void CancelGetMembranePotentials()
        {
            _cancellationTokenSource?.Cancel();
            _cancellationTokenSource = null;
        }
    }
}