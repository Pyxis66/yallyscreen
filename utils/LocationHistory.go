package utils

import (
	// "fmt"
	// "sort"

	"github.com/Pyxis66/YallyScreen/logger"
	// "github.com/Pyxis66/YallyScreen/octoprintApis"
	"github.com/Pyxis66/YallyScreen/octoprintApis/dataModels"
	// "github.com/Pyxis66/YallyScreen/uiWidgets"
)

type LocationHistory struct {
	Locations []dataModels.Location
}

func (this *LocationHistory) Length() int {
	return len(this.Locations)
}

func (this *LocationHistory) CurrentLocation() dataModels.Location {
	length := this.Length()
	if length < 1 {
		logger.Error("CurrentLocation() - length < 1")
		panic("PANIC!!! - LocationHistory.current() - locations is empty")
	}

	return this.Locations[length - 1]
}

func (this *LocationHistory) GoForward(folder string) {
	newLocation := string(this.CurrentLocation()) + "/" + folder
	this.Locations = append(this.Locations, dataModels.Location(newLocation))
}

func (this *LocationHistory) GoBack() {
	this.Locations = this.Locations[0 : len(this.Locations) - 1]
}

func (this *LocationHistory) IsRoot() bool {
	if len(this.Locations) > 1 {
		return false
	} else {
		return true
	}
}
