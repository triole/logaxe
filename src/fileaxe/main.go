package fileaxe

import (
	"fmt"

	"github.com/triole/logseal"
)

func (fa FileAxe) Run() {
	if !fa.Conf.Ls.Plain {
		fa.Lg.Debug(
			"start fileaxe",
			logseal.F{"conf": fmt.Sprintf("%+v", fa.Conf)},
		)
	}

	fileList := fa.Find(
		fa.Conf.Folder, fa.Conf.Matcher,
		fa.Conf.MinAge, fa.Conf.MaxAge, fa.Conf.Now,
	)
	switch fa.Conf.Action {
	case "ls":
		fa.list(fileList)
	case "ex":
		fa.exists(fileList)
	case "rt":
		fa.rotate(fileList)
	case "cp":
		fa.copy(fileList)
	case "mv":
		fa.move(fileList)
	case "tn":
		fa.truncate(fileList)
	case "rm":
		fa.remove(fileList)
	}
}
