package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// 定义地理位置信息结构体
type Location struct {
	Name      string  // 地点名称
	Longitude float64 // 地点经度
	Latitude  float64 // 地点纬度
	Distance  float64 //距离中心点距离
}

var ctx = context.Background()

// 创建redis客户端
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // Redis密码
	DB:       0,  // 数据库编号
})

// 设置普通类型键值对
func RedisSet(key string, value interface{}, second time.Duration) error {
	return rdb.Set(ctx, key, value, second).Err()
}

// 获取普通类型键值对
func RedisGet(name string) (string, error) {
	return rdb.Get(ctx, name).Result()
}

// 将地理位置信息存储到Redis中
func RedisGeoAdd(key string, locations ...*Location) error {
	var err error
	for _, location := range locations {
		geoLocation := &redis.GeoLocation{
			Name:      location.Name,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		}
		err = rdb.GeoAdd(ctx, key, geoLocation).Err()
		if err != nil {
			break
		}
	}
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 获取给定坐标指定范围内的点
func RedisGeoRadius(key string, longitude, latitude, radius float64) ([]*Location, error) {
	var locations []*Location
	// radius Default is km
	geoLocations, err := rdb.GeoRadius(ctx, key, longitude, latitude, &redis.GeoRadiusQuery{
		Radius:    radius,
		WithCoord: true,
		WithDist:  true,
		Unit:      "km",
	}).Result()
	if err != nil {
		return nil, err
	}
	for _, geoLocation := range geoLocations {
		log.Println("geoLocation.Name:", geoLocation.Name)
		log.Println("geoLocation.Longitude:", geoLocation.Longitude)
		log.Println("geoLocation.Latitude:", geoLocation.Latitude)
		log.Println("geoLocation.Distance:", geoLocation.Dist)
		locations = append(locations, &Location{
			Name:      geoLocation.Name,
			Longitude: geoLocation.Longitude,
			Latitude:  geoLocation.Latitude,
			Distance:  geoLocation.Dist,
		})
	}
	return locations, nil
}

// 关闭连接
func RedisClose() error {
	return rdb.Close()
}
