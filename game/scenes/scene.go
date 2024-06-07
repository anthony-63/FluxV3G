package scenes

const (
	SCENE_TYPE_STARTUP = iota
	SCENE_TYPE_DEBUG
	SCENE_TYPE_MENU
	SCENE_TYPE_GAME
)

var SceneList = []IScene{}

type IScene interface {
	Draw()
	Update(float64)
	GetType() int
}
