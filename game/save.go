package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fatih/color"
)

const SaveFile = "adventure-save.json"

type Save struct {
	Battles     int `json:"battles,omitempty"`
	BattlesWon  int `json:"battles_won,omitempty"`
	BattlesLost int `json:"battles_lost,omitempty"`

	HeroName string `json:"hero_name,omitempty"`
}

func LoadSave() *Save {
	ret := &Save{}

	data, err := ioutil.ReadFile(SaveFile)
	if err != nil {
		return ret
	}

	err = json.Unmarshal(data, ret)
	if err != nil {
		color.Red("Could not read save file: %s\n", err.Error())
	}
	return ret
}

func WriteSave(s *Save) {
	b, err := json.Marshal(s)
	if err != nil {
		color.Red("Could not encode save file: %s\n", err.Error())
	}
	ioutil.WriteFile(SaveFile, b, 0777)
}
