package parser

import (
	"encoding/xml"
	"fmt"
	"time"

	"gar-loader/internal/repository/postgres"
)

type AddrObjXML struct {
	XMLName    xml.Name `xml:"OBJECT"`
	ID         int64    `xml:"ID,attr"`
	ObjectID   int64    `xml:"OBJECTID,attr"`
	ObjectGUID string   `xml:"OBJECTGUID,attr"`
	ChangeID   int64    `xml:"CHANGEID,attr"`
	Name       string   `xml:"NAME,attr"`
	TypeName   string   `xml:"TYPENAME,attr"`
	Level      string   `xml:"LEVEL,attr"`
	OperTypeID string   `xml:"OPERTYPEID,attr"`
	PrevID     int64    `xml:"PREVID,attr"`
	NextID     int64    `xml:"NEXTID,attr"`
	UpdateDate string   `xml:"UPDATEDATE,attr"`
	StartDate  string   `xml:"STARTDATE,attr"`
	EndDate    string   `xml:"ENDDATE,attr"`
	IsActual   int16    `xml:"ISACTUAL,attr"`
	IsActive   int16    `xml:"ISACTIVE,attr"`
}

func MapAddrObj(raw AddrObjXML) (postgres.AddrObj, error) {
	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.AddrObj{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	startDate, err := parseDate(raw.StartDate)
	if err != nil {
		return postgres.AddrObj{}, fmt.Errorf("parse STARTDATE for id=%d: %w", raw.ID, err)
	}

	endDate, err := parseDate(raw.EndDate)
	if err != nil {
		return postgres.AddrObj{}, fmt.Errorf("parse ENDDATE for id=%d: %w", raw.ID, err)
	}

	return postgres.AddrObj{
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
