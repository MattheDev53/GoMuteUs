package discord

import (
	"github.com/bwmarrin/discordgo"
	cfg "github.com/MattheDev53/GoMuteUs/config"
	lg "github.com/charmbracelet/lipgloss"
	catppuccin "github.com/catppuccin/go"

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
var Members     []*discordgo.Member

const UserStr string = "[%4s][%02d]: %-32s\n"

func ListMembers() {
	for i := 0; i < len(Members); i++ {
		fmt.Printf(UserStr, StateName(MemberState[i]), i, Members[i].User.Username)
	}
}
func ListMembersInState(state State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] == state {
			fmt.Printf(UserStr, StateName(MemberState[i]), i, Members[i].User.Username)
		}
	}
}
func ListMembersNotInState(state State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] != state {
			fmt.Printf(UserStr, StateName(MemberState[i]), i, Members[i].User.Username)
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
		if UserOnline(i) {
			SetUserState(i, MemberState[i])
		} else {
			SetUserState(i, Offline)
		}
	}
}

func SelectMember() int {
	  var id int
		ListMembers()
		fmt.Printf("Select a User |> ")
		fmt.Scanln(&id)
		return id
}
func SelectMemberInState(s State) int {
	  var id int
		ListMembersInState(s)
		fmt.Printf("Select a User |> ")
		fmt.Scanln(&id)
		return id
}
func SelectMemberNotInState(s State) int {
	  var id int
		ListMembersNotInState(s)
		fmt.Printf("Select a User |> ")
		fmt.Scanln(&id)
		return id
}

func SetUserDead(id int) {
	if id == -1 { SelectMemberInState(Alive) }
	if id == -1 { return }
	SetUserState(id, Dead)
}

func SetUserState(id int, s State) {
	if id == -1 { SelectMemberNotInState(s) }
	if id == -1 { return }
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
		return lg.NewStyle().Bold(true).Foreground(lg.Color(catppuccin.Mocha.Green().Hex)).Render("LIVE")
	case Dead:
		return lg.NewStyle().Bold(true).Foreground(lg.Color(catppuccin.Mocha.Red().Hex)).Render("DEAD")
	case Offline:
		return lg.NewStyle().Bold(true).Foreground(lg.Color(catppuccin.Mocha.Peach().Hex)).Render("AWAY")
	}
	return lg.NewStyle().Bold(true).Foreground(lg.Color(catppuccin.Mocha.Mauve().Hex)).Render("UNKN")
}

func MuteUser(id int) {
	if id == -1 { SelectMemberNotInState(Offline) }
	if id == -1 { return }
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
	if id == -1 { SelectMemberNotInState(Offline) }
	if id == -1 { return }
	Session.GuildMemberMute(Conf.GuildID, Members[id].User.ID, false)
}

func UnmuteState(s State) {
	for i := 0; i < len(Members); i++ {
		if MemberState[i] != Offline && (MemberState[i] == s || s == -1) {
			UnmuteUser(i)
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
