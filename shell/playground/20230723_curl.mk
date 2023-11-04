#!/bin/zsh
# 基本
f1:
	curl 'https://api.warframe.market/v1/sister/weapons'


# HTTP レスポンスヘッダを表示する
# -I で、Headerのみ
I:
	curl -I 'https://api.warframe.market/v1/sister/weapons'


# -i ならば、Response Header, Body 両方
i:
	curl -i 'https://api.warframe.market/v1/sister/weapons'

# コンテンツ取得時の詳細なログを表示する
# -v オプションを指定すると、コンテンツを取得した際の詳細なログ (レスポンスヘッダの内容や HTTP ステータスコード等) を表示します。
v:
	curl -v 'https://api.warframe.market/v1/sister/weapons'
