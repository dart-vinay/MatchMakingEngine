package main

import(

)
type playerInterface interface{
	getID() string
	getRatingVector() []int64
}

type Player struct {
	UserID string `json:"UserID"`
	MailID string `json:"mailID"`
	RatingVector []int64 `json:"rating"`
}
