package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func ParseXMLStream[TRaw any, TModel any](
	xmlPath string,
	rowTag string,
	batchSize int,
	mapFn func(TRaw) (TModel, error),
	handler func([]TModel) error,
) error {
	if xmlPath == "" {
		return fmt.Errorf("xmlPath is empty")
	}
	if rowTag == "" {
		return fmt.Errorf("rowTag is empty")
	}
	if batchSize <= 0 {
		return fmt.Errorf("batchSize must be > 0")
	}
	if mapFn == nil {
		return fmt.Errorf("mapFn is nil")
	}
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}

	file, err := os.Open(xmlPath)
	if err != nil {
		return fmt.Errorf("open xml file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	batch := make([]TModel, 0, batchSize)

	flush := func() error {
		if len(batch) == 0 {
			return nil
		}

		if err := handler(batch); err != nil {
			return err
		}

		batch = batch[:0]
		return nil
	}

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read xml token: %w", err)
		}

		startElement, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		if startElement.Name.Local != rowTag {
			continue
		}

		var raw TRaw
		if err := decoder.DecodeElement(&raw, &startElement); err != nil {
			return fmt.Errorf("decode %s: %w", rowTag, err)
		}

		model, err := mapFn(raw)
		if err != nil {
			return err
		}

		batch = append(batch, model)

		if len(batch) >= batchSize {
			if err := flush(); err != nil {
				return err
			}
		}
	}

	return flush()
}
