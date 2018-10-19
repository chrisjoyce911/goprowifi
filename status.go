package goprowifi

import (
	"encoding/json"
	"fmt"
)

/*

----------------------
Settings object:
Video Mode based parameters:
Current SUB-Mode Video - 68:

Video - 0
TimeLapse Video - 1
Video+Photo - 2
Looping - 3
Video Resolution - 2:

4K SuperView - 2
4K - 1
2.7K SuperView - 5
2.7K - 4
2.7K 4:3 - 6
1440 - 7
1080 SuperView - 8
1080 - 9
960 - 10
720 SuperView - 11
720 - 12
WVGA - 13
Frame Rate - 3:

240 - 0
120 - 1
100 - 2
90 - 3
80 - 4
60 - 5
50 - 6
48 - 7
30 - 8
25 - 9
24 - 10
15 - 11
12.5 - 12
FOV video - 4:

Wide - 0
Medium - 1
Narrow - 2
Linear - 4
TimeLapse Video Interval - 5:

1/2 Seconds - 0
1 Second - 1
2 Seconds - 2
5 Seconds - 3
10 Seconds - 4
30 Seconds - 5
60 Seconds - 6
Looping Video Interval - 6:

Max - 0
5 Minutes - 1
20 Minutes - 2
60 Minutes - 3
120 Minutes - 4
Photo+Video Interval - 7:

1 Photo / 5 Seconds - 1
1 Photo / 10 Seconds - 2
1 Photo / 30 Seconds - 3
1 Photo / 60 Seconds - 4
Low Light - 8:

ON - 1
OFF - 0
Spot Meter - 9:

ON - 1
OFF - 0
Protune - 10:

ON - 1
OFF - 0
White Balance - 11:

Auto - 0
3000K - 1
4000K - 5
4800K - 6
5500K - 2
6000K - 7
6500K - 3
Native - 4
Color - 12:

GoPro Color - 0
Flat - 1
Manual Exposure (v4.00FW) - 73:

1/24 - 3
1/25 - 4
1/30 - 5
1/48 - 6
1/50 - 7
1/60 - 8
1/80 - 9
1/90 - 10
1/96 - 11
1/100 - 12
1/120 - 13
1/160 - 14
1/180 - 15
1/192 - 16
1/200 - 17
1/240 - 18
1/320 - 19
1/360 - 20
1/400 - 21
1/480 - 22
1/960 - 23
ISO Mode (v4.00FW) - 74:

Max - 0
Lock - 1
ISO Limit - 13:

6400 - 0
1600 - 1
400 - 2
3200 - 3
800 - 4
200 - 7
100 - 8
Sharpness - 14:

High - 0
Medium - 1
Low - 2
ON (HERO4 Session) - 3
OFF (HERO4 Session) - 4
EV Comp - 15:

-2.0 - 8
-1.5 - 7
-1.0 - 6
-0.5 - 5
0.0 - 4
0.5 - 3
1.0 - 2
1.5 - 1
2.0 - 0
Photo Mode based parameters:
Photo Sub Mode - 69:

Single - 0
Continuous - 1
Night - 2
Continuous Rate - 18:

3 Frames / Second - 0
5 Frames / Second - 1
10 Frames / Second - 2
Megapixels - 17:

12MP Wide - 0
7MP Wide - 1
7MP Med - 2
5MP Med - 3
Shutter - 19:

Auto - 0
2 Seconds - 1
5 Seconds - 2
10 Seconds - 3
15 Seconds - 4
20 Seconds - 5
30 Seconds - 6
Spot Meter - 20:

ON - 1
OFF - 0
Protune - 21:

ON - 1
OFF - 0
White Balance - 22:

Auto - 0
3000K - 1
4000K - 5
4800K - 6
5500K - 2
6000K - 7
6500K - 3
Native - 4
Color - 23:

GoPro Color - 0
Flat - 1
Sharpness - 25:

High - 0
Medium - 1
Low - 2
EV Comp - 26:

-2.0 - 8
-1.5 - 7
-1.0 - 6
-0.5 - 5
0.0 - 4
0.5 - 3
1.0 - 2
1.5 - 1
2.0 - 0
ISO Min (4.00FW) - 75:

800 - 0
400 - 1
200 - 2
100 - 3
ISO Limit (ISO Max in 4.00FW) - 24:

800 - 0
400 - 1
200 - 2
100 - 3
MultiShot Mode based parameters:
Default Multi-Shot Sub Mode - 27:

Burst - 0
Time Lapse - 1
Night Lapse - 2
Multi-Shot Sub Mode - 70:

Burst - 0
Time Lapse - 1
Night Lapse - 2
NightLapse/NightPhoto Shutter Exposure - 31:

Auto - 0
2 Seconds - 1
5 Seconds - 2
10 Seconds - 3
15 Seconds - 4
20 Seconds - 5
30 Seconds - 6
Burst Rate - 29:

3 Photos / 1 Second - 0
5 Photos / 1 Second - 1
10 Photos / 1 Second - 2
10 Photos / 2 Seconds - 3
10 Photos / 3 Seconds - 4
30 Photos / 1 Second - 5
30 Photos / 2 Seconds - 6
30 Photos / 3 Seconds - 7
30 Photos / 6 Seconds - 8
TimeLapse Interval - 30:

1 Photo / 0.5 Sec - 0
1 Photo / 1 Sec - 1
1 Photo / 2 Sec - 2
1 Photo / 5 Sec - 5
1 Photo / 10 Sec - 10
1 Photo / 30 Sec - 30
1 Photo / 60 Sec - 60
NightLapse Interval - 32:

Continuous - 0
4 Seconds - 4
5 Seconds - 5
10 Seconds - 10
15 Seconds - 15
20 Seconds - 20
30 Seconds - 30
1 Minute - 60
2 Minutes - 120
5 Minutes - 300
30 Minutes - 1800
60 Minutes - 3600
Megapixels - 28:

12MP Wide - 0
7MP Wide - 1
7MP Med - 2
5MP Med - 3
Spot Meter - 33:

ON - 1
OFF - 0
Protune - 34:

ON - 1
OFF - 0
White Balance - 35:

Auto - 0
3000K - 1
4000K - 5
4800K - 6
5500K - 2
6000K - 7
6500K - 3
Native - 4
Color - 36:

GoPro Color - 0
Flat - 1
Sharpness - 38:

High - 0
Medium - 1
Low - 2
EV Comp - 39:

-2.0 - 8
-1.5 - 7
-1.0 - 6
-0.5 - 5
0.0 - 4
0.5 - 3
1.0 - 2
1.5 - 1
2.0 - 0
ISO Min (4.00FW) - 76:

800 - 0
400 - 1
200 - 2
100 - 3
ISO Limit (ISO Max in 4.00FW) - 37:

800 - 0
400 - 1
200 - 2
100 - 3
Other Settings:
LCD Display - 72:

ON - 1

OFF - 0

Orientation - 52:

Auto - 0
UP - 1
DOWN - 2
Default Boot Mode - 53:

Video - 0
Photo - 1
MultiShot - 2
Quick Capture - 54:

OFF - 0
ON - 1
LED status - 55:

OFF - 0
2 LEDs (H4 Black/Silver) FULL LEDs (HERO4 Session) - 1
4 LEDs (H4 Black/Silver)- 2
Volume for beeps - 56:

	Mute - 2
	70% - 1
	100% - 0
	Video Format - 57:

	NTSC - 0
	PAL - 1

On Screen data (HDMI/LCD BacPac/LCD) - 58:

	ON - 1
	OFF - 0

LCD Brightness - 49:

	High - 0
	Medium - 1
	Low - 2
	LCD Lock - 50:

	On - 1
	Off - 0

LCD Timeout sleep - 51:

	Never sleep - 0
	1min sleep - 1
	2min sleep - 2
	3min sleep - 3
	Auto Power Off - 59:

	Never - 0
	1 min - 1
	2 mins - 2
	3 mins - 3
	5 mins - 4
*/

//CameraStatus current status and settings of the camera
type CameraStatus struct {
	Status   CurrentStatus   `json:"status"`
	Settings CurrentSettings `json:"settings"`
}

// CurrentStatus ...
type CurrentStatus struct {
	// 1 - Internal Battery is available:
	// 	0 = No Battery
	// 	1 = Battery is available
	InternalBattery int `json:"1"`

	// 2 - Internal Battery Level:
	// 	4 = Charging
	// 	3 = Full
	// 	2 = Halfway
	// 	1 = Low
	// 	0 = Empty
	InternalBatteryLevel int `json:"2"`

	Num3  int `json:"3"`
	Num4  int `json:"4"`
	Num6  int `json:"6"`
	Num8  int `json:"8"`
	Num9  int `json:"9"`
	Num10 int `json:"10"`
	Num11 int `json:"11"`

	// 13 - Current Recording Video Duration
	CurrentRecordingVideoDuration int `json:"13"`

	Num14 int    `json:"14"`
	Num15 int    `json:"15"`
	Num16 int    `json:"16"`
	Num17 int    `json:"17"`
	Num19 int    `json:"19"`
	Num20 int    `json:"20"`
	Num21 int    `json:"21"`
	Num22 int    `json:"22"`
	Num23 int    `json:"23"`
	Num24 int    `json:"24"`
	Num26 int    `json:"26"`
	Num27 int    `json:"27"`
	Num28 int    `json:"28"`
	Num29 string `json:"29"`

	// 30 - WiFi SSID
	WiFiSSID string `json:"30"`

	// 31 - Number of clients connected to the camera
	NumberOfClients int `json:"31"`

	// 32 - Streaming feed status:
	// 	0 = Not Streaming
	// 	1 = Streaming
	StreamingFeed int `json:"32"`

	// 33 - SD card inserted:
	// 	0 = SD card inserted
	// 	2 = No SD Card present
	SDcardInserted int `json:"33"`

	// 34 - Remaining Photos
	RemainingPhotos int `json:"34"`

	// 35 - Remaining Video Time
	RemainingVideoTime int `json:"35"`

	// 36 - Number of Batch Photos taken (Example: TimeLapse batches, burst batches, continouous photo batches...)
	NumberOfBatchPhotos int `json:"36"`

	// 37 - Number of videos shot
	NumberOfVideos int `json:"37"`

	// 38 - Number of ALL photos taken
	Num38 int `json:"38"`

	// 39 - Number of MultiShot pictures taken
	// 39 - Number of ALL videos taken
	// 	8 = Recording/Processing status:
	// 	0 = Not recording/Processing
	// 	1 = Recording/processing
	Num39 int    `json:"39"`
	Num40 string `json:"40"`
	Num41 int    `json:"41"`
	Num42 int    `json:"42"`

	// 43 - Current Mode:
	// 	Video - 0
	// 	Photo - 1
	// 	MultiShot - 2
	CurrentMode int `json:"43"`

	// 44 - Current SubMode
	// 	0 = Video/Single Pic/Burst
	// 	1 = TL Video/Continuous/TimeLapse
	// 	2 = Video+Photo/NightPhoto/NightLapse
	CurrentSubMode int `json:"44"`

	Num45 int `json:"45"`
	Num46 int `json:"46"`
	Num47 int `json:"47"`
	Num48 int `json:"48"`
	Num49 int `json:"49"`

	// 54 - Remaning free space on memorycard in bytes
	RemaningFreeSpace int `json:"54"`

	Num55 int `json:"55"`
	Num56 int `json:"56"`
	Num57 int `json:"57"`
	Num58 int `json:"58"`
	Num59 int `json:"59"`
	Num60 int `json:"60"`
	Num61 int `json:"61"`
	Num62 int `json:"62"`
	Num63 int `json:"63"`
	Num64 int `json:"64"`
	Num65 int `json:"65"`
	Num66 int `json:"66"`
	Num67 int `json:"67"`
	Num68 int `json:"68"`
	Num69 int `json:"69"`
	Num70 int `json:"70"`
	Num71 int `json:"71"`
	Num72 int `json:"72"`
	Num73 int `json:"73"`
	Num74 int `json:"74"`
}

// CurrentSettings ...
type CurrentSettings struct {
	Num1 int `json:"1"`

	// 2 - Video Resolutions
	// 	1 = 4K
	// 	4 = 2.7K: http://10.5.5.9/gp/gpControl/setting/2/4
	// 	6 = 2.7K 4:3: http://10.5.5.9/gp/gpControl/setting/2/6
	// 	7 = 1440p: http://10.5.5.9/gp/gpControl/setting/2/7
	// 	9 = 1080p: http://10.5.5.9/gp/gpControl/setting/2/9
	// 	10 = 960p: http://10.5.5.9/gp/gpControl/setting/2/10
	// 	12 = 720p: http://10.5.5.9/gp/gpControl/setting/2/12
	// 	17 = WVGA: http://10.5.5.9/gp/gpControl/setting/2/17
	VideoResolutions int `json:"2"`

	// 3 - Frame Rate
	// 	0 = 240fps:	http://10.5.5.9/gp/gpControl/setting/3/0
	// 	1 = 120fps:	http://10.5.5.9/gp/gpControl/setting/3/1
	// 	2 = 100fps:	http://10.5.5.9/gp/gpControl/setting/3/2
	// 	3 = 90fps:	http://10.5.5.9/gp/gpControl/setting/3/3
	// 	4 = 80fps:	http://10.5.5.9/gp/gpControl/setting/3/4
	// 	5 = 60fps:	http://10.5.5.9/gp/gpControl/setting/3/5
	// 	6 = 50fps:	http://10.5.5.9/gp/gpControl/setting/3/6
	// 	7 = 48fps:	http://10.5.5.9/gp/gpControl/setting/3/7
	// 	8 = 30fps:	http://10.5.5.9/gp/gpControl/setting/3/8
	// 	9 = 25fps:	http://10.5.5.9/gp/gpControl/setting/3/9
	FrameRate int `json:"3"`

	// 4 - Field of View
	// 	0 = Wide: http://10.5.5.9/gp/gpControl/setting/4/0
	// 	1 = Medium: http://10.5.5.9/gp/gpControl/setting/4/1
	// 	2 = Narrow: http://10.5.5.9/gp/gpControl/setting/4/2
	// 	3 = SuperView: http://10.5.5.9/gp/gpControl/4/3
	// 	4 = Linear: http://10.5.5.9/gp/gpControl/setting/4/4
	FOV int `json:"4"`

	// 5 - Video Timelapse Interval:
	// 	0 = 0.5: http://10.5.5.9/gp/gpControl/setting/5/0
	// 	1 = 1: http://10.5.5.9/gp/gpControl/setting/5/1
	// 	2 = 2: http://10.5.5.9/gp/gpControl/setting/5/2
	// 	3 = 5: http://10.5.5.9/gp/gpControl/setting/5/3
	// 	4 = 10: http://10.5.5.9/gp/gpControl/setting/5/4
	// 	5 = 30: http://10.5.5.9/gp/gpControl/setting/5/5
	// 	6 = 60: http://10.5.5.9/gp/gpControl/setting/5/6
	VideoTimelapseInterval int `json:"5"`

	// 6 - Video Looping Duration:
	// 	0 = Max: http://10.5.5.9/gp/gpControl/setting/6/0
	// 	1 = 5Min: http://10.5.5.9/gp/gpControl/setting/6/1
	// 	2 = 20Min: http://10.5.5.9/gp/gpControl/setting/6/2
	// 	3 = 60Min: http://10.5.5.9/gp/gpControl/setting/6/3
	// 	4 = 120Min: http://10.5.5.9/gp/gpControl/setting/6/4
	VideoLoopingDuration int `json:"6"`

	// 7 - Video+Photo Interval:
	// 	1 = 5: http://10.5.5.9/gp/gpControl/setting/7/1
	// 	2 = 10: http://10.5.5.9/gp/gpControl/setting/7/2
	// 	3 = 30: http://10.5.5.9/gp/gpControl/setting/7/3
	// 	4 = 60Min: http://10.5.5.9/gp/gpControl/setting/7/4
	VideoPhotoInterval int `json:"7"`

	// 8 - Low Light
	// 	0 = OFF: http://10.5.5.9/gp/gpControl/setting/8/0
	// 	1 = ON: http://10.5.5.9/gp/gpControl/setting/8/1
	LowLight int `json:"8"`

	// 9 - Spot Meter:
	// 	0 = off: http://10.5.5.9/gp/gpControl/setting/9/0
	// 	1 = on: http://10.5.5.9/gp/gpControl/setting/9/1
	SpotMeter int `json:"9"`

	// https://github.com/KonradIT/goprowifihack/blob/master/HERO5/HERO5-Commands.md
	Num10 int `json:"10"`
	Num11 int `json:"11"`
	Num12 int `json:"12"`
	Num13 int `json:"13"`
	Num14 int `json:"14"`
	Num15 int `json:"15"`
	Num16 int `json:"16"`
	Num17 int `json:"17"`
	Num18 int `json:"18"`
	Num19 int `json:"19"`
	Num20 int `json:"20"`
	Num21 int `json:"21"`
	Num22 int `json:"22"`
	Num23 int `json:"23"`
	Num24 int `json:"24"`
	Num25 int `json:"25"`
	Num26 int `json:"26"`
	Num27 int `json:"27"`
	Num28 int `json:"28"`
	Num29 int `json:"29"`
	Num30 int `json:"30"`
	Num31 int `json:"31"`
	Num32 int `json:"32"`
	Num33 int `json:"33"`
	Num34 int `json:"34"`
	Num35 int `json:"35"`
	Num36 int `json:"36"`
	Num37 int `json:"37"`
	Num38 int `json:"38"`
	Num39 int `json:"39"`
	Num40 int `json:"40"`
	Num41 int `json:"41"`
	Num42 int `json:"42"`
	Num43 int `json:"43"`
	Num44 int `json:"44"`
	Num45 int `json:"45"`
	Num46 int `json:"46"`
	Num47 int `json:"47"`
	Num48 int `json:"48"`
	Num50 int `json:"50"`
	Num51 int `json:"51"`
	Num52 int `json:"52"`
	Num54 int `json:"54"`
	Num57 int `json:"57"`
	Num58 int `json:"58"`
	Num59 int `json:"59"`
	Num60 int `json:"60"`
	Num61 int `json:"61"`
	Num62 int `json:"62"`
	Num63 int `json:"63"`
	Num64 int `json:"64"`
	Num65 int `json:"65"`
	Num66 int `json:"66"`
	Num67 int `json:"67"`
	Num68 int `json:"68"`
	Num69 int `json:"69"`
	Num70 int `json:"70"`
	Num71 int `json:"71"`
	Num72 int `json:"72"`
	Num73 int `json:"73"`
	Num74 int `json:"74"`
	Num75 int `json:"75"`
	Num76 int `json:"76"`
	Num77 int `json:"77"`
	Num78 int `json:"78"`
	Num79 int `json:"79"`
	Num80 int `json:"80"`
	Num81 int `json:"81"`
	Num82 int `json:"82"`
	Num83 int `json:"83"`
	Num84 int `json:"84"`
	Num85 int `json:"85"`
	Num86 int `json:"86"`
	Num87 int `json:"87"`
	Num88 int `json:"88"`
	Num89 int `json:"89"`
	Num91 int `json:"91"`
	Num92 int `json:"92"`
	Num93 int `json:"93"`
	Num94 int `json:"94"`
	Num95 int `json:"95"`
	Num96 int `json:"96"`
	Num97 int `json:"97"`
	Num98 int `json:"98"`
	Num99 int `json:"99"`
}

// GetCameraStatus ...
// This URL contains a JSON with the camera parameters:
//
// http://10.5.5.9/gp/gpControl/status
func (g *Client) GetCameraStatus() (status CurrentStatus, httpStatus int, err error) {

	bodyBytes, httpStatus, err := g.request("http://10.5.5.9/gp/gpControl/status")

	reply := new(CameraStatus)
	err = json.Unmarshal(bodyBytes, &reply)
	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s", err)
	}

	status = reply.Status
	fmt.Printf("[cameraStatus] WiFiSSID : %s\n", status.WiFiSSID)

	return
}
