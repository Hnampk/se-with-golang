package solid

type Drone struct {
}

type Vec3 struct {
}

type Target struct {
}

type RGBA struct {
}

// NavigateTo applies any required changes to the drone's speed
// vector so that its eventual position matches dst.
func (d *Drone) NavigateTo(dst Vec3) error {
	//...

	return nil
}

// Position returns the current drone position vector.
func (d *Drone) Position() Vec3 {
	//...

	return Vec3{}
}

// Position returns the current drone speed vector.
func (d *Drone) Speed() Vec3 {
	//...
	return Vec3{}
}

// DetectTargets captures an image of the drone's field of view (FoV) using
// the on-board camera and feeds it to a pre-trained SSD MobileNet V1
// neural
// network to detect and classify interesting nearby targets. For more info
// on this model see:
//
// https://github.com/tensorflow/models/tree/master/research/object_detection
// func (d *Drone) DetectTargets() ([]*Target, error) {
// 	//...
// 	return nil, nil
// }

/* ==============
===== FIXED =====
- better move the above DetectTargets() as function of type MobileNet instead, to keep the Single Responsible Principle
- define MobileNet type, which contains the implementation for our target detector
=================
============== */

// CaptureImage records and returns an image of the drone's field of
// view using the on-board drone camera.
func (d *Drone) CaptureImage() (*RGBA, error) {
	//...

	return nil, nil
}

// MobileNet performs target detection for drones using the
// SSD MobileNet V1 NN.
// For more info on this model see:
//
// https://github.com/tensorflow/models/tree/master/research/object_detection
type MobileNet struct {
	// various attributes...
}

// DetectTargets captures an image of the drone's field of view and feeds
// it to a neural network to detect and classify interesting nearby
// targets.
func (mn *MobileNet) DetectTargets(d *Drone) ([]*Target, error) {

	return nil, nil
}
