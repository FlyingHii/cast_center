package upwork

import (
	"bufio"
	"fmt"
	"os"

	"github.com/upwork/golang-sdk/api"
	"github.com/upwork/golang-sdk/api/auth"
)

const CfgFile = "config.json" // update the path to your config file, or provide properties directl
func getClient() {
	client := api.Setup(api.ReadConfig(CfgFile))
	// we need an access token/secret pair in case we haven't received it yet
	if !client.HasAccessToken() {
		aurl := client.GetAuthorizationUrl("")

		// read verifier
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Visit the authorization url and provide oauth_verifier for further authorization")
		fmt.Println(aurl)
		fmt.Print("Paste oauth_verifier here: ")
		verifier, _ := reader.ReadString('\n')
		verifier = verifier[:len(verifier)-1]

		_, err := auth.New(client).GetAccessToken(verifier)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Access token received, you can start using API methods")
		fmt.Println("Please store access token and secret in `config.json`")
		fmt.Println("config.json:")
		fmt.Printf("%#v\n", client.Config)
	} else {
		fmt.Println("You already have access token, you can start using API methods")
	}
}

func GetUserInfo() {
	client := api.Setup(api.ReadConfig(CfgFile))

	userInfo, err := auth.New(client).GetUserInfo()
	if err != nil {
		fmt.Println("Error getting user info:", err)
	}
	fmt.Println(userInfo)
}
