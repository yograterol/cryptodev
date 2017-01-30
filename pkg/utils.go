package pkg

import (
	"github.com/cavaliercoder/grab"
	"fmt"
	"os"
	"time"
	"os/exec"
	"log"
	"path/filepath"
	"github.com/asdine/storm"
)

var (
	folderCryptodev string
	folderCryptodevBin string
	folderCryptodevData string
	downloadDest string
	db *storm.DB
)

func init() {
	homePath := os.Getenv("HOME")
	folderCryptodev = filepath.Join(homePath, ".cryptodev")
	folderCryptodevBin = filepath.Join(folderCryptodev, "bin")
	folderCryptodevData = filepath.Join(folderCryptodev, "data")
	downloadDest = filepath.Join("/tmp", "cryptodev")

	createDir(folderCryptodev)
	createDir(folderCryptodevBin)
	createDir(folderCryptodevData)
	createDir(downloadDest)

	dbInstance, err := storm.Open(filepath.Join(folderCryptodev, "cryptodev.db"))
	if err != nil {
		log.Fatal("Can't load the database.")
		os.Exit(1)
	}
	db = dbInstance
}

func createDir(target string) {
	if _, err := os.Stat(target); os.IsNotExist(err) {
		os.Mkdir(target, 0755)
	}
}

func downloadTarball(url string, dest string) string {
	// Code from: http://cavaliercoder.com/blog/downloading-large-files-in-go.html
	respCh, err := grab.GetAsync(dest, url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", url, err)
		os.Exit(1)
	}

	resp := <-respCh

	for !resp.IsComplete() {
		fmt.Printf("\033[1ADownloading: %s Progress %d / %d bytes (%d%%)\033[K\n",
			resp.Filename, resp.BytesTransferred(), resp.Size, int(100 * resp.Progress()))
		time.Sleep(200 * time.Millisecond)
	}

	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", url, resp.Error)
		os.Exit(1)
	}

	return resp.Filename
}

func ungzip(source, target string) {
	cmdName := "tar"
	cmdArgs := []string{
		"xvf",
		source,
		"-C",
		target,
		"--strip-components=1"}

	if _, err := os.Stat(target); err == nil {
		os.RemoveAll(target)
	}

	err := os.Mkdir(target, 0755)

	if err != nil {
		log.Fatal("There was an error creating the target folder.")
		os.Exit(1)
	}

	errExtracting := exec.Command(cmdName, cmdArgs...).Run()

	if errExtracting != nil {
		fmt.Println(errExtracting)
		log.Fatal("There was an error extracting the files.")
		os.Exit(1)
	}
}
