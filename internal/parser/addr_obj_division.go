package parser

import "gar-loader/internal/repository/postgres"

type AddrObjDivisionXML struct {
	ID       int64 `xml:"ID,attr"`
	ParentID int64 `xml:"PARENTID,attr"`
	ChildID  int64 `xml:"CHILDID,attr"`
	ChangeID int64 `xml:"CHANGEID,attr"`
}

func MapAddrObjDivision(raw AddrObjDivisionXML) (postgres.AddrObjDivision, error) {
	return postgres.AddrObjDivision{
		ID:       raw.ID,
		ParentID: raw.ParentID,
		ChildID:  raw.ChildID,
		ChangeID: raw.ChangeID,
	}, nil
}
