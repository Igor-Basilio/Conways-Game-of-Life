Implementation of Conways Game of Life using Golang + Raylib 

The rules of the game are as follows : ( slightly different from the ones on the wikipedia page )

1. Any alive cell with one or less neighbours dies
2. Any alive cell with four or more neighbours dies
3. Any dead cell with exactly three neighbours lives

If none of those rules apply then the cell moves on to the next generation.

Controls : 

  1. Left Control to pause the game and be able to select and click cells with the mouse.

Running : 

  1. All you need is to have Go installed, https://go.dev/doc/install
  2. Change into the main directory of the project
  3. Run : go run ./main
