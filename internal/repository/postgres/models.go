package postgres

import "time"

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
	IsActive   int16
	IsActual   int16
}
type AddrObjDivision struct {
	ID       int64
	ParentID int64
	ChildID  int64
	ChangeID int64
}
type AdmHierarchy struct {
	ID          int64
	ObjectID    int64
	ParentObjID *int64
	ChangeID    int64
	RegionCode  *int32
	PrevID      *int64
	NextID      *int64
	UpdateDate  time.Time
	StartDate   time.Time
	EndDate     time.Time
	IsActive    int16
	Path        string
}
type Apartment struct {
	ID         int64
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64
	Number     string
	ApartType  int32
	OperTypeID int32
	PrevID     *int64
	NextID     *int64
	UpdateDate time.Time
	StartDate  time.Time
	EndDate    time.Time
	IsActual   int16
	IsActive   int16
}
type Carplace struct {
	ID         int64
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64
	Number     string
	OperTypeID int32
	PrevID     *int64
	NextID     *int64
	UpdateDate time.Time
	StartDate  time.Time
	EndDate    time.Time
}
type ChangeHistory struct {
	ChangeID    int64
	ObjectID    int64
	AdrObjectID string
	OperTypeID  int32
	NDocID      *int64
	ChangeDate  time.Time
}
type House struct {
	ID         int64
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64

	HouseNum string
	AddNum1  string
	AddNum2  string

	HouseType  *int32
	AddType1   *int32
	AddType2   *int32
	OperTypeID int32

	PrevID *int64
	NextID *int64

	UpdateDate time.Time
	StartDate  time.Time
	EndDate    time.Time

	IsActual int16
	IsActive int16
}
type MunHierarchy struct {
	ID          int64
	ObjectID    int64
	ParentObjID *int64
	ChangeID    int64
	Oktmo       string
	PrevID      *int64
	NextID      *int64
	UpdateDate  time.Time
	StartDate   time.Time
	EndDate     time.Time
	IsActive    int16
	Path        string
}
type NormativeDoc struct {
	ID         int64
	Name       string
	Date       time.Time
	Number     string
	Type       int32
	Kind       int32
	UpdateDate time.Time
	OrgName    string
	RegNum     string
	RegDate    *time.Time
	AccDate    *time.Time
	Comment    string
}
type ReestrObject struct {
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64
	IsActive   int16
	LevelID    int32
	CreateDate time.Time
	UpdateDate time.Time
}
type Room struct {
	ID         int64
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64
	Number     string
	RoomType   int32
	OperTypeID int32
	PrevID     *int64
	NextID     *int64
	UpdateDate time.Time
	StartDate  time.Time
	EndDate    time.Time
	IsActual   int16
	IsActive   int16
}
type Stead struct {
	ID         int64
	ObjectID   int64
	ObjectGUID string
	ChangeID   int64
	Number     string
	OperTypeID int32
	PrevID     *int64
	NextID     *int64
	UpdateDate time.Time
	StartDate  time.Time
	EndDate    time.Time
	IsActual   int16
	IsActive   int16
}
