# Pokémon GO

Pokémon GO is a free-to-play location-based, augmented reality game developed and published by Niantic for iOS and Android devices. It was initially released in selected countries in July 2016. In the game, players use the smart device's GPS and camera to capture, battle, and train virtual creatures, called Pokémon, who appear on the screen as if they were in the same real-world location as the player.

This project aims primarily at providing the value tables used in the game, and an API the retrieve them. Current features include:
- Cheatsheet (Not Currently Public):
  - A table containing all Pokémons and their information
  - Will be public soon
- CP / HP Calculator:
  - Given the "IVs (Individual Values)" and "Level" of a Pokémon, calculate its "CP (Combat Power)" and "HP (Hit Points). This is basically used to test the "IV Calculator"
  - The source code is available [here](cmd/cp_hp_calc/)
- IV Calculator:
  - Given the "CP (Combat Power)", "HP (Hit Points)", "Stardust to Power Up", and "Candy to Power Up" of a Pokémon, calculate its possible "IVs (Individual Values)" and "Level" combinations. This is probably the most wanted feature.
  - The source code is available [here](cmd/iv_calc/)
- Multi-lingual Ready:
  - The library Supported languages include: Japanese (日本語), English, French (Français), German (Deutsch), Italian (Italiano), Korean (한국어), Spanish (Español), Simplified Chinese (简体中文（官方译名）), Traditional Chinese (繁體中文（官方譯名）), PRC Chinese (简体中文（中国译名）), Hong Kong Chinese (繁體中文（香港譯名）), Taiwan Chinese (繁體中文（台灣譯名）)
