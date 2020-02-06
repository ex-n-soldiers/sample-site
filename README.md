# サンプルサイト起動方法
下記のGoコマンドを実行する。ファイアウォール経由の許可を求められる場合はアクセスを許可する。
```
cd sample-site
go run application.go
```

さらに3日分の蓄積サイトにアクセスする場合は、-dオプションを指定して実行する。
```
go run application.go -d 1
go run application.go -d 2
go run application.go -d 3
```

# サンプルサイトへのアクセス方法
下記のURLへアクセスする。
```
http://localhost:5000/
```
