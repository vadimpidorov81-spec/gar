package main

import (
	"context"
	"fmt"
	"gar-loader/internal/downloader"
	"gar-loader/internal/parser"

	"gar-loader/internal/repository/postgres"
	"log"
	"os"
)

type importJob func(ctx context.Context, repo *postgres.Repository, xmlPath string) error

var jobs = map[string]importJob{
	"AS_ADDR_OBJ_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.AddrObjXML, postgres.AddrObj](
			xmlPath,
			"OBJECT",
			1000,
			parser.MapAddrObj,
			func(items []postgres.AddrObj) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.AddrObjUpsertConfig)
			},
		)
	},
	"AS_ADDR_OBJ_DIVISION_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.AddrObjDivisionXML, postgres.AddrObjDivision](
			xmlPath,
			"ITEM",
			1000,
			parser.MapAddrObjDivision,
			func(items []postgres.AddrObjDivision) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.AddrObjDivisionUpsertConfig)
			},
		)
	},
	"AS_ADM_HIERARCHY_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.AdmHierarchyXML, postgres.AdmHierarchy](
			xmlPath,
			"ITEM",
			1000,
			parser.MapAdmHierarchy,
			func(items []postgres.AdmHierarchy) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.AdmHierarchyUpsertConfig)
			},
		)
	},
	"AS_APARTMENTS_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.ApartmentXML, postgres.Apartment](
			xmlPath,
			"APARTMENT",
			1000,
			parser.MapApartment,
			func(items []postgres.Apartment) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.ApartmentUpsertConfig)
			},
		)
	},
	"AS_CARPLACES_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.CarplaceXML, postgres.Carplace](
			xmlPath,
			"CARPLACE",
			1000,
			parser.MapCarplace,
			func(items []postgres.Carplace) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.CarplaceUpsertConfig)
			},
		)
	},
	"AS_CHANGE_HISTORY_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.ChangeHistoryXML, postgres.ChangeHistory](
			xmlPath,
			"ITEM",
			1000,
			parser.MapChangeHistory,
			func(items []postgres.ChangeHistory) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.ChangeHistoryUpsertConfig)
			},
		)
	},
	"AS_HOUSES_2": func(ctx context.Context, repo *postgres.Repository, xmlPath string) error {
		return parser.ParseXMLStream[parser.HouseXML, postgres.House](
			xmlPath,
			"HOUSE",
			1000,
			parser.MapHouse,
			func(items []postgres.House) error {
				return postgres.UpsertBatch(ctx, repo, items, postgres.HouseUpsertConfig)
			},
		)
	},
}

func DeleteFolder(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("папка не найдена: %s", path)
		}
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("это не папка: %s", path)
	}
	return os.RemoveAll(path)
}

func main() {
	dstDir := os.Getenv("DST_DIR")
	xmlDir := os.Getenv("XML_DIR")
	extractFolder := os.Getenv("EXTRACT_FOLDER")
	downloadsFolder := os.Getenv("DOWNLOADS_FOLDER")
	//DELETE folders
	defer func() {
		err := DeleteFolder(downloadsFolder)
		if err != nil {
			log.Println("delete downloads folder:", err)
		}
		err = DeleteFolder(extractFolder)
		if err != nil {
			log.Println("delete extracted folder:", err)
		}
	}()

	ctx := context.Background()
	_, deltaURL, err := downloader.GetArchiveURLs(ctx, nil)
	url := deltaURL
	if err != nil {
		log.Fatal(err)
	}
	d := downloader.New(nil)
	//download
	log.Println("downloading zip...")
	zipPath, err := d.Download(ctx, url, downloadsFolder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("downloaded:", zipPath)
	//	unzip
	log.Println("extracting folder...")
	err = downloader.UnzipRegion(zipPath, dstDir, "16")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("unzipped:", zipPath)

	db, err := postgres.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := postgres.NewRepository(db)

	err = run(ctx, repo, xmlDir, "AS_ADDR_OBJ_2")
	if err != nil {
		log.Fatal(err, "...ASS_ADDR_OBJ_2")
	}
	err = run(ctx, repo, xmlDir, "AS_ADDR_OBJ_DIVISION_2")
	if err != nil {
		log.Fatal(err, "...AS_ADDR_OBJ_DIVISION_2")
	}
	err = run(ctx, repo, xmlDir, "AS_ADM_HIERARCHY_2")
	if err != nil {
		log.Fatal(err, "...AS_ADM_HIERARCHY_2")
	}
	err = run(ctx, repo, xmlDir, "AS_APARTMENTS_2")
	if err != nil {
		log.Fatal(err, "...AS_APARTMENTS_2")
	}
	err = run(ctx, repo, xmlDir, "AS_CARPLACES_2")
	if err != nil {
		log.Fatal(err, "...AS_CARPLACES_2")
	}
	err = run(ctx, repo, xmlDir, "AS_CHANGE_HISTORY_2")
	if err != nil {
		log.Fatal(err, "...AS_CHANGE_HISTORY_2")
	}
	err = run(ctx, repo, xmlDir, "AS_HOUSES_2")
	if err != nil {
		log.Fatal(err, "...AS_HOUSES_2")
	}
}

func run(ctx context.Context, repo *postgres.Repository, xmlDir string, prefix string) error {
	job, ok := jobs[prefix]
	if !ok {
		return fmt.Errorf("unknown import prefix: %s", prefix)
	}

	xmlPath, err := downloader.FindFileByPrefix(xmlDir, prefix)
	if err != nil {
		return err
	}

	return job(ctx, repo, xmlPath)
}
