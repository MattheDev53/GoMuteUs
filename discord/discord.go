package discord

import (
	"github.com/bwmarrin/discordgo"
	cfg "github.com/MattheDev53/GoMuteUs/config"

	"fmt"
)

type State int

const (
	Alive State = iota
	Dead
	Offline
)

var Session     *discordgo.Session
var MemberState [99]State
var Conf        cfg.Config

var	Members, err = Session.GuildMembers(Conf.GuildID, "0", 99)

func ListAllMembers() {
	for i := 0; i < len(Members); i++ {
		fmt.Printf("%02d: %32s [%s]\n", i, Members[i].User.Username, StateName(MemberState[i]))
	}
}
func ListInState(state State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] == state {
			fmt.Printf("%02d: %32s [%s]\n", i, Members[i].User.Username, StateName(MemberState[i]))
		}
	}
}
func ListNotInState(state State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] != state {
			fmt.Printf("%02d: %32s [%s]\n", i, Members[i].User.Username, StateName(MemberState[i]))
		}
	}
}

func InitializeUserState() {
	for i := 0; i < len(Members); i++ {
		MemberState[i] = Alive
	}
}

func StateName(s State) string {
	switch s {
	case Alive:
		return "LIVE"
	case Dead:
		return "DEAD"
	case Offline:
		return "AWAY"
	}
	return "UNKN"
}

func MuteUser(id int) {
	if id == -1 {
		ListNotInState(Offline)
		fmt.Printf("Select a User |> ")
		fmt.Scanln(&id)
	}
	Session.GuildMemberMute(Conf.GuildID, Members[id].User.ID, true)
}
