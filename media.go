package goprowifi

import (
	"encoding/json"
	"fmt"
)

// Media list of the contents of the SD card
type Media struct {
	Device         string `json:"device"`
	EpochTimestamp string `json:"epoch_timestamp"`
	FileName       string `json:"file_name"`
	MediaType      string `json:"media_type"`
	Ls             string `json:"ls"`
	S              string `json:"s"`
}

type cameraMediaList struct {
	ID    string        `json:"id"`
	Media []cameraMedia `json:"media"`
}

// cameraMedia ...
type cameraMedia struct {
	Device string  `json:"d"`
	Files  []Files `json:"fs"`
}

// Files ...
type Files struct {
	Ls             string `json:"ls"`
	EpochTimestamp string `json:"mod"`
	FileName       string `json:"n"`
	S              string `json:"s"`
}

// GetMediaList ist of the contents of the SD card
//
// http://10.5.5.9:8080/gp/gpMediaList
func (g *Client) getMediaList() (mediaList []*Media, httpStatus int, err error) {

	bodyBytes, httpStatus, err := g.request("http://10.5.5.9:8080/gp/gpMediaList")

	reply := new(cameraMediaList)
	err = json.Unmarshal(bodyBytes, &reply)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	for _, cm := range reply.Media {
		for _, dm := range cm.Files {
			nm := new(Media)
			nm.Device = cm.Device
			nm.EpochTimestamp = dm.EpochTimestamp
			nm.FileName = dm.FileName
			nm.Ls = dm.Ls
			nm.S = dm.S
			nm.MediaType = nm.FileName[len(nm.FileName)-3:]
			mediaList = append(mediaList, nm)
		}
	}

	return
}

// Media List
//http://10.5.5.9:8080/gp/gpMediaList

// {
//     "id": "3806895087272624399",
//     "media": [
//         {
//             "d": "100GOPRO",
//             "fs": [
//                 {
//                     "n": "GOPR1130.MP4",
//                     "mod": "1539705058",
//                     "ls": "11170499",
//                     "s": "188100018"
//                 },
//                 {
//                     "n": "GOPR1147.JPG",
//                     "mod": "1539960090",
//                     "s": "3843146"
//                 }
//             ]
//         }
//     ]
// }

// Media Thumbnail:
// To get the thumbnail of a video: http://10.5.5.9/gp/gpMediaMetadata?p=100GOPRO/GOPR1148.MP4
// To get the thumbnail of a photo: http://10.5.5.9/gp/gpMediaMetadata?p=100GOPRO/GOPR1147.JPG

// http://www.golangprograms.com/golang-download-image-from-given-url.html
