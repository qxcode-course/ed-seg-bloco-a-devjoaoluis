package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func drawSquare(pen *Pen, side float64) {

	if side < 5 {
		return
	}

	pen.SetLineWidth(1)
	pen.SetRGB(float64(randInt(0, 255)),
		float64(randInt(0, 255)),
		float64(randInt(0, 255)))
	pen.Walk(side)
	pen.Right(90)
	drawSquare(pen, side-5)
}

func fillCanvas(pen *Pen) {
	pen.SetRGB(0, 0, 0)

}

func main() {
	pen := NewPen(500, 500) // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)   // muda a cor do pincel para vermelho
	pen.SetPosition(0, 500) // move o pincel para x 250, y 500
	pen.SetHeading(90)      // coloca o pincel apontando para cima
	// pen.Walk(100)             // anda 100 pixels
	// pen.Left(30)              // dobra 30 graus para esquerda
	// pen.Walk(180)             // anda 100 pixels
	// pen.DrawCircle(50)        // desenha um círculo de raio 50
	// pen.Right(60)             // gira para direita 60 graus
	// pen.Walk(150)
	// pen.FillSquare()
	pen.SetRGB(0, 0, 0)
	pen.Walk(500)
	pen.FillSquare(500, 500)
	pen.SetPosition(10, 500)

	//Desenha quadrado
	drawSquare(pen, 500)

	pen.SavePNG("square.png")
	fmt.Println("PNG file created successfully.")
}

// pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
// 	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
// 	pen.SetPosition(250, 500) // move o pincel para x 250, y 500
// 	pen.SetHeading(90)        // coloca o pincel apontando para cima
// 	pen.Walk(100)             // anda 100 pixels
// 	pen.Left(30)              // dobra 30 graus para esquerda
// 	pen.Walk(100)             // anda 100 pixels
// 	pen.DrawCircle(50)        // desenha um círculo de raio 50
// 	pen.Right(60)             // gira para direita 60 graus
// 	pen.Walk(150)
// 	for range 10 {
// 		pen.Up()
// 		pen.Walk(30) // anda sem riscar
// 		pen.Down()

// 		pen.DrawCircle(10) //desenha um circulo pequeno

// 		pen.Up()
// 		pen.Walk(-30) // volta sem riscar
// 		pen.Down()

// 		pen.Left(36) // gira
// 	}

// 	pen.SavePNG("tree.png")
// 	fmt.Println("PNG file created successfully.")
