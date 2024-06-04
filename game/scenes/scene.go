package scenes

var SceneList = []IScene{}

type IScene interface {
	Draw()
	Update(float64)
}
