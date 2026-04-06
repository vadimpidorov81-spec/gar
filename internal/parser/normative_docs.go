package parser

import (
	"encoding/xml"
	"fmt"
	"time"

	"gar-loader/internal/repository/postgres"
)

type NormativeDocXML struct {
	XMLName    xml.Name `xml:"NORMDOC"`
	ID         int64    `xml:"ID,attr"`
	Name       string   `xml:"NAME,attr"`
	Date       string   `xml:"DATE,attr"`
	Number     string   `xml:"NUMBER,attr"`
	Type       string   `xml:"TYPE,attr"`
	Kind       string   `xml:"KIND,attr"`
	UpdateDate string   `xml:"UPDATEDATE,attr"`
	OrgName    string   `xml:"ORGNAME,attr"`
	RegNum     string   `xml:"REGNUM,attr"`
	RegDate    string   `xml:"REGDATE,attr"`
	AccDate    string   `xml:"ACCDATE,attr"`
	Comment    string   `xml:"COMMENT,attr"`
}

func MapNormativeDoc(raw NormativeDocXML) (postgres.NormativeDoc, error) {
	docDate, err := parseDate(raw.Date)
	if err != nil {
		return postgres.NormativeDoc{}, fmt.Errorf("parse DATE for id=%d: %w", raw.ID, err)
	}

	updateDate, err := parseDate(raw.UpdateDate)
	if err != nil {
		return postgres.NormativeDoc{}, fmt.Errorf("parse UPDATEDATE for id=%d: %w", raw.ID, err)
	}

	docType, err := parseInt32(raw.Type)
	if err != nil {
		return postgres.NormativeDoc{}, fmt.Errorf("parse TYPE for id=%d: %w", raw.ID, err)
	}

	docKind, err := parseInt32(raw.Kind)
	if err != nil {
		return postgres.NormativeDoc{}, fmt.Errorf("parse KIND for id=%d: %w", raw.ID, err)
	}

	regDate, err := parseNullableDate(raw.RegDate)
	if err != nil {
		return postgres.NormativeDoc{}, fmt.Errorf("parse REGDATE for id=%d: %w", raw.ID, err)
	}

	accDate, err := parseNullableDate(raw.AccDate)
	if err != nil {
		return postgres.NormativeDoc{}, fmt.Errorf("parse ACCDATE for id=%d: %w", raw.ID, err)
	}

	return postgres.NormativeDoc{
		ID:         raw.ID,
		Name:       raw.Name,
		Date:       docDate,
		Number:     raw.Number,
		Type:       docType,
		Kind:       docKind,
		UpdateDate: updateDate,
		OrgName:    raw.OrgName,
		RegNum:     raw.RegNum,
		RegDate:    regDate,
		AccDate:    accDate,
		Comment:    raw.Comment,
	}, nil
}

func parseNullableDate(value string) (*time.Time, error) {
	if value == "" {
		return nil, nil
	}

	v, err := parseDate(value)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
