package main

import (
	"encoding/json"
	"os"
)

const stateFile = "state.json"

type State map[string]bool

func LoadState() (State, error) {
	state := make(State)
	
	data, err := os.ReadFile(stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return state, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &state)
	return state, err
}

func SaveState(state State) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}
