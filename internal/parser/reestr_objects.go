package parser

import (
	"encoding/xml"
	"fmt"

	"gar-loader/internal/repository/postgres"
)

type ReestrObjectXML struct {
	XMLName    xml.Name `xml:"OBJECT"`
	ObjectID   int64    `xml:"OBJECTID,attr"`
	ObjectGUID string   `xml:"OBJECTGUID,attr"`
	ChangeID   int64    `xml:"CHANGEID,attr"`
	IsActive   int16    `xml:"ISACTIVE,attr"`
	LevelID    string   `xml:"LEVELID,attr"`
	CreateDate string   `xml:"CREATEDATE,attr"`
	UpdateDate string   `xml:"UPDATEDATE,attr"`
}

func MapReestrObject(raw ReestrObjectXML) (postgres.ReestrObject, error) {
	levelID, err := parseInt32(raw.LevelID)
	if err != nil {
		return postgres.ReestrObject{}, fmt.Errorf("parse LEVELID for objectid=%d: %w", raw.ObjectID, err)
	}

	createDate, err := parseDate(raw.CreateDate)
	if err != nil {
		return postgres.ReestrObject{}, fmt.Errorf("parse CREATEDATE for objectid=%d: %w", raw.ObjectID, err)
	}

	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.ReestrObject{}, fmt.Errorf("parse UPDATEDATE for objectid=%d: %w", raw.ObjectID, err)
	}

	return postgres.ReestrObject{
		ObjectID:   raw.ObjectID,
		ObjectGUID: raw.ObjectGUID,
		ChangeID:   raw.ChangeID,
		IsActive:   raw.IsActive,
		LevelID:    levelID,
		CreateDate: createDate,
		UpdateDate: updateDate,
	}, nil
}
