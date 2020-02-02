package main

import (
    "flag"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

var (
    day int
)

func main() {
    // オプション取得
    flag.IntVar(&day, "d", 0, "select day from 1 to 3. 0 (default) shows lists with multiple pages.")
    flag.Parse()

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

    switch day {
    case 0:
        // 複数ページの商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static")))
    case 1:
        // 2020/01/01の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200102")))
    case 2:
        // 2020/01/02の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200103")))
    case 3:
        // 2020/01/03の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200104")))
    default:
        // エラー
        log.Fatal("execute \"go run main.go\" or \"go run main.go -d <day>\" (<day> must be from 1 to 5).")
    }

    log.Println("Listening...")
    // 3000ポートでサーバーを立ち上げる
    http.ListenAndServe(":3000", nil)
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
