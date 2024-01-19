package interpolation

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type WeatherData struct {
	Latitude    float64
	Longitude   float64
	Temperature float64
}

func Interpolation() {
	// Создаем набор данных с метеоданными
	data := []WeatherData{
		{40.7128, -74.0060, 25.0},
		{34.0522, -118.2437, 30.0},
		{51.5074, -0.1278, 20.0},
		{48.8566, 2.3522, 15.0},
	}

	// Инициализируем срезы для координат и данных
	var latitudes []float64
	var longitudes []float64
	var temperatures []float64

	// Заполняем срезы значениями из набора данных
	for _, d := range data {
		latitudes = append(latitudes, d.Latitude)
		longitudes = append(longitudes, d.Longitude)
		temperatures = append(temperatures, d.Temperature)
	}

	// Создаем матрицу расстояний между точками
	distances := make([]float64, len(latitudes)*len(latitudes))
	for i := 0; i < len(latitudes); i++ {
		for j := 0; j < len(latitudes); j++ {
			distance := math.Sqrt(math.Pow(latitudes[i]-latitudes[j], 2) + math.Pow(longitudes[i]-longitudes[j], 2))
			distances[i*len(latitudes)+j] = distance
		}
	}

	// Создаем матрицу весов на основе расстояний
	weights := make([]float64, len(distances))
	floats.Scale(-1, distances)
	floats.Exp(distances)
	floats.Scale(-1, distances)
	for i := 0; i < len(weights); i++ {
		weights[i] = distances[i] / floats.Sum(distances)
	}

	// Создаем матрицу данных
	var dataMatrix mat.Dense
	dataMatrix.SetRawData(len(latitudes), 1, temperatures)

	// Создаем матрицу координат
	coords := make([]float64, len(latitudes)*2)
	for i := 0; i < len(latitudes); i++ {
		coords[i*2] = latitudes[i]
		coords[i*2+1] = longitudes[i]
	}
	var coordMatrix mat.Dense
	coordMatrix.SetRawData(len(latitudes), 2, coords)

	// Создаем интерполятор
	interpolator := &mat.Dense{}
	interpolator.Mul(coordMatrix.T(), dataMatrix)
	interpolator.Mul(interpolator, coordMatrix)
	interpolator.Mul(interpolator, mat.NewDense(len(weights), len(weights), weights))

	// Задаем координаты точки, для которой хотим выполнить интерполяцию
	targetLatitude := 45.0
	targetLongitude := -90.0

	// Выполняем интерполяцию для заданных координат
	var targetMatrix mat.Dense
	targetMatrix.SetRawData(1, 2, []float64{targetLatitude, targetLongitude})
	interpolatedTemperature := mat.Dot(targetMatrix, interpolator)

	// Выводим результат на экран
	fmt.Printf("Интерполированная температура: %.2f\n", interpolatedTemperature.At(0, 0))
}