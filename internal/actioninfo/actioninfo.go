package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, action := range dataset {
		err := dp.Parse(action)
		if err != nil {
			log.Printf("action: %s, parse failed: %s", action, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action: %s, gettign action info failed: %s", action, err)
			continue
		}

		fmt.Println(info)
	}
}
