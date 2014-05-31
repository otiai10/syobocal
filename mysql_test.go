package animapi_test

import "github.com/otiai10/animapi"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"
import "testing"
import . "github.com/otiai10/mint"

func TestAnimapi_DB(t *testing.T) {
	mysqlClient := animapi.DB("./my.conf")
	Expect(t, mysqlClient.Err).ToBe(nil)

	mysqlClient = animapi.DB("./notfound.conf")
	Expect(t, mysqlClient.Err.Error()).ToBe("open ./notfound.conf: no such file or directory")

	mysqlClient = animapi.DB("./my.conf", "test")
	Expect(t, mysqlClient.Err).ToBe(nil)
}

func TestAnimapi_DB_FindPrograms(t *testing.T) {
	since, _ := animapi.Since("-4h")
	c := "./my.conf"

	// Find
	programs := animapi.DB(c, "test").FindPrograms(since)
	Expect(t, programs).TypeOf("[]model.Program")
	Expect(t, len(programs)).ToBe(0)

	// Add
	programs = getSamplePrograms()
	e := animapi.DB(c, "test").AddPrograms(programs)
	Expect(t, e).ToBe(nil)

	// Find
	programs = animapi.DB(c, "test").FindPrograms(since)
	Expect(t, len(programs)).ToBe(1)

	animapi.DB(c, "test").TearDown()
}

func getSamplePrograms() []model.Program {
	bytes := []byte(sampleResponse)
	response, _ := infrastructure.ConvertBytes2Response(bytes)
	return model.CreateProgramsFromSyobocalResponse(response)
}

var sampleResponse = `
<?xml version="1.0" encoding="UTF-8"?>
    <TitleLookupResponse>
        <Result>
            <Code>200</Code>
            <Message></Message>
        </Result>
    <TitleItems>
        <TitleItem id="3252">
            <TID>3252</TID>
            <LastUpdate>2014-05-25 00:03:32</LastUpdate>
            <Title>ニセコイ</Title>
            <ShortTitle></ShortTitle>
            <TitleYomi>にせこい</TitleYomi>
            <TitleEN></TitleEN>
            <Comment>
*リンク
-[[公式 http://www.nisekoi.jp/]]
-[[TOKYO MX http://s.mxtv.jp/nisekoi/]]
-[[BS11デジタル http://www.bs11.jp/anime/2468/]]
-[[ニコニコチャンネル http://ch.nicovideo.jp/nisekoi]]

*メモ
-全20話

*スタッフ
:原作:古味直志
:掲載誌:週刊少年ジャンプ(集英社)
:総監督:新房昭之
:監督:龍輪直征
:シリーズ構成:東冨耶子、新房昭之
:キャラクターデザイン:杉山延寛
:総作画監督:杉山延寛、潮月一也、西澤真也(#9～)
:美術監督:内藤健
:美術設定:大原盛仁
:色彩設計:滝沢いづみ
:ビジュアルエフェクト:酒井基
:撮影監督:江上怜
:編集:松原理恵
:音楽スーパーバイザー:神前暁(MONACA)
:音楽:千葉&quot;naotyu-&quot;直樹(#1～#6)、石濱翔(MONACA)(#1～#6)、菊谷知樹(#7～)
:音楽制作:アニプレックス
:音響監督:亀山俊樹
:音響制作:グルーヴ
:アニメーション制作:SHAFT
:製作協力:ブシロード、ムービック
:製作:アニプレックス、SHAFT、集英社、MBS

*オープニングテーマ1「CLICK」
:作詞・作曲・編曲:kz(livetune)
:歌:ClariS
:使用話数:#2～#14
-#1はオープニング映像をスタッフロールの後に放送

*オープニングテーマ2「STEP」
:作詞・作曲・編曲:kz(livetune)
:歌:ClariS
:使用話数:#15～#19
-#14はエンディングテーマとして使用
-#20はオープニングなし

*エンディングテーマ1「Heart Pattern」
:作詞・作曲:渡辺翔
:編曲:渡辺和紀
:歌:桐崎千棘(東山奈央)
:使用話数:#2～#7

*エンディングテーマ2「リカバーデコレーション」
:作詞・作曲:渡辺翔
:編曲:森谷敏紀
:歌:小野寺小咲(花澤香菜)
:使用話数:#8、#10～#13

*エンディングテーマ3「TRICK BOX」
:作詞・作曲:渡辺翔
:編曲:倉内達矢
:歌:鶫誠士郎(小松未可子)
:使用話数:#15、#17

*エンディングテーマ4「はなごのみ」
:作詞・作曲:渡辺翔
:編曲:前口渉
:歌:橘万里花(阿澄佳奈)
:使用話数:#18、#19
-#9、#16はエンディングテーマなし

*エンディングテーマ5「想像ダイアリー」
:作詞・作曲:渡辺翔
:編曲:清水哲平
:歌:桐崎千棘(東山奈央)、小野寺小咲(花澤香菜)、鶫誠士郎(小松未可子)、橘万里花(阿澄佳奈)
:使用話数:#20

*キャスト
:一条楽:内山昂輝
:桐崎千棘:東山奈央
:小野寺小咲:花澤香菜
:鶫誠士郎:小松未可子
:橘万里花:阿澄佳奈
:宮本るり:内山夕実
:舞子集:梶裕貴
:クロード:子安武人
:竜:檜山修之
:楽の父:緒方賢一
:千棘の父:江原正士
:小咲の母:大原さやか
:万里花の父:立木文彦
:本田:大地葉
:キョーコ:生天目仁美
:楽(子供時代):関根明良

*提供バックイラスト
:#1:こやまひろかず(TYPE-MOON)
:#2:ウエダハジメ
:#3:日向悠二
:#4:氷川へきる
:#5:カントク
:#6:あぼしまこ
:#7:ろびこ
:#8:まごまご
:#9:なもり
:#10:吉田明彦
:#11:あまからするめ
:#12:水玉子
:#13:ぽよよん♡ろっく
:#14:蒼樹うめ
:#15:岸田メル
:#16:黒星紅白
:#17:toi8
:#18:矢吹健太朗
:#19:KEI
:#20:古味直志
            </Comment>
            <Cat>1</Cat>
            <TitleFlag>0</TitleFlag>
            <FirstYear>2014</FirstYear>
            <FirstMonth>1</FirstMonth>
            <FirstEndYear>2014</FirstEndYear>
            <FirstEndMonth>5</FirstEndMonth>
            <FirstCh>TOKYO MX、とちぎテレビ、群馬テレビ</FirstCh>
            <Keywords></Keywords>
            <UserPoint>451</UserPoint>
            <UserPointRank>4</UserPointRank>
            <SubTitles>
*01*ヤクソク
*02*ソウグウ
*03*ニタモノ
*04*ホウモン
*05*スイエイ
*06*カシカリ
*07*ライバル
*08*シアワセ
*09*オンセン
*10*クジビキ
*11*オイワイ
*12*カクニン
*13*ホウカゴ
*14*シュラバ
*15*サンボン
*16*タイフウ
*17*エンニチ
*18*ウミベデ
*19*エンゲキ
*20*ホンバン
            </SubTitles>
        </TitleItem>
    </TitleItems>
</TitleLookupResponse>`
