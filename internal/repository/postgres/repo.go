package postgres

import (
	"context"
	"fmt"
	"gar-loader/internal/parser"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

const upsertAddrObjQuery = `
insert into gar.as_addr_obj (
    id,
    objectid,
    objectguid,
    changeid,
    name,
    typename,
    level,
    opertypeid,
    previd,
    nextid,
    updatedate,
    startdate,
    enddate,
    isactual,
    isactive
) values (
    $1, $2, $3, $4, $5,
    $6, $7, $8, $9, $10,
    $11, $12, $13, $14, $15
)
on conflict (id) do update set
    objectid   = excluded.objectid,
    objectguid = excluded.objectguid,
    changeid   = excluded.changeid,
    name       = excluded.name,
    typename   = excluded.typename,
    level      = excluded.level,
    opertypeid = excluded.opertypeid,
    previd     = excluded.previd,
    nextid     = excluded.nextid,
    updatedate = excluded.updatedate,
    startdate  = excluded.startdate,
    enddate    = excluded.enddate,
    isactual   = excluded.isactual,
    isactive   = excluded.isactive
`

func (r *Repository) UpsertAddrObjsBatch(ctx context.Context, items []parser.AddrObj) error {
	if len(items) == 0 {
		return nil
	}

	const batchSize = 1000

	for start := 0; start < len(items); start += batchSize {
		end := start + batchSize
		if end > len(items) {
			end = len(items)
		}

		if err := r.upsertAddrObjsChunk(ctx, items[start:end]); err != nil {
			return fmt.Errorf("upsert addr objs chunk [%d:%d]: %w", start, end, err)
		}
	}

	return nil
}

func (r *Repository) upsertAddrObjsChunk(ctx context.Context, items []parser.AddrObj) error {
	var batch pgx.Batch

	for _, item := range items {
		batch.Queue(
			upsertAddrObjQuery,
			item.ID,
			item.ObjectID,
			item.ObjectGUID,
			item.ChangeID,
			item.Name,
			item.TypeName,
			item.Level,
			item.OperTypeID,
			item.PrevID,
			item.NextID,
			item.UpdateDate,
			item.StartDate,
			item.EndDate,
			item.IsActual,
			item.IsActive,
		)
	}

	results := r.db.SendBatch(ctx, &batch)
	defer results.Close()

	for i := 0; i < len(items); i++ {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("exec batch item %d (id=%d): %w", i, items[i].ID, err)
		}
	}

	return nil
}
