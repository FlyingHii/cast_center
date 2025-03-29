package upwork

import (
	"bufio"
	"fmt"
	"os"

	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/auth"
)

const cfgFile = "config.json" // update the path to your config file, or provide properties directly in your code
func getClient() {
	client := api.Setup(api.ReadConfig(cfgFile))
	// we need an access token/secret pair in case we haven't received it yet
	if !client.HasAccessToken() {
		aurl := client.GetAuthorizationUrl("")

		// read verifier
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Visit the authorization url and provide oauth_verifier for further authorization")
		fmt.Println(aurl)
		verifier, _ := reader.ReadString('\n')

		// get access token
		token := client.GetAccessToken(verifier)
		fmt.Println(token)
	}

}

func GetUserInfo() {
	client := api.Setup(api.ReadConfig(cfgFile))

	userInfo, err := auth.New(client).GetUserInfo()
	if err != nil {
		fmt.Println("Error getting user info:", err)
	}
	fmt.Println(userInfo)
}
