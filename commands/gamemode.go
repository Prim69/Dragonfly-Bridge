package commands

import (
	"VBridge/session"
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"strings"
)

type Gamemode struct {
	GameMode string
	Player []cmd.Target `optional:""`
}

func (t Gamemode) Run(source cmd.Source, output *cmd.Output){
	p := source.(*player.Player)
	if !session.Get(p).HasFlag(session.Admin) {
		p.Message(NoPermission)
		return
	}
	var gm world.GameMode
	switch strings.ToLower(t.GameMode) {
		case "survival", "0", "s":
			gm = world.GameModeSurvival{}
		case "creative", "1", "c":
			gm = world.GameModeCreative{}
		case "adventure", "2", "a":
			gm = world.GameModeAdventure{}
		case "spectator", "3", "sp":
			gm = world.GameModeSpectator{}
		default:
			output.Error("§cInvalid Gamemode!")
			return
	}

	if len(t.Player) > 0 {
		if target, ok := t.Player[0].(*player.Player); ok {
			target.SetGameMode(gm)
			target.Message("§b" + p.Name() + "§7 has set your gamemode to §b" + t.GameMode + "!")
			p.Message(fmt.Sprintf("§7You have set §b%s's §7gamemode to §b%s!", target.Name(), t.GameMode))
			return
		}
	}

	p.SetGameMode(gm)
	output.Print("§7Your gamemode has been set to §b" + t.GameMode + "!")
}