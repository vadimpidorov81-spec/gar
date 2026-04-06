package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"gar-loader/internal/repository/postgres"
)

type ApartmentXML struct {
	XMLName    xml.Name `xml:"APARTMENT"`
	ID         int64    `xml:"ID,attr"`
	ObjectID   int64    `xml:"OBJECTID,attr"`
	ObjectGUID string   `xml:"OBJECTGUID,attr"`
	ChangeID   int64    `xml:"CHANGEID,attr"`
	Number     string   `xml:"NUMBER,attr"`
	ApartType  string   `xml:"APARTTYPE,attr"`
	OperTypeID string   `xml:"OPERTYPEID,attr"`
	PrevID     string   `xml:"PREVID,attr"`
	NextID     string   `xml:"NEXTID,attr"`
	UpdateDate string   `xml:"UPDATEDATE,attr"`
	StartDate  string   `xml:"STARTDATE,attr"`
	EndDate    string   `xml:"ENDDATE,attr"`
	IsActual   int16    `xml:"ISACTUAL,attr"`
	IsActive   int16    `xml:"ISACTIVE,attr"`
}

func MapApartment(raw ApartmentXML) (postgres.Apartment, error) {
	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	startDate, err := parseDate(raw.StartDate)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse STARTDATE for id=%d: %w", raw.ID, err)
	}

	endDate, err := parseDate(raw.EndDate)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse ENDDATE for id=%d: %w", raw.ID, err)
	}

	apartType, err := parseInt32(raw.ApartType)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse APARTTYPE for id=%d: %w", raw.ID, err)
	}

	operTypeID, err := parseInt32(raw.OperTypeID)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse OPERTYPEID for id=%d: %w", raw.ID, err)
	}

	prevID, err := parseNullableInt64(raw.PrevID)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse PREVID for id=%d: %w", raw.ID, err)
	}

	nextID, err := parseNullableInt64(raw.NextID)
	if err != nil {
		return postgres.Apartment{}, fmt.Errorf("parse NEXTID for id=%d: %w", raw.ID, err)
	}

	return postgres.Apartment{
		ID:         raw.ID,
		ObjectID:   raw.ObjectID,
		ObjectGUID: raw.ObjectGUID,
		ChangeID:   raw.ChangeID,
		Number:     raw.Number,
		ApartType:  apartType,
		OperTypeID: operTypeID,
		PrevID:     prevID,
		NextID:     nextID,
		UpdateDate: updateDate,
		StartDate:  startDate,
		EndDate:    endDate,
		IsActual:   raw.IsActual,
		IsActive:   raw.IsActive,
	}, nil
}

func parseInt32(value string) (int32, error) {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(v), nil
}
