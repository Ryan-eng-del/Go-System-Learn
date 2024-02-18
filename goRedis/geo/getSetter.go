package geo

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)

func Add() {	
	ctx := context.Background()
	locations :=  []*redis.GeoLocation{
		{Name: "beijing", Longitude: 116.41667, Latitude: 39.91667},
		{Name: "shenzhen", Longitude: 114.06667, Latitude: 22.61667},
		{Name: "guangzhou", Longitude: 113.23333, Latitude: 23.16666},
	}

	Client.GeoAdd(ctx, "locations", locations...)
	gens, _:= Client.GeoPos(ctx, "locations", "shenzhen").Result()

	for _, gen := range gens {
		fmt.Println(gen)
	}	

	fmt.Println(Client.GeoHash(ctx, "locations", "shenzhen").Result())

	fmt.Println(Client.GeoDist(ctx, "locations", "shenzhen", "guangzhou", "km").Result())


}