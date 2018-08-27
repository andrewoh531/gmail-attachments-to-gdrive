package clients

import (
	"time"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"fmt"
)

func Retrieve(accessToken string, refreshToken string) {
	// TODO obtain credentials in a deployable way (cannot check in credentials.json)
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := gmail.New(config.Client(context.Background(), &oauth2.Token{
		AccessToken: accessToken,
		TokenType: "Bearer",
		RefreshToken: refreshToken,
	}))

	user := "me"
	r, err := srv.Users.Labels.List(user).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}

	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}
}

func getStartDate() time.Time {
	oneWeekDuration := time.Duration(24 * 7) * time.Hour
	oneWeekAgo := time.Now().Add(-oneWeekDuration)
	return sundayOfThatWeek(oneWeekAgo)
}

func sundayOfThatWeek(date time.Time) time.Time {
	weekday := date.Weekday()
	return date.Add(- time.Duration(int(weekday *  24)) * time.Hour)
}

