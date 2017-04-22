package main

type Pilot struct {
	ID       string
	Name     string
	Licensed bool
	Address  string
	Phone    string
}

func NewPilot() *Pilot {
	var i = new(Pilot)

	//i.Id = "a.assign(metadata.InstanceID)"
	i.Name = "Fred Smith"
	i.Licensed = true
	i.Address = "98 Wallaby Way, Sydney AUS"
	i.Phone = "+23 0903 91203"

	return i
}
