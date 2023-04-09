package internals

import (
	"context"
	"log"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

// Set up a new GitHub client with an access token
var ctx = context.Background()
var ts = oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: "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"},
)
var tc = oauth2.NewClient(ctx, ts)
var client = github.NewClient(tc)

func GetCrateUrl(fileName string) string {
	// Get the contents of the "crates" folder
	crateUrl := ""
	_, dirContents, _, err := client.Repositories.GetContents(ctx, "justjordant", "cargo-crates", "Crates", &github.RepositoryContentGetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the contents of the "crates" folder
	for _, content := range dirContents {
		//  if content.GetType() == "dir" {
		//  // If the content is a directory, get its contents as well
		//  _, subDirContents, _, err := client.Repositories.GetContents(ctx, "justjordant", "cargo-crates", content.GetPath(), &github.RepositoryContentGetOptions{})
		//  if err != nil {
		//  panic(err)
		//  }

		//  // Print the names of the files in the subdirectory
		//  for _, subContent := range subDirContents {
		//  fmt.Println(subContent.GetName())
		//  }
		//  } else {
		// If the content is a file, print its name
		//  }
		//  fmt.Println(content.GetName())
		//  fmt.Println(content.GetType())
		//  fmt.Println(content.GetURL())
		//  fmt.Println(content.GetGitURL())
		//  fmt.Println(content.GetPath())
		//  fmt.Println("\n")
		if fileName == content.GetName() {
			crateUrl = content.GetURL()
		}
	}
	return crateUrl
}

func downloadCargoYaml(fileName string) {
}
