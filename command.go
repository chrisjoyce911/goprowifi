package goprowifi

import (
	"errors"
	"fmt"
)

// sendCommand basic handler to send all commands
func (g *Client) sendCommand(commandURL string) (httpStatus int, err error) {

	bodyBytes, httpStatus, err := g.request(commandURL)

	if err != nil {
		err = fmt.Errorf("Can't unmarshal retrieve response %s %s", err, string(bodyBytes))
	}

	return
}

// CaptureStart operate the shutter Shutter
// Shutter
// Trigger: http://10.5.5.9/gp/gpControl/command/shutter?p=1
func (g *Client) CaptureStart() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/shutter?p=1")
	if status != 200 {
		err = errors.New("error starting starting capture")
	}
	return
}

// CaptureStop operate the shutter Shutter
// Shutter
// Stop (Video/Timelapse): http://10.5.5.9/gp/gpControl/command/shutter?p=0
func (g *Client) CaptureStop() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/shutter?p=0")
	if status != 200 {
		err = errors.New("error stopping starting capture")
	}
	return
}

// ModeVideo sets mode to Video (VIDEO)
// Video (VIDEO): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=0
func (g *Client) ModeVideo() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=0")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModeVideoTimeLapse sets mode to TimeLapse Video (VIDEO)
// TimeLapse Video (VIDEO): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=1
func (g *Client) ModeVideoTimeLapse() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=1")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModeVideoPhoto sets mode to Video + Photo (VIDEO)
// Video + Photo (VIDEO): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=2
func (g *Client) ModeVideoPhoto() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=2")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModeVideoLooping sets mode to Looping (VIDEO)
// Looping (VIDEO): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=3
func (g *Client) ModeVideoLooping() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=0&sub_mode=3")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModePhotoSingle sets mode to Single (PHOTO)
// Single (PHOTO): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=1&sub_mode=1
func (g *Client) ModePhotoSingle() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=1&sub_mode=1")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModePhotoNight sets mode to Night (PHOTO)
// Night (PHOTO): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=1&sub_mode=2
func (g *Client) ModePhotoNight() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=1&sub_mode=2")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModeMultiShotBurst sets mode to Burst (MultiShot)
// Burst (MultiShot): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=2&sub_mode=0
func (g *Client) ModeMultiShotBurst() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=2&sub_mode=0")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModeMultiShotTimelapse sets mode to Timelapse (MultiShot)
// Timelapse (MultiShot): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=2&sub_mode=1
func (g *Client) ModeMultiShotTimelapse() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=2&sub_mode=1")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// ModeMultiShotNightLapse sets mode to NightLapse (MultiShot)
// NightLapse (MultiShot): http://10.5.5.9/gp/gpControl/command/sub_mode?mode=2&sub_mode=2
func (g *Client) ModeMultiShotNightLapse() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/sub_mode?mode=2&sub_mode=2")
	if status != 200 {
		err = errors.New("error changing operation modes")
	}
	return
}

// TagMoment add highlight tag
// Sends tag command: http://10.5.5.9/gp/gpControl/command/storage/tag_moment
func (g *Client) TagMoment() (err error) {
	status, err := g.sendCommand("http://10.5.5.9/gp/gpControl/command/storage/tag_moment")
	if status != 200 {
		err = errors.New("error tagging highlight")
	}
	return
}

// http://10.5.5.9/gp/gpControl/command/mode?p=3

// 				{
// 					"version": "5.00",
// 					"path_info": "command/mode",
// 					"query_string": "p=3",
// 					"error_code": -1,
// 					"error_msg": "500 Internal Server Error\r\n",
// 					"function": "gpcontrol_cgi_handler",
// 					"line": 186
// 				}
