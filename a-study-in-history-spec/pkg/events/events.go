package events

import "github.com/aaegamysta/a-study-in-history/spec/gen"

type Type int64

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

type EventsCollection struct {
	Type   Type
	Day    int64
	Month  int64
	Events []Event
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

func MapEventsCollectionToGRPC(coll EventsCollection) *gen.EventsCollection {
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
