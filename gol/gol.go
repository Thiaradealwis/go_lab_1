package main

import "fmt"

func calculateNextState(p golParams, world [][]byte) [][]byte {
	//turns := p.turns
	image_width := p.imageWidth
	image_height := p.imageHeight
	new_world := make([][]byte, image_height)
	for i := 0; i < image_height; i++ {
		new_world[i] = make([]byte, image_width)
	}

	for y := 0; y < image_height; y++ {
		for x := 0; x < image_width; x++ {

			var live = world[(y+image_height-1)%image_height][(x+image_width-1)%image_width] +
				world[(y+image_height-1)%image_height][(x+image_width)%image_width] +
				world[(y+image_height-1)%image_height][(x+image_width+1)%image_width] +
				world[(y+image_height)%image_height][(x+image_width-1)%image_width] +
				world[(y+image_height)%image_height][(x+image_width+1)%image_width] +
				world[(y+image_height+1)%image_height][(x+image_width-1)%image_width] +
				world[(y+image_height+1)%image_height][(x+image_width)%image_width] +
				world[(y+image_height+1)%image_height][(x+image_width+1)%image_width]
			live = live / 255

			if world[y][x] == 255 {
				if live < 2 {
					new_world[y][x] = 0
				}
				if live > 3 {
					new_world[y][x] = 0
				}
				if live == 2 {
					new_world[y][x] = world[y][x]
				}
				if live == 3 {
					new_world[y][x] = world[y][x]
				}
			} else {
				if live == 3 {
					new_world[y][x] = 255
				} else {
					new_world[y][x] = world[y][x]
				}
			}

		}
	}
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	image_width := p.imageWidth
	image_height := p.imageHeight
	var alive []cell

	for y := 0; y < image_height; y++ {
		for x := 0; x < image_width; x++ {
			if world[y][x] == 255 {
				p := cell{x, y}
				//print("(", x, ",", y, ")")
				alive = append(alive, p)
			}

		}
	}
	fmt.Println("alive", alive)
	return alive
}
