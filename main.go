package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

type CartItem struct {
	ItemID int `json:"item_id"`
	Qty    int `json:"qty"`
	Sku    int `json:"sku"`
	Weight int `json:"weight"`
	Height int `json:"height"`
	Length int `json:"length"`
	Width  int `json:"width"`
}

type Warehouse struct {
	FirstMile string `json:"first_mile"`
	ID        int    `json:"id"`
}

type Unit struct {
	ReqID      int         `json:"req_id"`
	SellerID   int         `json:"seller_id"`
	CityCodeTo int64       `json:"city_code_to"`
	CartItem   CartItem    `json:"cart_item"`
	Warehouses []Warehouse `json:"warehouses"`
	UserID     int         `json:"user_id"`
}

type Request struct {
	Units []Unit `json:"units"`
}

func randomIdDropOffSeller() int {
	dropOffSellers := []int{
		3231761093,
		1055235757,
		1047545381,
		1047545381,
		1047545381,
		3240281476,
		3229314690,
		3593512463,
		3768978177,
		3609642704,
		4966933857,
		1059365900,
		3530297616,
		3599849138,
		3593512463,
		3611034752,
		3609642704,
		3659969906,
		3619097265,
		3687211344,
		3703409599,
	}
	return dropOffSellers[rand.Intn(len(dropOffSellers))]
}

func randomIdDropOffWarehouse() int {
	dropOffWarehouses := []int{
		450000179068,
		450000163448,
		450000161452,
		450000163433,
		450000230841,
		450000181172,
		450000198021,
		450002786011,
		450003302082,
		450002786020,
		450000182706,
		410000157716,
		410000165702,
		430000639310,
		430000640154,
		430000640563,
		430000641191,
		430000642551,
		430000643238,
		430000643246,
		430000644041,
		430000644042,
		430000647144,
		430000647148,
		430000647601,
		430000647723,
		430000649740,
		430000651120,
		430000651140,
		430000651142,
		430000651542,
		430000651673,
		430000653030,
		430000658493,
		430000662587,
		450000127721,
		450000139913,
		450000144194,
		450000144543,
		450000161452,
		450000161573,
		450000162720,
		450000162722,
		450000162753,
		450000163433,
		450000163448,
		450000166685,
		450000167142,
		450000169170,
		450000171536,
		450000178550,
		450000178555,
		450000178924,
		450000179068,
		450000180501,
		450000180502,
		450000180504,
		450000180509,
	}
	return dropOffWarehouses[rand.Intn(len(dropOffWarehouses))]
}

func randomCartItem() CartItem {
	cartItems := []CartItem{
		{
			ItemID: 1,
			Qty:    1,
			Sku:    1,
			Weight: 1,
			Height: 1,
			Length: 1,
			Width:  1,
		},
		{
			ItemID: 2,
			Qty:    1,
			Sku:    1,
			Weight: 1,
			Height: 2,
			Length: 2,
			Width:  2,
		},
		{
			ItemID: 3,
			Qty:    1,
			Sku:    1,
			Weight: 1,
			Height: 3,
			Length: 3,
			Width:  3,
		},
	}
	return cartItems[rand.Intn(len(cartItems))]
}

func randomWarehouses() []Warehouse {
	var warehouses []Warehouse
	for warehouseCnt := 0; warehouseCnt < 2; warehouseCnt++ {
		warehouse := Warehouse{
			FirstMile: "DROPOFF",
			ID:        randomIdDropOffWarehouse(),
		}
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func randomUnits() []Unit {
	var units []Unit
	for unitCnt := 0; unitCnt < 5; unitCnt++ {
		unit := Unit{
			ReqID:      unitCnt,
			SellerID:   randomIdDropOffSeller(),
			CityCodeTo: 917480698071000000,
			CartItem:   randomCartItem(),
			Warehouses: randomWarehouses(),
			UserID:     1,
		}
		units = append(units, unit)
	}
	return units
}

func main() {
	rand.Seed(time.Now().Unix())

	var requests []Request
	for requestCnt := 0; requestCnt < 1000; requestCnt++ {
		request := Request{
			Units: randomUnits(),
		}
		requests = append(requests, request)
	}

	data, _ := json.Marshal(requests)
	f, _ := os.Create("ammo.json")
	f.Write(data)
}
