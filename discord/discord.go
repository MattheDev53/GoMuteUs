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
var Members []*discordgo.Member

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
	fmt.Println("Optimizing User List")
	for i := 0; i < len(Members); i++ {
		if UserOnline(i) {
			SetUserState(i, Alive)
		} else {
			SetUserState(i, Offline)
		}
	}
}
func OptimizeUsers() {
	fmt.Println("Optimizing User List")
	for i := 0; i < len(Members); i++ {
		if !UserOnline(i) {
			SetUserState(i, Offline)
		}
	}
}

func SetUserState(id int, s State) {
	if id == -1 {
		ListNotInState(Offline)
		fmt.Printf("Select a User |> ")
		fmt.Scanln(&id)
	}
	MemberState[id] = s
}

func ClearDead() {
	UnmuteState(Dead)
	for i := 0; i < len(Members); i++ {
		if MemberState[i] == Dead {
			SetUserState(i, Alive)
		}
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

func MuteState(s State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] != Offline && (MemberState[i] == s || s == -1) {
			MuteUser(i)
		}
	}
}

func UnmuteUser(id int) {
	if id == -1 {
		ListNotInState(Offline)
		fmt.Printf("Select a User |> ")
		fmt.Scanln(&id)
	}
	Session.GuildMemberMute(Conf.GuildID, Members[id].User.ID, false)
}

func UnmuteState(s State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] != Offline && (MemberState[i] == s || s == -1) {
			MuteUser(i)
		}
	}
}

func UserOnline(id int) bool {
	err := Session.GuildMemberDeafen(Conf.GuildID, Members[id].User.ID, false)
	if err != nil {
		return false
	}
	return true
}
