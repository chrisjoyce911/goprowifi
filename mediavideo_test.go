package goprowifi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetVideoInformation(t *testing.T) {
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
		wantReply      VideoInformation
		wantHTTPStatus int
		wantErr        error
	}{
		{
			name: "Hero5 OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"dur": "16",
						"c": "0",
						"avc_profile": "100",
						"profile": "41",
						"tag_count": 5,
						"tags": [
							6160,
							9080,
							11040,
							12880,
							14560
						],
						"w": "1920",
						"h": "1080",
						"fps": "90000",
						"fps_denom": "3600",
						"prog": "1",
						"subsample": "0"
					}`,
				},
			},
			args: args{
				folder: "100GOPRO",
				file:   "GOPR1148.MP4",
			},
			wantReply: VideoInformation{
				AvcProfile:    "100",
				C:             "0",
				VideoDuration: "16",
				FPS:           "90000",
				FpsDenom:      "3600",
				Hight:         "1080",
				Profile:       "41",
				Prog:          "1",
				Subsample:     "0",
				TagCount:      5,
				Tags:          []int{6160, 9080, 11040, 12880, 14560},
				Width:         "1920",
			},
			wantHTTPStatus: 200,
			wantErr:        nil,
			// wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hero := CreateTestRestClient(tt.fields.myDoer)
			gotReply, gotHTTPStatus, err := hero.GetVideoInformation(tt.args.folder, tt.args.file)

			if err == nil {
				assert.Equal(t, tt.wantHTTPStatus, gotHTTPStatus, "HTTPStatus did not match")
				assert.Equal(t, tt.wantReply, gotReply, "Reply did not match")
			}

		})
	}
}

func TestClient_GetExtendedVideoInformation(t *testing.T) {
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
		wantReply      ExtendedVideoInformation
		wantHTTPStatus int
		wantErr        error
	}{
		{
			name: "Hero5 OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"cre": "1539971179",
						"s": "90425681",
						"us": "0",
						"eis": "1",
						"pta": "0",
						"ao": "wind",
						"tr": "0",
						"mp": "0",
						"gumi": "1c131d69f60531ffa4c4e049c4a91dd2",
						"ls": "5367405",
						"cl": "0",
						"hc": "5",
						"hi": [
							6160,
							9080,
							11040,
							12880,
							14560
						],
						"dur": "16",
						"w": "1920",
						"h": "1080",
						"fps": "90000",
						"fps_denom": "3600",
						"prog": "1",
						"subsample": "0"
					}`,
				},
			},
			args: args{
				folder: "100GOPRO",
				file:   "GOPR1148.MP4",
			},
			wantReply: ExtendedVideoInformation{
				Ao:             "wind",
				Cl:             "0",
				EpochTimestamp: "1539971179",
				VideoDuration:  "16",
				EIS:            "1",
				FPS:            "90000",
				FpsDenom:       "3600",
				Gumi:           "1c131d69f60531ffa4c4e049c4a91dd2",
				Hight:          "1080",
				TagCount:       "5",
				Tags:           []int{6160, 9080, 11040, 12880, 14560},
				Ls:             "5367405",
				Mp:             "0",
				Prog:           "1",
				Pta:            "0",
				S:              "90425681",
				Subsample:      "0",
				Tr:             "0",
				Us:             "0",
				Width:          "1920",
			},
			wantHTTPStatus: 200,
			wantErr:        nil,
			// wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hero := CreateTestRestClient(tt.fields.myDoer)
			gotReply, gotHTTPStatus, err := hero.GetExtendedVideoInformation(tt.args.folder, tt.args.file)

			if err == nil {
				assert.Equal(t, tt.wantHTTPStatus, gotHTTPStatus, "HTTPStatus did not match")
				assert.Equal(t, tt.wantReply, gotReply, "Reply did not match")
			}

		})
	}
}
