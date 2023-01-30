
package main

import (
	"fmt"
	"time"
)

func returnGame(gameID string) playerGameInstance{
	fmt.Println(Games)
	for _, game := range Games{
		if game.gameID==gameID{
			return game
		}
	}
	emptygame:=playerGameInstance{
		gameID			: "nil",
		skills			: []string{},
		gameTime		: 0*time.Second,
		minPlayer		: 0,
		maxPlayer		: 0,
	}
	return emptygame
}

func returnPlayer(playerID string) Player{
	for _, player := range Players{
		if player.UserID==playerID{
			return player
		}
	}
	emptyPlayer:=Player{
		UserID			: "",
		MailID			: "",
		RatingVector	: []int64{},
	}
	return emptyPlayer
}