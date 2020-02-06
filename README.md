# サンプルサイト起動方法
下記のGoコマンドを実行する。ファイアウォール経由の許可を求められる場合はアクセスを許可する。
```
cd sample-site
go run application.go
```

# サンプルサイトへのアクセス方法
下記のURLへアクセスする。
```
http://localhost:5000/
```

# 蓄積サイトへのアクセス方法
下記のURLへアクセスする。
```
http://localhost:5000/old/
```
なお、サイトは下記の3つを蓄積している。
```
http://localhost:5000/old/20200102/
http://localhost:5000/old/20200103/
http://localhost:5000/old/20200104/
```
