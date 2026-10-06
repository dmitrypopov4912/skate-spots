package main

import (
	"context"
	"fmt"
	"skate-spots/internal/repository/memory"
	"skate-spots/internal/service"
)

func main() {
	ctx := context.Background()
	repo := memory.NewSpotRepository()
	spotService := service.NewSpotService(repo)
	_, err := spotService.CreateSpot(ctx, service.CreateSpotInput{Name: "Парк Горького", Lat: 55.7297, Lon: 37.6014})
	if err != nil {
		fmt.Printf("invalid spot")
	}
	list, err := spotService.ListSpots(ctx)
	if err != nil {
		fmt.Println("smth wrong with repo/list")
	}
	for _, val := range list {
		fmt.Println(val)
	}
}
