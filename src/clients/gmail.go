package clients

import (
	"time"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
	"golang.org/x/oauth2/google"
	"log"
	"fmt"
)

func Retrieve(gmailClientCredentials string, refreshToken string) {

	config, err := google.ConfigFromJSON([]byte(gmailClientCredentials), gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// This is somewhat of a hack as we aren't storing the access token and
	// relying on the refresh token to obtain a new session every time
	srv, err := gmail.New(config.Client(context.Background(), &oauth2.Token{
		AccessToken: "invalid-token-to-rely-on-refresh-token",
		TokenType: "Bearer",
		RefreshToken: refreshToken,
		Expiry: time.Unix(0, 0), // Already expired to force use of refresh token
	}))

	r, err := srv.Users.Labels.List("me").Do()

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

