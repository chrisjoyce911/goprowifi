package goprowifi

import (
	"encoding/json"
	"fmt"
)

// PhotoInformation Photo Information
type PhotoInformation struct {
	EpochTimestamp string `json:"cre"`
	// EIS enabled
	Eis   string `json:"eis"`
	Gumi  string `json:"gumi"`
	Hight string `json:"h"`
	Hc    string `json:"hc"`
	// RAW enabled
	Raw   string `json:"raw"`
	S     string `json:"s"`
	Tr    string `json:"tr"`
	Us    string `json:"us"`
	Width string `json:"w"`
	// WDR enabled
	Wdr string `json:"wdr"`
}

// GetPhotoInformation Extended Video Information
//
// http://10.5.5.9/gp/gpMediaMetadata?p=100GOPRO/GOPR1147.JPG&t=v4info
func (g *Client) GetPhotoInformation(folder string, file string) (reply PhotoInformation, httpStatus int, err error) {

	url := fmt.Sprintf("http://10.5.5.9:8080/gp/gpMediaMetadata?p=%s/%s&t=v4info", folder, file)
	bodyBytes, httpStatus, err := g.request(url)

	err = json.Unmarshal(bodyBytes, &reply)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	return
}

// PhotoEXIFInformation Photo EXIF Information
// Photo EXIF Information
// http://10.5.5.9/gp/gpMediaMetadata?p=/100GOPRO/GOPR1147.JPG&t=exif
type PhotoEXIFInformation struct {
	Exif Exif `json:"exif"`
}

//Exif ...
type Exif struct {
	ApertureValue            string `json:"ApertureValue"`
	ColorSpace               string `json:"ColorSpace"`
	ComponentsConfiguration  string `json:"ComponentsConfiguration"`
	CompressedBitsPerPixel   string `json:"CompressedBitsPerPixel"`
	Compression              string `json:"Compression"`
	Contrast                 string `json:"Contrast"`
	CustomRendered           string `json:"CustomRendered"`
	DateTime                 string `json:"DateTime"`
	DateTimeDigitized        string `json:"DateTimeDigitized"`
	DateTimeOriginal         string `json:"DateTimeOriginal"`
	DeviceSettingDescription string `json:"DeviceSettingDescription"`
	DigitalZoomRatio         string `json:"DigitalZoomRatio"`
	ExifVersion              string `json:"ExifVersion"`
	ExposureBiasValue        string `json:"ExposureBiasValue"`
	ExposureIndex            string `json:"ExposureIndex"`
	ExposureMode             string `json:"ExposureMode"`
	ExposureProgram          string `json:"ExposureProgram"`
	ExposureTime             string `json:"ExposureTime"`
	FNumber                  string `json:"FNumber"`
	FileSource               string `json:"FileSource"`
	Flash                    string `json:"Flash"`
	FlashPixVersion          string `json:"FlashPixVersion"`
	FocalLength              string `json:"FocalLength"`
	FocalLengthIn35mmFilm    string `json:"FocalLengthIn35mmFilm"`
	GPSAltitude              string `json:"GPSAltitude"`
	GPSAltitudeRef           string `json:"GPSAltitudeRef"`
	GPSDateStamp             string `json:"GPSDateStamp"`
	GPSLatitude              string `json:"GPSLatitude"`
	GPSLatitudeRef           string `json:"GPSLatitudeRef"`
	GPSLongitude             string `json:"GPSLongitude"`
	GPSLongitudeRef          string `json:"GPSLongitudeRef"`
	GPSTimeStamp             string `json:"GPSTimeStamp"`
	GainControl              string `json:"GainControl"`
	ISOSpeedRatings          string `json:"ISOSpeedRatings"`
	ImageDescription         string `json:"ImageDescription"`
	InteroperabilityIndex    string `json:"InteroperabilityIndex"`
	InteroperabilityVersion  string `json:"InteroperabilityVersion"`
	LightSource              string `json:"LightSource"`
	Make                     string `json:"Make"`
	MakerNote                string `json:"MakerNote"`
	MaxApertureValue         string `json:"MaxApertureValue"`
	MeteringMode             string `json:"MeteringMode"`
	Model                    string `json:"Model"`
	Orientation              string `json:"Orientation"`
	PixelXDimension          string `json:"PixelXDimension"`
	PixelYDimension          string `json:"PixelYDimension"`
	ResolutionUnit           string `json:"ResolutionUnit"`
	Saturation               string `json:"Saturation"`
	SceneCaptureType         string `json:"SceneCaptureType"`
	SceneType                string `json:"SceneType"`
	SensingMethod            string `json:"SensingMethod"`
	Sharpness                string `json:"Sharpness"`
	ShutterSpeedValue        string `json:"ShutterSpeedValue"`
	Software                 string `json:"Software"`
	SubjectDistance          string `json:"SubjectDistance"`
	SubjectDistanceRange     string `json:"SubjectDistanceRange"`
	WhiteBalance             string `json:"WhiteBalance"`
	XResolution              string `json:"XResolution"`
	YCbCrPositioning         string `json:"YCbCrPositioning"`
	YResolution              string `json:"YResolution"`
}

// GetPhotoEXIFInformation Photo EXIF Information
//
// http://10.5.5.9/gp/gpMediaMetadata?p=100GOPRO/GOPR1147.JPG&t=exif
func (g *Client) GetPhotoEXIFInformation(folder string, file string) (reply PhotoEXIFInformation, httpStatus int, err error) {

	url := fmt.Sprintf("http://10.5.5.9:8080/gp/gpMediaMetadata?p=%s/%s&t=exif", folder, file)
	bodyBytes, httpStatus, err := g.request(url)

	err = json.Unmarshal(bodyBytes, &reply)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	return
}
