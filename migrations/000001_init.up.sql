create schema if not exists gar;

CREATE TABLE gar.as_addr_obj (
                                  id           bigint        NOT NULL,
                                  objectid     bigint        NOT NULL,
                                  objectguid   uuid          NOT NULL,
                                  changeid     bigint        NOT NULL,
                                  name         varchar(250)  NOT NULL,
                                  typename     varchar(50)   NOT NULL,
                                  level        varchar(10)   NOT NULL,
                                  opertypeid   varchar(2)    NOT NULL,
                                  previd       bigint                   DEFAULT 0,
                                  nextid       bigint                   DEFAULT 0,
                                  updatedate   date          NOT NULL,
                                  startdate    date          NOT NULL,
                                  enddate      date          NOT NULL,
                                  isactual     smallint      NOT NULL,
                                  isactive     smallint      NOT NULL,

                                  CONSTRAINT pk_as_addr_obj PRIMARY KEY (id),
                                  CONSTRAINT chk_as_addr_obj_isactual CHECK (isactual IN (0, 1)),
                                  CONSTRAINT chk_as_addr_obj_isactive CHECK (isactive IN (0, 1))
);
create table if not exists gar.as_addr_obj_division (
                                id         bigint           not null,
                                parentid   bigint           not null,
                                childid    bigint           not null,
                                changeid   bigint           not null,

                                constraint pk_as_addr_obj_division primary key (id)
);
create table if not exists gar.as_adm_hierarchy (
                                id           bigint    not null,
                                objectid     bigint    not null,
                                parentobjid  bigint,
                                changeid     bigint    not null,
                                regioncode   integer,
                                previd       bigint,
                                nextid       bigint,
                                updatedate   date      not null,
                                startdate    date      not null,
                                enddate      date      not null,
                                isactive     smallint  not null,
                                path         text      not null,

                                constraint pk_as_adm_hierarchy primary key (id),
                                constraint chk_as_adm_hierarchy_isactive check (isactive in (0, 1))
);
create table if not exists gar.as_apartments (
                                id           bigint        not null,
                                objectid     bigint        not null,
                                objectguid   uuid          not null,
                                changeid     bigint        not null,
                                number       varchar(50)   not null,
                                aparttype    integer       not null,
                                opertypeid   integer       not null,
                                previd       bigint,
                                nextid       bigint,
                                updatedate   date          not null,
                                startdate    date          not null,
                                enddate      date          not null,
                                isactual     smallint      not null,
                                isactive     smallint      not null,

    constraint pk_as_apartments primary key (id),
    constraint chk_as_apartments_isactual check (isactual in (0, 1)),
    constraint chk_as_apartments_isactive check (isactive in (0, 1))
);
create table if not exists gar.as_carplaces (
                            id           bigint       not null,
                            objectid     bigint       not null,
                            objectguid   uuid         not null,
                            changeid     bigint       not null,
                            number       varchar(50)  not null,
                            opertypeid   integer      not null,
                            previd       bigint,
                            nextid       bigint,
                            updatedate   date         not null,
                            startdate    date         not null,
                            enddate      date         not null,

    constraint pk_as_carplaces primary key (id)
);
create table if not exists gar.as_change_history (
                                                     changeid     bigint       not null,
                                                     objectid     bigint       not null,
                                                     adrobjectid  uuid         not null,
                                                     opertypeid   integer      not null,
                                                     ndocid       bigint,
                                                     changedate   date         not null,

                                                     constraint pk_as_change_history primary key (changeid)
);
create table if not exists gar.as_houses (
                                             id           bigint       not null,
                                             objectid     bigint       not null,
                                             objectguid   uuid         not null,
                                             changeid     bigint       not null,

                                             housenum     varchar(50),
    addnum1      varchar(50),
    addnum2      varchar(50),

    housetype    integer,
    addtype1     integer,
    addtype2     integer,
    opertypeid   integer      not null,

    previd       bigint,
    nextid       bigint,

    updatedate   date         not null,
    startdate    date         not null,
    enddate      date         not null,
    isactual     smallint     not null,
    isactive     smallint     not null,

    constraint pk_as_houses primary key (id),
    constraint chk_as_houses_isactual check (isactual in (0, 1)),
    constraint chk_as_houses_isactive check (isactive in (0, 1))
);

