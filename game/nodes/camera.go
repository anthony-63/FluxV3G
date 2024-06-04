package nodes

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	Position rl.Vector3

	Pitch float32
	Yaw   float32

	RlCamera rl.Camera3D
}

func NewCamera(position rl.Vector3) Camera {
	camera := Camera{}

	camera.RlCamera = rl.Camera3D{}
	camera.RlCamera.Position = position
	camera.RlCamera.Fovy = 70
	camera.RlCamera.Up = rl.Vector3{
		X: 0,
		Y: 1,
		Z: 0,
	}

	camera.RlCamera.Target = rl.Vector3Zero()
	camera.RlCamera.Projection = rl.CameraPerspective

	return camera
}

func (cam *Camera) Update(dt float64) {

}

func (cam *Camera) Draw() {

}
