# Pokémon GO

[Pokémon GO][pokemon-go] is a free-to-play location-based, augmented reality game developed and published by [Niantic][niantic] for [iOS][ios] and [Android][android] devices. It was initially released in selected countries in July 2016. In the game, players use the smart device's GPS and camera to capture, battle, and train virtual creatures, called Pokémon, who appear on the screen as if they were in the same real-world location as the player.

This project aims primarily at providing the value tables used in the game, and an API the retrieve them. Current features include:
- Cheatsheet (Not Currently Public):
  - A table containing all Pokémons and their information.
  - It will be available publicly soon.
- CP / HP Calculator:
  - Given the "IVs (Individual Values)" and "Level" of a Pokémon, calculate its "CP (Combat Power)" and "HP (Hit Points). This is basically used to test the "IV Calculator"
  - The source code is available [here][cp_hp_calc].
  - Binaries will be available for Windows, Mac, and Linux soon.
- IV Calculator:
  - Given the "CP (Combat Power)", "HP (Hit Points)", "Stardust to Power Up", and "Candy to Power Up" of a Pokémon, calculate its possible "IVs (Individual Values)" and "Level" combinations. This is probably the most wanted feature.
  - The source code is available [here][iv_calc].
  - Binaries will be available for Windows, Mac, and Linux soon.
- Multi-lingual Ready:
  - The library supports multiple languages: local names and translations are available. Here is a list of currently supported languages: Japanese (日本語), English, French (Français), German (Deutsch), Italian (Italiano), Korean (한국어), Spanish (Español), Simplified Chinese (简体中文（官方译名）), Traditional Chinese (繁體中文（官方譯名）), PRC Chinese (简体中文（中国译名）), Hong Kong Chinese (繁體中文（香港譯名）), Taiwan Chinese (繁體中文（台灣譯名）).
  - For instance, the Pokémon with ID #001 is named: [フシギダネ (日本語)](pokemon/internal/ja/pokemon.go#L8), [Bulbasaur (English)](pokemon/internal/en/pokemon.go#L8), [Bulbizarre (Français)](pokemon/internal/fr/pokemon.go#L8), [Bisasam (Deutsch)](pokemon/internal/de/pokemon.go#L8), [Bulbasaur (Italiano)](pokemon/internal/it/pokemon.go#L8), [이상해씨 (한국어)](pokemon/internal/ko/pokemon.go#L8), [Bulbasaur (Español)](pokemon/internal/es/pokemon.go#L8), [妙蛙种子 (简体中文（官方译名）)](pokemon/internal/zh-CHS/pokemon.go#L8), [妙蛙種子 (繁體中文（官方譯名）)](pokemon/internal/zh-CHT/pokemon.go#L8), [妙蛙种子 (简体中文（中国译名）)](pokemon/internal/zh-CN/pokemon.go#L8), [奇異種子 (繁體中文（香港譯名）)](pokemon/internal/zh-HK/pokemon.go#L8), [妙蛙種子 (繁體中文（台灣譯名）)](pokemon/internal/zh-TW/pokemon.go#L8).
  - For instance, the Pokémon with ID #100 is named: [ビリリダマ (日本語)](pokemon/internal/ja/pokemon.go#L107), [Voltorb (English)](pokemon/internal/en/pokemon.go#L107), [Voltorbe (Français)](pokemon/internal/fr/pokemon.go#L107), [Voltobal (Deutsch)](pokemon/internal/de/pokemon.go#L107), [Voltorb (Italiano)](pokemon/internal/it/pokemon.go#L107), [찌리리공 (한국어)](pokemon/internal/ko/pokemon.go#L107), [Voltorb (Español)](pokemon/internal/es/pokemon.go#L107), [霹雳电球 (简体中文（官方译名）)](pokemon/internal/zh-CHS/pokemon.go#L107), [霹靂電球 (繁體中文（官方譯名）)](pokemon/internal/zh-CHT/pokemon.go#L107), [雷电球 (简体中文（中国译名）)](pokemon/internal/zh-CN/pokemon.go#L107), [霹靂蛋 (繁體中文（香港譯名）)](pokemon/internal/zh-HK/pokemon.go#L107), [雷電球 (繁體中文（台灣譯名）)](pokemon/internal/zh-TW/pokemon.go#L107).

[pokemon-go]: http://www.pokemongo.com/
[ios]: https://itunes.apple.com/app/pokemon-go/id1094591345
[android]: https://play.google.com/store/apps/details?id=com.nianticlabs.pokemongo
[niantic]: https://nianticlabs.com/
[cp_hp_calc]: cmd/cp_hp_calc/
[iv_calc]: cmd/iv_calc/
