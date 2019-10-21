package main

import (
          "bufio" //for input
	  "flag"  //for using flag
	  "fmt"  //input output
	  "log"  //for logging errors etc
	  "os"   //file handling
	  "strconv"  //to convert strings
           "github.com/dghubble/go-twitter/twitter"
	  "github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter twitterhandle: ")
    x, _ := reader.ReadString('\n')
	input := flag.String("user", x, "twitter handle")
	flag.Parse()

	
	creds := Credentials{
		AccessToken:       "",
		AccessTokenSecret: "",
		ConsumerKey:       "",
		ConsumerSecret:    "",
	} //Enter credentials here

	client, err := getClient(&creds, input)
	if err != nil {
		log.Println("Error getting the Twitter Client")
		log.Println(err)
	}
	_ = client

}

func getClient(creds *Credentials, input *string) (*twitter.Client, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	user, resp, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: *input,
	})

	if err != nil {
		log.Println(err)
		fmt.Printf("User not found.\n")
		return client, nil
	}
	_ = resp
	fmt.Printf("\n%+v\n%+v\n%+v\n", user.Name, user.Description, user.FollowersCount)
	fmt.Printf("%+v\n%+v\n", user.FriendsCount, user.Location)

	f, err := os.Create("ClientDetails")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return client, nil
	}

	
	details := []string{"Twitter Handle : ", user.ScreenName, "Name :", user.Name, "Description : ", user.Description, "Followers : ", strconv.Itoa(user.FollowersCount), "Following : ", strconv.Itoa(user.FriendsCount), "Location : ", user.Location, "URL : ", user.URL, "Created At : ", user.CreatedAt}

	for k, v := range details {
		if k%2 == 0 {
			fmt.Fprintln(f)
		fmt.Fprintf(f, v)
		if err != nil {
			fmt.Println(err)
			return client, nil
		}
	}

	fmt.Fprintln(f)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return client, nil
	}

	fmt.Println("\nAbove details stored in file:ClientDetails")
    }
	return client, nil
}
