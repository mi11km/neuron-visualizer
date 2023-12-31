using UnityEngine;


namespace Interfaces
{
    public class Player : MonoBehaviour
    {
        private const float Speed = 0.1f; // for keyboard input

        private void Start()
        {
        }

        private void Update()
        {
            KeyboardInput();
        }

        // プレイヤーを指定した位置が目の前になるように配置する
        public void RepositionInFrontOf(float x, float y, float z)
        {
            transform.position = new Vector3(x, y, z - 50.0f);
            transform.LookAt(new Vector3(x, y, z));
        }


        // キーボード入力によるプレイヤーの移動
        private void KeyboardInput()
        {
            var tf = transform;
            var xMouse = Input.GetAxis("Mouse X");
            var yMouse = Input.GetAxis("Mouse Y");
            var newRotation = tf.localEulerAngles;
            newRotation.y += xMouse * 0.5f;
            newRotation.x -= yMouse * 0.5f;
            tf.localEulerAngles = newRotation;

            var velocity = Vector3.zero;
            if (Input.GetKey(KeyCode.W)) velocity.z = Speed;
            if (Input.GetKey(KeyCode.A)) velocity.x = -Speed;
            if (Input.GetKey(KeyCode.S)) velocity.z = -Speed;
            if (Input.GetKey(KeyCode.D)) velocity.x = Speed;
            if (Input.GetKey(KeyCode.E)) velocity.y = Speed;
            if (Input.GetKey(KeyCode.Q)) velocity.y = -Speed;
            tf.position += tf.rotation * velocity;
        }
    }
}