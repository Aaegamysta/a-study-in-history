package wikipedia

import (
	"context"
	"fmt"
	"io"

	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"github.com/buger/jsonparser"
)

type firstPage struct {
	description string
	path        string
	width       int64
	height      int64
}

func (i *Impl) parseEvents(_ context.Context, coll *events.Collection, r io.ReadCloser) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	bytes, dataType, offset, err := jsonparser.Get(bytes, coll.Type.String())
	if err != nil {
		return err
	}
	if dataType != jsonparser.Array {
		return fmt.Errorf("error happened while parsing %s events on %d-%d: object parsed is not an array",
			coll.Type, coll.Month, coll.Day,
		)
	}
	if offset == 0 || bytes == nil {
		return fmt.Errorf("error happened while parsing %s events on %d-%d: array is empty", coll.Type, coll.Month, coll.Day)
	}
	//nolint staticcheck // parsingErr is used as a way to return error from the callback
	offset, err = jsonparser.ArrayEach(bytes, func(b []byte, _ jsonparser.ValueType, _ int, parsingErr error) {
		title, parseTitleErr := jsonparser.GetString(b, "text")
		if parseTitleErr != nil {
			//nolint errcheck // parsingErr is used as a way to return error from the callback
			parsingErr = parseTitleErr
			return
		}
		// not all events like holidays have no year
		year, _ := jsonparser.GetInt(b, "year")
		parsedFirstPage := parseFirstPage(b)
		coll.Events = append(coll.Events, events.Event{
			Type:        coll.Type,
			Day:         coll.Day,
			Month:       coll.Month,
			Year:        year,
			ID:          events.GenerateEventID(coll.Type, coll.Month, coll.Month, year, title),
			Title:       title,
			Description: parsedFirstPage.description,
			Thumbnail: events.Thumbnail{
				Path:   parsedFirstPage.path,
				Width:  parsedFirstPage.width,
				Height: parsedFirstPage.height,
			},
		})
	})
	if offset == 0 {
		return fmt.Errorf("error happened while parsing %s events on %d-%d: array is empty", coll.Type, coll.Month, coll.Day)
	}
	if err != nil {
		return err
	}
	return nil
}

func parseFirstPage(bytes []byte) firstPage {
	//nolint: dogsled // first page and its contents extract and thumbnail might not present for some events
	b, _, _, _ := jsonparser.Get(bytes, "pages", "[0]")
	extract, _ := jsonparser.GetString(b, "extract")
	source, _ := jsonparser.GetString(b, "thumbnail", "source")
	width, _ := jsonparser.GetInt(b, "thumbnail", "width")
	height, _ := jsonparser.GetInt(b, "thumbnail", "height")

	return firstPage{
		description: extract,
		path:        source,
		width:       width,
		height:      height,
	}
}
