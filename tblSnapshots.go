package main

import (
	_ "fmt"
	"time"
)

type Snapshots struct {
	Id     int
	LinkId int
	Tstamp time.Time
	User   string
	Post   string
	Src    []byte
}

func (s *Snapshots) Save() (int, error) {

}

func (s *Snapshots) Open(Id int) error {

}

func SearchSnapshot(LinkId int) Snapshots {

}
