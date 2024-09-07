package events

import (
	//nolint: gosec // MD5 is used for generating unique IDs, not for cryptographic purposes
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"

	"github.com/aaegamysta/a-study-in-history/spec/gen"
)

type Type int32

const (
	Unspecified Type = iota
	Historical
	Birth
	Death
	Holiday
)

func (t Type) String() string {
	switch t {
	case Unspecified:
		return "Unspecified"
	case Historical:
		return "events"
	case Birth:
		return "births"
	case Death:
		return "deaths"
	case Holiday:
		return "holidays"
	default:
		return "Unknown"
	}
}

func TypeFromString(s string) Type {
	switch s {
	case "Unspecified":
		return Unspecified
	case "events":
		return Historical
	case "births":
		return Birth
	case "deaths":
		return Death
	case "holidays":
		return Holiday
	default:
		return Unspecified
	}
}

type Thumbnail struct {
	Path   string
	Width  int64
	Height int64
}

type Event struct {
	Type        Type
	Day         int64
	Month       int64
	Year        int64
	ID          string
	Title       string
	Description string
	Thumbnail   Thumbnail
}

func (e Event) FormatForPosting() string {
	return fmt.Sprintf("On this day, %d-%d-%d, %s", e.Day, e.Month, e.Year, e.Title)
}

type Collection struct {
	Type   Type
	Day    int64
	Month  int64
	Events []Event
}

func (e1 *Collection) Equals(e2 *Collection) bool {
	if e1 == nil {
		return false
	}
	if e1.Day != e2.Day || e1.Month != e2.Month {
		return false
	}
	if len(e1.Events) != len(e2.Events) {
		return false
	}
	for i := range len(e1.Events) {
		if !reflect.DeepEqual(e1.Events[i].Type, e2.Events[i].Type) {
			return false
		}
	}
	return true
}

func mapEventToGRPC(e Event) *gen.Event {
	return &gen.Event{
		Type:        gen.Type(e.Type),
		Day:         e.Day,
		Month:       e.Month,
		Year:        e.Year,
		Title:       e.Title,
		Description: e.Description,
		Thumbnail: &gen.Thumbnail{
			Url:    e.Thumbnail.Path,
			Width:  e.Thumbnail.Width,
			Height: e.Thumbnail.Height,
		},
	}
}

func mapGRPCToEvent(e *gen.Event) Event {
	return Event{
		Type:        Type(e.GetType()),
		Day:         e.GetDay(),
		Month:       e.GetMonth(),
		Year:        e.GetYear(),
		Title:       e.GetTitle(),
		Description: e.GetDescription(),
		Thumbnail: Thumbnail{
			Path:   e.GetThumbnail().GetUrl(),
			Width:  e.GetThumbnail().GetWidth(),
			Height: e.GetThumbnail().GetHeight(),
		},
	}
}

func MapEventsCollectionToGRPC(coll Collection) *gen.EventsCollection {
	grpcColl := gen.EventsCollection{
		Events: make([]*gen.Event, 0),
		Type:   gen.Type(coll.Type),
		Day:    coll.Day,
		Month:  coll.Month,
	}
	for _, e := range coll.Events {
		grpcColl.Events = append(grpcColl.Events, mapEventToGRPC(e))
	}
	return &grpcColl
}

func MapGRPCToEventsCollection(grpcColl *gen.EventsCollection) Collection {
	coll := Collection{
		Events: make([]Event, 0),
		Type:   Type(grpcColl.GetType()),
		Day:    grpcColl.GetDay(),
		Month:  grpcColl.GetMonth(),
	}
	for _, e := range grpcColl.GetEvents() {
		coll.Events = append(coll.Events, mapGRPCToEvent(e))
	}
	return coll
}

// nolint: gosec // MD5 is used for generating unique IDs, not for cryptographic purposes
func GenerateEventID(typing Type, month, day, year int64, title string) string {
	h := md5.New()
	h.Write([]byte(typing.String()))
	h.Write([]byte(strconv.Itoa(int(month))))
	h.Write([]byte(strconv.Itoa(int(day))))
	h.Write([]byte(strconv.Itoa(int(year))))
	h.Write([]byte(title))
	sum := h.Sum(nil)
	ID := hex.EncodeToString(sum)
	return ID
}
