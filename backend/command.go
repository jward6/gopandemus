package backend

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"strings"
)

var errInvalidCommand = errors.New("Invalid Command Type.")

type execute func(c context.Context) (result string, err error)

func errorExecute(c context.Context) (result string, err error) {
	return "", errInvalidCommand
}

func createCommandHandle(input string) execute {
	s := strings.Split(input, " ")
	cmd := s[0]
	args := s[1:]

	switch cmd {
	case "draw":
		//player := args[0]
		//numOfCards := args[1]
		return func(c context.Context) (string, error) {
			return "Draw Card Command", nil
		}
	case "infect":
		return func(c context.Context) (string, error) {
			return "Infect Command", nil
		}
	case "treat":
		return func(c context.Context) (string, error) {
			return "Treat Disease Command", nil
		}
	case "place":
		player := args[0]
		city := args[1]
		return func(c context.Context) (string, error) {
			return fmt.Sprintf("Place Player, %s, at %s", player, city), nil
		}
	}

	return errorExecute
}
