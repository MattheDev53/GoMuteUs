# GoMuteUs

Goofy little rewrite of [AMuteUs](https://github.com/MattheDev53/AMuteUs)

## The Problem

Have you ever played the hit game Among Us with your friends,
but some of your friends don't have consoles/PCs and only have phones so you gotta
individually mute each of them in the server you all are in?

Then you wrote a script in python and then everything was great... except for the fact that

1. It crashes **constantly**
2. It's pretty slow due to just using the raw rest API and wait timers
3. The config file kept getting looked at by git
4. You get the point, it has a lot of problems

So the only real solutions at this point are to either
dust off my rusty python knowledge
or
rewrite the whole thing from scratch in a language that I'm more comfortable using

## The Solution

Rewrite AMuteUs from scratch in [Go](https://go.dev),
fully utilizing [Nix](https://nixos.org)'s potential,
and using [an actual Discord API Library](https://github.com/bwmarrin/discordgo)
because I don't want to use the Rest API again.

So far, it's alright. It isn't fully in a state where I want it, but it is functional.

## Ok How do I use it tho

### Building

```sh
nix build
```

### Running

```sh
nix run github:MattheDev53/GoMuteUs
```

OR

```sh
git clone git@github.com:MattheDev53/GoMuteUs.git
cd GoMuteUs
nix run
```

### Windows???

Use [WSL](https://learn.microsoft.com/en-us/windows/wsl/install)
or
compile it yourself.
Leave a PR if you've got a better idea

### Command Reference

The commands are all numbers so one can enter them on just a numpad.
If you want text commands, send patches I guess.

```
0  - Mute Alive Players
00 - Mute Non-Offline Players
01 - Mute A Specific Player

1  - Unmute Alive Players
10 - Unmute Non-Offline Players
11 - Unmute A Specific Player

7  - List Offline Players
70 - Refresh the list of Offline Players
71 - Add a Player to the list of Offline Players

8  - List Non-Offline Players
80 - I don't actually know
81 - Add a Player to the list of Alive Players

9  - List Dead Players
90 - Clear the list of Dead Players
91 - Add a Player to the list of Dead Players

+  - End Command (used to run multiple commands)
```
