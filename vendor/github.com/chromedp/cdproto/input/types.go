package input

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Code generated by cdproto-gen. DO NOT EDIT.

// TouchPoint [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#type-TouchPoint
type TouchPoint struct {
	X                  float64 `json:"x"`                                     // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y                  float64 `json:"y"`                                     // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX            float64 `json:"radiusX,omitempty,omitzero"`            // X radius of the touch area (default: 1.0).
	RadiusY            float64 `json:"radiusY,omitempty,omitzero"`            // Y radius of the touch area (default: 1.0).
	RotationAngle      float64 `json:"rotationAngle,omitempty,omitzero"`      // Rotation angle (default: 0.0).
	Force              float64 `json:"force,omitempty,omitzero"`              // Force (default: 1.0).
	TangentialPressure float64 `json:"tangentialPressure,omitempty,omitzero"` // The normalized tangential pressure, which has a range of [-1,1] (default: 0).
	TiltX              float64 `json:"tiltX,omitempty,omitzero"`              // The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0)
	TiltY              float64 `json:"tiltY,omitempty,omitzero"`              // The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
	Twist              int64   `json:"twist,omitempty,omitzero"`              // The clockwise rotation of a pen stylus around its own major axis, in degrees in the range [0,359] (default: 0).
	ID                 float64 `json:"id,omitempty,omitzero"`                 // Identifier used to track touch sources between events, must be unique within an event.
}

// GestureSourceType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#type-GestureSourceType
type GestureSourceType string

// String returns the GestureSourceType as string value.
func (t GestureSourceType) String() string {
	return string(t)
}

// GestureSourceType values.
const (
	GestureDefault GestureSourceType = "default"
	GestureTouch   GestureSourceType = "touch"
	GestureMouse   GestureSourceType = "mouse"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *GestureSourceType) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch GestureSourceType(s) {
	case GestureDefault:
		*t = GestureDefault
	case GestureTouch:
		*t = GestureTouch
	case GestureMouse:
		*t = GestureMouse
	default:
		return fmt.Errorf("unknown GestureSourceType value: %v", s)
	}
	return nil
}

// MouseButton [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#type-MouseButton
type MouseButton string

// String returns the MouseButton as string value.
func (t MouseButton) String() string {
	return string(t)
}

// MouseButton values.
const (
	None    MouseButton = "none"
	Left    MouseButton = "left"
	Middle  MouseButton = "middle"
	Right   MouseButton = "right"
	Back    MouseButton = "back"
	Forward MouseButton = "forward"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *MouseButton) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch MouseButton(s) {
	case None:
		*t = None
	case Left:
		*t = Left
	case Middle:
		*t = Middle
	case Right:
		*t = Right
	case Back:
		*t = Back
	case Forward:
		*t = Forward
	default:
		return fmt.Errorf("unknown MouseButton value: %v", s)
	}
	return nil
}

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#type-TimeSinceEpoch
type TimeSinceEpoch time.Time

// Time returns the TimeSinceEpoch as time.Time value.
func (t TimeSinceEpoch) Time() time.Time {
	return time.Time(t)
}

// MarshalJSON satisfies [json.Marshaler].
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	v := float64(time.Time(t).UnixNano() / int64(time.Second))
	return strconv.AppendFloat(make([]byte, 0, 20), v, 'f', -1, 64), nil
}

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *TimeSinceEpoch) UnmarshalJSON(buf []byte) error {
	f, err := strconv.ParseFloat(string(buf), 64)
	if err != nil {
		return err
	}
	*t = TimeSinceEpoch(time.Unix(0, int64(f*float64(time.Second))))
	return nil
}

// DragDataItem [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#type-DragDataItem
type DragDataItem struct {
	MimeType string `json:"mimeType"`                   // Mime type of the dragged data.
	Data     string `json:"data"`                       // Depending of the value of mimeType, it contains the dragged link, text, HTML markup or any other data.
	Title    string `json:"title,omitempty,omitzero"`   // Title associated with a link. Only valid when mimeType == "text/uri-list".
	BaseURL  string `json:"baseURL,omitempty,omitzero"` // Stores the base URL for the contained markup. Only valid when mimeType == "text/html".
}

// DragData [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#type-DragData
type DragData struct {
	Items              []*DragDataItem `json:"items"`
	Files              []string        `json:"files,omitempty,omitzero"` // List of filenames that should be included when dropping
	DragOperationsMask int64           `json:"dragOperationsMask"`       // Bit field representing allowed drag operations. Copy = 1, Link = 2, Move = 16
}

// Modifier input key modifier type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchKeyEvent
type Modifier int64

// Int64 returns the Modifier as int64 value.
func (t Modifier) Int64() int64 {
	return int64(t)
}

// Modifier values.
const (
	ModifierNone  Modifier = 0
	ModifierAlt   Modifier = 1
	ModifierCtrl  Modifier = 2
	ModifierMeta  Modifier = 4
	ModifierShift Modifier = 8
)

// String satisfies the [fmt.Stringer] interface.
func (t Modifier) String() string {
	switch t {
	case ModifierNone:
		return "None"
	case ModifierAlt:
		return "Alt"
	case ModifierCtrl:
		return "Ctrl"
	case ModifierMeta:
		return "Meta"
	case ModifierShift:
		return "Shift"
	}
	return fmt.Sprintf("Modifier(%d)", t)
}

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *Modifier) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	switch Modifier(v) {
	case ModifierNone:
		*t = ModifierNone
	case ModifierAlt:
		*t = ModifierAlt
	case ModifierCtrl:
		*t = ModifierCtrl
	case ModifierMeta:
		*t = ModifierMeta
	case ModifierShift:
		*t = ModifierShift
	default:
		return fmt.Errorf("unknown Modifier value: %v", s)
	}
	return nil
}

// ModifierCommand is an alias for ModifierMeta.
const ModifierCommand Modifier = ModifierMeta

// DispatchDragEventType type of the drag event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchDragEvent
type DispatchDragEventType string

// String returns the DispatchDragEventType as string value.
func (t DispatchDragEventType) String() string {
	return string(t)
}

// DispatchDragEventType values.
const (
	DragEnter  DispatchDragEventType = "dragEnter"
	DragOver   DispatchDragEventType = "dragOver"
	Drop       DispatchDragEventType = "drop"
	DragCancel DispatchDragEventType = "dragCancel"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *DispatchDragEventType) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch DispatchDragEventType(s) {
	case DragEnter:
		*t = DragEnter
	case DragOver:
		*t = DragOver
	case Drop:
		*t = Drop
	case DragCancel:
		*t = DragCancel
	default:
		return fmt.Errorf("unknown DispatchDragEventType value: %v", s)
	}
	return nil
}

// KeyType type of the key event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchKeyEvent
type KeyType string

// String returns the KeyType as string value.
func (t KeyType) String() string {
	return string(t)
}

// KeyType values.
const (
	KeyDown    KeyType = "keyDown"
	KeyUp      KeyType = "keyUp"
	KeyRawDown KeyType = "rawKeyDown"
	KeyChar    KeyType = "char"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *KeyType) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch KeyType(s) {
	case KeyDown:
		*t = KeyDown
	case KeyUp:
		*t = KeyUp
	case KeyRawDown:
		*t = KeyRawDown
	case KeyChar:
		*t = KeyChar
	default:
		return fmt.Errorf("unknown KeyType value: %v", s)
	}
	return nil
}

// MouseType type of the mouse event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchMouseEvent
type MouseType string

// String returns the MouseType as string value.
func (t MouseType) String() string {
	return string(t)
}

// MouseType values.
const (
	MousePressed  MouseType = "mousePressed"
	MouseReleased MouseType = "mouseReleased"
	MouseMoved    MouseType = "mouseMoved"
	MouseWheel    MouseType = "mouseWheel"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *MouseType) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch MouseType(s) {
	case MousePressed:
		*t = MousePressed
	case MouseReleased:
		*t = MouseReleased
	case MouseMoved:
		*t = MouseMoved
	case MouseWheel:
		*t = MouseWheel
	default:
		return fmt.Errorf("unknown MouseType value: %v", s)
	}
	return nil
}

// DispatchMouseEventPointerType pointer type (default: "mouse").
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchMouseEvent
type DispatchMouseEventPointerType string

// String returns the DispatchMouseEventPointerType as string value.
func (t DispatchMouseEventPointerType) String() string {
	return string(t)
}

// DispatchMouseEventPointerType values.
const (
	Mouse DispatchMouseEventPointerType = "mouse"
	Pen   DispatchMouseEventPointerType = "pen"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *DispatchMouseEventPointerType) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch DispatchMouseEventPointerType(s) {
	case Mouse:
		*t = Mouse
	case Pen:
		*t = Pen
	default:
		return fmt.Errorf("unknown DispatchMouseEventPointerType value: %v", s)
	}
	return nil
}

// TouchType type of the touch event. TouchEnd and TouchCancel must not
// contain any touch points, while TouchStart and TouchMove must contains at
// least one.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Input#method-dispatchTouchEvent
type TouchType string

// String returns the TouchType as string value.
func (t TouchType) String() string {
	return string(t)
}

// TouchType values.
const (
	TouchStart  TouchType = "touchStart"
	TouchEnd    TouchType = "touchEnd"
	TouchMove   TouchType = "touchMove"
	TouchCancel TouchType = "touchCancel"
)

// UnmarshalJSON satisfies [json.Unmarshaler].
func (t *TouchType) UnmarshalJSON(buf []byte) error {
	s := string(buf)
	s = strings.TrimSuffix(strings.TrimPrefix(s, `"`), `"`)

	switch TouchType(s) {
	case TouchStart:
		*t = TouchStart
	case TouchEnd:
		*t = TouchEnd
	case TouchMove:
		*t = TouchMove
	case TouchCancel:
		*t = TouchCancel
	default:
		return fmt.Errorf("unknown TouchType value: %v", s)
	}
	return nil
}
