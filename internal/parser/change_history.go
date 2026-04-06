package parser

import (
	"encoding/xml"
	"fmt"

	"gar-loader/internal/repository/postgres"
)

type ChangeHistoryXML struct {
	XMLName     xml.Name `xml:"ITEM"`
	ChangeID    int64    `xml:"CHANGEID,attr"`
	ObjectID    int64    `xml:"OBJECTID,attr"`
	AdrObjectID string   `xml:"ADROBJECTID,attr"`
	OperTypeID  string   `xml:"OPERTYPEID,attr"`
	NDocID      string   `xml:"NDOCID,attr"`
	ChangeDate  string   `xml:"CHANGEDATE,attr"`
}

func MapChangeHistory(raw ChangeHistoryXML) (postgres.ChangeHistory, error) {
	changeDate, err := parseDate(raw.ChangeDate)
	if err != nil {
		return postgres.ChangeHistory{}, fmt.Errorf("parse CHANGEDATE for changeid=%d: %w", raw.ChangeID, err)
	}

	operTypeID, err := parseInt32(raw.OperTypeID)
	if err != nil {
		return postgres.ChangeHistory{}, fmt.Errorf("parse OPERTYPEID for changeid=%d: %w", raw.ChangeID, err)
	}

	nDocID, err := parseNullableInt64(raw.NDocID)
	if err != nil {
		return postgres.ChangeHistory{}, fmt.Errorf("parse NDOCID for changeid=%d: %w", raw.ChangeID, err)
	}

	return postgres.ChangeHistory{
		ChangeID:    raw.ChangeID,
		ObjectID:    raw.ObjectID,
		AdrObjectID: raw.AdrObjectID,
		OperTypeID:  operTypeID,
		NDocID:      nDocID,
		ChangeDate:  changeDate,
	}, nil
}
