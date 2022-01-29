# kifuwarabe-uec13

第１３回UEC杯コンピュータ囲碁大会向けきふわらべ  

GoGo をベースとしています  
📖 [GoGo](https://github.com/bleu48/GoGo)  

# Set up

## Set up - 1 - Get a GUI

きふわらべは 思考エンジンだけですので、GUI が別途必要です。  
ここでは CgfGoBan で動かす例を示します。  

![20220130go67a2.png](doc/img/20220130go67a2.png)  

📖 [CgfGoBan](http://www.yss-aya.com/cgfgoban_j.html)  

👆 CgfGoBan のページから英語版に飛んでください  

![20220130go68a2.png](doc/img/20220130go68a2.png)  

👆 `Download CgfGoBan 1.08 2015/03/16` をクリックして `cgfgoban108.zip` ファイルをダウンロードしてきてください  

![20220130go69.png](doc/img/20220130go69.png)  

👆 がんばって好きなところに解凍してください  

![20220130go70.png](doc/img/20220130go70.png)  

👆 最初の中身はこんな感じです  

## Set up - 2 - Make directories

![202201go1.png](./doc/img/202201go1.png)  

👆 CgfGoBan の `cgfgoban.exe` ファイルと同じディレクトリーに、  
以下の図を参考に `📂input/game_conf.toml` と `📂output`（空ディレクトリー） を置いてください  
`game_conf.toml` は、この Git hub の `📂input` の中に置いてあります  

```plain
📂 cgfgoban108
|
+----📂 input
|    |
|    +----📄 game_conf.toml (設定ファイル)
|
+----📂 output （空っぽ）
|
+----📄 cgfgoban.exe

```

game_conf.toml テキストファイルの内容は以下の通りです  

```toml
[Game]
# デフォルト値です。 CgfGoBan などの GUI はこのファイルを見ません

# Komi - コミ☆（＾～＾）
Komi = 6.5

# BoardSize - 何路盤。
BoardSize = 19

# MaxMoves - 最大手数。
MaxMoves = 1000

BoardData = '''
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
'''
```

## Set up - 3 - Download exe

次に、以下から きふわらべの `kifuwarabe-uec13.exe` ファイルを入手してください。  

![20220130go71a2.png](doc/img/20220130go71a2.png)  

📖 [Kifuwarabe-uec13 Release 0.0.1](https://github.com/muzudho/kifuwarabe-uec13/releases/tag/0.0.1)

👆 .zip に圧縮されているので、解凍してください。  
もし ウィルス扱いされてダウンロードできないなら、 以下の方法で、ソースをビルドして kifuwarabe-uec13.exe を作ってください。  

## Set up - 4 - Download source

きふわらべの .exe をソースからビルドして作る方法です。  
きふわらべのソースコードは、２つの方法で入手できます  

![20220130go71a3.png](./doc/img/20220130go71a3.png)  

📖 [Kifuwarabe-uec13 Release 0.0.1](https://github.com/muzudho/kifuwarabe-uec13/releases/tag/0.0.1)

👆 1つ目は、 リリースページの src をクリックする方法です。  

![20220130go73a2.png](./doc/img/20220130go73a2.png)  

👆 2つ目は、 `GitHub Desktop` を使ってダウンロードする方法です。 **がんばって** `GitHub Desktop` をインストールしておいてください。  

![GitHub Desktop](./doc/img/20220130go74.png)  

👆 そのあと がんばって ローカルPCに `kifuwarabe-uec13のリポジトリ―` を作ってください  
何がなにやら分からないかも知れませんが **がんばって** ください  

## Set up - 5 - Editor

![20220130go72.png](./doc/img/20220130go72.png)  

📖 [Visual Studio Code](https://code.visualstudio.com/)  

👆 プログラムを編集するツールとしては Visual Studio Code で十分でしょう。 **がんばって** インストールしてください  
**がんばって** `kifuwarabe-uec13のリポジトリ―` の📂フォルダーを開いてください。  
上図画面下のコマンドを打つところは **ターミナル** という名前です。 **がんばって** ターミナルを表示してください。  

## Set up - 6 - Programming Language

📖 [Go Downloads](https://go.dev/dl/)  

👆 Go言語を **がんばって** インストールしてください  

## Set up - 7 - Build

Visual Studio Code に戻りまして。  

![20220130go75a2.png](doc/img/20220130go75a2.png)  

```shell
go build
```

👆 ターミナルに上記のコマンドを打鍵してください。 `kifuwarabe-uec13.exe` 実行ファイルが作成されます。  
お疲れさまでした。  

# Run

ターミナルで以下のコマンドを打鍵すると、ターミナルで実行することもできます。  

```shell
# GTP形式の思考エンジンとして対局するなら
kifuwarabe-uec13

# 動作テストするなら
kifuwarabe-uec13 SelfPlay
```

👆 拡張子を省いた実行ファイル名を打鍵してください

## Set up - 8 - CgfGoBan

![20220130go76.png](doc/img/20220130go76.png)  

👆 CgfGoBan を開いて、 `[設定] - [GTPの設定]` と進んでください  

![20220130go77.png](doc/img/20220130go77.png)  

👆 きふわらべの .exe へのフルパスを入れてください  

![20220130go78.png](doc/img/20220130go78.png)  

👆 `[対局] - [対局開始]` と進んでください  

![20220130go79.png](doc/img/20220130go79.png)  

👆 Black と White のどちらか片方を `Computer(GTP)` にしてください。  

![20220130go80.png](doc/img/20220130go80.png)  

👆 これで きふわらべと対局できました。  
お疲れさまでした。  
