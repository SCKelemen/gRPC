package UpdateManager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UpdateManager struct {
	Endpoint string
	Release  ReleaseStream
	Version  string
}

func (mgr UpdateManager) CheckForUpdates() {
	url := fmt.Sprintf("%s/updates", mgr.Endpoint)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var schedule UpdateSchedule
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal(content, &schedule)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	DownloadUpdates(getStream(schedule, mgr.Release))
}

func (mgr UpdateManager) DownloadUpdates(stream ReleaseStream) {

}

func (mgr UpdateManager) ApplyUpdates() {

}

type UpdateSchedule struct {
	Production  Release `json:"production"`
	Staging     Release `json:"staging"`
	Development Release `json:"development"`
}

type Release struct {
	Source   string `json:"source"`
	Checksum string `json:"checksum"`
}

type ReleaseStream int

const (
	// Production is the stable release
	Production ReleaseStream = iota
	// Staging is the beta/testing release
	Staging
	// Development is the YOLO release
	Development
)

func getStream(schedule UpdateSchedule, stream ReleaseStream) Release {
	switch stream {
	case Staging:
		return schedule.Staging
	case Development:
		return schedule.Development
	default:
		return schedule.Production
	}
}

func New(url string) (UpdateManager, error) {
	mgr := UpdateManager{Endpoint: url}
	return mgr, nil
}
