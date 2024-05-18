package main

import (
	"time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	pos    rl.Vector2
	status bool
	color  rl.Color
	space  rl.Rectangle
}

const (
	CELL_SIZE              = float32(20)
	SCREEN_WIDTH           = 800
	SCREEN_HEIGHT          = 600
	GENERATIONS_PER_SECOND = 10
)

var (
	SELECTED_CELL = false
	PAUSE_GAME = false
	CURRENT_CELL rl.Vector2
	DEBOUNCE = false
)

func main() {

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Conways_GOL")

	cells := make(map[rl.Vector2]Cell)
	current_cells := make(map[rl.Vector2]Cell)

	for j := 0; j < SCREEN_HEIGHT/int(CELL_SIZE); j++ {
		for idx := 0; idx < SCREEN_WIDTH/int(CELL_SIZE); idx++ {

			pos := rl.NewVector2(
				float32(CELL_SIZE*float32(idx)),
				float32(CELL_SIZE*float32(j)))

			cells[pos] = Cell{
				pos,
				false,
				rl.Blue,
				rl.Rectangle{pos.X, pos.Y, CELL_SIZE, CELL_SIZE}}

		}

	}

	rl.SetTargetFPS(GENERATIONS_PER_SECOND)

	ticker := time.NewTicker( 500 * time.Millisecond )
	Interrupt := make(chan bool)
	
	go func() {

		for {
			select {
			case <- Interrupt :
				return
			case <- ticker.C :	
				DEBOUNCE = !DEBOUNCE	
			}
		}
		
	}()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		
		rl.ClearBackground(rl.DarkGray)
		
		if rl.IsKeyPressed(rl.KeyLeftControl) {
			PAUSE_GAME = !PAUSE_GAME			
		}

		if PAUSE_GAME {
		
			for _, c := range cells {
 				
				m_status := MouseAction(cells, c, ticker )

				if SELECTED_CELL {

					rl.DrawRectangleLinesEx(c.space, 2, rl.Blue)
					

				}
				
				cells[rl.Vector2{X: c.pos.X, Y: c.pos.Y}] = Cell{
					rl.Vector2{X: c.pos.X, Y: c.pos.Y},
					m_status, c.color, c.space}
				
			}

		} else {
						
			for _, c := range cells {																
				u_status := checkCells(cells, &c)

				current_cells[rl.Vector2{X: c.pos.X, Y: c.pos.Y}] = Cell{
					rl.Vector2{X: c.pos.X, Y: c.pos.Y},
					u_status, c.color, c.space}
				
			}

			for _, c := range current_cells {

				cells[rl.Vector2{X: c.pos.X, Y: c.pos.Y}] = Cell{
					rl.Vector2{X: c.pos.X, Y: c.pos.Y},
					c.status, c.color, c.space}

			}

		}

		for _, c := range cells {

			if c.status != false {			
				rl.DrawRectangleRec(c.space, c.color)
			}

			rl.DrawRectangleLinesEx(c.space, 0.5, rl.Color{ 255, 255, 255, 64 } )
			
		}


		rl.EndDrawing()

	}

	rl.CloseWindow()
	Interrupt <- true 

}

func MouseAction(cells map[rl.Vector2]Cell, c Cell, ticker *time.Ticker) bool {

	var mouseOnCell bool
	var status bool = c.status
	mouse_pos := rl.GetMousePosition()

	if rl.CheckCollisionPointRec(mouse_pos, c.space) {
		mouseOnCell = true
	} else {
		mouseOnCell = false
	}

	if mouseOnCell {

		SELECTED_CELL = true

		if rl.IsMouseButtonDown(rl.MouseLeftButton) && ( c.pos != CURRENT_CELL ||
		   c.pos == CURRENT_CELL && DEBOUNCE ) {
			status = !status
			CURRENT_CELL = c.pos
			DEBOUNCE = false
			ticker.Reset( 500 * time.Millisecond ) 
		}

	} else {

		SELECTED_CELL = false
		
	}

	return status

}

func checkCells(cells map[rl.Vector2]Cell, c *Cell) bool {

	counter := 0
	status := c.status
	
	pos_c := rl.Vector2{X: c.pos.X - CELL_SIZE, Y: c.pos.Y - CELL_SIZE}

	cell, ok := cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X + CELL_SIZE, Y: c.pos.Y - CELL_SIZE}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X + CELL_SIZE, Y: c.pos.Y + CELL_SIZE}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X + CELL_SIZE, Y: c.pos.Y}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X, Y: c.pos.Y - CELL_SIZE}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X - CELL_SIZE, Y: c.pos.Y}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X, Y: c.pos.Y + CELL_SIZE}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	pos_c = rl.Vector2{X: c.pos.X - CELL_SIZE, Y: c.pos.Y + CELL_SIZE}

	cell, ok = cells[pos_c]

	if ok && cell.status {
		counter++
	}

	if counter <= 1 && c.status == true {
		status = false
	} else if counter >= 4 && c.status == true {
		status = false
	} else if counter == 3 && c.status == false {
		status = true
	}

	return status
	
}
