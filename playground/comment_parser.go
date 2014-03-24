package main

import "regexp"
import "fmt"
import "strings"

func main() {
	re, err := regexp.Compile("[\n]{2}")
	if err != nil {
		panic(err)
	}
	blocks := re.Split(getStr(), -1)
	for _, block := range blocks {
		h, v := extractHeaderAndValue(block)
		fmt.Println("HEAD\t" + h)
		fmt.Println("VALUE\t" + v)
	}
}

func extractHeaderAndValue(str string) (header, value string) {
	_r, _e := regexp.Compile("^([^\n]+)\n(.*)")
	if _e != nil {
		panic(_e)
	}
	submatch := _r.FindSubmatch([]byte(str))
	header = string(submatch[1])
	rplcr := strings.NewReplacer(header+"\n", "")
	value = rplcr.Replace(str)
	return header, value
}
func getStr() (str string) {
	str = "*リンク\n"
	str = str + "-[[公式 http://inarikonkon.jp/]]\n"
	str = str + "-[[TOKYO MX http://s.mxtv.jp/inarikonkon/]]\n"
	str = str + "-[[BS11デジタル http://www.bs11.jp/anime/2422/]]\n"
	str = str + "\n"
	str = str + "*スタッフ\n"
	str = str + ":原作:よしだもろへ\n"
	str = str + ":掲載誌:ヤングエース(KADOKAWA(角川書店))\n"
	str = str + ":監督:高橋亨\n"
	str = str + ":助監督:岡本英樹\n"
	str = str + ":シリーズ構成:待田堂子\n"
	str = str + ":キャラクターデザイン・総作画監督:高品有桂\n"
	str = str + ":メインアニメーター:沈宏\n"
	str = str + ":プロップデザイン:奥田万つ里\n"
	str = str + ":色彩設計:大内綾\n"
	str = str + ":美術監督:大西穰\n"
	str = str + ":美術設定:坂本竜\n"
	str = str + ":撮影監督:津田涼介\n"
	str = str + ":編集:西山茂\n"
	str = str + ":音響監督:たなかかずや\n"
	str = str + ":音響効果:出雲範子\n"
	str = str + ":音響制作:ダックスプロダクション\n"
	str = str + ":音楽:妹尾武\n"
	str = str + ":音楽制作:flyingDOG\n"
	str = str + ":アニメーション制作:プロダクションアイムズ\n"
	str = str + ":製作:[[いなり、こんこん、恋いろは。製作委員会]]\n"
	str = str + "\n"
	str = str + "*オープニングテーマ「今日に恋色」\n"
	str = str + ":作詞・作曲・編曲:kz(livetune)\n"
	str = str + ":歌:May'n\n"
	str = str + ":使用話数:#1～#9\n"
	str = str + "-#10はエンディングテーマとして使用\n"
	str = str + "\n"
	str = str + "*エンディングテーマ1「SAVED.」\n"
	str = str + ":作詞・作曲:鈴木祥子\n"
	str = str + ":編曲:山本隆二\n"
	str = str + ":歌:坂本真綾\n"
	str = str + ":使用話数:#1～#8\n"
	str = str + "-#10はオープニングテーマとして使用\n"
	str = str + "\n"
	str = str + "*エンディングテーマ2「誰よりも大切な人へ」\n"
	str = str + ":作詞・作曲・編曲:妹尾武\n"
	str = str + ":歌:宇迦之御魂之神(桑島法子)\n"
	str = str + ":使用話数:#9\n"
	str = str + "\n"
	str = str + "*挿入歌「涙はらはら -Reprise-」\n"
	str = str + ":作詞・作曲:つじあやの\n"
	str = str + ":編曲:曽我淳一\n"
	str = str + ":歌:伏見いなり(大空直美)\n"
	str = str + ":使用話数:#10\n"
	str = str + "\n"
	str = str + "*キャスト\n"
	str = str + ":伏見いなり:大空直美\n"
	str = str + ":宇迦之御魂之神(うか様):桑島法子\n"
	str = str + ":丹波橋紅司:岡本寛志\n"
	str = str + ":伏見燈日:上田燿司\n"
	str = str + ":墨染朱美:野水伊織\n"
	str = str + ":三条京子:池辺久美子\n"
	str = str + ":丸太町ちか:佐土原かおり\n"
	str = str + ":大年神:子安武人\n"
	str = str + ":大宮能売神:三上枝織\n"
	str = str + ":天照大御神:磯辺万沙子\n"
	str = str + ":シシ:日野聡\n"
	str = str + ":ロロ:花江夏樹\n"
	str = str + ":コン:原紗友里\n"
	str = str + ":伏見葛葉:仲みのり\n"
	str = str + ":伏見初午:岡哲也\n"
	str = str + ":橋本創:北纓裕紀\n"
	str = str + ":出町柳一:三輪瑛\n"
	str = str + ":木幡先生:サエキトモ\n"
	str = str + ":神大市比売:真田アサミ\n"
	str = str + ":須佐之男:黒田崇矢\n"
	return
}
