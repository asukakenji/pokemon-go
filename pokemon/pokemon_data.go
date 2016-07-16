package pokemon

import (
	t "github.com/asukakenji/pokemon-go/type"
)

type _pokemon struct {
	id              Pokemon
	evolveFrom      Pokemon
	baseCombatPower int16
	baseHitPoints   int16
	type1           t.Type
	type2           t.Type
	weight          float64 // in kg
	height          float64 // in m
	candyType       Pokemon
	candyToPowerUp  int16
	candyToEvolve   int16
}

const (
	p_ Pokemon = None
	t_ t.Type  = t.None
)

var pokemons = [...]*_pokemon{
	nil,
	&_pokemon{Bulbasaur, p_, -1, -1, t.Grass, t.Poison, 6.9, 0.7, Bulbasaur, 1, 25},              // 001: Bulbasaur (フシギダネ), SD:200
	&_pokemon{Ivysaur, Bulbasaur, -1, -1, t.Grass, t.Poison, 13.0, 1.0, Bulbasaur, 1, 100},       // 002: Ivysaur (フシギソウ)
	&_pokemon{Venusaur, Ivysaur, -1, -1, t.Grass, t.Poison, 100.0, 2.0, Bulbasaur, 1, 0},         // 003: Venusaur (フシギバナ)
	&_pokemon{Charmander, p_, 12, 10, t.Fire, t_, 8.5, 0.6, Charmander, 1, 25},                   // 004: Charmander (ヒトカゲ), SD: 200, CP: [12,,,35,,,], HP: [10,,,13,,,]
	&_pokemon{Charmeleon, Charmander, -1, -1, t.Fire, t_, 19.0, 1.1, Charmander, 1, 100},         // 005: Charmeleon (リザード)
	&_pokemon{Charizard, Charmeleon, -1, -1, t.Fire, t.Flying, 90.5, 1.7, Charmander, 1, 0},      // 006: Charizard (リザードン)
	&_pokemon{Squirtle, p_, -1, -1, t.Water, t_, 9.0, 0.5, Squirtle, 1, 25},                      // 007: Squirtle (ゼニガメ), SD: 800, CP: [,,,161,,,], HP: [,,,34,,,]
	&_pokemon{Wartortle, Squirtle, -1, -1, t.Water, t_, 22.5, 1.0, Squirtle, 1, 100},             // 008: Wartortle (カメール)
	&_pokemon{Blastoise, Wartortle, -1, -1, t.Water, t_, 85.5, 1.6, Squirtle, 1, 0},              // 009: Blastoise (カメックス)
	&_pokemon{Caterpie, p_, 10, 10, t.Bug, t_, 2.9, 0.3, Caterpie, 1, 12},                        // 010: Caterpie (キャタピー), SD: 200, CP: [10,,,], HP: [10,,,]
	&_pokemon{Metapod, Caterpie, -1, -1, t.Bug, t_, 9.9, 0.7, Caterpie, 1, 50},                   // 011: Metapod (トランセル)
	&_pokemon{Butterfree, Metapod, -1, -1, t.Bug, t.Flying, 32.0, 1.1, Caterpie, 1, 0},           // 012: Butterfree (バタフリー)
	&_pokemon{Weedle, p_, 10, 10, t.Bug, t.Poison, 3.2, 0.3, Weedle, 1, 12},                      // 013: Weedle (ビードル), SD: 200, CP: [,,,14,15,,,], HP: [,,,13,14,,,]
	&_pokemon{Kakuna, Weedle, -1, -1, t.Bug, t.Poison, 10.0, 0.6, Weedle, 1, 50},                 // 014: Kakuna (コクーン), SD: 600, CP: [,,,64,69,,,], HP: [,,,34,36,,,]
	&_pokemon{Beedrill, Kakuna, -1, -1, t.Bug, t.Poison, 29.5, 1.0, Weedle, 1, 0},                // 015: Beedrill (スピアー), SD: 800, CP: [,,,240,258,,,], HP: [,,,50,52,,,]
	&_pokemon{Pidgey, p_, 10, 10, t.Normal, t.Flying, 1.8, 0.3, Pidgey, 1, 12},                   // 016: Pidgey (ポッポ), SD: 200, CP: [10,,,23,,,], HP: [10,,,14,,,]
	&_pokemon{Pidgeotto, Pidgey, 15, 12, t.Normal, t.Flying, 30.0, 1.1, Pidgey, 1, 50},           // 017: Pidgeotto (ピジョン), SD: 15:200 / 117:400 CP: [15,,,110,117(125),,,], HP: [12,,,32,34,,,]
	&_pokemon{Pidgeot, Pidgeotto, -1, -1, t.Normal, t.Flying, 39.5, 1.5, Pidgey, 1, 0},           // 018: Pidgeot (ピジョット), SD: 234:400 / 263:600, CP:[,,,204,234,263,,,352,,,], HP: [,,,44,47,50,,,58,,,]
	&_pokemon{Rattata, p_, -1, -1, t.Normal, t_, 3.5, 0.3, Rattata, 1, 25},                       // 019: Rattata (コラッタ), SD: 200, CP: [,,,21,,,], HP: [,,,11,,,]
	&_pokemon{Raticate, Rattata, -1, -1, t.Normal, t_, 18.5, 0.7, Rattata, 1, 0},                 // 020: Raticate (ラッタ), SD: 800, CP: [,,,296,,,], HP: [,,,43,,,]
	&_pokemon{Spearow, p_, 10, 10, t.Normal, t.Flying, 2.0, 0.3, Spearow, 1, 50},                 // 021: Spearow (オニスズメ), SD: 10:200 / 400, CP: [10,,,63,,,], HP: [10,,,24,,,]
	&_pokemon{Fearow, Spearow, -1, -1, t.Normal, t.Flying, 38.0, 1.2, Spearow, 1, 0},             // 022: Fearow (オニドリル)
	&_pokemon{Ekans, p_, -1, -1, t.Poison, t_, 6.9, 2.0, Ekans, 1, 50},                           // 023: Ekans (アーボ)
	&_pokemon{Arbok, Ekans, -1, -1, t.Poison, t_, 65.0, 3.5, Ekans, 1, 0},                        // 024: Arbok (アーボック)
	&_pokemon{Pikachu, p_, 11, 10, t.Electric, t_, 6.0, 0.4, Pikachu, 1, 50},                     // 025: Pikachu (ピカチュウ), 172 -> 025, SD: 200, CP: [11,,,], HP: [10,,,]
	&_pokemon{Raichu, Pikachu, -1, -1, t.Electric, t_, 30.0, 0.8, Pikachu, 1, 0},                 // 026: Raichu (ライチュウ)
	&_pokemon{Sandshrew, p_, -1, -1, t.Ground, t_, 12.0, 0.6, Sandshrew, 1, 50},                  // 027: Sandshrew (サンド)
	&_pokemon{Sandslash, Sandshrew, -1, -1, t.Ground, t_, 29.5, 1.0, Sandshrew, 1, 0},            // 028: Sandslash (サンドパン)
	&_pokemon{Nidoran_female, p_, -1, -1, t.Poison, t_, 7.0, 0.4, Nidoran_female, 1, 25},         // 029: Nidoran♀ (ニドラン♀)
	&_pokemon{Nidorina, Nidoran_female, -1, -1, t.Poison, t_, 20.0, 0.8, Nidoran_female, 1, 100}, // 030: Nidorina (ニドリーナ)
	&_pokemon{Nidoqueen, Nidorina, -1, -1, t.Poison, t.Ground, 60.0, 1.3, Nidoran_female, 1, 0},  // 031: Nidoqueen (ニドクイン)
	&_pokemon{Nidoran_male, p_, -1, -1, t.Poison, t_, 9.0, 0.5, Nidoran_male, 1, 25},             // 032: Nidoran♂ (ニドラン♂)
	&_pokemon{Nidorino, Nidoran_male, -1, -1, t.Poison, t_, 19.5, 0.9, Nidoran_male, 1, 100},     // 033: Nidorino (ニドリーノ)
	&_pokemon{Nidoking, Nidorino, -1, -1, t.Poison, t.Ground, 62.0, 1.4, Nidoran_male, 1, 0},     // 034: Nidoking (ニドキング)
	&_pokemon{Clefairy, p_, -1, -1, t.Fairy, t_, 7.5, 0.6, Clefairy, 1, 50},                      // 035: Clefairy (ピッピ), 173 -> 035
	&_pokemon{Clefable, Clefairy, -1, -1, t.Fairy, t_, 40.0, 1.3, Clefairy, 1, 0},                // 036: Clefable (ピクシー)
	&_pokemon{Vulpix, p_, -1, -1, t.Fire, t_, 9.9, 0.6, Vulpix, 1, 50},                           // 037: Vulpix (ロコン)
	&_pokemon{Ninetales, Vulpix, -1, -1, t.Fire, t_, 19.9, 1.1, Vulpix, 1, 0},                    // 038: Ninetales (キュウコン)
	&_pokemon{Jigglypuff, p_, -1, -1, t.Normal, t.Fairy, 5.5, 0.5, Jigglypuff, 1, 50},            // 039: Jigglypuff (プリン), 174 -> 039, SD: 600, CP: [,,,110,,,], HP: [,,,69,,,]
	&_pokemon{Wigglytuff, Jigglypuff, -1, -1, t.Normal, t.Fairy, 12.0, 1.0, Jigglypuff, 1, 0},    // 040: Wigglytuff (プクリン)
	&_pokemon{Zubat, p_, -1, -1, t.Poison, t.Flying, 7.5, 0.8, Zubat, 1, 50},                     // 041: Zubat (ズバット)
	&_pokemon{Golbat, Zubat, -1, -1, t.Poison, t.Flying, 55.0, 1.6, Zubat, 1, 0},                 // 042: Golbat (ゴルバット)
	&_pokemon{Oddish, p_, 13, 10, t.Grass, t.Poison, 5.4, 0.5, Oddish, 1, 25},                    // 043: Oddish (ナゾノクサ), SD: 200, CP: [13,,,], HP: [10,,,]
	&_pokemon{Gloom, Oddish, -1, -1, t.Grass, t.Poison, 8.6, 0.8, Oddish, 1, 100},                // 044: Gloom (クサイハナ)
	&_pokemon{Vileplume, Gloom, -1, -1, t.Grass, t.Poison, 18.6, 1.2, Oddish, 1, 0},              // 045: Vileplume (ラフレシア)
	&_pokemon{Paras, p_, -1, -1, t.Bug, t.Grass, 5.4, 0.3, Paras, 1, 50},                         // 046: Paras (パラス)
	&_pokemon{Parasect, Paras, -1, -1, t.Bug, t.Grass, 29.5, 1.0, Paras, 1, 0},                   // 047: Parasect (パラセクト)
	&_pokemon{Venonat, p_, -1, -1, t.Bug, t.Poison, 30.0, 1.0, Venonat, 1, 50},                   // 048: Venonat (コンパン)
	&_pokemon{Venomoth, Venonat, -1, -1, t.Bug, t.Poison, 12.5, 1.5, Venonat, 1, 0},              // 049: Venomoth (モルフォン)
	&_pokemon{Diglett, p_, -1, -1, t.Ground, t_, 0.8, 0.2, Diglett, 1, 50},                       // 050: Diglett (ディグダ)
	&_pokemon{Dugtrio, Diglett, -1, -1, t.Ground, t_, 33.3, 0.7, Diglett, 1, 0},                  // 051: Dugtrio (ダグトリオ)
	&_pokemon{Meowth, p_, -1, -1, t.Normal, t_, 4.2, 0.4, Meowth, 1, 50},                         // 052: Meowth (ニャース)
	&_pokemon{Persian, Meowth, -1, -1, t.Normal, t_, 32.0, 1.0, Meowth, 1, 0},                    // 053: Persian (ペルシアン)
	&_pokemon{Psyduck, p_, -1, -1, t.Water, t_, 19.6, 0.8, Psyduck, 1, 50},                       // 054: Psyduck (コダック)
	&_pokemon{Golduck, Psyduck, -1, -1, t.Water, t_, 76.6, 1.7, Psyduck, 1, 0},                   // 055: Golduck (ゴルダック)
	&_pokemon{Mankey, p_, -1, -1, t.Fighting, t_, 28.0, 0.5, Mankey, 1, 50},                      // 056: Mankey (マンキー)
	&_pokemon{Primeape, Mankey, -1, -1, t.Fighting, t_, 32.0, 1.0, Mankey, 1, 0},                 // 057: Primeape (オコリザル)
	&_pokemon{Growlithe, p_, -1, -1, t.Fire, t_, 19.0, 0.7, Growlithe, 1, 50},                    // 058: Growlithe (ガーディ)
	&_pokemon{Arcanine, Growlithe, -1, -1, t.Fire, t_, 155.0, 1.9, Growlithe, 1, 0},              // 059: Arcanine (ウインディ)
	&_pokemon{Poliwag, p_, -1, -1, t.Water, t_, 12.4, 0.6, Poliwag, 1, 25},                       // 060: Poliwag (ニョロモ)
	&_pokemon{Poliwhirl, Poliwag, -1, -1, t.Water, t_, 20.0, 1.0, Poliwag, 1, 100},               // 061: Poliwhirl (ニョロゾ)
	&_pokemon{Poliwrath, Poliwhirl, -1, -1, t.Water, t.Fighting, 54.0, 1.3, Poliwag, 1, 0},       // 062: Poliwrath (ニョロボン)
	&_pokemon{Abra, p_, -1, -1, t.Psychic, t_, 19.5, 0.9, Abra, 1, 25},                           // 063: Abra (ケーシィ), SD: 400, CP: [,,,54,,,], HP: [,,,13,,,]
	&_pokemon{Kadabra, Abra, -1, -1, t.Psychic, t_, 56.5, 1.3, Abra, 1, 100},                     // 064: Kadabra (ユンゲラー)
	&_pokemon{Alakazam, Kadabra, -1, -1, t.Psychic, t_, 48.0, 1.5, Abra, 1, 0},                   // 065: Alakazam (フーディン)
	&_pokemon{Machop, p_, -1, -1, t.Fighting, t_, 19.5, 0.8, Machop, 1, 25},                      // 066: Machop (ワンリキー)
	&_pokemon{Machoke, Machop, -1, -1, t.Fighting, t_, 70.5, 1.5, Machop, 1, 100},                // 067: Machoke (ゴーリキー)
	&_pokemon{Machamp, Machoke, -1, -1, t.Fighting, t_, 130.0, 1.6, Machop, 1, 0},                // 068: Machamp (カイリキー)
	&_pokemon{Bellsprout, p_, -1, -1, t.Grass, t.Poison, 4.0, 0.7, Bellsprout, 1, 25},            // 069: Bellsprout (マダツボミ), SD: 400, CP: [,,,104,,,], HP: [,,,28,,,]
	&_pokemon{Weepinbell, Bellsprout, -1, -1, t.Grass, t.Poison, 6.4, 1.0, Bellsprout, 1, 100},   // 070: Weepinbell (ウツドン)
	&_pokemon{Victreebel, Weepinbell, -1, -1, t.Grass, t.Poison, 15.5, 1.7, Bellsprout, 1, 0},    // 071: Victreebel (ウツボット)
	&_pokemon{Tentacool, p_, -1, -1, t.Water, t.Poison, 45.5, 0.9, Tentacool, 1, 50},             // 072: Tentacool (メノクラゲ)
	&_pokemon{Tentacruel, Tentacool, -1, -1, t.Water, t.Poison, 55.0, 1.6, Tentacool, 1, 0},      // 073: Tentacruel (ドククラゲ)
	&_pokemon{Geodude, p_, -1, -1, t.Rock, t.Ground, 20.0, 0.4, Geodude, 1, 25},                  // 074: Geodude (イシツブテ)
	&_pokemon{Graveler, Geodude, -1, -1, t.Rock, t.Ground, 105.0, 1.0, Geodude, 1, 100},          // 075: Graveler (ゴローン)
	&_pokemon{Golem, Graveler, -1, -1, t.Rock, t.Ground, 300.0, 1.4, Geodude, 1, 0},              // 076: Golem (ゴローニャ)
	&_pokemon{Ponyta, p_, -1, -1, t.Fire, t_, 30.0, 1.0, Ponyta, 1, 50},                          // 077: Ponyta (ポニータ)
	&_pokemon{Rapidash, Ponyta, -1, -1, t.Fire, t_, 95.0, 1.7, Ponyta, 1, 0},                     // 078: Rapidash (ギャロップ)
	&_pokemon{Slowpoke, p_, -1, -1, t.Water, t.Psychic, 36.0, 1.2, Slowpoke, 1, 50},              // 079: Slowpoke (ヤドン)
	&_pokemon{Slowbro, Slowpoke, -1, -1, t.Water, t.Psychic, 78.5, 1.6, Slowpoke, 1, 0},          // 080: Slowbro (ヤドラン)
	&_pokemon{Magnemite, p_, -1, -1, t.Electric, t.Steel, 6.0, 0.3, Magnemite, 1, 50},            // 081: Magnemite (コイル), SD: 200, CP: [,,,38,,,], HP: [,,,10,,,]
	&_pokemon{Magneton, Magnemite, -1, -1, t.Electric, t.Steel, 60.0, 1.0, Magnemite, 1, 0},      // 082: Magneton (レアコイル)
	&_pokemon{Farfetch_d, p_, -1, -1, t.Normal, t.Flying, 15.0, 0.8, Farfetch_d, 1, 0},           // 083: Farfetch'd (カモネギ)
	&_pokemon{Doduo, p_, -1, -1, t.Normal, t.Flying, 39.2, 1.4, Doduo, 1, 50},                    // 084: Doduo (ドードー)
	&_pokemon{Dodrio, Doduo, -1, -1, t.Normal, t.Flying, 85.2, 1.8, Doduo, 1, 0},                 // 085: Dodrio (ドードリオ)
	&_pokemon{Seel, p_, -1, -1, t.Water, t_, 90.0, 1.1, Seel, 1, 50},                             // 086: Seel (パウワウ)
	&_pokemon{Dewgong, Seel, -1, -1, t.Water, t.Ice, 120.0, 1.7, Seel, 1, 0},                     // 087: Dewgong (ジュゴン)
	&_pokemon{Grimer, p_, -1, -1, t.Poison, t_, 30.0, 0.9, Grimer, 1, 50},                        // 088: Grimer (ベトベター)
	&_pokemon{Muk, Grimer, -1, -1, t.Poison, t_, 30.0, 1.2, Grimer, 1, 0},                        // 089: Muk (ベトベトン)
	&_pokemon{Shellder, p_, -1, -1, t.Water, t_, 4.0, 0.3, Shellder, 1, 50},                      // 090: Shellder (シェルダー)
	&_pokemon{Cloyster, Shellder, -1, -1, t.Water, t.Ice, 132.5, 1.5, Shellder, 1, 0},            // 091: Cloyster (パルシェン)
	&_pokemon{Gastly, p_, -1, -1, t.Ghost, t.Poison, 0.1, 1.3, Gastly, 1, 25},                    // 092: Gastly (ゴース)
	&_pokemon{Haunter, Gastly, -1, -1, t.Ghost, t.Poison, 0.1, 1.6, Gastly, 1, 100},              // 093: Haunter (ゴースト)
	&_pokemon{Gengar, Haunter, -1, -1, t.Ghost, t.Poison, 40.5, 1.5, Gastly, 1, 0},               // 094: Gengar (ゲンガー)
	&_pokemon{Onix, p_, -1, -1, t.Rock, t.Ground, 210.0, 8.8, Onix, 1, 0},                        // 095: Onix (イワーク)
	&_pokemon{Drowzee, p_, -1, -1, t.Psychic, t_, 32.4, 1.0, Drowzee, 1, 50},                     // 096: Drowzee (スリープ)
	&_pokemon{Hypno, Drowzee, -1, -1, t.Psychic, t_, 75.6, 1.6, Drowzee, 1, 0},                   // 097: Hypno (スリーパー)
	&_pokemon{Krabby, p_, -1, -1, t.Water, t_, 6.5, 0.4, Krabby, 1, 50},                          // 098: Krabby (クラブ)
	&_pokemon{Kingler, Krabby, -1, -1, t.Water, t_, 60.0, 1.3, Krabby, 1, 0},                     // 099: Kingler (キングラー)
	&_pokemon{Voltorb, p_, -1, -1, t.Electric, t_, 10.4, 0.5, Voltorb, 1, 50},                    // 100: Voltorb (ビリリダマ)
	&_pokemon{Electrode, Voltorb, -1, -1, t.Electric, t_, 66.6, 1.2, Voltorb, 1, 0},              // 101: Electrode (マルマイン)
	&_pokemon{Exeggcute, p_, -1, -1, t.Grass, t.Psychic, 2.5, 0.4, Exeggcute, 1, 50},             // 102: Exeggcute (タマタマ)
	&_pokemon{Exeggutor, Exeggcute, -1, -1, t.Grass, t.Psychic, 120.0, 2.0, Exeggcute, 1, 0},     // 103: Exeggutor (ナッシー)
	&_pokemon{Cubone, p_, -1, -1, t.Ground, t_, 6.5, 0.4, Cubone, 1, 50},                         // 104: Cubone (カラカラ)
	&_pokemon{Marowak, Cubone, -1, -1, t.Ground, t_, 45.0, 1.0, Cubone, 1, 0},                    // 105: Marowak (ガラガラ)
	&_pokemon{Hitmonlee, p_, -1, -1, t.Fighting, t_, 49.8, 1.5, Hitmonlee, 1, 0},                 // 106: Hitmonlee (サワムラー)
	&_pokemon{Hitmonchan, p_, -1, -1, t.Fighting, t_, 50.2, 1.4, Hitmonchan, 1, 0},               // 107: Hitmonchan (エビワラー)
	&_pokemon{Lickitung, p_, -1, -1, t.Normal, t_, 65.5, 1.2, Lickitung, 1, 0},                   // 108: Lickitung (ベロリンガ)
	&_pokemon{Koffing, p_, -1, -1, t.Poison, t_, 1.0, 0.6, Koffing, 1, 50},                       // 109: Koffing (ドガース), SD: 600, CP: [,,,143,,,], HP: [,,,23,,,]
	&_pokemon{Weezing, Koffing, -1, -1, t.Poison, t_, 9.5, 1.2, Koffing, 1, 0},                   // 110: Weezing (マタドガス)
	&_pokemon{Rhyhorn, p_, -1, -1, t.Ground, t.Rock, 115.0, 1.0, Rhyhorn, 1, 50},                 // 111: Rhyhorn (サイホーン)
	&_pokemon{Rhydon, Rhyhorn, -1, -1, t.Ground, t.Rock, 120.0, 1.9, Rhyhorn, 1, 0},              // 112: Rhydon (サイドン)
	&_pokemon{Chansey, p_, -1, -1, t.Normal, t_, 34.6, 1.1, Chansey, 1, 0},                       // 113: Chansey (ラッキー)
	&_pokemon{Tangela, p_, -1, -1, t.Grass, t_, 35.0, 1.0, Tangela, 1, 0},                        // 114: Tangela (モンジャラ)
	&_pokemon{Kangaskhan, p_, -1, -1, t.Normal, t_, 80.0, 2.2, Kangaskhan, 1, 0},                 // 115: Kangaskhan (ガルーラ)
	&_pokemon{Horsea, p_, -1, -1, t.Water, t_, 8.0, 0.4, Horsea, 1, 50},                          // 116: Horsea (タッツー), SD: 600, CP: [,,,93,,,], HP: [,,,20,,,]
	&_pokemon{Seadra, Horsea, -1, -1, t.Water, t_, 25.0, 1.2, Horsea, 1, 0},                      // 117: Seadra (シードラ)
	&_pokemon{Goldeen, p_, -1, -1, t.Water, t_, 15.0, 0.6, Goldeen, 1, 50},                       // 118: Goldeen (トサキント)
	&_pokemon{Seaking, Goldeen, -1, -1, t.Water, t_, 39.0, 1.3, Goldeen, 1, 0},                   // 119: Seaking (アズマオウ)
	&_pokemon{Staryu, p_, 11, 10, t.Water, t_, 34.5, 0.8, Staryu, 1, 50},                         // 120: Staryu (ヒトデマン), SD: 11:200 / 60:400, CP: [11,,,60,,,], HP: [10,,,15,,,]
	&_pokemon{Starmie, Staryu, -1, -1, t.Water, t.Psychic, 80.0, 1.1, Staryu, 1, 0},              // 121: Starmie (スターミー)
	&_pokemon{Mr_Mime, p_, -1, -1, t.Psychic, t.Fairy, 54.5, 1.3, Mr_Mime, 1, 0},                 // 122: Mr. Mime (バリヤード)
	&_pokemon{Scyther, p_, -1, -1, t.Bug, t.Flying, 56.0, 1.5, Scyther, 1, 0},                    // 123: Scyther (ストライク)
	&_pokemon{Jynx, p_, -1, -1, t.Ice, t.Psychic, 40.6, 1.4, Jynx, 1, 0},                         // 124: Jynx (ルージュラ)
	&_pokemon{Electabuzz, p_, -1, -1, t.Electric, t_, 30.0, 1.1, Electabuzz, 1, 0},               // 125: Electabuzz (エレブー)
	&_pokemon{Magmar, p_, -1, -1, t.Fire, t_, 44.5, 1.3, Magmar, 1, 0},                           // 126: Magmar (ブーバー)
	&_pokemon{Pinsir, p_, -1, -1, t.Bug, t_, 55.0, 1.5, Pinsir, 1, 0},                            // 127: Pinsir (カイロス)
	&_pokemon{Tauros, p_, -1, -1, t.Normal, t_, 88.4, 1.4, Tauros, 1, 0},                         // 128: Tauros (ケンタロス)
	&_pokemon{Magikarp, p_, -1, -1, t.Water, t_, 10.0, 0.9, Magikarp, 1, 400},                    // 129: Magikarp (コイキング)
	&_pokemon{Gyarados, Magikarp, -1, -1, t.Water, t.Flying, 235.0, 6.5, Magikarp, 1, 0},         // 130: Gyarados (ギャラドス)
	&_pokemon{Lapras, p_, -1, -1, t.Water, t.Ice, 220.0, 2.5, Lapras, 1, 0},                      // 131: Lapras (ラプラス)
	&_pokemon{Ditto, p_, -1, -1, t.Normal, t_, 4.0, 0.3, Ditto, 1, 0},                            // 132: Ditto (メタモン)
	&_pokemon{Eevee, p_, -1, -1, t.Normal, t_, 6.5, 0.3, Eevee, 1, 25},                           // 133: Eevee (イーブイ)
	&_pokemon{Vaporeon, Eevee, -1, -1, t.Water, t_, 29.0, 1.0, Eevee, 1, 0},                      // 134: Vaporeon (シャワーズ)
	&_pokemon{Jolteon, Eevee, -1, -1, t.Electric, t_, 24.5, 0.8, Eevee, 1, 0},                    // 135: Jolteon (サンダース)
	&_pokemon{Flareon, Eevee, -1, -1, t.Fire, t_, 25.0, 0.9, Eevee, 1, 0},                        // 136: Flareon (ブースター)
	&_pokemon{Porygon, p_, -1, -1, t.Normal, t_, 36.5, 0.8, Porygon, 1, 0},                       // 137: Porygon (ポリゴン)
	&_pokemon{Omanyte, p_, -1, -1, t.Rock, t.Water, 7.5, 0.4, Omanyte, 1, 50},                    // 138: Omanyte (オムナイト)
	&_pokemon{Omastar, Omanyte, -1, -1, t.Rock, t.Water, 35.0, 1.0, Omanyte, 1, 0},               // 139: Omastar (オムスター)
	&_pokemon{Kabuto, p_, -1, -1, t.Rock, t.Water, 11.5, 0.5, Kabuto, 1, 50},                     // 140: Kabuto (カブト)
	&_pokemon{Kabutops, Kabuto, -1, -1, t.Rock, t.Water, 40.5, 1.3, Kabuto, 1, 0},                // 141: Kabutops (カブトプス)
	&_pokemon{Aerodactyl, p_, -1, -1, t.Rock, t.Flying, 59.0, 1.8, Aerodactyl, 1, 0},             // 142: Aerodactyl (プテラ)
	&_pokemon{Snorlax, p_, -1, -1, t.Normal, t_, 460.0, 2.1, Snorlax, 1, 0},                      // 143: Snorlax (カビゴン)
	&_pokemon{Articuno, p_, -1, -1, t.Ice, t.Flying, 55.4, 1.7, Articuno, 1, 0},                  // 144: Articuno (フリーザー)
	&_pokemon{Zapdos, p_, -1, -1, t.Electric, t.Flying, 52.6, 1.6, Zapdos, 1, 0},                 // 145: Zapdos (サンダー)
	&_pokemon{Moltres, p_, -1, -1, t.Fire, t.Flying, 60.0, 2.0, Moltres, 1, 0},                   // 146: Moltres (ファイヤー)
	&_pokemon{Dratini, p_, -1, -1, t.Dragon, t_, 3.3, 1.8, Dratini, 1, 25},                       // 147: Dratini (ミニリュウ) SD: 64:400 / 172:800, CP: [,,,64,,,172,,,], HP: [,,,20,,,33,,,]
	&_pokemon{Dragonair, Dratini, -1, -1, t.Dragon, t_, 16.5, 4.0, Dratini, 1, 100},              // 148: Dragonair (ハクリュー)
	&_pokemon{Dragonite, Dragonair, -1, -1, t.Dragon, t.Flying, 210.0, 2.2, Dratini, 1, 0},       // 149: Dragonite (カイリュー)
	&_pokemon{Mewtwo, p_, -1, -1, t.Psychic, t_, 122.0, 2.0, Mewtwo, 1, 0},                       // 150: Mewtwo (ミュウツー)
	&_pokemon{Mew, p_, -1, -1, t.Psychic, t_, 4.0, 0.4, Mew, 1, 0},                               // 151: Mew (ミュウ)
}
