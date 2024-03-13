package hub

import (
	"encoding/json"
	"fmt"
	"os"
)

func Export() error {
	err := exportTagCategory("tagcategories.json")
	if err != nil {
		return err
	}
	err = exportTag("tags.json")
	if err != nil {
		return err
	}
	err = exportApplication("applications.json")
	if err != nil {
		return err
	}
	
	
	return err
}

func handleEmptyList() {
	fmt.Printf("empty response, skipping")
}

func saveData(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func exportApplication(filename string) error {
	list, err := RichClient.Application.List()
	if err != nil {
		return err
	}
	if len(list) < 1 {
		handleEmptyList()
		return nil
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	err = saveData(filename, data)
	return err
}

func exportTagCategory(filename string) error {
	list, err := RichClient.TagCategory.List()
	if err != nil {
		return err
	}
	if len(list) < 1 {
		handleEmptyList()
		return nil
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	err = saveData(filename, data)
	return err
}

func exportTag(filename string) error {
	list, err := RichClient.Tag.List()
	if err != nil {
		return err
	}
	if len(list) < 1 {
		handleEmptyList()
		return nil
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	err = saveData(filename, data)
	return err
}
