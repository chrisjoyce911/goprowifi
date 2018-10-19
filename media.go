package goprowifi

// MediaList ist of the contents of the SD card
type MediaList struct {
	ID    string `json:"id"`
	Media []struct {
		Device string `json:"d"`
		Files  []struct {
			Ls             string `json:"ls"`
			EpochTimestamp string `json:"mod"`
			FileName       string `json:"n"`
			S              string `json:"s"`
		} `json:"fs"`
	} `json:"media"`
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
