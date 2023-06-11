package redis

import (
	"context"
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
// redis命令行远程连接命令：redis-cli -h 47.113.230.181 -p 6379 -a 123456
var rdb = redis.NewClient(&redis.Options{
	Addr:     "47.113.230.181:6379",
	Password: "123456", // Redis密码
	DB:       0,        // 数据库编号
})

// 设置普通类型键值对，second为0时永久有效
func RedisSet(key string, value interface{}, second time.Duration) error {
	return rdb.Set(ctx, key, value, second).Err()
}

// 获取普通类型键值对
func RedisGet(name string) (string, error) {
	return rdb.Get(ctx, name).Result()
}

// 删除普通类型键值对
func RedisDel(name ...string) error {
	return rdb.Del(ctx, name...).Err()
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
		locations = append(locations, &Location{
			Name:      geoLocation.Name,
			Longitude: geoLocation.Longitude,
			Latitude:  geoLocation.Latitude,
			Distance:  geoLocation.Dist,
		})
	}
	return locations, nil
}

// 保存多个key-value键值对
// redis.RedisHMset("user1", "name", "wang", "age", 21, "sex", 0, "city", "Beijing")
func RedisHMset(key string, values ...interface{}) error {
	return rdb.HMSet(ctx, key, values).Err()
}

// 获取指定key中指定field的值
// value, err := redis.RedisHGet("user1", "name")
func RedisHGet(key string, field string) (string, error) {
	return rdb.HGet(ctx, key, field).Result()
}

// 关闭连接
func RedisClose() error {
	return rdb.Close()
}
