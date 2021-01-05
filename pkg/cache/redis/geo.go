package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

type GeoEnum string

type GeoUnit string

/*
长度计量单位
*/
const (
	KM        GeoUnit = "km"
	M         GeoUnit = "m"
	MI        GeoUnit = "mi"
	FT        GeoUnit = "ft"
	WITHCOORD GeoEnum = "WITHCOORD"
	WITHDIST  GeoEnum = "WITHDIST"
	WITHHASH  GeoEnum = "WITHHASH"
	ASC       GeoEnum = "ASC"
	DESC      GeoEnum = "DESC"
)

type Geo struct {
	ops Ops
}

type GeoPosition struct {
	Longitude float64
	Latitude  float64
	Dist      float64
	Member    string
	GeoHash   int64
}

type GeoDistInfo struct {
	Unit GeoUnit
	Dist float64
}

type GeoRadiusQuery struct {
	Longitude   float64
	Latitude    float64
	Radius      float64
	Unit        GeoUnit
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      GeoEnum
	Store     string
	StoreDist string
}

type GeoRadiusByMemberQuery struct {
	Member      string
	Radius      float64
	Unit        GeoUnit
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      GeoEnum
	Store     string
	StoreDist string
}

/*
添加一个地理信息
*/
func (geo *Geo) GeoAdd(key string, longitude float64, latitude float64, member string) int64 {

	location := redis.GeoLocation{
		Longitude: longitude,
		Latitude:  latitude,
		Name:      member,
	}
	res := geo.ops.GeoAddArr(key, &location)
	return res.Val()
}

/*
添加一批地理信息
*/
func (geo *Geo) GeoAddArr(key string, geoPosition []GeoPosition) int64 {
	var geoList = make([]*redis.GeoLocation, 2)
	for _, x := range geoPosition {
		geoEle := redis.GeoLocation{
			Longitude: x.Longitude,
			Latitude:  x.Latitude,
			Name:      x.Member,
		}
		geoList = append(geoList, &geoEle)
	}
	res := geo.ops.GeoAddArr(key, geoList...)
	return res.Val()
}

/*
根据地理位置名称获取经纬度
*/
func (geo *Geo) GeoPos(key string, member string) (err error, geoRes GeoPosition) {

	resList := geo.ops.GeoPos(key, member)
	resEle := resList.Val()[0]
	if resEle == nil {
		return errors.New("this member dont have any info"), GeoPosition{}
	}
	return nil, GeoPosition{
		Longitude: resEle.Longitude,
		Latitude:  resEle.Latitude,
		Member:    member,
	}
}

/*
获取一批地理位置的经纬度
*/
func (geo *Geo) GeoPosArr(key string, members []string) (err error, geoRes []GeoPosition) {
	resList := geo.ops.GeoPos(key, members...)
	if len(resList.Val()) == 0 {
		return errors.New("not find any geo info"), make([]GeoPosition, 0)
	}
	resGeoList := make([]GeoPosition, 2)
	resListVal := resList.Val()
	for i, x := range members {
		resValEle := resListVal[i]
		if resValEle != nil {
			resGeoList = append(resGeoList, GeoPosition{Longitude: resValEle.Longitude, Latitude: resValEle.Latitude, Member: x})
		}
	}
	return nil, resGeoList
}

func getUnit(unit GeoUnit) string {

	var res string = ""
	switch unit {
	case KM:
		res = "km"
	case M:
		res = "m"
	case FT:
		res = "ft"
	case MI:
		res = "mi"
	}
	return res
}
func GetSort(sort GeoEnum) string {
	var res string = ""
	switch sort {
	case ASC:
		res = "ASC"
	case DESC:
		res = "DESC"
	default:
		res = "ASC"
	}
	return res
}

/**
获取两个坐标的直线距离
*/
func (geo *Geo) GeoDist(key string, member1, member2 string, unit GeoUnit) (error, GeoDistInfo) {
	unitStr := getUnit(unit)
	if unitStr == "" {
		return errors.New("error unit"), GeoDistInfo{}
	}
	res := geo.ops.GeoDist(key, member1, member2, unitStr)
	return nil, GeoDistInfo{Unit: unit, Dist: res.Val()}
}

/**
经纬度中心距离计算
*/
func (geo *Geo) GeoRadius(key string, query GeoRadiusQuery) (error, []GeoPosition) {
	unitStr := getUnit(query.Unit)
	if unitStr == "" {
		return errors.New("error unit"), make([]GeoPosition, 0)
	}
	res := geo.ops.GeoRadius(key, query.Longitude, query.Latitude, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        unitStr,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        GetSort(query.Sort),
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	})
	geoList := make([]GeoPosition, query.Count)
	for _, x := range res.Val() {
		geoList = append(geoList, GeoPosition{
			Member:    x.Name,
			Longitude: x.Longitude,
			Latitude:  x.Latitude,
			Dist:      x.Dist,
			GeoHash:   x.GeoHash,
		})
	}
	return nil, geoList
}

/**
地理标识中心距离计算
*/
func (geo *Geo) GeoRadiusByMember(key string, query GeoRadiusByMemberQuery) (error, []GeoPosition) {
	unitStr := getUnit(query.Unit)
	if unitStr == "" {
		return errors.New("error unit"), make([]GeoPosition, 0)
	}
	res := geo.ops.GeoRadiusByMember(key, query.Member, &redis.GeoRadiusQuery{
		Radius:      query.Radius,
		Unit:        unitStr,
		WithCoord:   query.WithCoord,
		WithDist:    query.WithDist,
		WithGeoHash: query.WithGeoHash,
		Count:       query.Count,
		Sort:        GetSort(query.Sort),
		Store:       query.Store,
		StoreDist:   query.StoreDist,
	})
	geoList := make([]GeoPosition, query.Count)
	for _, x := range res.Val() {
		geoList = append(geoList, GeoPosition{
			Member:    x.Name,
			Longitude: x.Longitude,
			Latitude:  x.Latitude,
			Dist:      x.Dist,
			GeoHash:   x.GeoHash,
		})
	}
	return nil, geoList
}
