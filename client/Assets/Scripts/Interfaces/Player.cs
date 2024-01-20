using System;
using UnityEngine;

namespace Interfaces
{
    [RequireComponent(typeof(CharacterController))]
    public class Player : MonoBehaviour
    {
        [SerializeField] private Camera centerEyeCamera;

        private CharacterController _controller;
        private const float Speed = 0.1f; // for keyboard input
        private const float MSpeed = 20f; // for keyboard input
        private Vector3 _movement;
        private Vector3 _moveDir = Vector3.zero;
        private float _moveH;
        private float _moveV;

        private void Awake()
        {
            _controller = GetComponent<CharacterController>();
        }

        private void Update()
        {
            KeyboardInput();
            MetaQuestInput();
        }

        // プレイヤーを指定した位置が目の前になるように配置する
        public void RepositionInFrontOf(Vector3 position, float distance = 5.0f)
        {
            var x = position.x;
            var y = position.y;
            var z = position.z;
            transform.position = new Vector3(x, y, z - distance);
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

        private void MetaQuestInput()
        {
            if (OVRInput.Get(OVRInput.Button.SecondaryThumbstickUp) ||
                OVRInput.Get(OVRInput.Button.SecondaryThumbstickDown) ||
                OVRInput.Get(OVRInput.Button.SecondaryThumbstickLeft) ||
                OVRInput.Get(OVRInput.Button.SecondaryThumbstickRight))
            {
                _moveH = OVRInput.Get(OVRInput.RawAxis2D.RThumbstick).x;
                _moveV = OVRInput.Get(OVRInput.RawAxis2D.RThumbstick).y;
                _movement = new Vector3(_moveH, 0, _moveV);
                var cameraTransform = centerEyeCamera.transform;
                _moveDir = cameraTransform.forward * _movement.z + cameraTransform.right * _movement.x;
                _controller.Move(MSpeed * Time.deltaTime * _moveDir);
            }

            if (OVRInput.Get(OVRInput.Button.Two))
                transform.Translate(MSpeed * Time.deltaTime * centerEyeCamera.transform.up);
            if (OVRInput.Get(OVRInput.Button.One))
                transform.Translate(MSpeed * Time.deltaTime * -1 * centerEyeCamera.transform.up);

            if (OVRInput.Get(OVRInput.Button.PrimaryThumbstickLeft))
                transform.Rotate(0, 2 * MSpeed * Time.deltaTime, 0);
            if (OVRInput.Get(OVRInput.Button.PrimaryThumbstickRight))
                transform.Rotate(0, -2 * MSpeed * Time.deltaTime, 0);
            if (OVRInput.Get(OVRInput.Button.PrimaryThumbstickUp)) transform.Rotate(2 * MSpeed * Time.deltaTime, 0, 0);
            if (OVRInput.Get(OVRInput.Button.PrimaryThumbstickDown))
                transform.Rotate(-2 * MSpeed * Time.deltaTime, 0, 0);
        }
    }
}