package FSNotify

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

type NotifyFile struct {
	watch *fsnotify.Watcher
}

func NewNotifyFile() *NotifyFile {
	w := new(NotifyFile)
	w.watch, _ = fsnotify.NewWatcher()
	return w
}

//Authority Dir
func (this *NotifyFile) WatchDir(dir string) {
	//Use Walk to traverse all subdirectories under the directory
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//Determine whether it is a directory, monitoring directory, and files under the directory are within the monitoring range, no need to add
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = this.watch.Add(path)
			if err != nil {
				return err
			}
			fmt.Println("Monitoring : ", path)
		}
		return nil
	})

	go this.WatchEvent() //Multi thread
}

func (this *NotifyFile) WatchEvent() {
	for {
		select {
		case ev := <-this.watch.Events:
			{
				if ev.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("Create File : ", ev.Name)
					//Get the information of the newly created file, if it is a directory, add it to the monitoring
					file, err := os.Stat(ev.Name)
					if err == nil && file.IsDir() {
						this.watch.Add(ev.Name)
						fmt.Println("Add Monitoring : ", ev.Name)
					}
				}

				if ev.Op&fsnotify.Write == fsnotify.Write {
					//fmt.Println("Write File : ", ev.Name)
				}

				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("Delete File : ", ev.Name)
					//If the deleted file is a directory, remove monitoring
					fi, err := os.Stat(ev.Name)
					if err == nil && fi.IsDir() {
						this.watch.Remove(ev.Name)
						fmt.Println("Delete Monitoring : ", ev.Name)
					}
				}

				if ev.Op&fsnotify.Rename == fsnotify.Rename {
					//If the renamed file is a directory, remove the monitoring. Note that os.Stat cannot be used to determine whether it is a directory.
					//Because after renaming, go can no longer find the original file to obtain information, so simply remove it directly
					fmt.Println("Rename File : ", ev.Name)
					this.watch.Remove(ev.Name)
				}
				if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					fmt.Println("Change Authority : ", ev.Name)
				}
			}
		case err := <-this.watch.Errors:
			{
				fmt.Println("error : ", err)
				return
			}
		}
	}
}
