package parser

import (
	"encoding/xml"
	"fmt"

	"gar-loader/internal/repository/postgres"
)

type CarplaceXML struct {
	XMLName    xml.Name `xml:"CARPLACE"`
	ID         int64    `xml:"ID,attr"`
	ObjectID   int64    `xml:"OBJECTID,attr"`
	ObjectGUID string   `xml:"OBJECTGUID,attr"`
	ChangeID   int64    `xml:"CHANGEID,attr"`
	Number     string   `xml:"NUMBER,attr"`
	OperTypeID string   `xml:"OPERTYPEID,attr"`
	PrevID     string   `xml:"PREVID,attr"`
	NextID     string   `xml:"NEXTID,attr"`
	UpdateDate string   `xml:"UPDATEDATE,attr"`
	StartDate  string   `xml:"STARTDATE,attr"`
	EndDate    string   `xml:"ENDDATE,attr"`
}

func MapCarplace(raw CarplaceXML) (postgres.Carplace, error) {
	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.Carplace{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	startDate, err := parseDate(raw.StartDate)
	if err != nil {
		return postgres.Carplace{}, fmt.Errorf("parse STARTDATE for id=%d: %w", raw.ID, err)
	}

	endDate, err := parseDate(raw.EndDate)
	if err != nil {
		return postgres.Carplace{}, fmt.Errorf("parse ENDDATE for id=%d: %w", raw.ID, err)
	}

	operTypeID, err := parseInt32(raw.OperTypeID)
	if err != nil {
		return postgres.Carplace{}, fmt.Errorf("parse OPERTYPEID for id=%d: %w", raw.ID, err)
	}

	prevID, err := parseNullableInt64(raw.PrevID)
	if err != nil {
		return postgres.Carplace{}, fmt.Errorf("parse PREVID for id=%d: %w", raw.ID, err)
	}

	nextID, err := parseNullableInt64(raw.NextID)
	if err != nil {
		return postgres.Carplace{}, fmt.Errorf("parse NEXTID for id=%d: %w", raw.ID, err)
	}

	return postgres.Carplace{
		ID:         raw.ID,
		ObjectID:   raw.ObjectID,
		ObjectGUID: raw.ObjectGUID,
		ChangeID:   raw.ChangeID,
		Number:     raw.Number,
		OperTypeID: operTypeID,
		PrevID:     prevID,
		NextID:     nextID,
		UpdateDate: updateDate,
		StartDate:  startDate,
		EndDate:    endDate,
	}, nil
}
