package _type

import (
	e "github.com/asukakenji/pokemon-go/eff"
)

// Official Table: http://www.pokemon.co.jp/whats/summary/
var effectivenesses = [19][19]e.Effectiveness{
	/* ATK -> -  無  炎  水  草   電  氷  闘  毒   地  飛   超  虫  岩   霊  竜  悪   鋼  妖 */
	/* - */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_},
	/* 無 */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EO, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EX, e.E_, e.E_, e.E_, e.E_},
	/* 炎 */ {e.E_, e.E_, e.EA, e.EO, e.EA, e.E_, e.EA, e.E_, e.E_, e.EO, e.E_, e.E_, e.EA, e.EO, e.E_, e.E_, e.E_, e.EA, e.EA},
	/* 水 */ {e.E_, e.E_, e.EA, e.EA, e.EO, e.EO, e.EA, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EA, e.E_},
	/* 草 */ {e.E_, e.E_, e.EO, e.EA, e.EA, e.EA, e.EO, e.E_, e.EO, e.EA, e.EO, e.E_, e.EO, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_},
	/* 電 */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.EA, e.E_, e.E_, e.E_, e.EO, e.EA, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EA, e.E_},
	/* 氷 */ {e.E_, e.E_, e.EO, e.E_, e.E_, e.E_, e.EA, e.EO, e.E_, e.E_, e.E_, e.E_, e.E_, e.EO, e.E_, e.E_, e.E_, e.EO, e.E_},
	/* 闘 */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EO, e.EO, e.EA, e.EA, e.E_, e.E_, e.EA, e.E_, e.EO},
	/* 毒 */ {e.E_, e.E_, e.E_, e.E_, e.EA, e.E_, e.E_, e.EA, e.EA, e.EO, e.E_, e.EO, e.EA, e.E_, e.E_, e.E_, e.E_, e.E_, e.EA},
	/* 地 */ {e.E_, e.E_, e.E_, e.EO, e.EO, e.EX, e.EO, e.E_, e.EA, e.E_, e.E_, e.E_, e.E_, e.EA, e.E_, e.E_, e.E_, e.E_, e.E_},
	/* 飛 */ {e.E_, e.E_, e.E_, e.E_, e.EA, e.EO, e.EO, e.EA, e.E_, e.EX, e.E_, e.E_, e.EA, e.EO, e.E_, e.E_, e.E_, e.E_, e.E_},
	/* 超 */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EA, e.E_, e.E_, e.E_, e.EA, e.EO, e.E_, e.EO, e.E_, e.EO, e.E_, e.E_},
	/* 虫 */ {e.E_, e.E_, e.EO, e.E_, e.EA, e.E_, e.E_, e.EA, e.E_, e.EA, e.EO, e.E_, e.E_, e.EO, e.E_, e.E_, e.E_, e.E_, e.E_},
	/* 岩 */ {e.E_, e.EA, e.EA, e.EO, e.EO, e.E_, e.E_, e.EO, e.EA, e.EO, e.EA, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EO, e.E_},
	/* 霊 */ {e.E_, e.EX, e.E_, e.E_, e.E_, e.E_, e.E_, e.EX, e.EA, e.E_, e.E_, e.E_, e.EA, e.E_, e.EO, e.E_, e.EO, e.E_, e.E_},
	/* 竜 */ {e.E_, e.E_, e.EA, e.EA, e.EA, e.EA, e.EO, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EO, e.E_, e.E_, e.EO},
	/* 悪 */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EO, e.E_, e.E_, e.E_, e.EX, e.EO, e.E_, e.EA, e.E_, e.EA, e.E_, e.EO},
	/* 鋼 */ {e.E_, e.EA, e.EO, e.E_, e.EA, e.E_, e.EA, e.EO, e.EX, e.EO, e.EA, e.EA, e.EA, e.EA, e.E_, e.EA, e.E_, e.EA, e.EA},
	/* 妖 */ {e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.E_, e.EA, e.EO, e.E_, e.E_, e.E_, e.EA, e.E_, e.E_, e.EX, e.EA, e.EO, e.E_},
}
