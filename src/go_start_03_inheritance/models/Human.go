package models

/* This is structure of human */
type Human struct {
	Width  int
	Height int
	Gender string
	Name   string
}

type Energy struct {
	Intensity int    `intensity`
	Manifest  string `manifest`
}

func (e Energy) Explode() string {
	return e.Manifest
}

type Magic struct {
	Energy
	Level int
}
