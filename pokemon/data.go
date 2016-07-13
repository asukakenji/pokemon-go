package pokemon

import (
	"github.com/asukakenji/pokemon-go/typ"
)

type _pokemon struct {
	id              Pokemon
	baseCombatPower int16
	baseHitPoints   int16
	type1           typ.Type
	type2           typ.Type
	weight          float64 // in kg
	height          float64 // in m
	candyType       Pokemon
	candyToPowerUp  int16
	candyToEvolve   int16
	moves           []int16
}

var pokemons = [...]*_pokemon{
	nil,
	&_pokemon{Bulbasaur, -1, -1, typ.Grass, typ.Poison, 6.9, 0.7, Bulbasaur, 1, 25, nil},          // 001: Bulbasaur (フシギダネ), SD:200
	&_pokemon{Ivysaur, -1, -1, typ.Grass, typ.Poison, 13.0, 1.0, Bulbasaur, 1, -1, nil},           // 002: Ivysaur (フシギソウ)
	&_pokemon{Venusaur, -1, -1, typ.Grass, typ.Poison, 100.0, 2.0, Bulbasaur, 1, 0, nil},          // 003: Venusaur (フシギバナ)
	&_pokemon{Charmander, 12, 10, typ.Fire, typ.None, 8.5, 0.6, Charmander, 1, 25, nil},           // 004: Charmander (ヒトカゲ), SD: 200, CP: [12,,,35,,,], HP: [10,,,13,,,]
	&_pokemon{Charmeleon, -1, -1, typ.Fire, typ.None, 19.0, 1.1, Charmander, 1, -1, nil},          // 005: Charmeleon (リザード)
	&_pokemon{Charizard, -1, -1, typ.Fire, typ.Flying, 90.5, 1.7, Charmander, 1, 0, nil},          // 006: Charizard (リザードン)
	&_pokemon{Squirtle, -1, -1, typ.Water, typ.None, 9.0, 0.5, Squirtle, 1, 25, nil},              // 007: Squirtle (ゼニガメ), SD: 800, CP: [,,,161,,,], HP: [,,,34,,,]
	&_pokemon{Wartortle, -1, -1, typ.Water, typ.None, 22.5, 1.0, Squirtle, 1, -1, nil},            // 008: Wartortle (カメール)
	&_pokemon{Blastoise, -1, -1, typ.Water, typ.None, 85.5, 1.6, Squirtle, 1, 0, nil},             // 009: Blastoise (カメックス)
	&_pokemon{Caterpie, 10, 10, typ.Bug, typ.None, 2.9, 0.3, Caterpie, 1, 12, nil},                // 010: Caterpie (キャタピー), SD: 200, CP: [10,,,], HP: [10,,,]
	&_pokemon{Metapod, -1, -1, typ.Bug, typ.None, 9.9, 0.7, Caterpie, 1, -1, nil},                 // 011: Metapod (トランセル)
	&_pokemon{Butterfree, -1, -1, typ.Bug, typ.Flying, 32.0, 1.1, Caterpie, 1, 0, nil},            // 012: Butterfree (バタフリー)
	&_pokemon{Weedle, 10, 10, typ.Bug, typ.Poison, 3.2, 0.3, Weedle, 1, 12, nil},                  // 013: Weedle (ビードル), SD: 200, CP: [,,,14,15,,,], HP: [,,,13,14,,,]
	&_pokemon{Kakuna, -1, -1, typ.Bug, typ.Poison, 10.0, 0.6, Weedle, 1, 50, nil},                 // 014: Kakuna (コクーン), SD: 600, CP: [,,,64,69,,,], HP: [,,,34,36,,,]
	&_pokemon{Beedrill, -1, -1, typ.Bug, typ.Poison, 29.5, 1.0, Weedle, 1, 0, nil},                // 015: Beedrill (スピアー), SD: 800, CP: [,,,240,258,,,], HP: [,,,50,52,,,], Poison Jab (Poison, 0, 15), X-Scissor (Bug, 3, 30)
	&_pokemon{Pidgey, 10, 10, typ.Normal, typ.Flying, 1.8, 0.3, Pidgey, 1, 12, nil},               // 016: Pidgey (ポッポ), SD: 200, CP: [10,,,23,,,], HP: [10,,,14,,,], Quick Attack (Normal, 0, 10), Aerial Ace (Flying, 4, 25), Twister (Dragon, 5, 15)
	&_pokemon{Pidgeotto, 15, 12, typ.Normal, typ.Flying, 30.0, 1.1, Pidgey, 1, 50, nil},           // 017: Pidgeotto (ピジョン), SD: 15:200 / 117:400 CP: [15,,,110,117(125),,,], HP: [12,,,32,34,,,]
	&_pokemon{Pidgeot, -1, -1, typ.Normal, typ.Flying, 39.5, 1.5, Pidgey, 1, 0, nil},              // 018: Pidgeot (ピジョット), SD: 234:400 / 263:600, CP:[,,,204,234,263,,,352,,,], HP: [,,,44,47,50,,,58,,,], Wing Attack (Flying, 0, 12), Hurricane (Flying, 1, 60)
	&_pokemon{Rattata, -1, -1, typ.Normal, typ.None, 3.5, 0.3, Rattata, 1, 25, nil},               // 019: Rattata (コラッタ), SD: 200, CP: [,,,21,,,], HP: [,,,11,,,]
	&_pokemon{Raticate, -1, -1, typ.Normal, typ.None, 18.5, 0.7, Rattata, 1, 0, nil},              // 020: Raticate (ラッタ), SD: 800, CP: [,,,296,,,], HP: [,,,43,,,], Quick Attack (Normal, 0, 10), Dig (Ground, 3, 45)
	&_pokemon{Spearow, 10, 10, typ.Normal, typ.Flying, 2.0, 0.3, Spearow, 1, 50, nil},             // 021: Spearow (オニスズメ), SD: 10:200 / 400, CP: [10,,,63,,,], HP: [10,,,24,,,]
	&_pokemon{Fearow, -1, -1, typ.Normal, typ.Flying, 38.0, 1.2, Spearow, 1, 0, nil},              // 022: Fearow (オニドリル)
	&_pokemon{Ekans, -1, -1, typ.Poison, typ.None, 6.9, 2.0, Ekans, 1, -1, nil},                   // 023: Ekans (アーボ)
	&_pokemon{Arbok, -1, -1, typ.Poison, typ.None, 65.0, 3.5, Ekans, 1, 0, nil},                   // 024: Arbok (アーボック)
	&_pokemon{Pikachu, 11, 10, typ.Electric, typ.None, 6.0, 0.4, Pikachu, 1, 50, nil},             // 025: Pikachu (ピカチュウ), 172 -> 025, SD: 200, CP: [11,,,], HP: [10,,,]
	&_pokemon{Raichu, -1, -1, typ.Electric, typ.None, 30.0, 0.8, Pikachu, 1, 0, nil},              // 026: Raichu (ライチュウ)
	&_pokemon{Sandshrew, -1, -1, typ.Ground, typ.None, 12.0, 0.6, Sandshrew, 1, -1, nil},          // 027: Sandshrew (サンド)
	&_pokemon{Sandslash, -1, -1, typ.Ground, typ.None, 29.5, 1.0, Sandshrew, 1, 0, nil},           // 028: Sandslash (サンドパン)
	&_pokemon{Nidoran_female, -1, -1, typ.Poison, typ.None, 7.0, 0.4, Nidoran_female, 1, -1, nil}, // 029: Nidoran♀ (ニドラン♀)
	&_pokemon{Nidorina, -1, -1, typ.Poison, typ.None, 20.0, 0.8, Nidoran_female, 1, -1, nil},      // 030: Nidorina (ニドリーナ)
	&_pokemon{Nidoqueen, -1, -1, typ.Poison, typ.Ground, 60.0, 1.3, Nidoran_female, 1, 0, nil},    // 031: Nidoqueen (ニドクイン)
	&_pokemon{Nidoran_male, -1, -1, typ.Poison, typ.None, 9.0, 0.5, Nidoran_male, 1, -1, nil},     // 032: Nidoran♂ (ニドラン♂)
	&_pokemon{Nidorino, -1, -1, typ.Poison, typ.None, 19.5, 0.9, Nidoran_male, 1, -1, nil},        // 033: Nidorino (ニドリーノ)
	&_pokemon{Nidoking, -1, -1, typ.Poison, typ.Ground, 62.0, 1.4, Nidoran_male, 1, 0, nil},       // 034: Nidoking (ニドキング)
	&_pokemon{Clefairy, -1, -1, typ.Fairy, typ.None, 7.5, 0.6, Clefairy, 1, -1, nil},              // 035: Clefairy (ピッピ), 173 -> 035
	&_pokemon{Clefable, -1, -1, typ.Fairy, typ.None, 40.0, 1.3, Clefairy, 1, 0, nil},              // 036: Clefable (ピクシー)
	&_pokemon{Vulpix, -1, -1, typ.Fire, typ.None, 9.9, 0.6, Vulpix, 1, -1, nil},                   // 037: Vulpix (ロコン)
	&_pokemon{Ninetales, -1, -1, typ.Fire, typ.None, 19.9, 1.1, Vulpix, 1, 0, nil},                // 038: Ninetales (キュウコン)
	&_pokemon{Jigglypuff, -1, -1, typ.Normal, typ.Fairy, 5.5, 0.5, Jigglypuff, 1, 50, nil},        // 039: Jigglypuff (プリン), 174 -> 039, SD: 600, CP: [,,,110,,,], HP: [,,,69,,,]
	&_pokemon{Wigglytuff, -1, -1, typ.Normal, typ.Fairy, 12.0, 1.0, Jigglypuff, 1, 0, nil},        // 040: Wigglytuff (プクリン)
	&_pokemon{Zubat, -1, -1, typ.Poison, typ.Flying, 7.5, 0.8, Zubat, 1, -1, nil},                 // 041: Zubat (ズバット)
	&_pokemon{Golbat, -1, -1, typ.Poison, typ.Flying, 55.0, 1.6, Zubat, 1, 0, nil},                // 042: Golbat (ゴルバット)
	&_pokemon{Oddish, 13, 10, typ.Grass, typ.Poison, 5.4, 0.5, Oddish, 1, 25, nil},                // 043: Oddish (ナゾノクサ), SD: 200, CP: [13,,,], HP: [10,,,]
	&_pokemon{Gloom, -1, -1, typ.Grass, typ.Poison, 8.6, 0.8, Oddish, 1, -1, nil},                 // 044: Gloom (クサイハナ)
	&_pokemon{Vileplume, -1, -1, typ.Grass, typ.Poison, 18.6, 1.2, Oddish, 1, 0, nil},             // 045: Vileplume (ラフレシア)
	&_pokemon{Paras, -1, -1, typ.Bug, typ.Grass, 5.4, 0.3, Paras, 1, -1, nil},                     // 046: Paras (パラス)
	&_pokemon{Parasect, -1, -1, typ.Bug, typ.Grass, 29.5, 1.0, Paras, 1, 0, nil},                  // 047: Parasect (パラセクト)
	&_pokemon{Venonat, -1, -1, typ.Bug, typ.Poison, 30.0, 1.0, Venonat, 1, -1, nil},               // 048: Venonat (コンパン)
	&_pokemon{Venomoth, -1, -1, typ.Bug, typ.Poison, 12.5, 1.5, Venonat, 1, 0, nil},               // 049: Venomoth (モルフォン)
	&_pokemon{Diglett, -1, -1, typ.Ground, typ.None, 0.8, 0.2, Diglett, 1, -1, nil},               // 050: Diglett (ディグダ)
	&_pokemon{Dugtrio, -1, -1, typ.Ground, typ.None, 33.3, 0.7, Diglett, 1, 0, nil},               // 051: Dugtrio (ダグトリオ)
	&_pokemon{Meowth, -1, -1, typ.Normal, typ.None, 4.2, 0.4, Meowth, 1, -1, nil},                 // 052: Meowth (ニャース)
	&_pokemon{Persian, -1, -1, typ.Normal, typ.None, 32.0, 1.0, Meowth, 1, 0, nil},                // 053: Persian (ペルシアン)
	&_pokemon{Psyduck, -1, -1, typ.Water, typ.None, 19.6, 0.8, Psyduck, 1, -1, nil},               // 054: Psyduck (コダック)
	&_pokemon{Golduck, -1, -1, typ.Water, typ.None, 76.6, 1.7, Psyduck, 1, 0, nil},                // 055: Golduck (ゴルダック)
	&_pokemon{Mankey, -1, -1, typ.Fighting, typ.None, 28.0, 0.5, Mankey, 1, -1, nil},              // 056: Mankey (マンキー)
	&_pokemon{Primeape, -1, -1, typ.Fighting, typ.None, 32.0, 1.0, Mankey, 1, 0, nil},             // 057: Primeape (オコリザル)
	&_pokemon{Growlithe, -1, -1, typ.Fire, typ.None, 19.0, 0.7, Growlithe, 1, -1, nil},            // 058: Growlithe (ガーディ)
	&_pokemon{Arcanine, -1, -1, typ.Fire, typ.None, 155.0, 1.9, Growlithe, 1, 0, nil},             // 059: Arcanine (ウインディ)
	&_pokemon{Poliwag, -1, -1, typ.Water, typ.None, 12.4, 0.6, Poliwag, 1, -1, nil},               // 060: Poliwag (ニョロモ)
	&_pokemon{Poliwhirl, -1, -1, typ.Water, typ.None, 20.0, 1.0, Poliwag, 1, -1, nil},             // 061: Poliwhirl (ニョロゾ)
	&_pokemon{Poliwrath, -1, -1, typ.Water, typ.Fighting, 54.0, 1.3, Poliwag, 1, 0, nil},          // 062: Poliwrath (ニョロボン)
	&_pokemon{Abra, -1, -1, typ.Psychic, typ.None, 19.5, 0.9, Abra, 1, 25, nil},                   // 063: Abra (ケーシィ), SD: 400, CP: [,,,54,,,], HP: [,,,13,,,]
	&_pokemon{Kadabra, -1, -1, typ.Psychic, typ.None, 56.5, 1.3, Abra, 1, -1, nil},                // 064: Kadabra (ユンゲラー)
	&_pokemon{Alakazam, -1, -1, typ.Psychic, typ.None, 48.0, 1.5, Abra, 1, 0, nil},                // 065: Alakazam (フーディン)
	&_pokemon{Machop, -1, -1, typ.Fighting, typ.None, 19.5, 0.8, Machop, 1, -1, nil},              // 066: Machop (ワンリキー)
	&_pokemon{Machoke, -1, -1, typ.Fighting, typ.None, 70.5, 1.5, Machop, 1, -1, nil},             // 067: Machoke (ゴーリキー)
	&_pokemon{Machamp, -1, -1, typ.Fighting, typ.None, 130.0, 1.6, Machop, 1, 0, nil},             // 068: Machamp (カイリキー)
	&_pokemon{Bellsprout, -1, -1, typ.Grass, typ.Poison, 4.0, 0.7, Bellsprout, 1, 25, nil},        // 069: Bellsprout (マダツボミ), SD: 400, CP: [,,,104,,,], HP: [,,,28,,,]
	&_pokemon{Weepinbell, -1, -1, typ.Grass, typ.Poison, 6.4, 1.0, Bellsprout, 1, -1, nil},        // 070: Weepinbell (ウツドン)
	&_pokemon{Victreebel, -1, -1, typ.Grass, typ.Poison, 15.5, 1.7, Bellsprout, 1, 0, nil},        // 071: Victreebel (ウツボット)
	&_pokemon{Tentacool, -1, -1, typ.Water, typ.Poison, 45.5, 0.9, Tentacool, 1, -1, nil},         // 072: Tentacool (メノクラゲ)
	&_pokemon{Tentacruel, -1, -1, typ.Water, typ.Poison, 55.0, 1.6, Tentacool, 1, 0, nil},         // 073: Tentacruel (ドククラゲ)
	&_pokemon{Geodude, -1, -1, typ.Rock, typ.Ground, 20.0, 0.4, Geodude, 1, -1, nil},              // 074: Geodude (イシツブテ)
	&_pokemon{Graveler, -1, -1, typ.Rock, typ.Ground, 105.0, 1.0, Geodude, 1, -1, nil},            // 075: Graveler (ゴローン)
	&_pokemon{Golem, -1, -1, typ.Rock, typ.Ground, 300.0, 1.4, Geodude, 1, 0, nil},                // 076: Golem (ゴローニャ)
	&_pokemon{Ponyta, -1, -1, typ.Fire, typ.None, 30.0, 1.0, Ponyta, 1, -1, nil},                  // 077: Ponyta (ポニータ)
	&_pokemon{Rapidash, -1, -1, typ.Fire, typ.None, 95.0, 1.7, Ponyta, 1, 0, nil},                 // 078: Rapidash (ギャロップ)
	&_pokemon{Slowpoke, -1, -1, typ.Water, typ.Psychic, 36.0, 1.2, Slowpoke, 1, -1, nil},          // 079: Slowpoke (ヤドン)
	&_pokemon{Slowbro, -1, -1, typ.Water, typ.Psychic, 78.5, 1.6, Slowpoke, 1, 0, nil},            // 080: Slowbro (ヤドラン)
	&_pokemon{Magnemite, -1, -1, typ.Electric, typ.Steel, 6.0, 0.3, Magnemite, 1, 50, nil},        // 081: Magnemite (コイル), SD: 200, CP: [,,,38,,,], HP: [,,,10,,,]
	&_pokemon{Magneton, -1, -1, typ.Electric, typ.Steel, 60.0, 1.0, Magnemite, 1, 0, nil},         // 082: Magneton (レアコイル)
	&_pokemon{Farfetch_d, -1, -1, typ.Normal, typ.Flying, 15.0, 0.8, Farfetch_d, 1, 0, nil},       // 083: Farfetch'd (カモネギ)
	&_pokemon{Doduo, -1, -1, typ.Normal, typ.Flying, 39.2, 1.4, Doduo, 1, -1, nil},                // 084: Doduo (ドードー)
	&_pokemon{Dodrio, -1, -1, typ.Normal, typ.Flying, 85.2, 1.8, Doduo, 1, 0, nil},                // 085: Dodrio (ドードリオ)
	&_pokemon{Seel, -1, -1, typ.Water, typ.None, 90.0, 1.1, Seel, 1, -1, nil},                     // 086: Seel (パウワウ)
	&_pokemon{Dewgong, -1, -1, typ.Water, typ.Ice, 120.0, 1.7, Seel, 1, 0, nil},                   // 087: Dewgong (ジュゴン)
	&_pokemon{Grimer, -1, -1, typ.Poison, typ.None, 30.0, 0.9, Grimer, 1, -1, nil},                // 088: Grimer (ベトベター)
	&_pokemon{Muk, -1, -1, typ.Poison, typ.None, 30.0, 1.2, Grimer, 1, 0, nil},                    // 089: Muk (ベトベトン)
	&_pokemon{Shellder, -1, -1, typ.Water, typ.None, 4.0, 0.3, Shellder, 1, -1, nil},              // 090: Shellder (シェルダー)
	&_pokemon{Cloyster, -1, -1, typ.Water, typ.Ice, 132.5, 1.5, Shellder, 1, 0, nil},              // 091: Cloyster (パルシェン)
	&_pokemon{Gastly, -1, -1, typ.Ghost, typ.Poison, 0.1, 1.3, Gastly, 1, -1, nil},                // 092: Gastly (ゴース)
	&_pokemon{Haunter, -1, -1, typ.Ghost, typ.Poison, 0.1, 1.6, Gastly, 1, -1, nil},               // 093: Haunter (ゴースト)
	&_pokemon{Gengar, -1, -1, typ.Ghost, typ.Poison, 40.5, 1.5, Gastly, 1, 0, nil},                // 094: Gengar (ゲンガー)
	&_pokemon{Onix, -1, -1, typ.Rock, typ.Ground, 210.0, 8.8, Onix, 1, 0, nil},                    // 095: Onix (イワーク)
	&_pokemon{Drowzee, -1, -1, typ.Psychic, typ.None, 32.4, 1.0, Drowzee, 1, -1, nil},             // 096: Drowzee (スリープ)
	&_pokemon{Hypno, -1, -1, typ.Psychic, typ.None, 75.6, 1.6, Drowzee, 1, 0, nil},                // 097: Hypno (スリーパー)
	&_pokemon{Krabby, -1, -1, typ.Water, typ.None, 6.5, 0.4, Krabby, 1, -1, nil},                  // 098: Krabby (クラブ)
	&_pokemon{Kingler, -1, -1, typ.Water, typ.None, 60.0, 1.3, Krabby, 1, 0, nil},                 // 099: Kingler (キングラー)
	&_pokemon{Voltorb, -1, -1, typ.Electric, typ.None, 10.4, 0.5, Voltorb, 1, -1, nil},            // 100: Voltorb (ビリリダマ)
	&_pokemon{Electrode, -1, -1, typ.Electric, typ.None, 66.6, 1.2, Voltorb, 1, 0, nil},           // 101: Electrode (マルマイン)
	&_pokemon{Exeggcute, -1, -1, typ.Grass, typ.Psychic, 2.5, 0.4, Exeggcute, 1, -1, nil},         // 102: Exeggcute (タマタマ)
	&_pokemon{Exeggutor, -1, -1, typ.Grass, typ.Psychic, 120.0, 2.0, Exeggcute, 1, 0, nil},        // 103: Exeggutor (ナッシー)
	&_pokemon{Cubone, -1, -1, typ.Ground, typ.None, 6.5, 0.4, Cubone, 1, -1, nil},                 // 104: Cubone (カラカラ)
	&_pokemon{Marowak, -1, -1, typ.Ground, typ.None, 45.0, 1.0, Cubone, 1, 0, nil},                // 105: Marowak (ガラガラ)
	&_pokemon{Hitmonlee, -1, -1, typ.Fighting, typ.None, 49.8, 1.5, Hitmonlee, 1, 0, nil},         // 106: Hitmonlee (サワムラー)
	&_pokemon{Hitmonchan, -1, -1, typ.Fighting, typ.None, 50.2, 1.4, Hitmonchan, 1, 0, nil},       // 107: Hitmonchan (エビワラー)
	&_pokemon{Lickitung, -1, -1, typ.Normal, typ.None, 65.5, 1.2, Lickitung, 1, 0, nil},           // 108: Lickitung (ベロリンガ)
	&_pokemon{Koffing, -1, -1, typ.Poison, typ.None, 1.0, 0.6, Koffing, 1, 50, nil},               // 109: Koffing (ドガース), SD: 600, CP: [,,,143,,,], HP: [,,,23,,,]
	&_pokemon{Weezing, -1, -1, typ.Poison, typ.None, 9.5, 1.2, Koffing, 1, 0, nil},                // 110: Weezing (マタドガス)
	&_pokemon{Rhyhorn, -1, -1, typ.Ground, typ.Rock, 115.0, 1.0, Rhyhorn, 1, -1, nil},             // 111: Rhyhorn (サイホーン)
	&_pokemon{Rhydon, -1, -1, typ.Ground, typ.Rock, 120.0, 1.9, Rhyhorn, 1, 0, nil},               // 112: Rhydon (サイドン)
	&_pokemon{Chansey, -1, -1, typ.Normal, typ.None, 34.6, 1.1, Chansey, 1, 0, nil},               // 113: Chansey (ラッキー)
	&_pokemon{Tangela, -1, -1, typ.Grass, typ.None, 35.0, 1.0, Tangela, 1, 0, nil},                // 114: Tangela (モンジャラ)
	&_pokemon{Kangaskhan, -1, -1, typ.Normal, typ.None, 80.0, 2.2, Kangaskhan, 1, 0, nil},         // 115: Kangaskhan (ガルーラ)
	&_pokemon{Horsea, -1, -1, typ.Water, typ.None, 8.0, 0.4, Horsea, 1, 50, nil},                  // 116: Horsea (タッツー), SD: 600, CP: [,,,93,,,], HP: [,,,20,,,]
	&_pokemon{Seadra, -1, -1, typ.Water, typ.None, 25.0, 1.2, Horsea, 1, 0, nil},                  // 117: Seadra (シードラ)
	&_pokemon{Goldeen, -1, -1, typ.Water, typ.None, 15.0, 0.6, Goldeen, 1, -1, nil},               // 118: Goldeen (トサキント)
	&_pokemon{Seaking, -1, -1, typ.Water, typ.None, 39.0, 1.3, Goldeen, 1, 0, nil},                // 119: Seaking (アズマオウ)
	&_pokemon{Staryu, 11, 10, typ.Water, typ.None, 34.5, 0.8, Staryu, 1, 50, nil},                 // 120: Staryu (ヒトデマン), SD: 11:200 / 60:400, CP: [11,,,60,,,], HP: [10,,,15,,,]
	&_pokemon{Starmie, -1, -1, typ.Water, typ.Psychic, 80.0, 1.1, Staryu, 1, 0, nil},              // 121: Starmie (スターミー)
	&_pokemon{Mr_Mime, -1, -1, typ.Psychic, typ.Fairy, 54.5, 1.3, Mr_Mime, 1, 0, nil},             // 122: Mr. Mime (バリヤード)
	&_pokemon{Scyther, -1, -1, typ.Bug, typ.Flying, 56.0, 1.5, Scyther, 1, 0, nil},                // 123: Scyther (ストライク)
	&_pokemon{Jynx, -1, -1, typ.Ice, typ.Psychic, 40.6, 1.4, Jynx, 1, 0, nil},                     // 124: Jynx (ルージュラ)
	&_pokemon{Electabuzz, -1, -1, typ.Electric, typ.None, 30.0, 1.1, Electabuzz, 1, 0, nil},       // 125: Electabuzz (エレブー)
	&_pokemon{Magmar, -1, -1, typ.Fire, typ.None, 44.5, 1.3, Magmar, 1, 0, nil},                   // 126: Magmar (ブーバー)
	&_pokemon{Pinsir, -1, -1, typ.Bug, typ.None, 55.0, 1.5, Pinsir, 1, 0, nil},                    // 127: Pinsir (カイロス)
	&_pokemon{Tauros, -1, -1, typ.Normal, typ.None, 88.4, 1.4, Tauros, 1, 0, nil},                 // 128: Tauros (ケンタロス)
	&_pokemon{Magikarp, -1, -1, typ.Water, typ.None, 10.0, 0.9, Magikarp, 1, -1, nil},             // 129: Magikarp (コイキング)
	&_pokemon{Gyarados, -1, -1, typ.Water, typ.Flying, 235.0, 6.5, Magikarp, 1, 0, nil},           // 130: Gyarados (ギャラドス)
	&_pokemon{Lapras, -1, -1, typ.Water, typ.Ice, 220.0, 2.5, Lapras, 1, 0, nil},                  // 131: Lapras (ラプラス)
	&_pokemon{Ditto, -1, -1, typ.Normal, typ.None, 4.0, 0.3, Ditto, 1, 0, nil},                    // 132: Ditto (メタモン)
	&_pokemon{Eevee, -1, -1, typ.Normal, typ.None, 6.5, 0.3, Eevee, 1, -1, nil},                   // 133: Eevee (イーブイ)
	&_pokemon{Vaporeon, -1, -1, typ.Water, typ.None, 29.0, 1.0, Eevee, 1, 0, nil},                 // 134: Vaporeon (シャワーズ)
	&_pokemon{Jolteon, -1, -1, typ.Electric, typ.None, 24.5, 0.8, Eevee, 1, 0, nil},               // 135: Jolteon (サンダース)
	&_pokemon{Flareon, -1, -1, typ.Fire, typ.None, 25.0, 0.9, Eevee, 1, 0, nil},                   // 136: Flareon (ブースター)
	&_pokemon{Porygon, -1, -1, typ.Normal, typ.None, 36.5, 0.8, Porygon, 1, 0, nil},               // 137: Porygon (ポリゴン)
	&_pokemon{Omanyte, -1, -1, typ.Rock, typ.Water, 7.5, 0.4, Omanyte, 1, -1, nil},                // 138: Omanyte (オムナイト)
	&_pokemon{Omastar, -1, -1, typ.Rock, typ.Water, 35.0, 1.0, Omanyte, 1, 0, nil},                // 139: Omastar (オムスター)
	&_pokemon{Kabuto, -1, -1, typ.Rock, typ.Water, 11.5, 0.5, Kabuto, 1, -1, nil},                 // 140: Kabuto (カブト)
	&_pokemon{Kabutops, -1, -1, typ.Rock, typ.Water, 40.5, 1.3, Kabuto, 1, 0, nil},                // 141: Kabutops (カブトプス)
	&_pokemon{Aerodactyl, -1, -1, typ.Rock, typ.Flying, 59.0, 1.8, Aerodactyl, 1, 0, nil},         // 142: Aerodactyl (プテラ)
	&_pokemon{Snorlax, -1, -1, typ.Normal, typ.None, 460.0, 2.1, Snorlax, 1, 0, nil},              // 143: Snorlax (カビゴン)
	&_pokemon{Articuno, -1, -1, typ.Ice, typ.Flying, 55.4, 1.7, Articuno, 1, 0, nil},              // 144: Articuno (フリーザー)
	&_pokemon{Zapdos, -1, -1, typ.Electric, typ.Flying, 52.6, 1.6, Zapdos, 1, 0, nil},             // 145: Zapdos (サンダー)
	&_pokemon{Moltres, -1, -1, typ.Fire, typ.Flying, 60.0, 2.0, Moltres, 1, 0, nil},               // 146: Moltres (ファイヤー)
	&_pokemon{Dratini, -1, -1, typ.Dragon, typ.None, 3.3, 1.8, Dratini, 1, 25, nil},               // 147: Dratini (ミニリュウ) SD: 64:400 / 172:800, CP: [,,,64,,,172,,,], HP: [,,,20,,,33,,,]
	&_pokemon{Dragonair, -1, -1, typ.Dragon, typ.None, 16.5, 4.0, Dratini, 1, -1, nil},            // 148: Dragonair (ハクリュー)
	&_pokemon{Dragonite, -1, -1, typ.Dragon, typ.Flying, 210.0, 2.2, Dratini, 1, 0, nil},          // 149: Dragonite (カイリュー)
	&_pokemon{Mewtwo, -1, -1, typ.Psychic, typ.None, 122.0, 2.0, Mewtwo, 1, 0, nil},               // 150: Mewtwo (ミュウツー)
	&_pokemon{Mew, -1, -1, typ.Psychic, typ.None, 4.0, 0.4, Mew, 1, 0, nil},                       // 151: Mew (ミュウ)
}
