package syncer

import (
	"fmt"
	"github.com/zloylos/grsync"
	"time"
)

func syncer() {
	task := grsync.NewTask(
		"username@server.com:/source/folder",
		"/home/user/destination",
		grsync.RsyncOptions{},
	)
	go func() {
		for {
			state := task.State()
			fmt.Printf(
				"progress: %.2f / rem. %d / tot. %d / sp. %s \n",
				state.Progress,
				state.Remain,
				state.Total,
				state.Speed,
			)
			<-time.After(time.Second)
		}
	}()

	if err := task.Run(); err != nil {
		panic(err)
	}

	fmt.Println("well done")
	fmt.Println(task.Log())
}
