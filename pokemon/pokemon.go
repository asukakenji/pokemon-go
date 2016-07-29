package pokemon

import (
	"math"
	"strings"

	"github.com/asukakenji/pokemon-go/generic"
	"github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/move"

	"github.com/asukakenji/pokemon-go/pokemon/internal/de"
	"github.com/asukakenji/pokemon-go/pokemon/internal/en"
	"github.com/asukakenji/pokemon-go/pokemon/internal/es"
	"github.com/asukakenji/pokemon-go/pokemon/internal/fr"
	"github.com/asukakenji/pokemon-go/pokemon/internal/it"
	"github.com/asukakenji/pokemon-go/pokemon/internal/ja"
	"github.com/asukakenji/pokemon-go/pokemon/internal/ko"
	"github.com/asukakenji/pokemon-go/pokemon/internal/zh-CHS"
	"github.com/asukakenji/pokemon-go/pokemon/internal/zh-CHT"
	"github.com/asukakenji/pokemon-go/pokemon/internal/zh-CN"
	"github.com/asukakenji/pokemon-go/pokemon/internal/zh-HK"
	"github.com/asukakenji/pokemon-go/pokemon/internal/zh-TW"

	typ "github.com/asukakenji/pokemon-go/type"
	"github.com/asukakenji/pokemon-go/weak"
)

// Pokemon
type Pokemon int16

const (
	None           Pokemon = iota
	Bulbasaur              // 001: フシギダネ
	Ivysaur                // 002: フシギソウ
	Venusaur               // 003: フシギバナ -> メガフシギバナ (Mega Venusaur)
	Charmander             // 004: ヒトカゲ
	Charmeleon             // 005: リザード
	Charizard              // 006: リザードン -> メガリザードンＸ (Mega Charizard X) / メガリザードンＹ (Mega Charizard Y)
	Squirtle               // 007: ゼニガメ
	Wartortle              // 008: カメール
	Blastoise              // 009: カメックス -> メガカメックス (Mega Blastoise)
	Caterpie               // 010: キャタピー
	Metapod                // 011: トランセル
	Butterfree             // 012: バタフリー
	Weedle                 // 013: ビードル
	Kakuna                 // 014: コクーン
	Beedrill               // 015: スピアー -> メガスピアー (Mega Beedrill)
	Pidgey                 // 016: ポッポ
	Pidgeotto              // 017: ピジョン
	Pidgeot                // 018: ピジョット -> メガピジョット (Mega Pidgeot)
	Rattata                // 019: コラッタ
	Raticate               // 020: ラッタ
	Spearow                // 021: オニスズメ
	Fearow                 // 022: オニドリル
	Ekans                  // 023: アーボ
	Arbok                  // 024: アーボック
	Pikachu                // 025: ピカチュウ
	Raichu                 // 026: ライチュウ
	Sandshrew              // 027: サンド
	Sandslash              // 028: サンドパン
	Nidoran_female         // 029: ニドラン♀ (Nidoran♀)
	Nidorina               // 030: ニドリーナ
	Nidoqueen              // 031: ニドクイン
	Nidoran_male           // 032: ニドラン♂ (Nidoran♂)
	Nidorino               // 033: ニドリーノ
	Nidoking               // 034: ニドキング
	Clefairy               // 035: ピッピ
	Clefable               // 036: ピクシー
	Vulpix                 // 037: ロコン
	Ninetales              // 038: キュウコン
	Jigglypuff             // 039: プリン
	Wigglytuff             // 040: プクリン
	Zubat                  // 041: ズバット
	Golbat                 // 042: ゴルバット
	Oddish                 // 043: ナゾノクサ
	Gloom                  // 044: クサイハナ
	Vileplume              // 045: ラフレシア
	Paras                  // 046: パラス
	Parasect               // 047: パラセクト
	Venonat                // 048: コンパン
	Venomoth               // 049: モルフォン
	Diglett                // 050: ディグダ
	Dugtrio                // 051: ダグトリオ
	Meowth                 // 052: ニャース
	Persian                // 053: ペルシアン
	Psyduck                // 054: コダック
	Golduck                // 055: ゴルダック
	Mankey                 // 056: マンキー
	Primeape               // 057: オコリザル
	Growlithe              // 058: ガーディ
	Arcanine               // 059: ウインディ
	Poliwag                // 060: ニョロモ
	Poliwhirl              // 061: ニョロゾ
	Poliwrath              // 062: ニョロボン
	Abra                   // 063: ケーシィ
	Kadabra                // 064: ユンゲラー
	Alakazam               // 065: フーディン -> メガフーディン (Mega Alakazam)
	Machop                 // 066: ワンリキー
	Machoke                // 067: ゴーリキー
	Machamp                // 068: カイリキー
	Bellsprout             // 069: マダツボミ
	Weepinbell             // 070: ウツドン
	Victreebel             // 071: ウツボット
	Tentacool              // 072: メノクラゲ
	Tentacruel             // 073: ドククラゲ
	Geodude                // 074: イシツブテ
	Graveler               // 075: ゴローン
	Golem                  // 076: ゴローニャ
	Ponyta                 // 077: ポニータ
	Rapidash               // 078: ギャロップ
	Slowpoke               // 079: ヤドン
	Slowbro                // 080: ヤドラン -> メガヤドラン (Mega Slowbro)
	Magnemite              // 081: コイル
	Magneton               // 082: レアコイル
	Farfetch_d             // 083: カモネギ (Farfetch'd)
	Doduo                  // 084: ドードー
	Dodrio                 // 085: ドードリオ
	Seel                   // 086: パウワウ
	Dewgong                // 087: ジュゴン
	Grimer                 // 088: ベトベター
	Muk                    // 089: ベトベトン
	Shellder               // 090: シェルダー
	Cloyster               // 091: パルシェン
	Gastly                 // 092: ゴース
	Haunter                // 093: ゴースト
	Gengar                 // 094: ゲンガー -> メガゲンガー (Mega Gengar)
	Onix                   // 095: イワーク
	Drowzee                // 096: スリープ
	Hypno                  // 097: スリーパー
	Krabby                 // 098: クラブ
	Kingler                // 099: キングラー
	Voltorb                // 100: ビリリダマ
	Electrode              // 101: マルマイン
	Exeggcute              // 102: タマタマ
	Exeggutor              // 103: ナッシー
	Cubone                 // 104: カラカラ
	Marowak                // 105: ガラガラ
	Hitmonlee              // 106: サワムラー
	Hitmonchan             // 107: エビワラー
	Lickitung              // 108: ベロリンガ
	Koffing                // 109: ドガース
	Weezing                // 110: マタドガス
	Rhyhorn                // 111: サイホーン
	Rhydon                 // 112: サイドン
	Chansey                // 113: ラッキー
	Tangela                // 114: モンジャラ
	Kangaskhan             // 115: ガルーラ -> メガガルーラ (Mega Kangaskhan)
	Horsea                 // 116: タッツー
	Seadra                 // 117: シードラ
	Goldeen                // 118: トサキント
	Seaking                // 119: アズマオウ
	Staryu                 // 120: ヒトデマン
	Starmie                // 121: スターミー
	Mr_Mime                // 122: バリヤード (Mr. Mime)
	Scyther                // 123: ストライク
	Jynx                   // 124: ルージュラ
	Electabuzz             // 125: エレブー
	Magmar                 // 126: ブーバー
	Pinsir                 // 127: カイロス -> メガカイロス (Mega Pinsir)
	Tauros                 // 128: ケンタロス
	Magikarp               // 129: コイキング
	Gyarados               // 130: ギャラドス -> メガギャラドス (Mega Gyarados)
	Lapras                 // 131: ラプラス
	Ditto                  // 132: メタモン
	Eevee                  // 133: イーブイ
	Vaporeon               // 134: シャワーズ
	Jolteon                // 135: サンダース
	Flareon                // 136: ブースター
	Porygon                // 137: ポリゴン
	Omanyte                // 138: オムナイト
	Omastar                // 139: オムスター
	Kabuto                 // 140: カブト
	Kabutops               // 141: カブトプス
	Aerodactyl             // 142: プテラ -> メガプテラ (Mega Aerodactyl)
	Snorlax                // 143: カビゴン
	Articuno               // 144: フリーザー
	Zapdos                 // 145: サンダー
	Moltres                // 146: ファイヤー
	Dratini                // 147: ミニリュウ
	Dragonair              // 148: ハクリュー
	Dragonite              // 149: カイリュー
	Mewtwo                 // 150: ミュウツー -> メガミュウツーＸ (Mega Mewtwo X) / メガミュウツーＹ (Mega Mewtwo Y)
	Mew                    // 151: ミュウ
)

//go:generate stringer -type=Pokemon

func PackageName(l lang.Language) string {
	switch l {
	case lang.Japanese:
		return "ポケモン"
	case lang.English:
		return "Pokémon"
	case lang.French:
		return "Pokémon"
	case lang.German:
		return "Pokémon"
	case lang.Italian:
		return "Pokémon"
	case lang.Korean:
		return "포켓몬"
	case lang.Spanish:
		return "Pokémon"
	case lang.ChineseSimplified:
		return "精灵宝可梦"
	case lang.ChineseTraditional:
		return "精靈寶可夢"
	case lang.ChineseChina:
		return "神奇宝贝"
	case lang.ChineseHongKong:
		return "寵物小精靈"
	case lang.ChineseTaiwan:
		return "神奇寶貝"
	default:
		return ""
	}
}

type IndividualValues struct {
	Stamina int // 0 <= x <= 15
	Attack  int // 0 <= x <= 15
	Defense int // 0 <= x <= 15
}

func ByCodeName(codeName string) Pokemon {
	result := None
	All().ForEach(func(p Pokemon) {
		if p.String() == codeName {
			result = p
		}
	})
	return result
}

func (p Pokemon) Id() int {
	return int(p)
}

func (p Pokemon) EvolveFrom() Pokemon {
	return p.self().evolveFrom
}

func (p Pokemon) EvolveTo() Iterable {
	predicate := func(p2 Pokemon) bool {
		return p2.EvolveFrom() == p
	}
	return All().Filter(predicate)
}

func (p Pokemon) LocalName(l lang.Language) string {
	result := ""
	switch l {
	case lang.Japanese:
		result = ja.Pokemon(p).String()
	case lang.English:
		result = en.Pokemon(p).String()
	case lang.French:
		result = fr.Pokemon(p).String()
	case lang.German:
		result = de.Pokemon(p).String()
	case lang.Italian:
		result = it.Pokemon(p).String()
	case lang.Korean:
		result = ko.Pokemon(p).String()
	case lang.Spanish:
		result = es.Pokemon(p).String()
	case lang.ChineseSimplified:
		result = zhCHS.Pokemon(p).String()
	case lang.ChineseTraditional:
		result = zhCHT.Pokemon(p).String()
	case lang.ChineseChina:
		result = zhCN.Pokemon(p).String()
	case lang.ChineseHongKong:
		result = zhHK.Pokemon(p).String()
	case lang.ChineseTaiwan:
		result = zhTW.Pokemon(p).String()
	}
	result = strings.Replace(result, "_female_", "♀", -1)
	result = strings.Replace(result, "_male_", "♂", -1)
	result = strings.Replace(result, "_apostrophe_", "'", -1)
	result = strings.Replace(result, "_full_stop_", ".", -1)
	result = strings.Replace(result, "_space_", " ", -1)
	result = strings.Replace(result, "_fullwidth_digit_three_", "３", -1)
	return result
}

func (p Pokemon) BaseStamina() int {
	return int(p.self().baseStamina)
}

func (p Pokemon) BaseAttack() int {
	return int(p.self().baseAttack)
}

func (p Pokemon) BaseDefense() int {
	return int(p.self().baseDefense)
}

func (p Pokemon) CombatPower(iv IndividualValues, level float32) int {
	if level < 1.0 || level > 40.0 {
		panic("Invalid level")
	}
	stamina := float64(p.BaseStamina() + iv.Stamina)
	attack := float64(p.BaseAttack() + iv.Attack)
	defense := float64(p.BaseDefense() + iv.Defense)
	multiplier := getLevelItem(level).combatPowerMultiplier
	if result := int(attack * math.Sqrt(defense*stamina) * multiplier * multiplier / 10.0); result > 10 {
		return result
	}
	return 10
}

func (p Pokemon) HitPoints(iv IndividualValues, level float32) int {
	if level < 1.0 || level > 40.0 {
		panic("Invalid level")
	}
	stamina := float64(p.BaseStamina() + iv.Stamina)
	multiplier := getLevelItem(level).combatPowerMultiplier
	if result := int(stamina * multiplier); result > 10 {
		return result
	}
	return 10
}

func (p Pokemon) Type1() typ.Type {
	return p.self().type1
}

func (p Pokemon) Type2() typ.Type {
	return p.self().type2
}

func (p Pokemon) Weight() float64 {
	return p.self().weight
}

func (p Pokemon) Height() float64 {
	return p.self().height
}

func (p Pokemon) StardustAndCandyToPowerUp(level float32) (int, Pokemon, int) {
	if level < 1.0 || level > 40.0 {
		panic("Invalid level")
	}
	levelItem0 := getLevelItem(level)
	return int(levelItem0.stardustToPowerUp), p.self().candyType, int(levelItem0.candyToPowerUp)
}

func (p Pokemon) CandyToEvolve() (Pokemon, int) {
	return p.self().candyType, int(p.self().candyToEvolve)
}

// Multiplier to be applied when this Pokemon is attacked by the specified type of moves
func (p Pokemon) Multiplier(moveType typ.Type) float64 {
	return p.Type1().AttackedBy(moveType).Multiplier() * p.Type2().AttackedBy(moveType).Multiplier()
}

func (p Pokemon) Weaknesses() weak.Iterable {
	mapper := func(t typ.Type) interface{} {
		return weak.Weakness{t, p.Multiplier(t)}
	}
	predicate := func(w weak.Weakness) bool {
		return w.Multiplier > 1.0
	}
	return weak.Wrap(typ.All().Map(mapper)).Filter(predicate).Sort(weak.ByMultiplier(generic.Descending))
}

func (p Pokemon) Moves() move.Iterable {
	return move.Slice(moves[p.Id()])
}

func (p Pokemon) StandardMoves() move.Iterable {
	return p.Moves().Filter(func(m move.Move) bool {
		_, ok := m.(move.StandardMove)
		return ok
	})
}

func (p Pokemon) SpecialMoves() move.Iterable {
	return p.Moves().Filter(func(m move.Move) bool {
		_, ok := m.(move.SpecialMove)
		return ok
	})
}

func (p Pokemon) self() *_pokemon {
	return pokemons[p.Id()]
}

func getLevelItem(level float32) levelItem {
	return levelTable[int((level-1.0)*2)]
}

func LevelsByStardust(stardust int) []float32 {
	result := make([]float32, 0)
	for _, li := range levelTable {
		if int(li.stardustToPowerUp) == stardust {
			result = append(result, li.level)
		}
	}
	return result
}

func LevelsByStardustAndCandy(stardust, candy int) []float32 {
	result := make([]float32, 0)
	for _, li := range levelTable {
		if int(li.stardustToPowerUp) == stardust && int(li.candyToPowerUp) == candy {
			result = append(result, li.level)
		}
	}
	return result
}
