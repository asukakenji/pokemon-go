package pokemon

type levelItem struct {
	stardustToPowerUp     int16
	candyToPowerUp        int8
	combatPowerMultiplier float64
}

var levelTable = [...]levelItem{
	{-1, -1, 0.0}, {-1, -1, 0.0}, // 0.0, 0.5
	{200, 1, 0.094}, {200, 1, 0.094}, {200, 1, 0.166}, {200, 1, 0.166}, // 1.0, 1.5, 2.0, 2.5
	{400, 1, 0.216}, {400, 1, 0.216}, {400, 1, 0.256}, {400, 1, 0.256}, // 3.0, 3.5, 4.0, 4.5
	{600, 1, 0.290}, {600, 1, 0.290}, {600, 1, 0.321}, {600, 1, 0.321}, // 5.0, 5.5, 6.0, 6.5
	{800, 1, 0.349}, {800, 1, 0.349}, {800, 1, 0.375}, {800, 1, 0.375}, // 7.0, 7.5, 8.0, 8.5
	{1000, 1, 0.400}, {1000, 1, 0.400}, {1000, 1, 0.423}, {1000, 1, 0.423}, // 9.0, 9.5, 10.0, 10.5
	{1300, 2, 0.443}, {1300, 2, 0.443}, {1300, 2, 0.463}, {1300, 2, 0.463}, // 11.0, 11.5, 12.0, 12.5
	{1600, 2, 0.482}, {1600, 2, 0.482}, {1600, 2, 0.500}, {1600, 2, 0.500}, // 13.0, 13.5, 14.0, 14.5
	{1900, 2, 0.517}, {1900, 2, 0.517}, {1900, 2, 0.534}, {1900, 2, 0.534}, // 15.0, 15.5, 16.0, 16.5
	{2200, 2, 0.551}, {2200, 2, 0.551}, {2200, 2, 0.567}, {2200, 2, 0.567}, // 17.0, 17.5, 18.0, 18.5
	{2500, 2, 0.582}, {2500, 2, 0.582}, {2500, 2, 0.597}, {2500, 2, 0.597}, // 19.0, 19.5, 20.0, 20.5
	{3000, 3, 0.612}, {3000, 3, 0.612}, {3000, 3, 0.627}, {3000, 3, 0.627}, // 21.0, 21.5, 22.0, 22.5
	{3500, 3, 0.641}, {3500, 3, 0.641}, {3500, 3, 0.654}, {3500, 3, 0.654}, // 23.0, 23.5, 24.0, 24.5
	{4000, 3, 0.668}, {4000, 3, 0.668}, {4000, 4, 0.681}, {4000, 4, 0.681}, // 25.0, 25.5, 26.0, 26.5
	{4500, 4, 0.694}, {4500, 4, 0.694}, {4500, 4, 0.707}, {4500, 4, 0.707}, // 27.0, 27.5, 28.0, 28.5
	{5000, 4, 0.719}, {5000, 4, 0.719}, {5000, 4, 0.732}, {5000, 4, 0.732}, // 29.0, 29.5, 30.0, 30.5
	{6000, 6, 0.738}, {6000, 6, 0.738}, {6000, 6, 0.744}, {6000, 6, 0.744}, // 31.0, 31.5, 32.0, 32.5
	{7000, 8, 0.750}, {7000, 8, 0.750}, {7000, 8, 0.756}, {7000, 8, 0.756}, // 33.0, 33.5, 34.0, 34.5
	{8000, 10, 0.762}, {8000, 10, 0.762}, {8000, 10, 0.767}, {8000, 10, 0.767}, // 35.0, 35.5, 36.0, 36.5
	{9000, 12, 0.773}, {9000, 12, 0.773}, {9000, 12, 0.779}, {9000, 12, 0.779}, // 37.0, 37.5, 38.0, 38.5
	{10000, 15, 0.785}, {10000, 15, 0.785}, {10000, 15, 0.79}, // 39.0, 39.5, 40.0
}
