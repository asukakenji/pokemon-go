package zhHK

// Pokemon
type Pokemon int16

const (
	_ Pokemon = iota
	奇異種子
	奇異草
	奇異花
	小火龍
	火恐龍
	噴火龍
	車厘龜
	卡美龜
	水箭龜
	綠毛蟲
	鐵甲蟲
	巴他蝶
	獨角蟲
	鐵殼蛹
	大針蜂
	波波
	比比鳥
	大比鳥
	小哥達
	哥達
	鬼雀
	魔雀
	阿柏蛇
	阿柏怪
	比卡超
	雷超
	穿山鼠
	穿山王
	尼美蘭
	尼美羅
	尼美后
	尼多郎
	尼多利
	尼多王
	皮皮
	皮可斯
	六尾
	九尾
	波波球
	肥波球
	波音蝠
	大口蝠
	行路草
	怪味花
	霸王花
	蘑菇蟲
	巨菇蟲
	毛毛蟲
	魔魯風
	地鼠
	三頭地鼠
	喵喵怪
	高竇貓
	傻鴨
	高超鴨
	猴怪
	火爆猴
	護主犬
	奉神犬
	蚊香蝌蚪
	蚊香蛙
	大力蛙
	卡斯
	尤基納
	富迪
	鐵腕
	大力
	怪力
	喇叭芽
	口呆花
	大食花
	大眼水母
	多腳水母
	小拳石
	滾動石
	滾動岩
	小火馬
	烈焰馬
	小呆獸
	大呆獸
	小磁怪
	三合一磁怪
	火蔥鴨
	多多
	多多利
	小海獅
	白海獅
	爛泥怪
	爛泥獸
	貝殼怪
	鐵甲貝
	鬼斯
	鬼斯通
	耿鬼
	大岩蛇
	食夢獸
	催眠獸
	大鉗蟹
	巨鉗蟹
	霹靂蛋
	雷霆蛋
	蛋蛋
	椰樹獸
	卡拉卡拉
	格拉格拉
	沙古拉
	比華拉
	大舌頭
	毒氣丸
	毒氣雙子
	鐵甲犀牛
	鐵甲暴龍
	吉利蛋
	長藤怪
	袋獸
	噴墨海馬
	飛刺海馬
	獨角金魚
	金魚王
	海星星
	寶石海星
	吸盤小丑
	飛天螳螂
	紅唇娃
	電擊獸
	鴨嘴火龍
	鉗刀甲蟲
	大隻牛
	鯉魚王
	鯉魚龍
	背背龍
	百變怪
	伊貝
	水伊貝
	雷伊貝
	火伊貝
	立方獸
	菊石獸
	多刺菊石獸
	萬年蟲
	鐮刀蟲
	化石飛龍
	卡比獸
	急凍鳥
	雷鳥
	火鳥
	迷你龍
	哈古龍
	啟暴龍
	超夢夢
	夢夢
)

//go:generate stringer -type=Pokemon
