package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 640
	gridSize     = 20
)

type Game struct {
	snake    []Vector2
	food     Vector2
	speed    int
	direction Vector2
	gameOver bool
	score    int
}

type Vector2 struct {
	X, Y int
}

func (v Vector2) Equals(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

func resetGame() *Game {
	game := &Game{
		snake:    []Vector2{{10, 10}, {9, 10}, {8, 10}},
		food:     Vector2{rand.Intn(screenWidth/gridSize), rand.Intn(screenHeight/gridSize)},
		speed:    10,
		direction: Vector2{1, 0},
		gameOver: false,
		score:    0,
	}
	return game
}

func updateFoodPosition(game *Game) {
	game.food = Vector2{rand.Intn(screenWidth/gridSize), rand.Intn(screenHeight/gridSize)}
}

func checkCollision(game *Game) bool {
	head := game.snake[0]

	// 检查蛇头是否碰到边界
	if head.X < 0 || head.X >= screenWidth/gridSize ||
		head.Y < 0 || head.Y >= screenHeight/gridSize {
		return true
	}

	// 检查蛇头是否碰到自己的身体
	for _, body := range game.snake[1:] {
		if body.Equals(head) {
			return true
		}
	}

	return false
}

func updateSnake(game *Game) {
	// 移动蛇的身体
	newHead := Vector2{
		X: game.snake[0].X + game.direction.X,
		Y: game.snake[0].Y + game.direction.Y,
	}
	game.snake = append([]Vector2{newHead}, game.snake...)

	// 如果吃到了食物，增加蛇的长度
	if game.snake[0].Equals(game.food) {
		game.score++
		updateFoodPosition(game)
	} else {
		// 否则，移除蛇尾
		game.snake = game.snake[:len(game.snake)-1]
	}
}

func drawGame(screen *ebiten.Image, game *Game) {
	// 绘制背景
	ebitenutil.DrawRect(screen, 0, 0, screenWidth, screenHeight, color.RGBA{0.2, 0.6, 1, 1})

	// 绘制蛇的身体
	for i, segment := range game.snake {
		var image *ebiten.Image
		if i == 0 {
			// 蛇头
			image = assets["head"]
		} else {
			// 蛇身
			image = assets["body"]
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Scale(float64(gridSize), float64(gridSize))
		op.GeoM.Translate(float64(segment.X*gridSize), float64(segment.Y*gridSize))
		if i == 0 {
			// 蛇头的方向旋转
			switch game.direction {
			case Vector2{1, 0}:
			/op.GeoM.Rotate(0) // 右
			case Vector2{-1, 0}:
			/op.GeoM.Rotate(180) // 左
			case Vector2{0, 1}:
			/op.GeoM.Rotate(90) // 上
			case Vector2{0, -1}:
			/op.GeoM.Rotate(270) // 下
			}
		}
		screen.DrawImage(image, op)
	}

	// 绘制食物
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Scale(float64(gridSize), float64(gridSize))
	op.GeoM.Translate(float64(game.food.X*gridSize), float64(game.food.Y*gridSize))
	screen.DrawImage(assets["food"], op)

	// 绘制分数
	ebitenutil.DebugPrint(screen, "Score: "+strconv.Itoa(game.score), 10, 10)

}

func update(screen *ebiten.Image, input *ebiten.InputState, ttime time.Time) error {
	if game.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Press 'R' to restart.", screenWidth/2-100, screenHeight/2)
		if input.KeyTriggered(ebiten.KeyR) {
			game = resetGame()
		}
		return nil
	}

	// 处理输入
	if input.KeyTriggered(ebiten.KeyUp) && game.direction.Y != 1 {
		game.direction = Vector2{0, -1}
	}
	if input.KeyTriggered(ebiten.KeyDown) && game.direction.Y != -1 {
		game.direction = Vector2{0, 1}
	}
	if input.KeyTriggered(ebiten.KeyLeft) && game.direction.X != 1 {
		game.direction = Vector2{-1, 0}
	}
	if input.KeyTriggered(ebiten.KeyRight) && game.direction.X != -1 {
		game.direction = Vector2{1, 0}
	}

	// 更新蛇的位置
	updateSnake(game)

	// 检查碰撞
	if checkCollision(game) {
		game.gameOver = true
	}

	// 绘制游戏
	drawGame(screen, game)

	return nil
}

func main() {
	game = resetGame()

	assets := make(map[string]*ebiten.Image)
	assets["head"] = ebiten.NewImageFromFile("assets/head.png")
	assets["body"] = ebiten.NewImageFromFile("assets/body.png")
	assets["food"] = ebiten.NewImageFromFile("assets/food.png")

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("贪吃蛇游戏")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
