package goprowifi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_getMediaList(t *testing.T) {
	type fields struct {
		myDoer Doer
	}

	tests := []struct {
		name           string
		fields         fields
		wantReply      []*Media
		wantHTTPStatus int
		wantErr        error
	}{
		{
			name: "Hero5 OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"id": "3806895087272624399",
						"media": [
							{
								"d": "100GOPRO",
								"fs": [
									{
										"n": "GOPR1130.MP4",
										"mod": "1539705058",
										"ls": "11170499",
										"s": "188100018"
									},
									{
										"n": "GOPR1147.JPG",
										"mod": "1539960090",
										"s": "3843146"
									}
								]
							}
						]
					}`,
				},
			},
			wantReply: []*Media{
				{
					Device:         "100GOPRO",
					EpochTimestamp: "1539705058",
					FileName:       "GOPR1130.MP4",
					MediaType:      "MP4",
					Ls:             "11170499",
					S:              "188100018",
				},
				{
					Device:         "100GOPRO",
					EpochTimestamp: "1539960090",
					FileName:       "GOPR1147.JPG",
					MediaType:      "JPG",
					Ls:             "",
					S:              "3843146",
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
			gotReply, gotHTTPStatus, err := hero.getMediaList()

			if err == nil {
				assert.Equal(t, tt.wantHTTPStatus, gotHTTPStatus, "HTTPStatus did not match")
				assert.Equal(t, tt.wantReply, gotReply, "Reply did not match")
			}

		})
	}
}
