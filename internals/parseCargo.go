package internals

import (
	"context"
	"fmt"
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

func ParseCargo() {

	// Get the contents of the "crates" folder
	_, dirContents, _, err := client.Repositories.GetContents(ctx, "justjordant", "cargo-crates", "Crates", &github.RepositoryContentGetOptions{})
	if err != nil {
		panic(err)
	}

	// Iterate through the contents of the "crates" folder
	for _, content := range dirContents {
		if content.GetType() == "dir" {
			// If the content is a directory, get its contents as well
			_, subDirContents, _, err := client.Repositories.GetContents(ctx, "justjordant", "cargo-crates", content.GetPath(), &github.RepositoryContentGetOptions{})
			if err != nil {
				panic(err)
			}

			// Print the names of the files in the subdirectory
			for _, subContent := range subDirContents {
				fmt.Println(subContent.GetName())
			}
		} else {
			// If the content is a file, print its name
			fmt.Println(content.GetName())
		}
	}
}
