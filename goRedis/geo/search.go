package geo

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)

func Search() {
	ctx := context.Background()
	locations :=  []*redis.GeoLocation{
		{Name: "beijing", Longitude: 116.41667, Latitude: 39.91667},
		{Name: "shenzhen", Longitude: 114.06667, Latitude: 22.61667},
		{Name: "guangzhou", Longitude: 113.23333, Latitude: 23.16666},
	}

	Client.GeoAdd(ctx, "locations", locations...)

	// 按照杭州这个成员为圆心， 去查找
	fmt.Println(Client.GeoSearch(ctx, "locations", &redis.GeoSearchQuery{
		Member: "shenzhen",
		Radius: 200,
		RadiusUnit: "km",
	}).Result())

		// 按照杭州这个成员为矩形中心， 去查找
	fmt.Println(Client.GeoSearch(ctx, "locations", &redis.GeoSearchQuery{
			Member: "shenzhen",
			RadiusUnit: "km",
			BoxWidth: 200,
			BoxHeight: 200,
			BoxUnit: "km",
		}).Result())
	


		// 按照经纬度为圆心， 去查找
	fmt.Println(Client.GeoSearch(ctx, "locations", &redis.GeoSearchQuery{
			Latitude: 39.9166,
			Longitude: 116.41667,
			Radius: 200,
			RadiusUnit: "km",
		}).Result())


	// 获取位置详细信息
	fmt.Println(Client.GeoSearchLocation(ctx, "locations", &redis.GeoSearchLocationQuery{
		GeoSearchQuery: redis.GeoSearchQuery{
			Member: "shenzhen",
			Radius: 200,
			RadiusUnit: "km",
			Sort: "DESC", // ASC DESC 默认不排序
			Count: 3,
			// CountAny: true,
		},
		WithCoord: true,
		WithDist: true,
		WithHash: true,
	}).Result())

	// 搜索并存储到一个新的 key 当中
	// Client.GeoSearchStore()
}