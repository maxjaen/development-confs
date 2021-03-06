// https://learnxinyminutes.com/docs/go/
// https://learn.go.dev/
// https://golang.org/
// https://blog.golang.org/slices-intro

package main

import (
	"fmt"
	"os"
	"os/user"
    "io/ioutil"
	"log"
)

func main() {
	homeDirConfigs := []string{".gitconfig"}
	// add more configs from home directory here...

	syncWithRepoFiles(homeDirConfigs)
}

func syncWithRepoFiles(files []string) {
	for _, filePath := range files {
		syncWithRepoFile(filePath)
	}
}

func syncWithRepoFile(filePath string) {
	fileInHomeDir := getUserHomeDir() + "/" + filePath
	fileInRepoDir := "./" + filePath

	syncFilePossible := filesExists([]string{fileInHomeDir, fileInRepoDir})
	if syncFilePossible {
		deleteFile(fileInHomeDir)
		copyFile(fileInRepoDir, fileInHomeDir)
	}
}

func filesExists(fileNames []string) bool {
	for _, fileName := range fileNames {
		if !fileExists(fileName)  {
			fmt.Println(fmt.Sprintf("File %s does not exist %s", fileName, toGreenStr("(skipping)")))
			return false
		}
	}

	return true
}

func fileExists(fileName string) bool {
    info, err := os.Stat(fileName)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func copyFile(src string, dst string) {
	fmt.Println(fmt.Sprintf("%s %s to %s", toBlueStr("Copy"), src, dst))

    data, err := ioutil.ReadFile(src)
    checkError(err)
    err = ioutil.WriteFile(dst, data, 0644)
    checkError(err)
}

func deleteFile(filePath string) {
	fmt.Println(fmt.Sprintf("%s %s", toRed("Delete"), filePath))

	err := os.Remove(filePath)
    checkError(err)
}

func getUserHomeDir() string {
	usr, err := user.Current()
    checkError(err)

	return usr.HomeDir
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func toBlueStr(str string) string {
	return toColorfulStr("\033[34m", str)
}

func toRed(str string) string {
	return toColorfulStr("\033[31m", str)
}

func toGreenStr(str string) string {
	return toColorfulStr("\033[32m", str)
}

func toColorfulStr(color string, str string) string {
	return color + str + "\033[0m"
}
