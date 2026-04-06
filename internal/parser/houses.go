package parser

import (
	"encoding/xml"
	"fmt"

	"gar-loader/internal/repository/postgres"
)

type HouseXML struct {
	XMLName    xml.Name `xml:"HOUSE"`
	ID         int64    `xml:"ID,attr"`
	ObjectID   int64    `xml:"OBJECTID,attr"`
	ObjectGUID string   `xml:"OBJECTGUID,attr"`
	ChangeID   int64    `xml:"CHANGEID,attr"`

	HouseNum string `xml:"HOUSENUM,attr"`
	AddNum1  string `xml:"ADDNUM1,attr"`
	AddNum2  string `xml:"ADDNUM2,attr"`

	HouseType  string `xml:"HOUSETYPE,attr"`
	AddType1   string `xml:"ADDTYPE1,attr"`
	AddType2   string `xml:"ADDTYPE2,attr"`
	OperTypeID string `xml:"OPERTYPEID,attr"`

	PrevID string `xml:"PREVID,attr"`
	NextID string `xml:"NEXTID,attr"`

	UpdateDate string `xml:"UPDATEDATE,attr"`
	StartDate  string `xml:"STARTDATE,attr"`
	EndDate    string `xml:"ENDDATE,attr"`

	IsActual int16 `xml:"ISACTUAL,attr"`
	IsActive int16 `xml:"ISACTIVE,attr"`
}

func MapHouse(raw HouseXML) (postgres.House, error) {
	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	startDate, err := parseDate(raw.StartDate)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse STARTDATE for id=%d: %w", raw.ID, err)
	}

	endDate, err := parseDate(raw.EndDate)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse ENDDATE for id=%d: %w", raw.ID, err)
	}

	houseType, err := parseNullableInt32(raw.HouseType)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse HOUSETYPE for id=%d: %w", raw.ID, err)
	}

	addType1, err := parseNullableInt32(raw.AddType1)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse ADDTYPE1 for id=%d: %w", raw.ID, err)
	}

	addType2, err := parseNullableInt32(raw.AddType2)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse ADDTYPE2 for id=%d: %w", raw.ID, err)
	}

	operTypeID, err := parseInt32(raw.OperTypeID)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse OPERTYPEID for id=%d: %w", raw.ID, err)
	}

	prevID, err := parseNullableInt64(raw.PrevID)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse PREVID for id=%d: %w", raw.ID, err)
	}

	nextID, err := parseNullableInt64(raw.NextID)
	if err != nil {
		return postgres.House{}, fmt.Errorf("parse NEXTID for id=%d: %w", raw.ID, err)
	}

	return postgres.House{
		ID:         raw.ID,
		ObjectID:   raw.ObjectID,
		ObjectGUID: raw.ObjectGUID,
		ChangeID:   raw.ChangeID,
		HouseNum:   raw.HouseNum,
		AddNum1:    raw.AddNum1,
		AddNum2:    raw.AddNum2,
		HouseType:  houseType,
		AddType1:   addType1,
		AddType2:   addType2,
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
