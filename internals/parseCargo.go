package internals

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v37/github"
	"github.com/justjordant/cargo/internals/initUtils"
	"golang.org/x/oauth2"
)

var GitEnv = os.Getenv("CARGO_TOKEN")

// Set up a new GitHub client with an access token
var ctx = context.Background()
var githubToken = oauth2.StaticTokenSource(
	&oauth2.Token{AccessToken: GitEnv},
)
var tc = oauth2.NewClient(ctx, githubToken)
var client = github.NewClient(tc)

func GetCrateUrl(fileName string) (string, string) {
	// Get the contents of the "crates" folder
	crateUrl := ""
	crateFile := ""
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
			crateFile = content.GetName()
		}
	}
	return crateUrl, crateFile
}

func DownloadCargoYaml(fileName string) (filePath string, err error) {
	url, crateName := GetCrateUrl(fileName)
	// Create the filePath
	out, err := os.Create(FormatFileName(crateName))
	if err != nil {
		return filePath, err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return filePath, err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return filePath, err
	}

	// TODO we should be returning the failPath to be able to look up the yaml configration for getting the zip url location.
	return filePath, nil
}

func FormatFileName(fileName string) string {
	return initUtils.CargoPath + fileName
}

func CheckLocalCrateDir(crateName string) {
	// TODO -  This will check local dir for crates that are trying to be installed.
	// ie cargo install kubectl this would look here first for any know tap like file with the instructions.
	log.Fatal("func not yet implamented!!")
}

// TODO - This would be a func that you would pass in url and api key
func CheckContainerDir() {
	// TODO This will checks local containers for any instructions on where to look for the instructions for a given crate.
	// ie cargo install <org-name> <package>
	log.Fatal("func not yet implamented!!")
}

func InitContainerYaml() {
	// TODO - Create a org conf file with instructions on how to look install bla bla, this would be where it would say where the VCS is located Github, Gitlab, GitTea
	log.Fatal("func not yet implamented!!")
}

func CompareSHA256(filename string, expectedSHA string) (bool, error) {
	// Read the contents of the file
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	// Calculate the SHA256 checksum of the file
	sha256Bytes := sha256.Sum256(fileBytes)
	actualSHA := hex.EncodeToString(sha256Bytes[:])

	// Compare the calculated SHA256 with the expected value
	if actualSHA != expectedSHA {
		return false, nil
	}

	return true, nil
}

func EncryptCred() {
	// TODO - How should this be encrypted? Bcrypt, RSA?
	log.Fatal("func not yet implamented!!")
}

func DecryptCred() {
	// TODO - this wher we will pass in the encrypted value from the yaml file and we will decrypt it for use.
	log.Fatal("func not yet implamented!!")
}
