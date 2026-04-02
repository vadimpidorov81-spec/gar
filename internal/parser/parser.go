package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

func ParseAddrObjsFromXML(xmlPath string) ([]AddrObj, error) {
	file, err := os.Open(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("open xml file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	result := make([]AddrObj, 0)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read xml token: %w", err)
		}

		startElement, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		if startElement.Name.Local != "OBJECT" {
			continue
		}

		var raw addrObjXML
		if err := decoder.DecodeElement(&raw, &startElement); err != nil {
			return nil, fmt.Errorf("decode OBJECT: %w", err)
		}

		addrObj, err := mapAddrObj(raw)
		if err != nil {
			return nil, err
		}

		result = append(result, addrObj)
	}

	return result, nil
}

func mapAddrObj(raw addrObjXML) (AddrObj, error) {
	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return AddrObj{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	startDate, err := parseDate(raw.StartDate)
	if err != nil {
		return AddrObj{}, fmt.Errorf("parse STARTDATE for id=%d: %w", raw.ID, err)
	}

	endDate, err := parseDate(raw.EndDate)
	if err != nil {
		return AddrObj{}, fmt.Errorf("parse ENDDATE for id=%d: %w", raw.ID, err)
	}

	return AddrObj{
		ID:         raw.ID,
		ObjectID:   raw.ObjectID,
		ObjectGUID: raw.ObjectGUID,
		ChangeID:   raw.ChangeID,
		Name:       raw.Name,
		TypeName:   raw.TypeName,
		Level:      raw.Level,
		OperTypeID: raw.OperTypeID,
		PrevID:     raw.PrevID,
		NextID:     raw.NextID,
		UpdateDate: updateDate,
		StartDate:  startDate,
		EndDate:    endDate,
		IsActual:   raw.IsActual,
		IsActive:   raw.IsActive,
	}, nil
}

func parseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}
