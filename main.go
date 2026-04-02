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

func main() {
	ctx := context.Background()
	_, deltaURL, err := downloader.GetArchiveURLs(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := run(ctx, deltaURL, "AS_ADDR_OBJ_"); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, url string, prefix string) error {
	extractFolder := os.Getenv("EXTRACT_FOLDER")
	downloadsFolder := os.Getenv("DOWNLOADS_FOLDER")
	dstDir := os.Getenv("DST_DIR")
	xmlDir := os.Getenv("XML_DIR")
	d := downloader.New(nil)

	//delete /download and /extracted

	defer func() {
		err := downloader.DeleteFolder(downloadsFolder)
		if err != nil {
			log.Println("delete downloads folder:", err)
		}
		err = downloader.DeleteFolder(extractFolder)
		if err != nil {
			log.Println("delete extracted folder:", err)
		}
	}()

	//download zip file
	log.Println("downloading zip...")
	fmt.Println(url)
	zipPath, err := d.Download(ctx, url, downloadsFolder)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("downloaded:", zipPath)

	//	unzip
	log.Println("extracting folder...")
	err = parser.Unzip(zipPath, dstDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("unzipped:", zipPath)

	// from XML to structs
	log.Println("parsing to go structs...")
	xmlPath, err := parser.FindFileByPrefix(xmlDir, prefix)
	if err != nil {
		log.Fatal(err)
	}
	addrObjs, err := parser.ParseAddrObjsFromXML(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("parsed rows:", len(addrObjs))

	//batch and upsert
	log.Println("batching and upserting")
	db, err := postgres.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := postgres.NewRepository(db)
	if err := repo.UpsertAddrObjsBatch(ctx, addrObjs); err != nil {
		log.Fatal(err)
	}
	return nil

}
