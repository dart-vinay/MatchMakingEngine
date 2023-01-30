
package main

import (
	"fmt"
	"reflect"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"github.com/bradfitz/slice"
	"math"
)




type pair struct{
	first gameRequest
	second gameRequest
}



func flush(){
	mutex1.Lock()
	AuxA=AuxA[:0]
	mutex1.Unlock()

	mutex2.Lock()
	AuxB=AuxB[:0]
	mutex2.Unlock()

	mutex3.Lock()
	AuxC=AuxC[:0]
	mutex3.Unlock()
}

func uptakeQueues(){
	RequestA = append(RequestA, AuxA...)
	RequestB = append(RequestB, AuxB...)
	RequestC = append(RequestC, AuxC...)
	flush()
}

func notFound(req gameRequest, provided []gameRequest) bool{

	for _,obj := range provided{
		if(reflect.DeepEqual(obj,req)){
			return true
		}
	}
	return false
}

func generateID() int64{
	id++
	return id
}



func ProvideChannels(pairs []pair){
	fmt.Println("Providing Meeting")
	fmt.Println(pairs)
	for _ , sl := range pairs{
		global<-sl
	}
}

func check(req gameRequest) MeetingDetails{
	for{
		select{
		case result:=<-global:
			fmt.Println("Listening to Global")
			fmt.Println(req)
			fmt.Println(result.first, result.second)
			if (reflect.DeepEqual(result.first,req) || reflect.DeepEqual(result.second,req)){
				fmt.Println("MAtch! MAtch! Match!")
				first:=result.first
				// second:=result.second
				meetingID:=generateID()
				meetingObj:=MeetingDetails{
					MeetingID : meetingID,
					GameID    : first.game.gameID,
					Message   : "Let's Play",
				}
				return meetingObj
			}
		}
	}
}


func runServiceForA(){
	length:=len(RequestA)
	fmt.Println("Running Service for A")
	Instance++;
	// fmt.Println(RequestA)
	// fmt.Println(reflect.TypeOf(RequestA[0].channelToWait))
	 // this is kind of a burden
	// var pairs []pair

	if(length>=2){
		fmt.Println(RequestA)
		fmt.Println("Request A in Action")
		var matched []gameRequest
		for i:=1;i<length;i++ {
			
			// fmt.Println("Entering the requst array")
			skillDifference := math.Abs(float64(RequestA[i].player.RatingVector[0]-RequestA[i-1].player.RatingVector[0]))
			if(skillDifference<Threshhold){
				// fmt.Println("Match Occured")
				if(!notFound(RequestA[i], matched)){
					fmt.Println("Using Instance ", Instance)
					fmt.Println(matched)
					fmt.Println("Woauu!!")
					var p pair
					p=pair{first:RequestA[i], second:RequestA[i-1]}
					global<-p
					// pairs = append(pairs, p)
					matched=append(matched, RequestA[i])
					matched=append(matched, RequestA[i-1])
				}
			}
		}
		RequestDummy:=[]gameRequest{};
		for i:=0;i<length;i++{
			flag:=0
			for _, val := range matched{
				if(reflect.DeepEqual(val,RequestA[i])){
					flag=1;
				}
			}
			if(flag==0){
				RequestDummy=append(RequestDummy,RequestA[i])
			}
		}
		RequestA=RequestDummy
		// ProvideChannels(pairs)
	}
}


func runEngine(){
	fmt.Println("Inside run Engine")
	uptakeQueues()
	// fmt.Println(RequestA)
	
	//SortAllTheSlices
	slice.Sort(RequestA[:], func(i, j int) bool {
		// skillCOunt = 
		return RequestA[i].player.RatingVector[0] < RequestA[j].player.RatingVector[0]
	})
	slice.Sort(RequestB[:], func(i, j int) bool {
		return RequestB[i].player.RatingVector[0] < RequestB[j].player.RatingVector[0]
	})
	slice.Sort(RequestC[:], func(i, j int) bool {
		return RequestC[i].player.RatingVector[0] < RequestC[j].player.RatingVector[0]
	})

	go runServiceForA()
	// go runServiceForB()
	// go runServiceForC()

}

func startMatchingEngine(){
	fmt.Println("Inside Start Enginer")
	for{
		time.Sleep(5*time.Second)
		runEngine()
		
	}
}

func updateAuxQueue(requestObj *gameRequest){
	fmt.Println("Inside Update Aux Queues ")
	gameID:=requestObj.game.gameID
	if(gameID=="g1"){
		mutex1.Lock()
		AuxA = append(AuxA,*requestObj)
		mutex1.Unlock()
	}
	if(gameID=="g2"){
		mutex2.Lock()
		AuxB = append(AuxB,*requestObj)
		mutex2.Unlock()
	}
	if(gameID=="g3"){
		mutex3.Lock()
		AuxC = append(AuxC,*requestObj)
		mutex3.Unlock()
	}
}

func fullfillRequest(playerID string, gameID string) *gameRequest{

	channelToFulFill := make(chan MeetingDetails)

	player:=returnPlayer(playerID)
	game:=returnGame(gameID)
	
	if(game.gameID=="" || player.UserID==""){
		requestObj := gameRequest{
			player 	:       player,
			game    :       game,
			channelToWait   :  channelToFulFill,
		}
		return &requestObj;
	}
	timeout:=game.gameTime/10
	if (timeout<25*time.Second){
		timeout=25*time.Second
	}
	
	ctx,_:=context.WithTimeout(context.Background(), timeout)
	requestObj := gameRequest{
		player 	:       player,
		game    :       game,
		ctx     :       ctx,
		channelToWait   :  channelToFulFill,
	}
	go func(){
		resultObject:=check(requestObj)
		channelToFulFill<-resultObject
	}()

	updateAuxQueue(&requestObj)

	return &requestObj;
}


func handleGameRequest(w http.ResponseWriter, r *http.Request){
	fmt.Println("Entering Handle Game Request")

	vars := mux.Vars(r)
	playerID := vars["playerID"]
	gameID := vars["gameID"]
	fmt.Println(playerID, gameID)

	go func(){
		response:=fullfillRequest(playerID, gameID)

		if(response.game.gameID=="" || response.player.UserID==""){
			details:=MeetingDetails{
				MeetingID : -1,
				GameID    : "",
				Message   : "Invalid User or Game",
			}
			fmt.Println(" ---------------- CANT FOUND A MATCH------------------- ")
			fmt.Println(details)
			json.NewEncoder(w).Encode(details)
		} else{
			select{
			case <-response.ctx.Done():
				details:=MeetingDetails{
					MeetingID : -1,
					GameID    : "",
					Message   : "TimeOut!",
				}
				fmt.Println(" ---------------- CANT FOUND A MATCH------------------- ")
				fmt.Println(details)
				json.NewEncoder(w).Encode(details)
			case answer:=<-response.channelToWait:
				fmt.Println(" ---------------- FOUND A MATCH ------------------- ")
				fmt.Println(answer)
				json.NewEncoder(w).Encode(answer)
			}
		}
	}()
}

// func testFunc(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	playerID := vars["playerID"]
// 	gameID := vars["gameID"]
// 	fmt.Println(playerID, gameID)
// }
