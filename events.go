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

// To be implemented further
type MouseEvent struct {
	AltKey    bool
	Button    int
	Buttons   int
	ClientX   int
	ClientY   int
	CtrlKey   bool
	MetaKey   bool
	MovementX int
	MovementY int
	OffsetX   int
	OffsetY   int
	PageX     int
	PageY     int
	ScreenX   int
	ScreenY   int
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
		i, _ := strconv.Atoi(value.(string))
		val = reflect.ValueOf(i)
		break
	case "bool":
		i, _ := strconv.ParseBool(value.(string))
		val = reflect.ValueOf(i)
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
	m := structs.Map(obj)

	for k := range m {
		jk := []rune(k)
		jk[0] = unicode.ToLower(jk[0])
		err := setField(obj, k, val.Get(string(jk)).String())
		if err != nil {
			return nil, err
		}
	}
	return obj, nil
}
