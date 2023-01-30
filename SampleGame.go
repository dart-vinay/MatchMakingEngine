package main

import(
	"time"
)

func generateGames(){
	Games = []playerGameInstance{
		playerGameInstance{
			gameID			: "g1",
			skills			: []string{"A","B"},
			gameTime		: 60*time.Second,
			minPlayer		: 2,
			maxPlayer		: 2,
		},
		playerGameInstance{
			gameID			: "g2",
			skills			: []string{"D","E"},
			gameTime		: 1200*time.Second,
			minPlayer		: 2,
			maxPlayer		: 2,
		},
		playerGameInstance{
			gameID			: "g3",
			skills			: []string{"B","C"},
			gameTime		: 400*time.Second,
			minPlayer		: 2,
			maxPlayer		: 2,
		},
	}
}

// GameA:=playerGameInstance{
// 	gameID			: "g1"
// 	skills			: []string{"A","B"}
// 	gameTime		: 60*time.Second
// 	minPlayer		: 2
// 	maxPLayer		: 2
// }
// GameB:=playerGameInstance{
// 	gameID			: "g2"
// 	skills			: []string{"D","E"}
// 	gameTime		: 1200*time.Second
// 	minPlayer		: 2
// 	maxPLayer		: 2
// }
// GameC:=playerGameInstance{
// 	gameID			: "g3"
// 	skills			: []string{"B","C"}
// 	gameTime		: 400*time.Second
// 	minPlayer		: 2
// 	maxPLayer		: 2
// }
