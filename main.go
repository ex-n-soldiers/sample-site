package main

import (
    "flag"
    "log"
    "net/http"
)

var (
    day int
)

func main() {
    // オプション取得
    flag.IntVar(&day, "d", 0, "select day from 1 to 5. 0 (default) shows lists with multiple pages.")
    flag.Parse()

    switch day {
    case 0:
        // 複数ページの商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static")))
    case 1:
        // 2020/01/01の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200101")))
    case 2:
        // 2020/01/02の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200102")))
    case 3:
        // 2020/01/03の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200103")))
    case 4:
        // 2020/01/04の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200104")))
    case 5:
        // 2020/01/05の商品リストを表示
        http.Handle("/", http.FileServer(http.Dir("static/old/20200105")))
    default:
        // エラー
        log.Fatal("execute \"go run main.go\" or \"go run main.go -d <day>\" (<day> must be from 1 to 5).")
    }

    log.Println("Listening...")
    // 3000ポートでサーバーを立ち上げる
    http.ListenAndServe(":3000", nil)
}
