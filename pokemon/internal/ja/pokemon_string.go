// Code generated by "stringer -type=Pokemon"; DO NOT EDIT

package ja

import "fmt"

const _Pokemon_name = "フシギダネフシギソウフシギバナヒトカゲリザードリザードンゼニガメカメールカメックスキャタピートランセルバタフリービードルコクーンスピアーポッポピジョンピジョットコラッタラッタオニスズメオニドリルアーボアーボックピカチュウライチュウサンドサンドパンニドラン_female_ニドリーナニドクインニドラン_male_ニドリーノニドキングピッピピクシーロコンキュウコンプリンプクリンズバットゴルバットナゾノクサクサイハナラフレシアパラスパラセクトコンパンモルフォンディグダダグトリオニャースペルシアンコダックゴルダックマンキーオコリザルガーディウインディニョロモニョロゾニョロボンケーシィユンゲラーフーディンワンリキーゴーリキーカイリキーマダツボミウツドンウツボットメノクラゲドククラゲイシツブテゴローンゴローニャポニータギャロップヤドンヤドランコイルレアコイルカモネギドードードードリオパウワウジュゴンベトベターベトベトンシェルダーパルシェンゴースゴーストゲンガーイワークスリープスリーパークラブキングラービリリダママルマインタマタマナッシーカラカラガラガラサワムラーエビワラーベロリンガドガースマタドガスサイホーンサイドンラッキーモンジャラガルーラタッツーシードラトサキントアズマオウヒトデマンスターミーバリヤードストライクルージュラエレブーブーバーカイロスケンタロスコイキングギャラドスラプラスメタモンイーブイシャワーズサンダースブースターポリゴンオムナイトオムスターカブトカブトプスプテラカビゴンフリーザーサンダーファイヤーミニリュウハクリューカイリューミュウツーミュウ"

var _Pokemon_index = [...]uint16{0, 15, 30, 45, 57, 69, 84, 96, 108, 123, 138, 153, 168, 180, 192, 204, 213, 225, 240, 252, 261, 276, 291, 300, 315, 330, 345, 354, 369, 389, 404, 419, 437, 452, 467, 476, 488, 497, 512, 521, 533, 545, 560, 575, 590, 605, 614, 629, 641, 656, 668, 683, 695, 710, 722, 737, 749, 764, 776, 791, 803, 815, 830, 842, 857, 872, 887, 902, 917, 932, 944, 959, 974, 989, 1004, 1016, 1031, 1043, 1058, 1067, 1079, 1088, 1103, 1115, 1127, 1142, 1154, 1166, 1181, 1196, 1211, 1226, 1235, 1247, 1259, 1271, 1283, 1298, 1307, 1322, 1337, 1352, 1364, 1376, 1388, 1400, 1415, 1430, 1445, 1457, 1472, 1487, 1499, 1511, 1526, 1538, 1550, 1562, 1577, 1592, 1607, 1622, 1637, 1652, 1667, 1679, 1691, 1703, 1718, 1733, 1748, 1760, 1772, 1784, 1799, 1814, 1829, 1841, 1856, 1871, 1880, 1895, 1904, 1916, 1931, 1943, 1958, 1973, 1988, 2003, 2018, 2027}

func (i Pokemon) String() string {
	i -= 1
	if i < 0 || i >= Pokemon(len(_Pokemon_index)-1) {
		return fmt.Sprintf("Pokemon(%d)", i+1)
	}
	return _Pokemon_name[_Pokemon_index[i]:_Pokemon_index[i+1]]
}
