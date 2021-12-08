// Copyright (c) Denis Dubo Chevalier
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	// Get env
	usersStr, exists := os.LookupEnv("ROULETTE_LIST")
	var users []string
	if !exists {
		log.Fatal("ROULETTE_LIST environment variable is not set")
	}
	users = strings.Split(usersStr, ",")
	if len(users) < 2 {
		log.Fatal("Need at least two users in ROULETTE_LIST environment variable to work")
	}

	// Parse arguments
	args := os.Args[1:]
	if len(args) != 0 {
		switch args[0] {
		case "-a":
			for _, arg := range args[1:] {
				for i := range users {
					if strings.ToLower(users[i]) == strings.ToLower(arg) {
						users = remove(users, i)
						break
					}
				}
			}
		case "-h":
			fmt.Println("Usage: roulette\n\nroulette draws two names out of a user list.\n\nEnvironment variable" +
				"ROULETTE_USERS must be set with a comma separated list of users.\n\nParameters:\n-h\tPrints this" +
				"help\n-a\tRemoves a set of user names from the list." +
				"Example: roulette -a foo bar")
			os.Exit(0)
		default:
			log.Fatal("Unknown flag: " + args[0])
		}
	}
	log.Print("choosing from: ", users)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	selected := []string{
		users[r.Intn(len(users))],
		users[r.Intn(len(users))],
	}
	for selected[1] == selected[0] {
		selected[1] = users[r.Intn(len(users))]
	}
	fmt.Println("selected: ", selected)
}
