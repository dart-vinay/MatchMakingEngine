
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)
type MeetingDetails struct{
	MeetingID int64 `json:"meetingID"`
	GameID string `json:"gameID"`
	Message string `json:"message"`
}

// type solution struct{
// 	MeetID string `json:"meetID"`
// }

type gameRequest struct {
	player Player
	game playerGameInstance
	ctx context.Context
	channelToWait <-chan MeetingDetails
}

func handleRequests() {
	fmt.Println("Handling Requests...")
	myRouter := mux.NewRouter().StrictSlash(true)
	//Requests
	myRouter.HandleFunc("/requestGame/{playerID}/{gameID}", handleGameRequest)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
