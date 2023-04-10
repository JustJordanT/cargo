package internals

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v37/github"
	"github.com/justjordant/cargo/internals/initUtils"
	"golang.org/x/oauth2"
)

var GitEnv = os.Getenv("CARGO_TOEKN")

// Set up a new GitHub client with an access token
var ctx = context.Background()
var ts = oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: GitEnv},
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
			crateUrl = content.GetDownloadURL()
		}
	}
	return crateUrl
}

func DownloadCargoYaml(crateUrl string) (err error) {
	// Create the filePath
	out, err := os.Create(initUtils.CargoPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(crateUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
