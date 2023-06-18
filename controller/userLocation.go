package controller

import (
	"log"

	"FurballCommunity_backend/utils/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 保存用户地理位置
/**
传入数据格式：
{
    "name": "user1",
    "Longitude":121.4736977219581604,
    "Latitude": 31.23036910904709629
}
**/
func SetUserLocation(c *gin.Context) {
	var location redis.Location
	c.BindJSON(&location)
	locations := []*redis.Location{
		{Name: location.Name, Longitude: location.Longitude, Latitude: location.Latitude},
	}
	err := redis.RedisGeoAdd("userGeo", locations...)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
}

// 获取指定经纬度中心半径50km内的点,返回的单位是 km
/**
	接口传入数据格式：
{
    "Longitude":121.4736977219581604,
    "Latitude": 31.23036910904709629
}
**/
func GetUserLocationRadius(c *gin.Context) {
	var location redis.Location
	c.BindJSON(&location)
	var locations []*redis.Location
	locations, err := redis.RedisGeoRadius("userGeo", location.Longitude, location.Latitude, 50)
	if err != nil {
		// 处理错误
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	for _, location := range locations {
		log.Println(location.Name, location.Longitude, location.Latitude)
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": locations, "msg": "success"})
}
