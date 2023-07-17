//go:generate goversioninfo
package main

import (
	utils "Anon/utils"
	"os"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"gopkg.in/toast.v1"
)

func main() {
	if len(os.Args) < 2 {
		utils.Error("Impossible to find file path.", true)
	}

	s := spinner.New(spinner.CharSets[34], 1*time.Second)
	s.Suffix = " Loading..."
	s.Start()

	if utils.IsDir(os.Args[1]) {
		utils.SpinError("Invalid file path.", true, s)
	}

	abs, err := filepath.Abs("./icons/anon.png")
	if err != nil {
		utils.SpinError("Invalid file path.", true, s)
	}

	notification := toast.Notification{
		AppID:   "Anonfiles",
		Title:   "Uploading...",
		Message: "The file is about to get uploaded to anonfiles",
		Icon:    abs,
	}
	_ = notification.Push()

	ok, data := utils.Upload(os.Args[1])
	if !ok {
		notification := toast.Notification{
			AppID:   "Anonfiles",
			Title:   "Upload Error",
			Message: "Cannot upload the selected file to anonfiles.com",
			Icon:    abs,
		}
		_ = notification.Push()
		utils.SpinError("Impossible to upload to anonfiles.com", true, s)
	} else {
		utils.SpinCheck("File uploaded to anonfiles.com", false, s)
	}

	notification = toast.Notification{
		AppID:   "Anonfiles",
		Title:   "Upload Successful",
		Message: "File uploaded successfully to anonfiles.com",
		Icon:    abs,
	}
	_ = notification.Push()

	err = clipboard.WriteAll(data.Data.File.URL.Full)
	if err != nil {
		utils.SpinError("Impossible to copy text to the clipboard", true, s)
	}
}
