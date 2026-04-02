package parser

import (
	"encoding/xml"
	"time"
)

type AddrObj struct {
	ID         int64
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64
	Name       string
	TypeName   string
	Level      string
	OperTypeID string
	PrevID     int64
	NextID     int64
	UpdateDate time.Time
	StartDate  time.Time
	EndDate    time.Time
	IsActual   int16
	IsActive   int16
}

type addrObjXML struct {
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
