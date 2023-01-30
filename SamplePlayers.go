package main

import(

)
func generatePlayers(){
	Players = []Player{
		Player{
			UserID			: "p1",
			MailID			: "xyz1.gmail.com",
			RatingVector	: []int64{100,120,200,130,110},
		},
		Player{
			UserID			: "p2",
			MailID			: "xyz2.gmail.com",
			RatingVector	: []int64{120,100,170,30,10},
		},
		Player{
			UserID			: "p3",
			MailID			: "xyz3.gmail.com",
			RatingVector	: []int64{140,110,20,30,270},
		},
		Player{
			UserID			: "p4",
			MailID			: "xyz4.gmail.com",
			RatingVector	: []int64{160,70,80,30,110},
		},
		Player{
			UserID			: "p5",
			MailID			: "xyz5.gmail.com",
			RatingVector	: []int64{90,20,210,190,110},
		},
	}
}

// PlayerA:=Player{
// 	UserID			: "p1"
// 	MailID			: "xyz1.gmail.com"
// 	RatingVector	: []int64{100,120,200,130,110}
// }
// PlayerB:=Player{
// 	UserID			: "p2"
// 	MailID			: "xyz2.gmail.com"
// 	RatingVector	: []int64{120,100,170,30,10}
// }
// PlayerC:=Player{
// 	UserID			: "p3"
// 	MailID			: "xyz3.gmail.com"
// 	RatingVector	: []int64{140,110,20,30,270}
// }
// PlayerD:=Player{
// 	UserID			: "p4"
// 	MailID			: "xyz4.gmail.com"
// 	RatingVector	: []int64{160,70,80,30,110}
// }
// PlayerE:=Player{
// 	UserID			: "p5"
// 	MailID			: "xyz5.gmail.com"
// 	RatingVector	: []int64{90,20,210,190,110}
// }
