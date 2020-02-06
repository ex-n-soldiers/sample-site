# サンプルサイト起動方法
下記のGoコマンドを実行する。ファイアウォール経由の許可を求められる場合はアクセスを許可する。
```
cd sample-site
go run main.go
```

3日分の蓄積サイトにアクセスする場合は、-dオプションを付与して起動する。
* 1日目のサイトの場合
```
go run main.go -d 1
```
* 2日目のサイトの場合
```
go run main.go -d 2
```
* 3日目のサイトの場合
```
go run main.go -d 3
```

# サンプルサイトへのアクセス方法
下記のURLへアクセスする。
```
http://localhost:5000/
```
