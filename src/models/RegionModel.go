package models

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"strconv"
	"net"
	"encoding/binary"
	"errors"
)

type RegionModel struct {
	ID       int    //id编号
	IP       string //ip段
	StartIP  uint32  //开始IP
	StopIP   uint32  //结束IP
	Country  string    //国家
	Province string    //省
	City     string    //市
	District string    //区
	Isp      string    //运营商
	Type     string //类型
	Desc     string //说明
}


var regionFileName = "../res/region.db"

var RegionArray []RegionModel

func InitRegionModel() {

	buf, err := ioutil.ReadFile(regionFileName)

	if err != nil {
		buf, err = ioutil.ReadFile("./res/region.db")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		return
	}


	arr := strings.Split(string(buf), "\n")
	RegionArray = make([]RegionModel, len(arr))

	for i, v := range arr {
		tempArr := strings.Split(v, ",")
		if ( len(tempArr) < 10 ) { continue }
		RegionArray[i] = RegionModel{}
		RegionArray[i].ID = i
		RegionArray[i].IP = tempArr[0]
		RegionArray[i].StartIP, _ = stringToUint32(tempArr[1])
		RegionArray[i].StopIP, _ = stringToUint32(tempArr[2])
		RegionArray[i].Country = tempArr[3]
		RegionArray[i].Province = tempArr[4]
		RegionArray[i].City = tempArr[5]
		RegionArray[i].District = tempArr[6]
		RegionArray[i].Isp = tempArr[7]
		RegionArray[i].Type = tempArr[8]
		RegionArray[i].Desc = tempArr[9]
	}

}

func GetRegionModel( id int )(RegionModel){
	if id >= len(RegionArray){
		id = len(RegionArray)
	}
	return RegionArray[id]
}

func FindRegionModel( ip string )(RegionModel){
	regionModel , err := binarySearch(RegionArray, ip2Long(ip))
	if err == nil {
		return regionModel
	}
	return regionModel
}

func ip2Long(ip string) uint32 {
	NetIP := net.ParseIP(ip)
	if NetIP == nil {
		return 0
	}
	return binary.BigEndian.Uint32(NetIP.To4())
}


func binarySearch( list []RegionModel, ip uint32 )(RegionModel, error){

	var list_len = uint32( len(list) )

	var low uint32 = 0;

	var high uint32 = list_len - 1

	for ;  ((low <= high) && (low <= list_len - 1) && (high <= list_len - 1)) ; {

		var middle uint32 = (high + low) >> 1;

		if ip >= list[middle].StartIP && ip <= list[middle].StopIP {

			return list[middle] , nil

		}else if (ip < list[middle].StartIP ){

			high = middle - 1;

		}else {

			low = middle + 1;
		}

	}

	return RegionModel{}, errors.New( "binarySearch nil" )
}


func stringToUint32(str string)(uint32, error){
	num , err := stringToUint64(str)
	return uint32(num), err
}

func stringToUint64(str string)(uint64, error){
	return strconv.ParseUint(str, 10, 32)
}
