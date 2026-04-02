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