package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/binxio/gcloud/config"
	"golang.org/x/oauth2/google"
	"log"
)

// print the project access token and expiration date
func printToken(credentials *google.Credentials) {
	token, err := credentials.TokenSource.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("current project %s\n", credentials.ProjectID)
	fmt.Printf("access token    %v\n", token.AccessToken)
	fmt.Printf("expires         %v\n", token.Expiry)
}

// print the select configuration properties and the credentials provided by them
func main() {
	name := flag.String("configuration", "", "name of the configuration to use")
	flag.Parse()

	if !config.IsGCloudOnPath() {
		log.Fatal("please install the Google Cloud SDK and put gcloud on the path")
	}

	active, err := config.GetCloudSDKConfig(*name)
	if err != nil {
		log.Fatal(err)
	}
	if s, err := json.MarshalIndent(active, "", "  "); err == nil {
		fmt.Printf("%s\n", string(s))
	} else {
		log.Fatal(err)
	}

	credentials, err := config.GetCloudSDKCredentials(*name)
	if err != nil {
		log.Fatal(err)
	}

	printToken(credentials)
}
