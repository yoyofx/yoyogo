package redis

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
	Unit      GeoUnit
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

	location := GeoPosition{
		Longitude: longitude,
		Latitude:  latitude,
		Member:    member,
	}
	res := geo.ops.GeoAddArr(key, location)
	return res
}

/*
添加一批地理信息
*/
func (geo *Geo) GeoAddArr(key string, geoPosition []GeoPosition) int64 {
	res := geo.ops.GeoAddArr(key, geoPosition...)
	return res
}

/*
根据地理位置名称获取经纬度
*/
func (geo *Geo) GeoPos(key string, member string) (err error, geoRes GeoPosition) {
	err, resList := geo.ops.GeoPos(key, member)
	if err != nil {
		return err, GeoPosition{}
	}
	return nil, resList[0]
}

/*
获取一批地理位置的经纬度
*/
func (geo *Geo) GeoPosArr(key string, members []string) (err error, geoRes []GeoPosition) {
	return geo.ops.GeoPos(key, members...)
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
	return geo.ops.GeoDist(key, member1, member2, unit)
}

/**
经纬度中心距离计算
*/
func (geo *Geo) GeoRadius(key string, query GeoRadiusQuery) (error, []GeoPosition) {
	return geo.ops.GeoRadius(key, query)
}

/**
地理标识中心距离计算
*/
func (geo *Geo) GeoRadiusByMember(key string, query GeoRadiusByMemberQuery) (error, []GeoPosition) {
	return geo.ops.GeoRadiusByMember(key, query.Member, query)
}
