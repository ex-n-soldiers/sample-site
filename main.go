package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 最終アクセス日時、更新日時を設定
	loc, _ := time.LoadLocation("Local")
	datetime := time.Date(2020, 1, 1, 0, 0, 0, 0, loc)
	changeFileUpdatedAt(filepath.Join("static", "images"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "pdf"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "old", "20200102", "images"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "old", "20200102", "pdf"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "old", "20200103", "images"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "old", "20200103", "pdf"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "old", "20200104", "images"), datetime)
	changeFileUpdatedAt(filepath.Join("static", "old", "20200104", "pdf"), datetime)

	// 2020/01/03以降のカメラ画像、PDFファイルの最終アクセス日時、更新日時を変更
	datetime = time.Date(2020, 1, 3, 0, 0, 0, 0, loc)
	os.Chtimes(filepath.Join("static", "old", "20200103", "images", "camera.jpg"), datetime, datetime)
	os.Chtimes(filepath.Join("static", "old", "20200103", "pdf", "camera.pdf"), datetime, datetime)
	os.Chtimes(filepath.Join("static", "old", "20200104", "images", "camera.jpg"), datetime, datetime)
	os.Chtimes(filepath.Join("static", "old", "20200104", "pdf", "camera.pdf"), datetime, datetime)

	// オプション取得
	var day int
	flag.IntVar(&day, "d", 0, "select day from 1 to 3. 0 (default) shows lists with multiple pages.")
	flag.Parse()

	// オプションにより表示する商品リストを変更
	switch day {
	case 0:
		http.HandleFunc("/", handler)
	case 1:
		http.Handle("/", http.FileServer(http.Dir("static/old/20200102")))
	case 2:
		http.Handle("/", http.FileServer(http.Dir("static/old/20200103")))
	case 3:
		http.Handle("/", http.FileServer(http.Dir("static/old/20200104")))
	default:
		log.Fatal("execute \"go run main.go\" or \"go run main.go -d <day>\" (<day> must be from 0 to 3).")
	}

	// サーバーを立ち上げる
	port := "5000"
	log.Println("Listening...")
	log.Println("Go to http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

// 指定したフォルダ配下の最終アクセス日時、更新日時を修正する
//
// dirPath フォルダのパス
// time    更新後の日時
func changeFileUpdatedAt(dirPath string, time time.Time) {
	files, _ := filepath.Glob(filepath.Join(dirPath, "*"))
	for _, f := range files {
		os.Chtimes(f, time, time)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// パス、クエリパラメータを取得
	urlPath := path.Join("static", r.URL.Path)
	query := r.URL.Query()

	// 表示するページ番号を設定
	var page string
	p := query["page"]
	if len(p) == 0 {
		page = "1"
	} else {
		page = p[0]
	}

	// 表示するファイルを仮設定
	ext := path.Ext(urlPath)
	if ext == "" {
		urlPath = path.Join(urlPath, "index.html")
		ext = path.Ext(urlPath)
	}

	// index.htmlが指定された場合は、ページ名を付与したファイルを表示
	if path.Base(urlPath) == "index.html" {
		if page == "1" || page == "2" || page == "3" {
			urlPath = strings.TrimRight(urlPath, ext) + page + ext
		}
	}

	http.ServeFile(w, r, urlPath)
}
