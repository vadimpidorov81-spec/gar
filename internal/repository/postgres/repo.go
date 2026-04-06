package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

type BatchUpsertConfig[T any] struct {
	Table           string
	Columns         []string
	ConflictColumns []string
	UpdateColumns   []string
	Values          func(T) []any
}

func UpsertBatch[T any](ctx context.Context, r *Repository, items []T, cfg BatchUpsertConfig[T]) error {
	if len(items) == 0 {
		return nil
	}

	if err := validateBatchConfig(cfg); err != nil {
		return err
	}

	query := buildUpsertQuery(
		cfg.Table,
		cfg.Columns,
		cfg.ConflictColumns,
		cfg.UpdateColumns,
	)

	const batchSize = 1000

	for start := 0; start < len(items); start += batchSize {
		end := start + batchSize
		if end > len(items) {
			end = len(items)
		}

		if err := execBatchChunk(ctx, r.db, items[start:end], query, cfg.Values); err != nil {
			return fmt.Errorf("upsert batch chunk [%d:%d]: %w", start, end, err)
		}
	}

	return nil
}

func validateBatchConfig[T any](cfg BatchUpsertConfig[T]) error {
	if cfg.Table == "" {
		return fmt.Errorf("table is empty")
	}
	if len(cfg.Columns) == 0 {
		return fmt.Errorf("columns are empty")
	}
	if len(cfg.ConflictColumns) == 0 {
		return fmt.Errorf("conflict columns are empty")
	}
	if len(cfg.UpdateColumns) == 0 {
		return fmt.Errorf("update columns are empty")
	}
	if cfg.Values == nil {
		return fmt.Errorf("values func is nil")
	}

	return nil
}

func execBatchChunk[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	items []T,
	query string,
	valuesFn func(T) []any,
) error {
	var batch pgx.Batch

	for _, item := range items {
		batch.Queue(query, valuesFn(item)...)
	}

	results := db.SendBatch(ctx, &batch)
	defer results.Close()

	for i := 0; i < len(items); i++ {
		if _, err := results.Exec(); err != nil {
			return fmt.Errorf("exec batch item %d: %w", i, err)
		}
	}

	return nil
}

func buildUpsertQuery(
	table string,
	columns []string,
	conflictColumns []string,
	updateColumns []string,
) string {
	placeholders := make([]string, 0, len(columns))
	for i := range columns {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	updates := make([]string, 0, len(updateColumns))
	for _, col := range updateColumns {
		updates = append(updates, fmt.Sprintf("%s = excluded.%s", col, col))
	}

	return fmt.Sprintf(`
insert into %s (
    %s
) values (
    %s
)
on conflict (%s) do update set
    %s
`,
		table,
		strings.Join(columns, ",\n    "),
		strings.Join(placeholders, ", "),
		strings.Join(conflictColumns, ", "),
		strings.Join(updates, ",\n    "),
	)
}
