package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"gar-loader/internal/repository/postgres"
)

type AdmHierarchyXML struct {
	XMLName     xml.Name `xml:"ITEM"`
	ID          int64    `xml:"ID,attr"`
	ObjectID    int64    `xml:"OBJECTID,attr"`
	ParentObjID string   `xml:"PARENTOBJID,attr"`
	ChangeID    int64    `xml:"CHANGEID,attr"`
	RegionCode  string   `xml:"REGIONCODE,attr"`
	PrevID      string   `xml:"PREVID,attr"`
	NextID      string   `xml:"NEXTID,attr"`
	UpdateDate  string   `xml:"UPDATEDATE,attr"`
	StartDate   string   `xml:"STARTDATE,attr"`
	EndDate     string   `xml:"ENDDATE,attr"`
	IsActive    int16    `xml:"ISACTIVE,attr"`
	Path        string   `xml:"PATH,attr"`
}

func MapAdmHierarchy(raw AdmHierarchyXML) (postgres.AdmHierarchy, error) {
	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	startDate, err := parseDate(raw.StartDate)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse STARTDATE for id=%d: %w", raw.ID, err)
	}

	endDate, err := parseDate(raw.EndDate)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse ENDDATE for id=%d: %w", raw.ID, err)
	}

	parentObjID, err := parseNullableInt64(raw.ParentObjID)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse PARENTOBJID for id=%d: %w", raw.ID, err)
	}

	regionCode, err := parseNullableInt32(raw.RegionCode)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse REGIONCODE for id=%d: %w", raw.ID, err)
	}

	prevID, err := parseNullableInt64(raw.PrevID)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse PREVID for id=%d: %w", raw.ID, err)
	}

	nextID, err := parseNullableInt64(raw.NextID)
	if err != nil {
		return postgres.AdmHierarchy{}, fmt.Errorf("parse NEXTID for id=%d: %w", raw.ID, err)
	}

	return postgres.AdmHierarchy{
		ID:          raw.ID,
		ObjectID:    raw.ObjectID,
		ParentObjID: parentObjID,
		ChangeID:    raw.ChangeID,
		RegionCode:  regionCode,
		PrevID:      prevID,
		NextID:      nextID,
		UpdateDate:  updateDate,
		StartDate:   startDate,
		EndDate:     endDate,
		IsActive:    raw.IsActive,
		Path:        raw.Path,
	}, nil
}

func parseNullableInt64(value string) (*int64, error) {
	if value == "" {
		return nil, nil
	}

	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func parseNullableInt32(value string) (*int32, error) {
	if value == "" {
		return nil, nil
	}

	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return nil, err
	}

	n := int32(v)
	return &n, nil
}
