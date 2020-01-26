package main

import (
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

func main() {
    // 最終アクセス日時、更新日時を2020/1/1 00:00:00に修正する
    loc, _ := time.LoadLocation("Local")
    changeFileUpdatedAt(filepath.Join("static", "images"), time.Date(2020, 1, 1, 0, 0, 0, 0, loc))

    //ディレクトリを指定する
    fs := http.FileServer(http.Dir("static"))
    //ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
    http.Handle("/", fs)

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
