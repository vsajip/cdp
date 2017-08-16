// Code generated by cdpgen. DO NOT EDIT.

package input

// SetIgnoreInputEventsArgs represents the arguments for SetIgnoreInputEvents in the Input domain.
type SetIgnoreInputEventsArgs struct {
	Ignore bool `json:"ignore"` // Ignores input events processing when set to true.
}

// NewSetIgnoreInputEventsArgs initializes SetIgnoreInputEventsArgs with the required arguments.
func NewSetIgnoreInputEventsArgs(ignore bool) *SetIgnoreInputEventsArgs {
	args := new(SetIgnoreInputEventsArgs)
	args.Ignore = ignore
	return args
}

// DispatchKeyEventArgs represents the arguments for DispatchKeyEvent in the Input domain.
type DispatchKeyEventArgs struct {
	// Type Type of the key event.
	//
	// Values: "keyDown", "keyUp", "rawKeyDown", "char".
	Type                  string         `json:"type"`
	Modifiers             *int           `json:"modifiers,omitempty"`             // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp             TimeSinceEpoch `json:"timestamp,omitempty"`             // Time at which the event occurred.
	Text                  *string        `json:"text,omitempty"`                  // Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events (default: "")
	UnmodifiedText        *string        `json:"unmodifiedText,omitempty"`        // Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
	KeyIdentifier         *string        `json:"keyIdentifier,omitempty"`         // Unique key identifier (e.g., 'U+0041') (default: "").
	Code                  *string        `json:"code,omitempty"`                  // Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Key                   *string        `json:"key,omitempty"`                   // Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	WindowsVirtualKeyCode *int           `json:"windowsVirtualKeyCode,omitempty"` // Windows virtual key code (default: 0).
	NativeVirtualKeyCode  *int           `json:"nativeVirtualKeyCode,omitempty"`  // Native virtual key code (default: 0).
	AutoRepeat            *bool          `json:"autoRepeat,omitempty"`            // Whether the event was generated from auto repeat (default: false).
	IsKeypad              *bool          `json:"isKeypad,omitempty"`              // Whether the event was generated from the keypad (default: false).
	IsSystemKey           *bool          `json:"isSystemKey,omitempty"`           // Whether the event was a system key event (default: false).
}

// NewDispatchKeyEventArgs initializes DispatchKeyEventArgs with the required arguments.
func NewDispatchKeyEventArgs(typ string) *DispatchKeyEventArgs {
	args := new(DispatchKeyEventArgs)
	args.Type = typ
	return args
}

// SetModifiers sets the Modifiers optional argument. Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
func (a *DispatchKeyEventArgs) SetModifiers(modifiers int) *DispatchKeyEventArgs {
	a.Modifiers = &modifiers
	return a
}

// SetTimestamp sets the Timestamp optional argument. Time at which the event occurred.
func (a *DispatchKeyEventArgs) SetTimestamp(timestamp TimeSinceEpoch) *DispatchKeyEventArgs {
	a.Timestamp = timestamp
	return a
}

// SetText sets the Text optional argument. Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events (default: "")
func (a *DispatchKeyEventArgs) SetText(text string) *DispatchKeyEventArgs {
	a.Text = &text
	return a
}

// SetUnmodifiedText sets the UnmodifiedText optional argument. Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
func (a *DispatchKeyEventArgs) SetUnmodifiedText(unmodifiedText string) *DispatchKeyEventArgs {
	a.UnmodifiedText = &unmodifiedText
	return a
}

// SetKeyIdentifier sets the KeyIdentifier optional argument. Unique key identifier (e.g., 'U+0041') (default: "").
func (a *DispatchKeyEventArgs) SetKeyIdentifier(keyIdentifier string) *DispatchKeyEventArgs {
	a.KeyIdentifier = &keyIdentifier
	return a
}

// SetCode sets the Code optional argument. Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
func (a *DispatchKeyEventArgs) SetCode(code string) *DispatchKeyEventArgs {
	a.Code = &code
	return a
}

// SetKey sets the Key optional argument. Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
func (a *DispatchKeyEventArgs) SetKey(key string) *DispatchKeyEventArgs {
	a.Key = &key
	return a
}

// SetWindowsVirtualKeyCode sets the WindowsVirtualKeyCode optional argument. Windows virtual key code (default: 0).
func (a *DispatchKeyEventArgs) SetWindowsVirtualKeyCode(windowsVirtualKeyCode int) *DispatchKeyEventArgs {
	a.WindowsVirtualKeyCode = &windowsVirtualKeyCode
	return a
}

// SetNativeVirtualKeyCode sets the NativeVirtualKeyCode optional argument. Native virtual key code (default: 0).
func (a *DispatchKeyEventArgs) SetNativeVirtualKeyCode(nativeVirtualKeyCode int) *DispatchKeyEventArgs {
	a.NativeVirtualKeyCode = &nativeVirtualKeyCode
	return a
}

// SetAutoRepeat sets the AutoRepeat optional argument. Whether the event was generated from auto repeat (default: false).
func (a *DispatchKeyEventArgs) SetAutoRepeat(autoRepeat bool) *DispatchKeyEventArgs {
	a.AutoRepeat = &autoRepeat
	return a
}

// SetIsKeypad sets the IsKeypad optional argument. Whether the event was generated from the keypad (default: false).
func (a *DispatchKeyEventArgs) SetIsKeypad(isKeypad bool) *DispatchKeyEventArgs {
	a.IsKeypad = &isKeypad
	return a
}

// SetIsSystemKey sets the IsSystemKey optional argument. Whether the event was a system key event (default: false).
func (a *DispatchKeyEventArgs) SetIsSystemKey(isSystemKey bool) *DispatchKeyEventArgs {
	a.IsSystemKey = &isSystemKey
	return a
}

// DispatchMouseEventArgs represents the arguments for DispatchMouseEvent in the Input domain.
type DispatchMouseEventArgs struct {
	// Type Type of the mouse event.
	//
	// Values: "mousePressed", "mouseReleased", "mouseMoved", "mouseWheel".
	Type      string         `json:"type"`
	X         float64        `json:"x"`                   // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y         float64        `json:"y"`                   // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Modifiers *int           `json:"modifiers,omitempty"` // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp TimeSinceEpoch `json:"timestamp,omitempty"` // Time at which the event occurred.
	// Button Mouse button (default: "none").
	//
	// Values: "none", "left", "middle", "right".
	Button     *string  `json:"button,omitempty"`
	ClickCount *int     `json:"clickCount,omitempty"` // Number of times the mouse button was clicked (default: 0).
	DeltaX     *float64 `json:"deltaX,omitempty"`     // X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY     *float64 `json:"deltaY,omitempty"`     // Y delta in CSS pixels for mouse wheel event (default: 0).
}

// NewDispatchMouseEventArgs initializes DispatchMouseEventArgs with the required arguments.
func NewDispatchMouseEventArgs(typ string, x float64, y float64) *DispatchMouseEventArgs {
	args := new(DispatchMouseEventArgs)
	args.Type = typ
	args.X = x
	args.Y = y
	return args
}

// SetModifiers sets the Modifiers optional argument. Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
func (a *DispatchMouseEventArgs) SetModifiers(modifiers int) *DispatchMouseEventArgs {
	a.Modifiers = &modifiers
	return a
}

// SetTimestamp sets the Timestamp optional argument. Time at which the event occurred.
func (a *DispatchMouseEventArgs) SetTimestamp(timestamp TimeSinceEpoch) *DispatchMouseEventArgs {
	a.Timestamp = timestamp
	return a
}

// SetButton sets the Button optional argument. Mouse button (default: "none").
//
// Values: "none", "left", "middle", "right".
func (a *DispatchMouseEventArgs) SetButton(button string) *DispatchMouseEventArgs {
	a.Button = &button
	return a
}

// SetClickCount sets the ClickCount optional argument. Number of times the mouse button was clicked (default: 0).
func (a *DispatchMouseEventArgs) SetClickCount(clickCount int) *DispatchMouseEventArgs {
	a.ClickCount = &clickCount
	return a
}

// SetDeltaX sets the DeltaX optional argument. X delta in CSS pixels for mouse wheel event (default: 0).
func (a *DispatchMouseEventArgs) SetDeltaX(deltaX float64) *DispatchMouseEventArgs {
	a.DeltaX = &deltaX
	return a
}

// SetDeltaY sets the DeltaY optional argument. Y delta in CSS pixels for mouse wheel event (default: 0).
func (a *DispatchMouseEventArgs) SetDeltaY(deltaY float64) *DispatchMouseEventArgs {
	a.DeltaY = &deltaY
	return a
}

// DispatchTouchEventArgs represents the arguments for DispatchTouchEvent in the Input domain.
type DispatchTouchEventArgs struct {
	// Type Type of the touch event.
	//
	// Values: "touchStart", "touchEnd", "touchMove".
	Type        string         `json:"type"`
	TouchPoints []TouchPoint   `json:"touchPoints"`         // Touch points.
	Modifiers   *int           `json:"modifiers,omitempty"` // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Timestamp   TimeSinceEpoch `json:"timestamp,omitempty"` // Time at which the event occurred.
}

// NewDispatchTouchEventArgs initializes DispatchTouchEventArgs with the required arguments.
func NewDispatchTouchEventArgs(typ string, touchPoints []TouchPoint) *DispatchTouchEventArgs {
	args := new(DispatchTouchEventArgs)
	args.Type = typ
	args.TouchPoints = touchPoints
	return args
}

// SetModifiers sets the Modifiers optional argument. Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
func (a *DispatchTouchEventArgs) SetModifiers(modifiers int) *DispatchTouchEventArgs {
	a.Modifiers = &modifiers
	return a
}

// SetTimestamp sets the Timestamp optional argument. Time at which the event occurred.
func (a *DispatchTouchEventArgs) SetTimestamp(timestamp TimeSinceEpoch) *DispatchTouchEventArgs {
	a.Timestamp = timestamp
	return a
}

// EmulateTouchFromMouseEventArgs represents the arguments for EmulateTouchFromMouseEvent in the Input domain.
type EmulateTouchFromMouseEventArgs struct {
	// Type Type of the mouse event.
	//
	// Values: "mousePressed", "mouseReleased", "mouseMoved", "mouseWheel".
	Type      string         `json:"type"`
	X         int            `json:"x"`         // X coordinate of the mouse pointer in DIP.
	Y         int            `json:"y"`         // Y coordinate of the mouse pointer in DIP.
	Timestamp TimeSinceEpoch `json:"timestamp"` // Time at which the event occurred.
	// Button Mouse button.
	//
	// Values: "none", "left", "middle", "right".
	Button     string   `json:"button"`
	DeltaX     *float64 `json:"deltaX,omitempty"`     // X delta in DIP for mouse wheel event (default: 0).
	DeltaY     *float64 `json:"deltaY,omitempty"`     // Y delta in DIP for mouse wheel event (default: 0).
	Modifiers  *int     `json:"modifiers,omitempty"`  // Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	ClickCount *int     `json:"clickCount,omitempty"` // Number of times the mouse button was clicked (default: 0).
}

// NewEmulateTouchFromMouseEventArgs initializes EmulateTouchFromMouseEventArgs with the required arguments.
func NewEmulateTouchFromMouseEventArgs(typ string, x int, y int, timestamp TimeSinceEpoch, button string) *EmulateTouchFromMouseEventArgs {
	args := new(EmulateTouchFromMouseEventArgs)
	args.Type = typ
	args.X = x
	args.Y = y
	args.Timestamp = timestamp
	args.Button = button
	return args
}

// SetDeltaX sets the DeltaX optional argument. X delta in DIP for mouse wheel event (default: 0).
func (a *EmulateTouchFromMouseEventArgs) SetDeltaX(deltaX float64) *EmulateTouchFromMouseEventArgs {
	a.DeltaX = &deltaX
	return a
}

// SetDeltaY sets the DeltaY optional argument. Y delta in DIP for mouse wheel event (default: 0).
func (a *EmulateTouchFromMouseEventArgs) SetDeltaY(deltaY float64) *EmulateTouchFromMouseEventArgs {
	a.DeltaY = &deltaY
	return a
}

// SetModifiers sets the Modifiers optional argument. Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
func (a *EmulateTouchFromMouseEventArgs) SetModifiers(modifiers int) *EmulateTouchFromMouseEventArgs {
	a.Modifiers = &modifiers
	return a
}

// SetClickCount sets the ClickCount optional argument. Number of times the mouse button was clicked (default: 0).
func (a *EmulateTouchFromMouseEventArgs) SetClickCount(clickCount int) *EmulateTouchFromMouseEventArgs {
	a.ClickCount = &clickCount
	return a
}

// SynthesizePinchGestureArgs represents the arguments for SynthesizePinchGesture in the Input domain.
type SynthesizePinchGestureArgs struct {
	X                 float64           `json:"x"`                           // X coordinate of the start of the gesture in CSS pixels.
	Y                 float64           `json:"y"`                           // Y coordinate of the start of the gesture in CSS pixels.
	ScaleFactor       float64           `json:"scaleFactor"`                 // Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	RelativeSpeed     *int              `json:"relativeSpeed,omitempty"`     // Relative pointer speed in pixels per second (default: 800).
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"` // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
}

// NewSynthesizePinchGestureArgs initializes SynthesizePinchGestureArgs with the required arguments.
func NewSynthesizePinchGestureArgs(x float64, y float64, scaleFactor float64) *SynthesizePinchGestureArgs {
	args := new(SynthesizePinchGestureArgs)
	args.X = x
	args.Y = y
	args.ScaleFactor = scaleFactor
	return args
}

// SetRelativeSpeed sets the RelativeSpeed optional argument. Relative pointer speed in pixels per second (default: 800).
func (a *SynthesizePinchGestureArgs) SetRelativeSpeed(relativeSpeed int) *SynthesizePinchGestureArgs {
	a.RelativeSpeed = &relativeSpeed
	return a
}

// SetGestureSourceType sets the GestureSourceType optional argument. Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (a *SynthesizePinchGestureArgs) SetGestureSourceType(gestureSourceType GestureSourceType) *SynthesizePinchGestureArgs {
	a.GestureSourceType = gestureSourceType
	return a
}

// SynthesizeScrollGestureArgs represents the arguments for SynthesizeScrollGesture in the Input domain.
type SynthesizeScrollGestureArgs struct {
	X                     float64           `json:"x"`                               // X coordinate of the start of the gesture in CSS pixels.
	Y                     float64           `json:"y"`                               // Y coordinate of the start of the gesture in CSS pixels.
	XDistance             *float64          `json:"xDistance,omitempty"`             // The distance to scroll along the X axis (positive to scroll left).
	YDistance             *float64          `json:"yDistance,omitempty"`             // The distance to scroll along the Y axis (positive to scroll up).
	XOverscroll           *float64          `json:"xOverscroll,omitempty"`           // The number of additional pixels to scroll back along the X axis, in addition to the given distance.
	YOverscroll           *float64          `json:"yOverscroll,omitempty"`           // The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
	PreventFling          *bool             `json:"preventFling,omitempty"`          // Prevent fling (default: true).
	Speed                 *int              `json:"speed,omitempty"`                 // Swipe speed in pixels per second (default: 800).
	GestureSourceType     GestureSourceType `json:"gestureSourceType,omitempty"`     // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
	RepeatCount           *int              `json:"repeatCount,omitempty"`           // The number of times to repeat the gesture (default: 0).
	RepeatDelayMs         *int              `json:"repeatDelayMs,omitempty"`         // The number of milliseconds delay between each repeat. (default: 250).
	InteractionMarkerName *string           `json:"interactionMarkerName,omitempty"` // The name of the interaction markers to generate, if not empty (default: "").
}

// NewSynthesizeScrollGestureArgs initializes SynthesizeScrollGestureArgs with the required arguments.
func NewSynthesizeScrollGestureArgs(x float64, y float64) *SynthesizeScrollGestureArgs {
	args := new(SynthesizeScrollGestureArgs)
	args.X = x
	args.Y = y
	return args
}

// SetXDistance sets the XDistance optional argument. The distance to scroll along the X axis (positive to scroll left).
func (a *SynthesizeScrollGestureArgs) SetXDistance(xDistance float64) *SynthesizeScrollGestureArgs {
	a.XDistance = &xDistance
	return a
}

// SetYDistance sets the YDistance optional argument. The distance to scroll along the Y axis (positive to scroll up).
func (a *SynthesizeScrollGestureArgs) SetYDistance(yDistance float64) *SynthesizeScrollGestureArgs {
	a.YDistance = &yDistance
	return a
}

// SetXOverscroll sets the XOverscroll optional argument. The number of additional pixels to scroll back along the X axis, in addition to the given distance.
func (a *SynthesizeScrollGestureArgs) SetXOverscroll(xOverscroll float64) *SynthesizeScrollGestureArgs {
	a.XOverscroll = &xOverscroll
	return a
}

// SetYOverscroll sets the YOverscroll optional argument. The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
func (a *SynthesizeScrollGestureArgs) SetYOverscroll(yOverscroll float64) *SynthesizeScrollGestureArgs {
	a.YOverscroll = &yOverscroll
	return a
}

// SetPreventFling sets the PreventFling optional argument. Prevent fling (default: true).
func (a *SynthesizeScrollGestureArgs) SetPreventFling(preventFling bool) *SynthesizeScrollGestureArgs {
	a.PreventFling = &preventFling
	return a
}

// SetSpeed sets the Speed optional argument. Swipe speed in pixels per second (default: 800).
func (a *SynthesizeScrollGestureArgs) SetSpeed(speed int) *SynthesizeScrollGestureArgs {
	a.Speed = &speed
	return a
}

// SetGestureSourceType sets the GestureSourceType optional argument. Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (a *SynthesizeScrollGestureArgs) SetGestureSourceType(gestureSourceType GestureSourceType) *SynthesizeScrollGestureArgs {
	a.GestureSourceType = gestureSourceType
	return a
}

// SetRepeatCount sets the RepeatCount optional argument. The number of times to repeat the gesture (default: 0).
func (a *SynthesizeScrollGestureArgs) SetRepeatCount(repeatCount int) *SynthesizeScrollGestureArgs {
	a.RepeatCount = &repeatCount
	return a
}

// SetRepeatDelayMs sets the RepeatDelayMs optional argument. The number of milliseconds delay between each repeat. (default: 250).
func (a *SynthesizeScrollGestureArgs) SetRepeatDelayMs(repeatDelayMs int) *SynthesizeScrollGestureArgs {
	a.RepeatDelayMs = &repeatDelayMs
	return a
}

// SetInteractionMarkerName sets the InteractionMarkerName optional argument. The name of the interaction markers to generate, if not empty (default: "").
func (a *SynthesizeScrollGestureArgs) SetInteractionMarkerName(interactionMarkerName string) *SynthesizeScrollGestureArgs {
	a.InteractionMarkerName = &interactionMarkerName
	return a
}

// SynthesizeTapGestureArgs represents the arguments for SynthesizeTapGesture in the Input domain.
type SynthesizeTapGestureArgs struct {
	X                 float64           `json:"x"`                           // X coordinate of the start of the gesture in CSS pixels.
	Y                 float64           `json:"y"`                           // Y coordinate of the start of the gesture in CSS pixels.
	Duration          *int              `json:"duration,omitempty"`          // Duration between touchdown and touchup events in ms (default: 50).
	TapCount          *int              `json:"tapCount,omitempty"`          // Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"` // Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
}

// NewSynthesizeTapGestureArgs initializes SynthesizeTapGestureArgs with the required arguments.
func NewSynthesizeTapGestureArgs(x float64, y float64) *SynthesizeTapGestureArgs {
	args := new(SynthesizeTapGestureArgs)
	args.X = x
	args.Y = y
	return args
}

// SetDuration sets the Duration optional argument. Duration between touchdown and touchup events in ms (default: 50).
func (a *SynthesizeTapGestureArgs) SetDuration(duration int) *SynthesizeTapGestureArgs {
	a.Duration = &duration
	return a
}

// SetTapCount sets the TapCount optional argument. Number of times to perform the tap (e.g. 2 for double tap, default: 1).
func (a *SynthesizeTapGestureArgs) SetTapCount(tapCount int) *SynthesizeTapGestureArgs {
	a.TapCount = &tapCount
	return a
}

// SetGestureSourceType sets the GestureSourceType optional argument. Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (a *SynthesizeTapGestureArgs) SetGestureSourceType(gestureSourceType GestureSourceType) *SynthesizeTapGestureArgs {
	a.GestureSourceType = gestureSourceType
	return a
}
