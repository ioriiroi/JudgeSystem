# JudgeSystem
競プロライクな自作ジャッジシステム



# 概要
Pythonのコードと入出力を指定のディレクトリに置くと、自動で全ての結果をジャッジする。


# 使い方
/testcodeフォルダ内にテストするPythonのコード(main.py)を置く。

/testcase/inフォルダ内に入力(.txt)を、/testcase/outフォルダ内に答え(.txt)をそれぞれ置く。

コマンドラインにて"go run ./JudgeSystem [入出力のあるフォルダ名] [実行したいPythonのコードがあるフォルダ名]"を実行することで自動的に全てのテストケースを調べることができる。

(引数の指定をしていない場合は/testcase /testcode内のファイルを使う)


# 結果の種類
AC : 出力と答えが一致

WA : 出力と答えが不一致

RE : 実行時エラー

TLE : 実行時間(2秒)オーバー

PythonなのでCEは未実装


# 注意
出力の末尾には改行が必要 (無いと不正解扱いになる)

テストケースのファイル名は入出力で統一する必要がある
