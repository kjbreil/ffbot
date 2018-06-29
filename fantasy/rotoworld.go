package fantasy

// PlayerObject contains the base information about a player from rotoworld
// types are subject to change as I figure out what is actually being returned
// since original code is in JS which doens't give a single clue at what is
// contained
type PlayerObject struct {
	url          string
	name         string
	position     string
	number       int
	team         string
	age          int
	heightWeight string
	college      string
	drafted      string
	latestNews   string
	latestImpact string
	playerPhoto  string
}

// GetPlayer gets a player from rotoworld
func GetPlayer(playerName string) (p PlayerObject) {
	return
}
