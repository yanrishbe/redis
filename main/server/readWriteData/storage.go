package readWriteData

import (
	"fmt"
	"github.com/yanrishbe/redis/main/server/entities"
	"regexp"
	"sort"
	"strings"
)

func AnswerToClient (cmd chan entities.Command) {

	for cmd := range cmd {
		if len(cmd.Fields) < 1 {
			cmd.Result <- "Please input a command you'd like to execute\n"
			continue
		}

		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expected at least 2 arguments\n"
			continue
		}

		if len(cmd.Fields) > 3 {
			cmd.Result <- "Too many arguments"
			continue
		}

		fmt.Println("Command:", cmd.Fields)

		// Executing commands
		switch strings.ToLower(cmd.Fields[0]) {

		// GET <KEY>
		case "get":
			_, state := entities.Data[cmd.Fields[1]]
			if !state {
				cmd.Result <- "state: absent"
			} else {
				cmd.Result <- "value: " + entities.Data[cmd.Fields[1]] + "; state: present"
			}

		// SET <KEY> <VALUE>
		case "set":
			match, _ := regexp.MatchString("^[\\w]+$", cmd.Fields[1])
			if !match {
				cmd.Result <- "Incorrect key, please, try again"
				continue
			} else if len(cmd.Fields) != 3 {
				cmd.Result <- "Expected value"
				continue
			}

			entities.Data[cmd.Fields[1]] = cmd.Fields[2]
			cmd.Result <- "Including key and value in database..."

		// DEL <KEY>
		case "del":
			_, state := entities.Data[cmd.Fields[1]]
			if !state {
				cmd.Result <- "state: ignored"
			} else {
				delete(entities.Data, cmd.Fields[1])
				cmd.Result <- "state: absent"
			}

		// KEYS <PATTERN>
		case "keys":
			keys := make([]string, 0)
			keyString := ""

			if strings.Contains(cmd.Fields[1], "*") {
				cmd.Fields[1] = strings.TrimRight(cmd.Fields[1], "*")

				for key := range entities.Data {
					if strings.HasPrefix(key, cmd.Fields[1]) {
						keys = append(keys, key)
					}
				}

				sort.Strings(keys)
				result := strings.Join(keys, ", ")

				if len(keys) == 0 {
					cmd.Result <- "There are no keys matching the pattern"
				} else {
					cmd.Result <- result
				}

			} else {

				for key := range entities.Data {
					if key == cmd.Fields[1] {
						keyString += key
					}
				}

				if len(keyString) == 0 {
					cmd.Result <- "There are no keys matching the pattern"
				} else {
					cmd.Result <- keyString
				}
			}

		default:
			cmd.Result <- "Invalid command " + cmd.Fields[0]
		}
	}
}
