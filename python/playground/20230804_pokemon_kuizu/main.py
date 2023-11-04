# -*- coding: utf-8 -*-
"""
Created on Wed Aug  2 19:27:36 2023

@author: U-CHAN


東海オンエアの企画をプログラムでゲームにしてみよう。ウェブスクレイピングの練習がてら。
https://www.youtube.com/watch?v=q8mbM5CeYEg&t=60s&ab_channel=%E6%9D%B1%E6%B5%B7%E3%82%AA%E3%83%B3%E3%82%A8%E3%82%A2

ポケモン３０匹の画像をランダムでダウンロード、表示して、正解だったら次の画像、失敗だったら最初からって感じの簡単なゲーム。
"""

import requests
import random
import cv2

#ポケモン図鑑の番号をランダムで生成する関数。
def generate_number():
    x = random.randint(1,898)
    return x

#答えと回答を比較して正誤を判定する関数。
def judge_answer(correct, answer):
    if correct == answer:
        return True
    else:
        return False

#pokeapiを利用して、ポケモンの画像、名前を取得する関数。
def catch_pokemon(number):
    url1 = 'https://pokeapi.co/api/v2/pokemon/{}/'.format(number)
    url2 = 'https://pokeapi.co/api/v2/pokemon-species/{}/'.format(number)

    r1 = requests.get(url1, timeout=5)
    pokemon_img = r1.json()
    imageurl = pokemon_img['sprites']['other']['official-artwork']['front_default']

    r2 = requests.get(url2, timeout=5)
    pokemon_name = r2.json()

    ir = requests.get(imageurl)
    image = ir.content

    file_name = '{}.jpg'.format(number)

    with open(file_name, "wb") as f:
        f.write(image)

    name = pokemon_name['names'][0]['name']

    return name
#----------------------------------------------------------------------------------------------
saido = 1
lastjudge = False

while saido == 1:

    i = 1

    print("これよりポケモンクイズを開始します。出題範囲は図鑑番号898までです。")
    print("画像を見て、ポケモンの名前を入力してください。")
    print("30問連続正解でクリアです。")

    while i <= 30:
        if i > 1 :
            print("現在{}問正解\n".format(i - 1))
        print("第{}問".format(i))

        n = generate_number()
        pokemon = catch_pokemon(n)
        img = cv2.imread("{}.jpg".format(n))
        cv2.imshow("{}.jpg".format(n),img)
        cv2.waitKey(1)

        if judge_answer(pokemon, input('このポケモンの名前は？>>')):
            print("\n正解です。")
            i = i + 1

            if i == 31:
                print("全問正解です。おめでとうございます。")
                cv2.destroyAllWindows()
                break

            print("次の問題へすすみます。")
            cv2.destroyAllWindows()

        else:
            print("\n不正解です。")
            print("正解は{}です。".format(pokemon))
            print("再挑戦をお待ちしております。")
            cv2.destroyAllWindows()
            break

    while lastjudge == False:
        print("再度挑戦しますか？")
        judge = input("y/n>>")

        if judge == "n" :
            saido = 0
            lastjudge = True
            print("プログラムを終了します。お疲れさまでした。")
        elif judge =="y":
            lastjudge = True
        else:
            print("yかnを入力してください。")

    lastjudge = False
