using UnityEngine;
using Interfaces;

public class Main : MonoBehaviour
{
    [SerializeField] private Camera centerEyeCamera;
    [SerializeField] private Menu menu;
    [SerializeField] private Player player;
    [SerializeField] private NeuronGenerator neuronGenerator;

    private GameObject _generatedNeuronObj;

    private void Start()
    {
        menu.SetMenuMessage("表示するニューロンを選択してください");
        menu.SetNeuronDropdownOptions(neuronGenerator.GetAvailableNeuronNames());
    }

    private void Update()
    {
        if (OVRInput.GetDown(OVRInput.Button.Start)) ToggleMenu();
        if (Input.GetKeyDown(KeyCode.Space)) ToggleMenu();
    }

    public async void OnSelectedNeuronName()
    {
        var neuronName = menu.GetNeuronDropdownSelectedText();
        if (neuronName == "") return;
        // 既に生成されているニューロンの場合は何もしない
        var neuronObj = neuronGenerator.FindGeneratedNeuron(neuronName);
        if (neuronObj != null) return;
        
        menu.SetMenuMessage("ニューロンを生成しています...");

        // ニューロンを1つだけ表示して、プレイヤーをニューロンの前に移動する
        neuronGenerator.DestroyAllNeurons();
        _generatedNeuronObj = await neuronGenerator.GenerateSingleNeuron(neuronName, new Vector3(0, 0, 0));
        var neuronPosition = _generatedNeuronObj.transform.position;
        player.RepositionInFrontOf(neuronPosition, 20.0f);
        menu.gameObject.SetActive(false);
        
        menu.SetMenuMessage($"ニューロン {neuronName} が表示されています");
    }

    public void StartSingleNeuronFiring()
    {
        if (_generatedNeuronObj == null) return;
        menu.ToggleNeuronFiringButtons();
        neuronGenerator.StartSingleNeuronFiring(_generatedNeuronObj);
    }

    public void StopSingleNeuronFiring()
    {
        if (_generatedNeuronObj == null) return;
        menu.ToggleNeuronFiringButtons();
        neuronGenerator.StopSingleNeuronFiring(_generatedNeuronObj);
    }

    private void ToggleMenu()
    {
        var playerPosition = player.transform.position;
        var centerCameraDirection = new Vector3(0, 0, centerEyeCamera.gameObject.transform.position.z);
        menu.ToggleMenu(playerPosition, -centerCameraDirection);
        if (!menu.gameObject.activeSelf) return;
        var menuPosition = menu.transform.position;
        player.RepositionInFrontOf(menuPosition, 0.4f);
    }
}