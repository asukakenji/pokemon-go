package pokemon

import (
	t "github.com/asukakenji/pokemon-go/type"
)

type _pokemon struct {
	id            Pokemon
	evolveFrom    Pokemon
	baseStamina   int16
	baseAttack    int16
	baseDefense   int16
	type1         t.Type
	type2         t.Type
	weight        float64 // in kg
	height        float64 // in m
	candyType     Pokemon
	candyToEvolve int16
}

const (
	p_ Pokemon = None
	t_ t.Type  = t.None
)

var pokemons = [...]*_pokemon{
	nil,
	&_pokemon{Bulbasaur, p_, 90, 126, 126, t.Grass, t.Poison, 6.9, 0.7, Bulbasaur, 25},               // 001: Bulbasaur (フシギダネ)
	&_pokemon{Ivysaur, Bulbasaur, 120, 156, 158, t.Grass, t.Poison, 13.0, 1.0, Bulbasaur, 100},       // 002: Ivysaur (フシギソウ)
	&_pokemon{Venusaur, Ivysaur, 160, 198, 200, t.Grass, t.Poison, 100.0, 2.0, Bulbasaur, 0},         // 003: Venusaur (フシギバナ)
	&_pokemon{Charmander, p_, 78, 128, 108, t.Fire, t_, 8.5, 0.6, Charmander, 25},                    // 004: Charmander (ヒトカゲ)
	&_pokemon{Charmeleon, Charmander, 116, 160, 140, t.Fire, t_, 19.0, 1.1, Charmander, 100},         // 005: Charmeleon (リザード)
	&_pokemon{Charizard, Charmeleon, 156, 212, 182, t.Fire, t.Flying, 90.5, 1.7, Charmander, 0},      // 006: Charizard (リザードン)
	&_pokemon{Squirtle, p_, 88, 112, 142, t.Water, t_, 9.0, 0.5, Squirtle, 25},                       // 007: Squirtle (ゼニガメ)
	&_pokemon{Wartortle, Squirtle, 118, 144, 176, t.Water, t_, 22.5, 1.0, Squirtle, 100},             // 008: Wartortle (カメール)
	&_pokemon{Blastoise, Wartortle, 158, 186, 222, t.Water, t_, 85.5, 1.6, Squirtle, 0},              // 009: Blastoise (カメックス)
	&_pokemon{Caterpie, p_, 90, 62, 66, t.Bug, t_, 2.9, 0.3, Caterpie, 12},                           // 010: Caterpie (キャタピー)
	&_pokemon{Metapod, Caterpie, 100, 56, 86, t.Bug, t_, 9.9, 0.7, Caterpie, 50},                     // 011: Metapod (トランセル)
	&_pokemon{Butterfree, Metapod, 120, 144, 144, t.Bug, t.Flying, 32.0, 1.1, Caterpie, 0},           // 012: Butterfree (バタフリー)
	&_pokemon{Weedle, p_, 80, 68, 64, t.Bug, t.Poison, 3.2, 0.3, Weedle, 12},                         // 013: Weedle (ビードル)
	&_pokemon{Kakuna, Weedle, 90, 62, 82, t.Bug, t.Poison, 10.0, 0.6, Weedle, 50},                    // 014: Kakuna (コクーン)
	&_pokemon{Beedrill, Kakuna, 130, 144, 130, t.Bug, t.Poison, 29.5, 1.0, Weedle, 0},                // 015: Beedrill (スピアー)
	&_pokemon{Pidgey, p_, 80, 94, 90, t.Normal, t.Flying, 1.8, 0.3, Pidgey, 12},                      // 016: Pidgey (ポッポ)
	&_pokemon{Pidgeotto, Pidgey, 126, 126, 122, t.Normal, t.Flying, 30.0, 1.1, Pidgey, 50},           // 017: Pidgeotto (ピジョン)
	&_pokemon{Pidgeot, Pidgeotto, 166, 170, 166, t.Normal, t.Flying, 39.5, 1.5, Pidgey, 0},           // 018: Pidgeot (ピジョット)
	&_pokemon{Rattata, p_, 60, 92, 86, t.Normal, t_, 3.5, 0.3, Rattata, 25},                          // 019: Rattata (コラッタ)
	&_pokemon{Raticate, Rattata, 110, 146, 150, t.Normal, t_, 18.5, 0.7, Rattata, 0},                 // 020: Raticate (ラッタ)
	&_pokemon{Spearow, p_, 80, 102, 78, t.Normal, t.Flying, 2.0, 0.3, Spearow, 50},                   // 021: Spearow (オニスズメ)
	&_pokemon{Fearow, Spearow, 130, 168, 146, t.Normal, t.Flying, 38.0, 1.2, Spearow, 0},             // 022: Fearow (オニドリル)
	&_pokemon{Ekans, p_, 70, 112, 112, t.Poison, t_, 6.9, 2.0, Ekans, 50},                            // 023: Ekans (アーボ)
	&_pokemon{Arbok, Ekans, 120, 166, 166, t.Poison, t_, 65.0, 3.5, Ekans, 0},                        // 024: Arbok (アーボック)
	&_pokemon{Pikachu, p_, 70, 124, 108, t.Electric, t_, 6.0, 0.4, Pikachu, 50},                      // 025: Pikachu (ピカチュウ), 172 -> 025
	&_pokemon{Raichu, Pikachu, 120, 200, 154, t.Electric, t_, 30.0, 0.8, Pikachu, 0},                 // 026: Raichu (ライチュウ)
	&_pokemon{Sandshrew, p_, 100, 90, 114, t.Ground, t_, 12.0, 0.6, Sandshrew, 50},                   // 027: Sandshrew (サンド)
	&_pokemon{Sandslash, Sandshrew, 150, 150, 172, t.Ground, t_, 29.5, 1.0, Sandshrew, 0},            // 028: Sandslash (サンドパン)
	&_pokemon{Nidoran_female, p_, 110, 100, 104, t.Poison, t_, 7.0, 0.4, Nidoran_female, 25},         // 029: Nidoran♀ (ニドラン♀)
	&_pokemon{Nidorina, Nidoran_female, 140, 132, 136, t.Poison, t_, 20.0, 0.8, Nidoran_female, 100}, // 030: Nidorina (ニドリーナ)
	&_pokemon{Nidoqueen, Nidorina, 180, 184, 190, t.Poison, t.Ground, 60.0, 1.3, Nidoran_female, 0},  // 031: Nidoqueen (ニドクイン)
	&_pokemon{Nidoran_male, p_, 92, 110, 94, t.Poison, t_, 9.0, 0.5, Nidoran_male, 25},               // 032: Nidoran♂ (ニドラン♂)
	&_pokemon{Nidorino, Nidoran_male, 122, 142, 128, t.Poison, t_, 19.5, 0.9, Nidoran_male, 100},     // 033: Nidorino (ニドリーノ)
	&_pokemon{Nidoking, Nidorino, 162, 204, 170, t.Poison, t.Ground, 62.0, 1.4, Nidoran_male, 0},     // 034: Nidoking (ニドキング)
	&_pokemon{Clefairy, p_, 140, 116, 124, t.Fairy, t_, 7.5, 0.6, Clefairy, 50},                      // 035: Clefairy (ピッピ), 173 -> 035
	&_pokemon{Clefable, Clefairy, 190, 178, 178, t.Fairy, t_, 40.0, 1.3, Clefairy, 0},                // 036: Clefable (ピクシー)
	&_pokemon{Vulpix, p_, 76, 106, 118, t.Fire, t_, 9.9, 0.6, Vulpix, 50},                            // 037: Vulpix (ロコン)
	&_pokemon{Ninetales, Vulpix, 146, 176, 194, t.Fire, t_, 19.9, 1.1, Vulpix, 0},                    // 038: Ninetales (キュウコン)
	&_pokemon{Jigglypuff, p_, 230, 98, 54, t.Normal, t.Fairy, 5.5, 0.5, Jigglypuff, 50},              // 039: Jigglypuff (プリン), 174 -> 039
	&_pokemon{Wigglytuff, Jigglypuff, 280, 168, 108, t.Normal, t.Fairy, 12.0, 1.0, Jigglypuff, 0},    // 040: Wigglytuff (プクリン)
	&_pokemon{Zubat, p_, 80, 88, 90, t.Poison, t.Flying, 7.5, 0.8, Zubat, 50},                        // 041: Zubat (ズバット)
	&_pokemon{Golbat, Zubat, 150, 164, 164, t.Poison, t.Flying, 55.0, 1.6, Zubat, 0},                 // 042: Golbat (ゴルバット), 042 -> 169
	&_pokemon{Oddish, p_, 90, 134, 130, t.Grass, t.Poison, 5.4, 0.5, Oddish, 25},                     // 043: Oddish (ナゾノクサ)
	&_pokemon{Gloom, Oddish, 120, 162, 158, t.Grass, t.Poison, 8.6, 0.8, Oddish, 100},                // 044: Gloom (クサイハナ), 044 -> 045/182
	&_pokemon{Vileplume, Gloom, 150, 202, 190, t.Grass, t.Poison, 18.6, 1.2, Oddish, 0},              // 045: Vileplume (ラフレシア)
	&_pokemon{Paras, p_, 70, 122, 120, t.Bug, t.Grass, 5.4, 0.3, Paras, 50},                          // 046: Paras (パラス)
	&_pokemon{Parasect, Paras, 120, 162, 170, t.Bug, t.Grass, 29.5, 1.0, Paras, 0},                   // 047: Parasect (パラセクト)
	&_pokemon{Venonat, p_, 120, 108, 118, t.Bug, t.Poison, 30.0, 1.0, Venonat, 50},                   // 048: Venonat (コンパン)
	&_pokemon{Venomoth, Venonat, 140, 172, 154, t.Bug, t.Poison, 12.5, 1.5, Venonat, 0},              // 049: Venomoth (モルフォン)
	&_pokemon{Diglett, p_, 20, 108, 86, t.Ground, t_, 0.8, 0.2, Diglett, 50},                         // 050: Diglett (ディグダ)
	&_pokemon{Dugtrio, Diglett, 70, 148, 140, t.Ground, t_, 33.3, 0.7, Diglett, 0},                   // 051: Dugtrio (ダグトリオ)
	&_pokemon{Meowth, p_, 80, 104, 94, t.Normal, t_, 4.2, 0.4, Meowth, 50},                           // 052: Meowth (ニャース)
	&_pokemon{Persian, Meowth, 130, 156, 146, t.Normal, t_, 32.0, 1.0, Meowth, 0},                    // 053: Persian (ペルシアン)
	&_pokemon{Psyduck, p_, 100, 132, 112, t.Water, t_, 19.6, 0.8, Psyduck, 50},                       // 054: Psyduck (コダック)
	&_pokemon{Golduck, Psyduck, 160, 194, 176, t.Water, t_, 76.6, 1.7, Psyduck, 0},                   // 055: Golduck (ゴルダック)
	&_pokemon{Mankey, p_, 80, 112, 96, t.Fighting, t_, 28.0, 0.5, Mankey, 50},                        // 056: Mankey (マンキー)
	&_pokemon{Primeape, Mankey, 130, 178, 150, t.Fighting, t_, 32.0, 1.0, Mankey, 0},                 // 057: Primeape (オコリザル)
	&_pokemon{Growlithe, p_, 110, 156, 110, t.Fire, t_, 19.0, 0.7, Growlithe, 50},                    // 058: Growlithe (ガーディ)
	&_pokemon{Arcanine, Growlithe, 180, 230, 180, t.Fire, t_, 155.0, 1.9, Growlithe, 0},              // 059: Arcanine (ウインディ)
	&_pokemon{Poliwag, p_, 80, 108, 98, t.Water, t_, 12.4, 0.6, Poliwag, 25},                         // 060: Poliwag (ニョロモ)
	&_pokemon{Poliwhirl, Poliwag, 130, 132, 132, t.Water, t_, 20.0, 1.0, Poliwag, 100},               // 061: Poliwhirl (ニョロゾ), 061 -> 062/186
	&_pokemon{Poliwrath, Poliwhirl, 180, 180, 202, t.Water, t.Fighting, 54.0, 1.3, Poliwag, 0},       // 062: Poliwrath (ニョロボン)
	&_pokemon{Abra, p_, 50, 110, 76, t.Psychic, t_, 19.5, 0.9, Abra, 25},                             // 063: Abra (ケーシィ)
	&_pokemon{Kadabra, Abra, 80, 150, 112, t.Psychic, t_, 56.5, 1.3, Abra, 100},                      // 064: Kadabra (ユンゲラー)
	&_pokemon{Alakazam, Kadabra, 110, 186, 152, t.Psychic, t_, 48.0, 1.5, Abra, 0},                   // 065: Alakazam (フーディン)
	&_pokemon{Machop, p_, 140, 118, 96, t.Fighting, t_, 19.5, 0.8, Machop, 25},                       // 066: Machop (ワンリキー)
	&_pokemon{Machoke, Machop, 160, 154, 144, t.Fighting, t_, 70.5, 1.5, Machop, 100},                // 067: Machoke (ゴーリキー)
	&_pokemon{Machamp, Machoke, 180, 198, 180, t.Fighting, t_, 130.0, 1.6, Machop, 0},                // 068: Machamp (カイリキー)
	&_pokemon{Bellsprout, p_, 100, 158, 78, t.Grass, t.Poison, 4.0, 0.7, Bellsprout, 25},             // 069: Bellsprout (マダツボミ)
	&_pokemon{Weepinbell, Bellsprout, 130, 190, 110, t.Grass, t.Poison, 6.4, 1.0, Bellsprout, 100},   // 070: Weepinbell (ウツドン)
	&_pokemon{Victreebel, Weepinbell, 160, 222, 152, t.Grass, t.Poison, 15.5, 1.7, Bellsprout, 0},    // 071: Victreebel (ウツボット)
	&_pokemon{Tentacool, p_, 80, 106, 136, t.Water, t.Poison, 45.5, 0.9, Tentacool, 50},              // 072: Tentacool (メノクラゲ)
	&_pokemon{Tentacruel, Tentacool, 160, 170, 196, t.Water, t.Poison, 55.0, 1.6, Tentacool, 0},      // 073: Tentacruel (ドククラゲ)
	&_pokemon{Geodude, p_, 80, 106, 118, t.Rock, t.Ground, 20.0, 0.4, Geodude, 25},                   // 074: Geodude (イシツブテ)
	&_pokemon{Graveler, Geodude, 110, 142, 156, t.Rock, t.Ground, 105.0, 1.0, Geodude, 100},          // 075: Graveler (ゴローン)
	&_pokemon{Golem, Graveler, 160, 176, 198, t.Rock, t.Ground, 300.0, 1.4, Geodude, 0},              // 076: Golem (ゴローニャ)
	&_pokemon{Ponyta, p_, 100, 168, 138, t.Fire, t_, 30.0, 1.0, Ponyta, 50},                          // 077: Ponyta (ポニータ)
	&_pokemon{Rapidash, Ponyta, 130, 200, 170, t.Fire, t_, 95.0, 1.7, Ponyta, 0},                     // 078: Rapidash (ギャロップ)
	&_pokemon{Slowpoke, p_, 180, 110, 110, t.Water, t.Psychic, 36.0, 1.2, Slowpoke, 50},              // 079: Slowpoke (ヤドン), 079 -> 080/199
	&_pokemon{Slowbro, Slowpoke, 190, 184, 198, t.Water, t.Psychic, 78.5, 1.6, Slowpoke, 0},          // 080: Slowbro (ヤドラン)
	&_pokemon{Magnemite, p_, 50, 128, 138, t.Electric, t.Steel, 6.0, 0.3, Magnemite, 50},             // 081: Magnemite (コイル)
	&_pokemon{Magneton, Magnemite, 100, 186, 180, t.Electric, t.Steel, 60.0, 1.0, Magnemite, 0},      // 082: Magneton (レアコイル), 082 -> 462
	&_pokemon{Farfetch_d, p_, 104, 138, 132, t.Normal, t.Flying, 15.0, 0.8, Farfetch_d, 0},           // 083: Farfetch'd (カモネギ)
	&_pokemon{Doduo, p_, 70, 126, 96, t.Normal, t.Flying, 39.2, 1.4, Doduo, 50},                      // 084: Doduo (ドードー)
	&_pokemon{Dodrio, Doduo, 120, 182, 150, t.Normal, t.Flying, 85.2, 1.8, Doduo, 0},                 // 085: Dodrio (ドードリオ)
	&_pokemon{Seel, p_, 130, 104, 138, t.Water, t_, 90.0, 1.1, Seel, 50},                             // 086: Seel (パウワウ)
	&_pokemon{Dewgong, Seel, 180, 156, 192, t.Water, t.Ice, 120.0, 1.7, Seel, 0},                     // 087: Dewgong (ジュゴン)
	&_pokemon{Grimer, p_, 160, 124, 110, t.Poison, t_, 30.0, 0.9, Grimer, 50},                        // 088: Grimer (ベトベター)
	&_pokemon{Muk, Grimer, 210, 180, 188, t.Poison, t_, 30.0, 1.2, Grimer, 0},                        // 089: Muk (ベトベトン)
	&_pokemon{Shellder, p_, 60, 120, 112, t.Water, t_, 4.0, 0.3, Shellder, 50},                       // 090: Shellder (シェルダー)
	&_pokemon{Cloyster, Shellder, 100, 196, 196, t.Water, t.Ice, 132.5, 1.5, Shellder, 0},            // 091: Cloyster (パルシェン)
	&_pokemon{Gastly, p_, 60, 136, 82, t.Ghost, t.Poison, 0.1, 1.3, Gastly, 25},                      // 092: Gastly (ゴース)
	&_pokemon{Haunter, Gastly, 90, 172, 118, t.Ghost, t.Poison, 0.1, 1.6, Gastly, 100},               // 093: Haunter (ゴースト)
	&_pokemon{Gengar, Haunter, 120, 204, 156, t.Ghost, t.Poison, 40.5, 1.5, Gastly, 0},               // 094: Gengar (ゲンガー)
	&_pokemon{Onix, p_, 70, 90, 186, t.Rock, t.Ground, 210.0, 8.8, Onix, 0},                          // 095: Onix (イワーク), 095 -> 208
	&_pokemon{Drowzee, p_, 120, 104, 140, t.Psychic, t_, 32.4, 1.0, Drowzee, 50},                     // 096: Drowzee (スリープ)
	&_pokemon{Hypno, Drowzee, 170, 162, 196, t.Psychic, t_, 75.6, 1.6, Drowzee, 0},                   // 097: Hypno (スリーパー)
	&_pokemon{Krabby, p_, 60, 116, 110, t.Water, t_, 6.5, 0.4, Krabby, 50},                           // 098: Krabby (クラブ)
	&_pokemon{Kingler, Krabby, 110, 178, 168, t.Water, t_, 60.0, 1.3, Krabby, 0},                     // 099: Kingler (キングラー)
	&_pokemon{Voltorb, p_, 80, 102, 124, t.Electric, t_, 10.4, 0.5, Voltorb, 50},                     // 100: Voltorb (ビリリダマ)
	&_pokemon{Electrode, Voltorb, 120, 150, 174, t.Electric, t_, 66.6, 1.2, Voltorb, 0},              // 101: Electrode (マルマイン)
	&_pokemon{Exeggcute, p_, 120, 110, 132, t.Grass, t.Psychic, 2.5, 0.4, Exeggcute, 50},             // 102: Exeggcute (タマタマ)
	&_pokemon{Exeggutor, Exeggcute, 190, 232, 164, t.Grass, t.Psychic, 120.0, 2.0, Exeggcute, 0},     // 103: Exeggutor (ナッシー)
	&_pokemon{Cubone, p_, 100, 102, 150, t.Ground, t_, 6.5, 0.4, Cubone, 50},                         // 104: Cubone (カラカラ)
	&_pokemon{Marowak, Cubone, 120, 140, 202, t.Ground, t_, 45.0, 1.0, Cubone, 0},                    // 105: Marowak (ガラガラ)
	&_pokemon{Hitmonlee, p_, 100, 148, 172, t.Fighting, t_, 49.8, 1.5, Hitmonlee, 0},                 // 106: Hitmonlee (サワムラー), 236 -> 106
	&_pokemon{Hitmonchan, p_, 100, 138, 204, t.Fighting, t_, 50.2, 1.4, Hitmonchan, 0},               // 107: Hitmonchan (エビワラー), 236 -> 107
	&_pokemon{Lickitung, p_, 180, 126, 160, t.Normal, t_, 65.5, 1.2, Lickitung, 0},                   // 108: Lickitung (ベロリンガ), 108 -> 463
	&_pokemon{Koffing, p_, 80, 136, 142, t.Poison, t_, 1.0, 0.6, Koffing, 50},                        // 109: Koffing (ドガース)
	&_pokemon{Weezing, Koffing, 130, 190, 198, t.Poison, t_, 9.5, 1.2, Koffing, 0},                   // 110: Weezing (マタドガス)
	&_pokemon{Rhyhorn, p_, 160, 110, 116, t.Ground, t.Rock, 115.0, 1.0, Rhyhorn, 50},                 // 111: Rhyhorn (サイホーン)
	&_pokemon{Rhydon, Rhyhorn, 210, 116, 160, t.Ground, t.Rock, 120.0, 1.9, Rhyhorn, 0},              // 112: Rhydon (サイドン), 112 -> 464
	&_pokemon{Chansey, p_, 500, 40, 60, t.Normal, t_, 34.6, 1.1, Chansey, 0},                         // 113: Chansey (ラッキー), 440 -> 113 -> 242
	&_pokemon{Tangela, p_, 130, 164, 152, t.Grass, t_, 35.0, 1.0, Tangela, 0},                        // 114: Tangela (モンジャラ), 114 -> 465
	&_pokemon{Kangaskhan, p_, 210, 142, 178, t.Normal, t_, 80.0, 2.2, Kangaskhan, 0},                 // 115: Kangaskhan (ガルーラ)
	&_pokemon{Horsea, p_, 60, 122, 100, t.Water, t_, 8.0, 0.4, Horsea, 50},                           // 116: Horsea (タッツー)
	&_pokemon{Seadra, Horsea, 110, 176, 150, t.Water, t_, 25.0, 1.2, Horsea, 0},                      // 117: Seadra (シードラ), 117 -> 230
	&_pokemon{Goldeen, p_, 90, 112, 126, t.Water, t_, 15.0, 0.6, Goldeen, 50},                        // 118: Goldeen (トサキント)
	&_pokemon{Seaking, Goldeen, 160, 172, 160, t.Water, t_, 39.0, 1.3, Goldeen, 0},                   // 119: Seaking (アズマオウ)
	&_pokemon{Staryu, p_, 60, 130, 128, t.Water, t_, 34.5, 0.8, Staryu, 50},                          // 120: Staryu (ヒトデマン)
	&_pokemon{Starmie, Staryu, 120, 194, 192, t.Water, t.Psychic, 80.0, 1.1, Staryu, 0},              // 121: Starmie (スターミー)
	&_pokemon{Mr_Mime, p_, 80, 154, 196, t.Psychic, t.Fairy, 54.5, 1.3, Mr_Mime, 0},                  // 122: Mr. Mime (バリヤード), 439 -> 122
	&_pokemon{Scyther, p_, 140, 176, 180, t.Bug, t.Flying, 56.0, 1.5, Scyther, 0},                    // 123: Scyther (ストライク), 123 -> 212
	&_pokemon{Jynx, p_, 130, 172, 134, t.Ice, t.Psychic, 40.6, 1.4, Jynx, 0},                         // 124: Jynx (ルージュラ), 238 -> 124
	&_pokemon{Electabuzz, p_, 130, 198, 160, t.Electric, t_, 30.0, 1.1, Electabuzz, 0},               // 125: Electabuzz (エレブー), 239 -> 125 -> 466
	&_pokemon{Magmar, p_, 130, 214, 158, t.Fire, t_, 44.5, 1.3, Magmar, 0},                           // 126: Magmar (ブーバー), 240 -> 126 -> 467
	&_pokemon{Pinsir, p_, 130, 184, 186, t.Bug, t_, 55.0, 1.5, Pinsir, 0},                            // 127: Pinsir (カイロス)
	&_pokemon{Tauros, p_, 150, 148, 184, t.Normal, t_, 88.4, 1.4, Tauros, 0},                         // 128: Tauros (ケンタロス)
	&_pokemon{Magikarp, p_, 40, 42, 84, t.Water, t_, 10.0, 0.9, Magikarp, 400},                       // 129: Magikarp (コイキング)
	&_pokemon{Gyarados, Magikarp, 190, 192, 196, t.Water, t.Flying, 235.0, 6.5, Magikarp, 0},         // 130: Gyarados (ギャラドス)
	&_pokemon{Lapras, p_, 260, 186, 190, t.Water, t.Ice, 220.0, 2.5, Lapras, 0},                      // 131: Lapras (ラプラス)
	&_pokemon{Ditto, p_, 96, 110, 110, t.Normal, t_, 4.0, 0.3, Ditto, 0},                             // 132: Ditto (メタモン)
	&_pokemon{Eevee, p_, 110, 114, 128, t.Normal, t_, 6.5, 0.3, Eevee, 25},                           // 133: Eevee (イーブイ), 133 -> 134/135/136/196/197/470/471/700
	&_pokemon{Vaporeon, Eevee, 260, 186, 168, t.Water, t_, 29.0, 1.0, Eevee, 0},                      // 134: Vaporeon (シャワーズ)
	&_pokemon{Jolteon, Eevee, 130, 192, 174, t.Electric, t_, 24.5, 0.8, Eevee, 0},                    // 135: Jolteon (サンダース)
	&_pokemon{Flareon, Eevee, 130, 238, 178, t.Fire, t_, 25.0, 0.9, Eevee, 0},                        // 136: Flareon (ブースター)
	&_pokemon{Porygon, p_, 130, 156, 158, t.Normal, t_, 36.5, 0.8, Porygon, 0},                       // 137: Porygon (ポリゴン), 137 -> 233 -> 474
	&_pokemon{Omanyte, p_, 70, 132, 160, t.Rock, t.Water, 7.5, 0.4, Omanyte, 50},                     // 138: Omanyte (オムナイト)
	&_pokemon{Omastar, Omanyte, 140, 180, 202, t.Rock, t.Water, 35.0, 1.0, Omanyte, 0},               // 139: Omastar (オムスター)
	&_pokemon{Kabuto, p_, 60, 148, 142, t.Rock, t.Water, 11.5, 0.5, Kabuto, 50},                      // 140: Kabuto (カブト)
	&_pokemon{Kabutops, Kabuto, 120, 190, 190, t.Rock, t.Water, 40.5, 1.3, Kabuto, 0},                // 141: Kabutops (カブトプス)
	&_pokemon{Aerodactyl, p_, 160, 182, 162, t.Rock, t.Flying, 59.0, 1.8, Aerodactyl, 0},             // 142: Aerodactyl (プテラ)
	&_pokemon{Snorlax, p_, 320, 180, 180, t.Normal, t_, 460.0, 2.1, Snorlax, 0},                      // 143: Snorlax (カビゴン), 446 -> 143
	&_pokemon{Articuno, p_, 180, 198, 242, t.Ice, t.Flying, 55.4, 1.7, Articuno, 0},                  // 144: Articuno (フリーザー)
	&_pokemon{Zapdos, p_, 180, 232, 194, t.Electric, t.Flying, 52.6, 1.6, Zapdos, 0},                 // 145: Zapdos (サンダー)
	&_pokemon{Moltres, p_, 180, 242, 194, t.Fire, t.Flying, 60.0, 2.0, Moltres, 0},                   // 146: Moltres (ファイヤー)
	&_pokemon{Dratini, p_, 82, 128, 110, t.Dragon, t_, 3.3, 1.8, Dratini, 25},                        // 147: Dratini (ミニリュウ)
	&_pokemon{Dragonair, Dratini, 122, 170, 152, t.Dragon, t_, 16.5, 4.0, Dratini, 100},              // 148: Dragonair (ハクリュー)
	&_pokemon{Dragonite, Dragonair, 182, 250, 212, t.Dragon, t.Flying, 210.0, 2.2, Dratini, 0},       // 149: Dragonite (カイリュー)
	&_pokemon{Mewtwo, p_, 212, 284, 202, t.Psychic, t_, 122.0, 2.0, Mewtwo, 0},                       // 150: Mewtwo (ミュウツー)
	&_pokemon{Mew, p_, 200, 220, 220, t.Psychic, t_, 4.0, 0.4, Mew, 0},                               // 151: Mew (ミュウ)
}
