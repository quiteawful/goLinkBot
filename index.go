package main

import (
	_ "database/sql"
	"fmt"
	"strings"
)

func help() {
	ctxIrc.WriteToChannel("Datenkrake. I nomnom links.")
	ctxIrc.WriteToChannel("!search [--tag|-t] <query>")
	ctxIrc.WriteToChannel("!linkinfo <id>")
	ctxIrc.WriteToChannel("!addtag <id> <tag> [<tag> [..]]")
}

func search(query string, tagonly bool) {
	links, err := LinksSearch(query, tagonly)
	if err != nil {
		fmt.Printf("index.Search: %s\n", err.Error())
		return
	}

	if len(links) == 0 {
		ctxIrc.WriteToChannel("Sorry, no links found :/")
		return
	}
	for _, l := range links {
		ctxIrc.WriteToChannel(fmt.Sprintf("[%s] Id: %d %s: %s", l.Tstamp.Format("02.01.2006 15:04:05"), l.Id, l.User, l.Url))
	}
	return
}

func searchtags(query string) {

}

func linkinfo(id int) bool {
	l := &TblLinks{}
	err := l.Open(id)
	if err != nil {
		fmt.Printf("index.linkinfo: Fehler beim Öffnen von Orm Objekt. %s\n", err.Error())
		ctxIrc.WriteToChannel("Sorry, no info found :(")
		return false
	}

	ctxIrc.WriteToChannel(fmt.Sprintf("Id: %d", l.Id))
	ctxIrc.WriteToChannel(fmt.Sprintf("Link: %s", l.Url))
	ctxIrc.WriteToChannel(fmt.Sprintf("User: %s", l.User))
	ctxIrc.WriteToChannel(fmt.Sprintf("Timestamp: %s", l.Tstamp.Format("02.01.2006 15:04:05")))
	ctxIrc.WriteToChannel(fmt.Sprintf("Original Message: %s", l.Post))
	ctxIrc.WriteToChannel(fmt.Sprintf("Tags: %s", strings.Join(l.GetTags(), ", ")))

	return true
}

func addTag(id int, tag, user string) bool {
	fmt.Printf("LinkId: %d, Tag: %s, User: %s\n", id, tag, user)
	t := &TblTags{}
	t.Tag = strings.TrimSpace(tag)
	t.Save()

	tht := &TblHasTags{}
	tht.LinkId = id
	tht.TagId = t.Id
	tht.User = user

	tht.Save()
	return true
}
