package browser

import (
	"fmt"
	"reflect"
	"strconv"
	"syscall/js"
	"unicode"

	"github.com/fatih/structs"
)

// HTML DOM events.
const (
	EventAbort          = "abort"
	EventAfterPrint     = "afterprint"
	EventAnimationEnd   = "animationend"
	EventAnimationStart = "animationstart"
	EventBeforePrint    = "beforeprint"
	EventBeforeUnload   = "beforeunload"
	EventBlur           = "blur"
	EventCanPlay        = "canplay"
	EventCanPlaythrough = "canplaythrough"
	EventChange         = "change"
	EventClick          = "click"
	EventContextMenu    = "contextmenu"
	EventCopy           = "copy"
	EventCut            = "cut"
	EventDblClick       = "dblclick"
	EventDrag           = "drag"
	EventDragEnd        = "dragend"
	EventDragEnter      = "dragenter"
	EventDragLeave      = "dragleave"
	EventDragOver       = "dragover"
	EventDragStart      = "dragstart"
	EventDrop           = "drop"
	EventDurationChange = "durationchange"
	EventEnded          = "ended"
	EventError          = "error"
	EventFocus          = "focus"
	EventFocusIn        = "focusin"
	EventFocusOut       = "focusout"
	EventHashChange     = "hashchange"
	EventInput          = "input"
	EventInvalid        = "invalid"
	EventKeyDown        = "keydown"
	EventKeyPress       = "keypress"
	EventKeyUp          = "keyup"
	EventLoad           = "load"
	EventLoadedData     = "loadeddata"
	EventLoadedMetadata = "loadedmetadata"
	EventLoadStart      = "loadstart"
	EventMessage        = "message"
	EventMouseDown      = "mousedown"
	EventMouseEnter     = "mouseenter"
	EventMouseLeave     = "mouseleave"
	EventMouseMove      = "mousemove"
	EventMouseOver      = "mouseover"
	EventMouseOut       = "mouseout"
	EventMouseUp        = "mouseup"
	EventMouseWheel     = "mousewheel"
	EventOffline        = "offline"
	EventOnline         = "online"
	EventOpen           = "open"
	EventPageHide       = "pagehide"
	EventPageShow       = "pageshow"
	EventPaste          = "paste"
	EventPause          = "pause"
	EventPlay           = "play"
	EventPlaying        = "playing"
	EventPopState       = "popstate"
	EventProgress       = "progress"
	EventRateChange     = "ratechange"
	EventResize         = "resize"
	EventReset          = "reset"
	EventScroll         = "scroll"
	EventSearch         = "search"
	EventSeeked         = "seeked"
	EventSeeking        = "seeking"
	EventSelect         = "select"
	EventShow           = "show"
	EventStalled        = "stalled"
	EventStorage        = "storage"
	EventSubmit         = "submit"
	EventSuspend        = "suspend"
	EventTimeUpdate     = "timeupdate"
	EventToggle         = "toggle"
	EventTouchCancel    = "touchcancel"
	EventTouchEnd       = "touchend"
	EventTouchMove      = "touchmove"
	EventTouchStart     = "touchstart"
	EventTransitionEnd  = "transitionend"
	EventUnload         = "unload"
	EventVolumechange   = "volumechange"
	EventWaiting        = "waiting"
	EventWheel          = "wheel"
)

// Possible Values of KeyboardEvent.Code
// Implement further
const (
	KeyboardKeyA = "KeyA"
	KeyboardKeyB = "KeyB"
	KeyboardKeyC = "KeyC"
	KeyboardKeyD = "KeyD"
	KeyboardKeyE = "KeyE"
	KeyboardKeyF = "KeyF"
	KeyboardKeyG = "KeyG"
	KeyboardKeyH = "KeyH"
	KeyboardKeyI = "KeyI"
	KeyboardKeyJ = "KeyJ"
	KeyboardKeyK = "KeyK"
	KeyboardKeyL = "KeyL"
	KeyboardKeyM = "KeyM"
	KeyboardKeyN = "KeyN"
	KeyboardKeyO = "KeyO"
	KeyboardKeyP = "KeyP"
	KeyboardKeyQ = "KeyQ"
	KeyboardKeyR = "KeyR"
	KeyboardKeyS = "KeyS"
	KeyboardKeyT = "KeyT"
	KeyboardKeyU = "KeyU"
	KeyboardKeyV = "KeyV"
	KeyboardKeyW = "KeyW"
	KeyboardKeyX = "KeyX"
	KeyboardKeyY = "KeyY"
	KeyboardKeyZ = "KeyZ"
)

// To be implemented further (create structs for all js.Value)
type MouseEvent struct {
	AltKey    bool
	Button    int
	Buttons   int
	ClientX   float64
	ClientY   float64
	CtrlKey   bool
	MetaKey   bool
	MovementX float64
	MovementY float64
	OffsetX   float64
	OffsetY   float64
	PageX     int
	PageY     int
	// RelatedTarget js.Value
	ScreenX  int
	ScreenY  int
	ShiftKey bool
}

// To be implemented further
type KeyboardEvent struct {
	AltKey      bool
	Code        string
	CtrlKey     bool
	IsComposing bool
	Location    int
	MetaKey     bool
	Repeat      bool
	ShiftKey    bool
}

//  Dynamically map Map value to  Struct values
func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()

	var val reflect.Value
	switch structFieldType.String() {
	case "int":
		i, _ := strconv.Atoi(value.(js.Value).String())
		val = reflect.ValueOf(i)
		break
	case "float64":
		i, _ := strconv.ParseFloat(value.(js.Value).String(), 64)
		val = reflect.ValueOf(i)
		break
	case "bool":
		i, _ := strconv.ParseBool(value.(js.Value).String())
		val = reflect.ValueOf(i)
		break
	case "string":
		val = reflect.ValueOf(value.(js.Value).String())
		break
	default:
		val = reflect.ValueOf(value)
		break
	}

	structFieldValue.Set(val)
	return nil
}

func NewMouseEvent(val js.Value) (*MouseEvent, error) {

	obj := &MouseEvent{}
	o, e := eventCreator(val, obj)
	return o.(*MouseEvent), e
}

func NewKeyboardEvent(val js.Value) (*KeyboardEvent, error) {

	obj := &KeyboardEvent{}
	o, e := eventCreator(val, obj)
	return o.(*KeyboardEvent), e

}

func eventCreator(val js.Value, obj interface{}) (interface{}, error) {
	m := structs.Map(obj)

	for k := range m {
		jk := []rune(k)
		jk[0] = unicode.ToLower(jk[0])
		err := setField(obj, k, val.Get(string(jk)))
		if err != nil {
			return nil, err
		}
	}
	return obj, nil
}
