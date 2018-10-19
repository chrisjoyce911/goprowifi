package goprowifi

import (
	"encoding/json"
	"fmt"
)

// VideoInformation Video Information
type VideoInformation struct {
	AvcProfile string `json:"avc_profile"`
	C          string `json:"c"`
	// Video duration
	VideoDuration string `json:"dur"`
	FPS           string `json:"fps"`
	FpsDenom      string `json:"fps_denom"`
	Hight         string `json:"h"`
	Profile       string `json:"profile"`
	Prog          string `json:"prog"`
	Subsample     string `json:"subsample"`
	// Number of hilight tags
	TagCount int `json:"tag_count"`
	// Array of tags in milliseconds
	Tags  []int  `json:"tags"`
	Width string `json:"w"`
}

// GetVideoInformation ...
//
//  http://10.5.5.9:8080/gp/gpMediaMetadata?p=100GOPRO/GOPR1148.MP4&t=videoinfo
func (g *Client) GetVideoInformation(folder string, file string) (reply VideoInformation, httpStatus int, err error) {

	url := fmt.Sprintf("http://10.5.5.9:8080/gp/gpMediaMetadata?p=%s/%s&t=videoinfo", folder, file)
	bodyBytes, httpStatus, err := g.request(url)

	err = json.Unmarshal(bodyBytes, &reply)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	return
}

// ExtendedVideoInformation Extended Video Information
type ExtendedVideoInformation struct {
	Ao             string `json:"ao"`
	Cl             string `json:"cl"`
	EpochTimestamp string `json:"cre"`
	// Video duration
	VideoDuration string `json:"dur"`
	EIS           string `json:"eis"`
	FPS           string `json:"fps"`
	FpsDenom      string `json:"fps_denom"`
	Gumi          string `json:"gumi"`
	Hight         string `json:"h"`
	// Number of hilight tags
	TagCount string `json:"hc"`
	// Array of tags in milliseconds
	Tags      []int  `json:"hi"`
	Ls        string `json:"ls"`
	Mp        string `json:"mp"`
	Prog      string `json:"prog"`
	Pta       string `json:"pta"`
	S         string `json:"s"`
	Subsample string `json:"subsample"`
	Tr        string `json:"tr"`
	Us        string `json:"us"`
	Width     string `json:"w"`
}

// GetExtendedVideoInformation Extended Video Information
//
// http://10.5.5.9/gp/gpMediaMetadata?p=/100GOPRO/GOPR1148.MP4&t=v4info
func (g *Client) GetExtendedVideoInformation(folder string, file string) (reply ExtendedVideoInformation, httpStatus int, err error) {

	url := fmt.Sprintf("http://10.5.5.9:8080/gp/gpMediaMetadata?p=%s/%s&t=v4info", folder, file)
	bodyBytes, httpStatus, err := g.request(url)

	err = json.Unmarshal(bodyBytes, &reply)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	return
}
