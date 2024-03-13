package hub

import (
	"encoding/json"
	"os"

	"github.com/konveyor/tackle2-hub/api"
)

func Import() error {
	err := importApplication("applications.json")
	if err != nil {
		return err
	}
	
	
	return err
}


func loadData(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func importApplication(filename string) error {
	list := []*api.Application{}
	data, err := loadData(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	for i := range list {
		// check for alreadsy existing before create? or continue if exists by id, but rather check all before and stop
		err := RichClient.Application.Create(list[i])
		if err != nil {
			return err
		}
	}
	
	return err
}
