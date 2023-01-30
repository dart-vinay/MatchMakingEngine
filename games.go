package main

import(
	"time"
)

type Game interface{
	addSkill(skill string)
	deleteSkill(pos int)
	getGameTime() time.Duration
	modifyGameTime(time time.Duration)
}

type PlayerGame interface{
	Game
	// getMinPlayers() int64
	// getMaxPlayers() int64
	// getGameID() string
	// getGameSkills() []string
}

type playerGameInstance struct{
	gameID string
	skills []string
	gameTime time.Duration
	minPlayer int64
	maxPlayer int64
}

type TeamGame interface{
	Game
	// getMinPlayers() int64
	// getMaxPlayers() int64
	// getGameID() string
	// getGameSkills() []string
	// getTeamsAllowed() int64
}

type teamGameInstance struct{
	gameID string
	skills []string
	gameTime time.Duration
	minPlayerPerTeam int64
	maxPlayerPerTeam int64
	teamsAllowed int64
}

func (instance *teamGameInstance) addSkill(skill string){
	instance.skills  = append(instance.skills,skill)
}

func (instance *playerGameInstance) addSkill(skill string){
	instance.skills = append(instance.skills,skill)
}

func (instance *teamGameInstance) deleteSkill(pos int){
	instance.skills[pos] = instance.skills[len(instance.skills)-1]
	instance.skills[len(instance.skills)-1]=""
	instance.skills = instance.skills[:len(instance.skills)-1]
}

func (instance *playerGameInstance) deleteSkill(pos int){
	instance.skills[pos] = instance.skills[len(instance.skills)-1]
	instance.skills[len(instance.skills)-1]=""
	instance.skills = instance.skills[:len(instance.skills)-1]
}

func (instance *teamGameInstance) getGameTime() time.Duration{
	return instance.gameTime
}

func (instance *playerGameInstance) getGameTime() time.Duration{
	return instance.gameTime
}

func (instance *teamGameInstance) modifyGameTime(time time.Duration) {
	instance.gameTime = time
}

func (instance *playerGameInstance) modifyGameTime(time time.Duration){
	instance.gameTime = time
}
