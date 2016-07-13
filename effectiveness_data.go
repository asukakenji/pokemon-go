package pokemon

// Official Table: http://www.pokemon.co.jp/whats/summary/
var table = [19][19]Effectiveness{
	/* ATK -> -  無  炎  水  草   電  氷  闘  毒   地  飛   超  虫  岩   霊  竜  悪   鋼  妖 */
	/* - */ {E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_},
	/* 無 */ {E_, E_, E_, E_, E_, E_, E_, EO, E_, E_, E_, E_, E_, E_, EX, E_, E_, E_, E_},
	/* 炎 */ {E_, E_, EA, EO, EA, E_, EA, E_, E_, EO, E_, E_, EA, EO, E_, E_, E_, EA, EA},
	/* 水 */ {E_, E_, EA, EA, EO, EO, EA, E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, EA, E_},
	/* 草 */ {E_, E_, EO, EA, EA, EA, EO, E_, EO, EA, EO, E_, EO, E_, E_, E_, E_, E_, E_},
	/* 電 */ {E_, E_, E_, E_, E_, EA, E_, E_, E_, EO, EA, E_, E_, E_, E_, E_, E_, EA, E_},
	/* 氷 */ {E_, E_, EO, E_, E_, E_, EA, EO, E_, E_, E_, E_, E_, EO, E_, E_, E_, EO, E_},
	/* 闘 */ {E_, E_, E_, E_, E_, E_, E_, E_, E_, E_, EO, EO, EA, EA, E_, E_, EA, E_, EO},
	/* 毒 */ {E_, E_, E_, E_, EA, E_, E_, EA, EA, EO, E_, EO, EA, E_, E_, E_, E_, E_, EA},
	/* 地 */ {E_, E_, E_, EO, EO, EX, EO, E_, EA, E_, E_, E_, E_, EA, E_, E_, E_, E_, E_},
	/* 飛 */ {E_, E_, E_, E_, EA, EO, EO, EA, E_, EX, E_, E_, EA, EO, E_, E_, E_, E_, E_},
	/* 超 */ {E_, E_, E_, E_, E_, E_, E_, EA, E_, E_, E_, EA, EO, E_, EO, E_, EO, E_, E_},
	/* 虫 */ {E_, E_, EO, E_, EA, E_, E_, EA, E_, EA, EO, E_, E_, EO, E_, E_, E_, E_, E_},
	/* 岩 */ {E_, EA, EA, EO, EO, E_, E_, EO, EA, EO, EA, E_, E_, E_, E_, E_, E_, EO, E_},
	/* 霊 */ {E_, EX, E_, E_, E_, E_, E_, EX, EA, E_, E_, E_, EA, E_, EO, E_, EO, E_, E_},
	/* 竜 */ {E_, E_, EA, EA, EA, EA, EO, E_, E_, E_, E_, E_, E_, E_, E_, EO, E_, E_, EO},
	/* 悪 */ {E_, E_, E_, E_, E_, E_, E_, EO, E_, E_, E_, EX, EO, E_, EA, E_, EA, E_, EO},
	/* 鋼 */ {E_, EA, EO, E_, EA, E_, EA, EO, EX, EO, EA, EA, EA, EA, E_, EA, E_, EA, EA},
	/* 妖 */ {E_, E_, E_, E_, E_, E_, E_, EA, EO, E_, E_, E_, EA, E_, E_, EX, EA, EO, E_},
}
