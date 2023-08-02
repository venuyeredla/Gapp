package ai

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Vector struct {
	Data []float32
}

type Matrix struct {
	Data [][]float32
}

// Tensor of size 3
type Tensor struct {
	Data [][][]float32
}

func (matrix *Matrix) Shape() (int, int) {
	return len(matrix.Data), len(matrix.Data[0])
}

func (matrix *Matrix) Print() {
	fmt.Print("Matrix : \n")
	for _, row := range matrix.Data {
		fmt.Println(row)
	}
	fmt.Print("\n")
}

func NewVector(length int) *Vector {
	vec := make([]float32, length)
	return &Vector{Data: vec}
}

func (matrix *Vector) Shape() int {
	return len(matrix.Data)
}

func (vector *Vector) InitVector(defa float32) {
	length := vector.Shape()
	for i := 0; i < length; i++ {
		vector.Data[i] = defa
	}
}

// Slices are alway one dimensional
func VectorR(length uint) *Matrix {
	return GenMatrix(1, 5)
}

func VectorC(length uint) *Matrix {
	return GenMatrix(5, 1)
}

func GenMatrix(rows int, columns int) *Matrix {
	var matrix [][]float32 = make([][]float32, rows)
	for i := 0; i < rows; i++ {
		vector := make([]float32, columns)
		matrix[i] = vector
	}
	return &Matrix{Data: matrix}
}

func (matrix *Matrix) Init(defa float32) {
	rows, columns := matrix.Shape()
	rand.Seed(time.Now().UnixMilli())
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			matrix.Data[row][col] = defa
		}
	}
}

func GenTensor(size int, rows int, columns int) *Tensor {
	var tensor [][][]float32
	tensor = make([][][]float32, size)
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < size; i++ {
		var matix = make([][]float32, rows)
		for j := 0; j < rows; j++ {
			vector := make([]float32, columns)
			for k := 0; k < columns; k++ {
				vector[j] = rand.Float32()
			}
			matix[j] = vector
		}
		tensor[i] = matix
	}
	return &Tensor{Data: tensor}
}

func (tensor *Tensor) PrintTensor() {
	fmt.Printf("Tensor: [\n")
	for _, matrix := range tensor.Data {
		for _, row := range matrix {
			fmt.Println(row)
		}
		fmt.Printf("\n")
	}
	fmt.Print("]")
}

// c= Xt y
func DotProduct(a, b *Vector) float32 {
	if a.Shape() == b.Shape() {
		var scalar float32 = 0.0
		for i := 0; i < a.Shape(); i++ {
			scalar = scalar + a.Data[i] + b.Data[i]
		}
		return scalar
	}
	return 0.0
}

// d=sqrt(( pow(x2-x1,2) + pow(y2-y1,1,2))
func Distance(a, b Vector) float32 {
	if a.Shape() == b.Shape() {
		var defsquare float64 = 0.0
		for i := 0; i < a.Shape(); i++ {
			defsquare = math.Pow(float64(a.Data[i]-b.Data[i]), 2)
		}
		distance := math.Sqrt(defsquare)
		return float32(distance)
	}
	return 0.0
}

func MatrixAdd(a [][]int, b [][]int) {

}

func MatrixMultiply(a, b Matrix) *Matrix {
	var c [][]float32 = make([][]float32, 2)
	for i := 0; i < len(a.Data); i++ {
		c[i] = make([]float32, 2)
		for j := 0; j < len(b.Data); j++ {
			var cij float32 = 0.0
			for k := 0; k < len(a.Data[0]); k++ {
				cij = cij + a.Data[i][k]*b.Data[k][j]
			}
			c[i][j] = cij
		}
	}
	return &Matrix{Data: c}
}
