# Pokémon GO Tools and Library in Go (Golang)

[Pokémon GO][pokemon-go] is a free-to-play location-based, augmented reality game developed and published by [Niantic][niantic] for [iOS][ios] and [Android][android] devices. It was initially released in selected countries in July 2016. In the game, players use the smart device's GPS and camera to capture, battle, and train virtual creatures, called Pokémon, who appear on the screen as if they were in the same real-world location as the player.

This project aims primarily at providing the data tables used in the game, an API the retrieve them, and tools to assist playing the game.

## Online Resources
- A list of Pokémons in different languages (available [here][pokemon])
- A table of Pokémons and their information (coming soon)
- A list of Pokémon types in different languages (available [here][type])
- A table of Pokémon types and effectivenesses (coming soon)
- A list of Pokémon moves in different languages (coming soon)
- A table of Pokémon moves and their information (coming soon)

## Executables
- IV Calculator Pro (development in progress):
  - Like "IV Calculator".
  - Also deals with "Power Up" and "Evolve".
  - Stores Pokémons in a database for follow ups.
- **IV Calculator**:
  - Given the "CP (Combat Power)", "HP (Hit Points)", "Stardust to Power Up", and "Candy to Power Up" of a Pokémon, calculate its possible "IVs (Individual Values)" and "Level" combinations.
  - This is probably the most wanted feature.
  - The source code is available [here][iv_calc].
  - Binaries (coming soon):
    - Windows (32-bit / x86)
    - Windows (64-bit / x64)
    - Mac (64-bit / x64)
    - Linux (32-bit / x86)
    - Linux (64-bit / x64)
- CP / HP Calculator:
  - Given the "IVs (Individual Values)" and "Level" of a Pokémon, calculate its "CP (Combat Power)" and "HP (Hit Points).
  - This is basically used to test the "IV Calculator" and compare implementations written by other developers.
  - The source code is available [here][cp_hp_calc].
  - Binaries (coming soon):
    - Windows (32-bit / x86)
    - Windows (64-bit / x64)
    - Mac (64-bit / x64)
    - Linux (32-bit / x86)
    - Linux (64-bit / x64)
- Cheatsheet Generator (coming soon):
  - Generate data tables.
  - Supported output formats:
    - Markdown (`*.md`)
    - HTML (`*.html`)

## APIs
- Package `pokemon`
  - This package deals with "Pokémon" in the game.
- Package `move`
  - This package deals with "Move" in the game.
- Package `weak`
  - This package deals with "Weakness" in the game.
- Package `type` (actually `_type`, since `type` is a keyword in Go (Golang))
  - This package deals with "Type" in the game.
- Package `eff`
  - This package deals with "Effectiveness" in the game.
- Package `lang`
  - This package deals with multi-lingual issues. Local names and translations are available. The following languages and currently supported: Japanese (日本語), English, French (Français), German (Deutsch), Italian (Italiano), Korean (한국어), Spanish (Español), Simplified Chinese (简体中文（官方译名）), Traditional Chinese (繁體中文（官方譯名）), PRC Chinese (简体中文（中国译名）), Hong Kong Chinese (繁體中文（香港譯名）), Taiwan Chinese (繁體中文（台灣譯名）).
  - For instance, the Pokémon with ID #001 is named: [フシギダネ (日本語)](pokemon/internal/ja/pokemon.go#L8), [Bulbasaur (English)](pokemon/internal/en/pokemon.go#L8), [Bulbizarre (Français)](pokemon/internal/fr/pokemon.go#L8), [Bisasam (Deutsch)](pokemon/internal/de/pokemon.go#L8), [Bulbasaur (Italiano)](pokemon/internal/it/pokemon.go#L8), [이상해씨 (한국어)](pokemon/internal/ko/pokemon.go#L8), [Bulbasaur (Español)](pokemon/internal/es/pokemon.go#L8), [妙蛙种子 (简体中文（官方译名）)](pokemon/internal/zh-CHS/pokemon.go#L8), [妙蛙種子 (繁體中文（官方譯名）)](pokemon/internal/zh-CHT/pokemon.go#L8), [妙蛙种子 (简体中文（中国译名）)](pokemon/internal/zh-CN/pokemon.go#L8), [奇異種子 (繁體中文（香港譯名）)](pokemon/internal/zh-HK/pokemon.go#L8), [妙蛙種子 (繁體中文（台灣譯名）)](pokemon/internal/zh-TW/pokemon.go#L8).
  - For instance, the Pokémon with ID #100 is named: [ビリリダマ (日本語)](pokemon/internal/ja/pokemon.go#L107), [Voltorb (English)](pokemon/internal/en/pokemon.go#L107), [Voltorbe (Français)](pokemon/internal/fr/pokemon.go#L107), [Voltobal (Deutsch)](pokemon/internal/de/pokemon.go#L107), [Voltorb (Italiano)](pokemon/internal/it/pokemon.go#L107), [찌리리공 (한국어)](pokemon/internal/ko/pokemon.go#L107), [Voltorb (Español)](pokemon/internal/es/pokemon.go#L107), [霹雳电球 (简体中文（官方译名）)](pokemon/internal/zh-CHS/pokemon.go#L107), [霹靂電球 (繁體中文（官方譯名）)](pokemon/internal/zh-CHT/pokemon.go#L107), [雷电球 (简体中文（中国译名）)](pokemon/internal/zh-CN/pokemon.go#L107), [霹靂蛋 (繁體中文（香港譯名）)](pokemon/internal/zh-HK/pokemon.go#L107), [雷電球 (繁體中文（台灣譯名）)](pokemon/internal/zh-TW/pokemon.go#L107).
- Package `generic`
  - This package deals with generic collections and common definitions in the project.

[pokemon-go]: http://www.pokemongo.com/
[ios]: https://itunes.apple.com/app/pokemon-go/id1094591345
[android]: https://play.google.com/store/apps/details?id=com.nianticlabs.pokemongo
[niantic]: https://nianticlabs.com/
[cp_hp_calc]: cmd/cp_hp_calc/
[iv_calc]: cmd/iv_calc/
[pokemon]: pokemon/README.md
[type]: type/README.md
