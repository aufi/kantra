package hub

import (
	"fmt"
	"os"

	"github.com/konveyor/tackle2-hub/binding"
)

const (
	HubBaseURL = "HUB_BASE_URL"
	Username = "HUB_USERNAME"
	Password = "HUB_PASSWORD"
)

var (
	// Setup Hub API client
	Client     *binding.Client
	RichClient *binding.RichClient
	Resources  []interface{}
)

func Connect() (err error) {
	if os.Getenv(HubBaseURL) == "" {
		return fmt.Errorf("cannot find Hub URL, specify it in %s env variable", HubBaseURL)
	}
	RichClient = binding.New(os.Getenv(HubBaseURL))
	RichClient.Client.Retry = 1
	err = RichClient.Login(os.Getenv(Username), os.Getenv(Password))
	return err
}
