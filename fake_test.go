package main

import (
	"encoding/xml"
	"testing"
)

func TestInit(t *testing.T) {
	var testFakeBus FakeBus

	if len(testFakeBus.BusInstances) != 0 {
		t.Fatalf("expected empty fakeBus.BusInstances, got %v", len(testFakeBus.BusInstances))
	}

	testFakeBus.Init()

	if len(testFakeBus.BusInstances) != 2 {
		t.Fatalf("expected len(fakeBus.BusInstances) is 2, got %v", len(testFakeBus.BusInstances))
	}
}

func TestGetBusRoute(t *testing.T) {
	var testFakeBus FakeBus
	Expected := BusRoute{xml.Name{}, 2, "수원",
		"201320974", "2018-1", "테스트용시내버스", 0}

	testFakeBus.Init()
	ret := testFakeBus.GetBusRoute()
	if ret != Expected {
		t.Fatalf("expected BusRoute: %v, got %v", Expected, ret)
	}
}

func TestGetBusRouteInfo(t *testing.T) {
	var testFakeBus FakeBus
	Expected := BusRouteInfoItem{xml.Name{}, 2, "00:00", "00:00",
		"03129", "202000004", "아주대학교입구", "수원", "201320974",
		"2018-1", "테스트용시내버스", "04238", "203000067",
		"아주대학교입구", "00:00", "00:00"}

	testFakeBus.Init()
	ret := testFakeBus.GetBusRouteInfo()
	if ret != Expected {
		t.Fatalf("expected BusRouteInfo: %v, got %v", Expected, ret)
	}
}

func TestGetBusRouteStationList(t *testing.T) {
	var testFakeBus FakeBus
	testFakeBus.Init()
	if ret := testFakeBus.GetBusRouteStationList(); len(ret) != 8 {
		t.Fatalf("expected length is %v, got %v", 8, len(ret))
	}
}

func TestFakeBusGetCurrentBusLocation(t *testing.T) {
	var testFakeBus FakeBus
	testFakeBus.Init()

	if ret := testFakeBus.GetCurrentBusLocation(); len(ret) != 2 {
		t.Fatalf("expected length is %v, got %v", 2, len(ret))
	}
}

func TestIsInBusList(t *testing.T) {
	var testFakeBus FakeBus
	testFakeBus.Init()

	tt := []struct {
		value  string
		result bool
	}{
		{"202000004", true},
		{"200000000", false},
	}

	for _, tc := range tt {
		if ret := testFakeBus.IsInBusList(tc.value); ret != tc.result {
			t.Fatalf("expected value: %v, got %v", tc.result, ret)
		}
	}

}

func TestGetBusArrival(t *testing.T) {
	var testFakeBus FakeBus
	Expected := BusArrival{
		xml.Name{},
		3, 7, -1, -1,
		"경기96희0820", "경기93범1004",
		3, 7, 23, 17,
		"201320974", "2018-1", "테스트용시내버스", 8,
	}
	testFakeBus.Init()
	ret := testFakeBus.GetBusArrival("202000038")
	if ret != Expected {
		t.Fatalf("expected BusArrival: %v, got %v", Expected, ret)
	}
}

func TestAbs(t *testing.T) {
	tt := []struct {
		value  int
		result int
	}{
		{1, 1},
		{-1, 1},
		{0, 0},
	}

	for _, tc := range tt {
		if ret := abs(tc.value); ret != tc.result {
			t.Fatalf("expected value: %v, got %v", tc.result, ret)
		}
	}
}
