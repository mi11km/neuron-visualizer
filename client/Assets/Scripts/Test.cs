using System.Collections;
using System.Collections.Generic;
using System.Threading;
using Cysharp.Net.Http;
using Grpc.Net.Client;
using Health.V1;
using UnityEngine;


public class Test : MonoBehaviour
{
    const string Endpoint = "http://localhost:8080";

    // Start is called before the first frame update
    async void Start()
    {
        // Initialize gRPC client
        using var httpHandler = new YetAnotherHttpHandler() { SkipCertificateVerification = true, Http2Only = true };
        using var channel = GrpcChannel.ForAddress(Endpoint, new GrpcChannelOptions() { HttpHandler = httpHandler });

        var healthCheckServiceClient = new HealthService.HealthServiceClient(channel);
        var response = healthCheckServiceClient.Check(new CheckRequest() { Service = "NeuronService" });
        Debug.Log(response.ToString());

        var streamResponse = healthCheckServiceClient.Watch(new WatchRequest() { Service = "NeruronServcie" });

        var cancellationTokenSource = new CancellationTokenSource();
        while (await streamResponse.ResponseStream.MoveNext(cancellationTokenSource.Token))
        {
            Debug.Log(streamResponse.ResponseStream.Current.ToString());
        }
    }

    // Update is called once per frame
    void Update()
    {
    }
}