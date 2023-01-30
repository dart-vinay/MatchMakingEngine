package main

// import "database/sql"
// import _ "github.com/go-sql-driver/mysql"

import (
	"fmt"
	"sync"
	"time"
)

var Players []Player
var Games []playerGameInstance

var RequestA []gameRequest
var RequestB []gameRequest
var RequestC []gameRequest

var AuxA []gameRequest
var AuxB []gameRequest
var AuxC []gameRequest

var mutex1 = &sync.Mutex{}
var mutex2 = &sync.Mutex{}
var mutex3 = &sync.Mutex{}

var Threshhold float64
var id int64
var global chan pair

var Instance int64
func main() {
	
	fmt.Println("Welcome to the Match Maker!")
	generatePlayers()
	generateGames()
	global= make(chan pair, 1000000);
	Threshhold=150
	id=0
	Instance=0
	go startMatchingEngine()
	handleRequests()
	

	time.Sleep(10*time.Minute)
}
