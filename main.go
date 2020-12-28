package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"log"
	"strings"
	"net"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
	"github.com/joho/godotenv"
	"github.com/jrm780/gotirc"
)

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func execute(command string) string {
  out, err := exec.Command("bash", "-c", command).Output()

  if err != nil {
    fmt.Println("%s", err)
  }

  output := string(out[:])
  return output
}

func changeVimColor(messageString []string, client *gotirc.Client, channel string, tags map[string]string, clientRPC *rpc.Session){


	if len(messageString) == 2{

		if !strings.ContainsAny(messageString[1], "<>:|") && strings.HasPrefix(messageString[0], "!color"){

			themesAvailble := themes()

			fmt.Println(messageString[1])
			_, found := Find(themesAvailble, messageString[1])
			if found {
				sendInput("<ESC>:color " + messageString[1] + "<CR>", clientRPC)
			}
			if !found{
				client.Say(channel, tags["display-name"] + " Tema não encontrado.")
			}
		}
	}
}

func listThemes(messageString []string, client *gotirc.Client, channel string, tags map[string]string){

	if messageString[0] == "!themes"{

		client.Say(channel, tags["display-name"] + " Veja os temas disponíveis em: https://github.com/edersonferreira/neovim-integration-twitch/blob/main/colors.txt")
	}
}

func move(messageString []string, tags map[string]string, clientRPC *rpc.Session){
	if messageString[0] == "!move" && len(messageString) == 2{
		movement := messageString[1]
		if !strings.ContainsAny(movement, "drRD:<>ZaioAIOuUvtyYcCsSxX!|-+~"){
			sendInput("<ESC>" + messageString[1], clientRPC)
		}
	}
}

func initRPC(port string) *rpc.Session {
	conn, err := net.Dial("tcp", "localhost:" + port)
	if err != nil {
                fmt.Println("fail to connect to server.")
	}
	clientRPC := rpc.NewSession(conn, true)

	return clientRPC
}

func sendInput(input string, clientRPC *rpc.Session){
	_, xerr := clientRPC.Send("nvim_input", input)
	if xerr != nil {
		fmt.Println(xerr)
		return
	}
}

func main() {


  errEnv := godotenv.Load()
  if errEnv != nil {
    log.Fatal("Error loading .env file")
  }

	token := os.Getenv("TOKEN")
	user := os.Getenv("USER")
	channel := os.Getenv("CHANNEL")
	port := os.Getenv("PORT")

	clientRPC := initRPC(port)

        options := gotirc.Options{
            Host:     "irc.chat.twitch.tv",
            Port:     6667,
            Channels: []string{"#" + channel},
        }

        client := gotirc.NewClient(options)
        
        // Whenever someone sends a message, log it
				client.OnChat(func(channel string, tags map[string]string, msg string) {
				
					fmt.Println(msg)
					messageString := strings.Split(msg, " ")

					go listThemes(messageString, client, channel, tags)
					go move(messageString, tags, clientRPC)
					go changeVimColor(messageString, client, channel, tags, clientRPC)
     })

		 client.Connect(user, token)
}

func themes() []string{

	data, err := ioutil.ReadFile("colors.txt")

	if err != nil {
		panic(err)
	}

	themes := strings.Fields(string(data))

	return themes
}
