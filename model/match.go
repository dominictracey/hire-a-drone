package model

type Match struct {
	Id         string
	HomeTeam   string
	VisitTeam  string
	HomeScore  int
	VisitScore int
}

func newMatch() *Match {
	var i = new(Match)

	i.Id = "a.assign(metadata.InstanceID)"
	i.HomeTeam = "New Zealand"
	i.VisitTeam = "Australia"
	i.HomeScore = 45
	i.VisitScore = 23

	return i
}
