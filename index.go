package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
	"github.com/joho/godotenv"
	"github.com/jrm780/gotirc"
)


func execute(command string) string {
  out, err := exec.Command("bash", "-c", command).Output()

  if err != nil {
    fmt.Println("%s", err)
  }

  output := string(out[:])
  return output
}

func changeVimColor(messageString []string){

	if len(messageString) == 2{

		if !strings.ContainsAny(messageString[1], "<>:|") && strings.HasPrefix(messageString[0], "!color"){
			execute("nvr --remote-send \":color " + messageString[1] + "<CR>\"")
		}
	}
}

func listThemes(messageString []string, client *gotirc.Client, channel string, tags map[string]string){

	if messageString[0] == "!themes"{

		client.Say(channel, tags["display-name"] + " Os temas disponíveis são: blue, dalton, darkblue, default, delek, desert, elflord, evening, industry, koehler, morning, murphy, pablo, peachpuff, ron, shine, slate, torte, zellner")

	}
}

func main() {

  errEnv := godotenv.Load()
  if errEnv != nil {
    log.Fatal("Error loading .env file")
  }

	token := os.Getenv("TOKEN")

        options := gotirc.Options{
            Host:     "irc.chat.twitch.tv",
            Port:     6667,
            Channels: []string{"#edersondeveloper"},
        }

        client := gotirc.NewClient(options)
        
        // Whenever someone sends a message, log it
				client.OnChat(func(channel string, tags map[string]string, msg string) {
				
					fmt.Println(msg)
					messageString := strings.Split(msg, " ")

					listThemes(messageString, client, channel, tags)
					changeVimColor(messageString)
     })

		 client.Connect("edersondeveloper", token)
}
