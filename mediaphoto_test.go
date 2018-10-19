package goprowifi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetPhotoInformation(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		folder string
		file   string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantReply      PhotoInformation
		wantHTTPStatus int
		wantErr        error
	}{
		{
			name: "Hero5 OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"cre": "1539960075",
						"s": "3843146",
						"hc": "0",
						"us": "0",
						"eis": "0",
						"wdr": "0",
						"raw": "0",
						"tr": "0",
						"gumi": "96f1af231bed1a6e5e81fb46f0738678",
						"w": "4000",
						"h": "3000"
					}`,
				},
			},
			args: args{
				folder: "100GOPRO",
				file:   "GOPR1147.JPG",
			},
			wantReply: PhotoInformation{
				EpochTimestamp: "1539960075",
				Eis:            "0",
				Gumi:           "96f1af231bed1a6e5e81fb46f0738678",
				Hight:          "3000",
				Hc:             "0",
				Raw:            "0",
				S:              "3843146",
				Tr:             "0",
				Us:             "0",
				Width:          "4000",
				Wdr:            "0",
			},
			wantHTTPStatus: 200,
			wantErr:        nil,
			// wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hero := CreateTestRestClient(tt.fields.myDoer)
			gotReply, gotHTTPStatus, err := hero.GetPhotoInformation(tt.args.folder, tt.args.file)

			if err == nil {
				assert.Equal(t, tt.wantHTTPStatus, gotHTTPStatus, "HTTPStatus did not match")
				assert.Equal(t, tt.wantReply, gotReply, "Reply did not match")
			}

		})
	}
}

func TestClient_GetPhotoEXIFInformation(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		folder string
		file   string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantReply      PhotoEXIFInformation
		wantHTTPStatus int
		wantErr        error
	}{
		{
			name: "Hero5 OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"exif": {
							"ImageDescription": "DCIM\\100GOPROGOPR1147.JPG",
							"Make": "GoPro",
							"Model": "HERO5 Black",
							"Orientation": "Top-left",
							"XResolution": "72",
							"YResolution": "72",
							"ResolutionUnit": "Inch",
							"Software": "HD5.02.02.60.00",
							"DateTime": "2018:10:19 14:41:31",
							"YCbCrPositioning": "Centered",
							"Compression": "JPEG compression",
							"ExposureTime": "1/100 sec.",
							"FNumber": "f/2.8",
							"ExposureProgram": "Normal program",
							"ISOSpeedRatings": "530",
							"ExifVersion": "Exif Version 2.21",
							"DateTimeOriginal": "2018:10:19 14:41:31",
							"DateTimeDigitized": "2018:10:19 14:41:31",
							"ComponentsConfiguration": "Y Cb Cr -",
							"CompressedBitsPerPixel": "14461165.980",
							"ShutterSpeedValue": "6.64 EV (1/99 sec.)",
							"ApertureValue": "2.97 EV (f/2.8)",
							"ExposureBiasValue": "0.00 EV",
							"MaxApertureValue": "2.97 EV (f/2.8)",
							"SubjectDistance": "0.0 m",
							"MeteringMode": "Unknown",
							"LightSource": "Unknown",
							"Flash": "No flash function",
							"FocalLength": "3.0 mm",
							"MakerNote": "0d0011400180864500004c4145383034323931363630313337300000000000000000000000000000000000000000000000000000000000000000efefefefefef",
							"FlashPixVersion": "FlashPix Version 1.0",
							"ColorSpace": "sRGB",
							"PixelXDimension": "4000",
							"PixelYDimension": "3000",
							"ExposureIndex": "0/0",
							"SensingMethod": "One-chip color area sensor",
							"FileSource": "DSC",
							"SceneType": "Directly photographed",
							"CustomRendered": "Normal process",
							"ExposureMode": "Auto exposure",
							"WhiteBalance": "Auto white balance",
							"DigitalZoomRatio": "1.000",
							"FocalLengthIn35mmFilm": "15",
							"SceneCaptureType": "Landscape",
							"GainControl": "Normal",
							"Contrast": "Normal",
							"Saturation": "Normal",
							"Sharpness": "Hard",
							"DeviceSettingDescription": "4 bytes undefined data",
							"SubjectDistanceRange": "Unknown",
							"GPSLatitudeRef": "S",
							"GPSLatitude": "27,  2, 31.3598400",
							"GPSLongitudeRef": "E",
							"GPSLongitude": "153,  8, 39.0991199",
							"GPSAltitudeRef": "Sea level",
							"GPSAltitude": "59.616",
							"GPSTimeStamp": "03:41:19.00",
							"GPSDateStamp": "2018:10:19",
							"InteroperabilityIndex": "R98",
							"InteroperabilityVersion": "0100"
						}
					}`,
				},
			},
			args: args{
				folder: "100GOPRO",
				file:   "GOPR1147.JPG",
			},
			wantReply: PhotoEXIFInformation{
				Exif{
					ApertureValue:            "2.97 EV (f/2.8)",
					ColorSpace:               "sRGB",
					ComponentsConfiguration:  "Y Cb Cr -",
					CompressedBitsPerPixel:   "14461165.980",
					Compression:              "JPEG compression",
					Contrast:                 "Normal",
					CustomRendered:           "Normal process",
					DateTime:                 "2018:10:19 14:41:31",
					DateTimeDigitized:        "2018:10:19 14:41:31",
					DateTimeOriginal:         "2018:10:19 14:41:31",
					DeviceSettingDescription: "4 bytes undefined data",
					DigitalZoomRatio:         "1.000",
					ExifVersion:              "Exif Version 2.21",
					ExposureBiasValue:        "0.00 EV",
					ExposureIndex:            "0/0",
					ExposureMode:             "Auto exposure",
					ExposureProgram:          "Normal program",
					ExposureTime:             "1/100 sec.",
					FNumber:                  "f/2.8",
					FileSource:               "DSC",
					Flash:                    "No flash function",
					FlashPixVersion:          "FlashPix Version 1.0",
					FocalLength:              "3.0 mm",
					FocalLengthIn35mmFilm:    "15",
					GPSAltitude:              "59.616",
					GPSAltitudeRef:           "Sea level",
					GPSDateStamp:             "2018:10:19",
					GPSLatitude:              "27,  2, 31.3598400",
					GPSLatitudeRef:           "S",
					GPSLongitude:             "153,  8, 39.0991199",
					GPSLongitudeRef:          "E",
					GPSTimeStamp:             "03:41:19.00",
					GainControl:              "Normal",
					ISOSpeedRatings:          "530",
					ImageDescription:         "DCIM\\100GOPROGOPR1147.JPG",
					InteroperabilityIndex:    "R98",
					InteroperabilityVersion:  "0100",
					LightSource:              "Unknown",
					Make:                     "GoPro",
					MakerNote:                "0d0011400180864500004c4145383034323931363630313337300000000000000000000000000000000000000000000000000000000000000000efefefefefef",
					MaxApertureValue:         "2.97 EV (f/2.8)",
					MeteringMode:             "Unknown",
					Model:                    "HERO5 Black",
					Orientation:              "Top-left",
					PixelXDimension:          "4000",
					PixelYDimension:          "3000",
					ResolutionUnit:           "Inch",
					Saturation:               "Normal",
					SceneCaptureType:         "Landscape",
					SceneType:                "Directly photographed",
					SensingMethod:            "One-chip color area sensor",
					Sharpness:                "Hard",
					ShutterSpeedValue:        "6.64 EV (1/99 sec.)",
					Software:                 "HD5.02.02.60.00",
					SubjectDistance:          "0.0 m",
					SubjectDistanceRange:     "Unknown",
					WhiteBalance:             "Auto white balance",
					XResolution:              "72",
					YCbCrPositioning:         "Centered",
					YResolution:              "72",
				},
			},
			wantHTTPStatus: 200,
			wantErr:        nil,
			// wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hero := CreateTestRestClient(tt.fields.myDoer)
			gotReply, gotHTTPStatus, err := hero.GetPhotoEXIFInformation(tt.args.folder, tt.args.file)

			if err == nil {
				assert.Equal(t, tt.wantHTTPStatus, gotHTTPStatus, "HTTPStatus did not match")
				assert.Equal(t, tt.wantReply, gotReply, "Reply did not match")
			}

		})
	}
}
