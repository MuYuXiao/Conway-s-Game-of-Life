package ui

import (
	"conway/game"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GameCanvas struct {
	board *game.Board
	cells [][]*canvas.Rectangle
}

func ShowGameWindows() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Conways")
	row, col := 50, 50
	board := game.NewBoard(row, col)
	board.RandBoard()
	//board逻辑
	gameCanvas := NewGameCanvas(&board)
	gameCanvas.initCanvas()
	//board渲染
	gridLayout := make([]fyne.CanvasObject, row*col)
	gridLayout = *renderBoard(&gridLayout, gameCanvas, col)
	//board容器
	boardContainer := container.New(layout.NewGridLayoutWithColumns(col), gridLayout...)
	/*
		按钮区
	*/
	var continueUpdating bool
	//开始按钮
	StartButton := widget.NewButton("Start", func() {
		continueUpdating = true
		go func() {
			ticker := time.NewTicker(30 * time.Millisecond)
			defer ticker.Stop()
			for range ticker.C {
				if !continueUpdating {
					return
				}
				board.Update()
				gameCanvas.UpdateCanvas(&board)
				boardContainer.Objects = container.New(layout.NewGridLayoutWithColumns(col), *renderBoard(&gridLayout, gameCanvas, col)...).Objects
				boardContainer.Refresh()
			}
		}()
	})
	//更新按钮
	UpdateButton := widget.NewButton("Update", func() {
		board.Update()
		gameCanvas.UpdateCanvas(&board)
		boardContainer.Objects = container.New(layout.NewGridLayoutWithColumns(col), *renderBoard(&gridLayout, gameCanvas, col)...).Objects
		boardContainer.Refresh()
	})
	//停止按钮
	ResetButton := widget.NewButton("Reset", func() {
		board = game.NewBoard(row, col)
		board.RandBoard()
		gameCanvas.UpdateCanvas(&board)
		boardContainer.Objects = container.New(layout.NewGridLayoutWithColumns(col), *renderBoard(&gridLayout, gameCanvas, col)...).Objects
		boardContainer.Refresh()
	})
	// 添加暂停逻辑的按钮
	PauseButton := widget.NewButton("Pause", func() {
		// 设置继续逻辑的标志为 false
		continueUpdating = false
	})
	//Button容器
	buttonsContainer := container.New(layout.NewHBoxLayout(),
		StartButton,
		UpdateButton,
		PauseButton,
		ResetButton,
	)
	//fyne渲染
	content := container.New(layout.NewVBoxLayout(),
		buttonsContainer,
		boardContainer,
	)
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// 新建board画布
func NewGameCanvas(b *game.Board) *GameCanvas {
	myCanvas := &GameCanvas{
		board: b,
		cells: make([][]*canvas.Rectangle, len(*b)),
	}
	for i := range myCanvas.cells {
		myCanvas.cells[i] = make([]*canvas.Rectangle, len((*b)[0]))
	}
	return myCanvas
}

func (g *GameCanvas) initCanvas() {
	black := color.Black
	white := color.White
	row, col := len(*(*g).board), len((*(*g).board)[0])
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if (*(*g).board)[x][y] {
				(*g).cells[x][y] = canvas.NewRectangle(white)
			} else {
				(*g).cells[x][y] = canvas.NewRectangle(black)
			}
		}
	}
}

func (g *GameCanvas) UpdateCanvas(b *game.Board) {
	g.board = b
	g.initCanvas()
}

func renderBoard(gridLayout *[]fyne.CanvasObject, gameCanvas *GameCanvas, col int) *[]fyne.CanvasObject {
	for i := range gameCanvas.cells {
		for j := range gameCanvas.cells[i] {
			(*gridLayout)[i*col+j] = gameCanvas.cells[i][j]
		}
	}
	return gridLayout
}
