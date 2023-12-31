using Cysharp.Net.Http;
using Grpc.Net.Client;
using Neuron.V1;
using System.Collections.Generic;
using System.Threading;


namespace Domain
{
    public class NeuronRepository
    {
        const string Endpoint = "http://localhost:8080";

        public NeuronRepository()
        {
        }

        public Neuron GetNeuron(string name)
        {
            // Initialize gRPC client
            using var httpHandler = new YetAnotherHttpHandler()
                { SkipCertificateVerification = true, Http2Only = true };
            using var channel =
                GrpcChannel.ForAddress(Endpoint, new GrpcChannelOptions() { HttpHandler = httpHandler });


            var neuronServiceClient = new NeuronService.NeuronServiceClient(channel);
            var response = neuronServiceClient.GetNeuronShape(new GetNeuronShapeRequest()
                { NeuronName = name });

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
            // Initialize gRPC client
            using var httpHandler = new YetAnotherHttpHandler()
                { SkipCertificateVerification = true, Http2Only = true };
            using var channel =
                GrpcChannel.ForAddress(Endpoint, new GrpcChannelOptions() { HttpHandler = httpHandler });
            var cancellationToken = new CancellationTokenSource();

            var neuronServiceClient = new NeuronService.NeuronServiceClient(channel);
            var response = neuronServiceClient.GetMembranePotentials(
                new GetMembranePotentialsRequest() { NeuronName = name });

            // TODO: キャンセル処理ちゃんと実装する
            while (response.ResponseStream.MoveNext(cancellationToken.Token).Result)
            {
                yield return response.ResponseStream.Current;
            }
        }
    }
}