// Code generated by "stringer -type=Pokemon"; DO NOT EDIT

package pokemon

import "fmt"

const _Pokemon_name = "NoneBulbasaurIvysaurVenusaurCharmanderCharmeleonCharizardSquirtleWartortleBlastoiseCaterpieMetapodButterfreeWeedleKakunaBeedrillPidgeyPidgeottoPidgeotRattataRaticateSpearowFearowEkansArbokPikachuRaichuSandshrewSandslashNidoran_femaleNidorinaNidoqueenNidoran_maleNidorinoNidokingClefairyClefableVulpixNinetalesJigglypuffWigglytuffZubatGolbatOddishGloomVileplumeParasParasectVenonatVenomothDiglettDugtrioMeowthPersianPsyduckGolduckMankeyPrimeapeGrowlitheArcaninePoliwagPoliwhirlPoliwrathAbraKadabraAlakazamMachopMachokeMachampBellsproutWeepinbellVictreebelTentacoolTentacruelGeodudeGravelerGolemPonytaRapidashSlowpokeSlowbroMagnemiteMagnetonFarfetch_dDoduoDodrioSeelDewgongGrimerMukShellderCloysterGastlyHaunterGengarOnixDrowzeeHypnoKrabbyKinglerVoltorbElectrodeExeggcuteExeggutorCuboneMarowakHitmonleeHitmonchanLickitungKoffingWeezingRhyhornRhydonChanseyTangelaKangaskhanHorseaSeadraGoldeenSeakingStaryuStarmieMr_MimeScytherJynxElectabuzzMagmarPinsirTaurosMagikarpGyaradosLaprasDittoEeveeVaporeonJolteonFlareonPorygonOmanyteOmastarKabutoKabutopsAerodactylSnorlaxArticunoZapdosMoltresDratiniDragonairDragoniteMewtwoMew"

var _Pokemon_index = [...]uint16{0, 4, 13, 20, 28, 38, 48, 57, 65, 74, 83, 91, 98, 108, 114, 120, 128, 134, 143, 150, 157, 165, 172, 178, 183, 188, 195, 201, 210, 219, 233, 241, 250, 262, 270, 278, 286, 294, 300, 309, 319, 329, 334, 340, 346, 351, 360, 365, 373, 380, 388, 395, 402, 408, 415, 422, 429, 435, 443, 452, 460, 467, 476, 485, 489, 496, 504, 510, 517, 524, 534, 544, 554, 563, 573, 580, 588, 593, 599, 607, 615, 622, 631, 639, 649, 654, 660, 664, 671, 677, 680, 688, 696, 702, 709, 715, 719, 726, 731, 737, 744, 751, 760, 769, 778, 784, 791, 800, 810, 819, 826, 833, 840, 846, 853, 860, 870, 876, 882, 889, 896, 902, 909, 916, 923, 927, 937, 943, 949, 955, 963, 971, 977, 982, 987, 995, 1002, 1009, 1016, 1023, 1030, 1036, 1044, 1054, 1061, 1069, 1075, 1082, 1089, 1098, 1107, 1113, 1116}

func (i Pokemon) String() string {
	if i < 0 || i >= Pokemon(len(_Pokemon_index)-1) {
		return fmt.Sprintf("Pokemon(%d)", i)
	}
	return _Pokemon_name[_Pokemon_index[i]:_Pokemon_index[i+1]]
}
