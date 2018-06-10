package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

/*
	Bus Number: 2018-1
	Bus RouteID: 201320974
*/

type FakeBus struct {
	BusInstances        []FakeBusIntance
	busRoute            BusRoute
	busRouteInfoItem    BusRouteInfoItem
	busRouteStationList BusRouteStationList
}

type FakeBusIntance struct {
	currentStationNo int
	plateNo          string
}

var fakeBus FakeBus

func (fb *FakeBus) Init() {
	*fb = FakeBus{
		[]FakeBusIntance{
			FakeBusIntance{0, "경기93범1004"},
			FakeBusIntance{4, "경기96희0820"},
		},
		BusRoute{xml.Name{}, 2, "수원", "201320974", "2018-1", "테스트용시내버스", 0},
		BusRouteInfoItem{xml.Name{}, 2, "00:00", "00:00", "03129", "202000004", "아주대학교입구", "수원", "201320974", "2018-1", "테스트용시내버스", "04238", "203000067", "아주대학교입구", "00:00", "00:00"},
		BusRouteStationList{
			BusRouteStation{xml.Name{}, "", 0, "04238", "", "203000067", "아주대학교입구", 0, 0, 1, ""},
			BusRouteStation{xml.Name{}, "", 0, "04237", "", "203000066", "아주대.아주대학교병원", 0, 0, 2, ""},
			BusRouteStation{xml.Name{}, "", 0, "03124", "", "202000039", "창현고교.아주대학교.유신고교", 0, 0, 3, ""},
			BusRouteStation{xml.Name{}, "", 0, "03117", "", "202000038", "효성초등학교", 0, 0, 4, ""},
			BusRouteStation{xml.Name{}, "", 0, "03119", "", "202000032", "효성초등학교", 0, 0, 5, ""},
			BusRouteStation{xml.Name{}, "", 0, "03125", "", "202000061", "창현고교.아주대학교.유신고교", 0, 0, 6, ""},
			BusRouteStation{xml.Name{}, "", 0, "03126", "", "202000005", "아주대.아주대학교병원", 0, 0, 7, ""},
			BusRouteStation{xml.Name{}, "", 0, "03129", "", "202000004", "아주대학교입구", 0, 0, 8, ""},
		},
	}
}

func (fb *FakeBus) Run() {
	fb.Init()

	for {
		for i := 0; i < len(fb.BusInstances); i++ {
			fb.BusInstances[i].currentStationNo = (fb.BusInstances[i].currentStationNo + 1) % 8
		}

		time.Sleep(1 * time.Minute)
	}
}

func (fb *FakeBus) GetBusRoute() BusRoute {
	return fb.busRoute
}

func (fb *FakeBus) GetBusRouteInfo() BusRouteInfoItem {
	return fb.busRouteInfoItem
}

func (fb *FakeBus) GetBusRouteStationList() BusRouteStationList {
	return fb.busRouteStationList
}

func (fb *FakeBus) GetCurrentBusLocation() BusLocationList {
	busLocationList := BusLocationList{}
	for i := 0; i < len(fb.BusInstances); i++ {
		busLocationList = append(busLocationList, BusLocation{
			xml.Name{},
			0,
			-1,
			fb.BusInstances[i].plateNo,
			-1,
			fb.busRouteStationList[fb.BusInstances[i].currentStationNo].StationID,
			fb.busRouteStationList[fb.BusInstances[i].currentStationNo].StationSeq,
		})
	}

	return busLocationList
}

func (fb *FakeBus) IsInBusList(stationID string) bool {
	for _, bs := range fb.busRouteStationList {
		if bs.StationID == stationID {
			return true
		}
	}
	return false
}

func (fb *FakeBus) GetBusArrival(stationID string) BusArrival {
	busArrival := BusArrival{
		xml.Name{},
		0,
		0,
		-1,
		-1,
		"",
		"",
		0,
		0,
		-1,
		-1,
		fb.busRoute.RouteID,
		fb.busRoute.RouteName,
		fb.busRoute.RouteTypeName,
		0,
	}

	for _, bs := range fb.busRouteStationList {
		if bs.StationID == stationID {
			busArrival.StaOrder = bs.StationSeq
			break
		}
	}

	LocationNos := []int{}

	for i := 0; i < len(fb.BusInstances); i++ {
		fmt.Println("current:", fb.busRouteStationList[fb.BusInstances[i].currentStationNo].StationSeq)
		LocationNos = append(LocationNos, busArrival.StaOrder-fb.busRouteStationList[fb.BusInstances[i].currentStationNo].StationSeq)
		if LocationNos[i] < 0 {
			LocationNos[i] = len(fb.busRouteStationList) - LocationNos[i]
			LocationNos[i]--
		} else if LocationNos[i] == 0 {
			LocationNos[i] = 1
		}

		fmt.Println("result:", LocationNos[i])
	}
	fmt.Println("")

	firstIndex := 0

	if LocationNos[0] > LocationNos[1] {
		firstIndex = 1
	}

	busArrival.LocationNo1 = LocationNos[firstIndex]
	busArrival.LocationNo2 = LocationNos[abs(firstIndex-1)]

	busArrival.PlateNo1 = fb.BusInstances[firstIndex].plateNo
	busArrival.PlateNo2 = fb.BusInstances[abs(firstIndex-1)].plateNo

	busArrival.PredictTime1 = LocationNos[firstIndex] * 1
	busArrival.PredictTime2 = LocationNos[abs(firstIndex-1)] * 1

	return busArrival
}

func abs(value int) int {
	if value < 0 {
		return value * (-1)
	}
	return value

}
