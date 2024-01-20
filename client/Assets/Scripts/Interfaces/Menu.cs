using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.UI;

namespace Interfaces
{
    public class Menu : MonoBehaviour
    {
        [SerializeField] private TMP_Dropdown neuronDropdown;
        [SerializeField] private Button startFiringButton;
        [SerializeField] private Button stopFiringButton;
        [SerializeField] private TextMeshProUGUI menuMessage;

        private void Awake()
        {
            gameObject.SetActive(false);
        }

        /// <summary>
        /// メニューの表示・非表示を切り替える
        /// </summary>
        /// <param name="position">メニューを表示する位置</param>
        /// <param name="direction">メニューを表示する向き</param>
        public void ToggleMenu(Vector3 position, Vector3 direction)
        {
            var self = gameObject;
            if (!self.activeSelf)
            {
                self.transform.position = new Vector3(position.x, position.y, position.z);
                self.transform.LookAt(direction);
            }

            self.SetActive(!self.activeSelf);
        }

        /*
         * ニューロン選択ドロップダウン
         */

        /// <summary>
        /// ニューロン選択ドロップダウンの選択されているテキストを返す
        /// </summary>
        /// <returns>ニューロン選択ドロップダウンで選択されているテキスト</returns>
        public string GetNeuronDropdownSelectedText()
        {
            return neuronDropdown.options[neuronDropdown.value].text;
        }

        /// <summary>
        /// ニューロン選択ドロップダウンの選択しを設定する
        /// </summary>
        /// <param name="neuronNames">ニューロン名のリスト</param>
        public void SetNeuronDropdownOptions(List<string> neuronNames)
        {
            neuronDropdown.ClearOptions();
            neuronDropdown.AddOptions(new List<string> {""});
            neuronDropdown.AddOptions(neuronNames);
        }
        
        /*
         * ニューロン発火ボタン
         */
        public void ToggleNeuronFiringButtons()
        {
            startFiringButton.gameObject.SetActive(!startFiringButton.gameObject.activeSelf);
            stopFiringButton.gameObject.SetActive(!stopFiringButton.gameObject.activeSelf);
        }
        
        /*
         * メニューメッセージ
         */
        public void SetMenuMessage(string message)
        {
            menuMessage.SetText(message);
        }
    }
}