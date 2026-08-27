package main

import (
	"fmt"

	"github.com/thiagozs/go-xutils/v2/geo"
)

func main() {
	latitude := -23.5505
	longitude := -46.6333

	fmt.Println("latitude válida:", geo.IsLatitude(latitude))
	fmt.Println("longitude válida:", geo.IsLongitude(longitude))
	fmt.Println("coordenadas válidas:", geo.AreCoordinates(latitude, longitude))
}
