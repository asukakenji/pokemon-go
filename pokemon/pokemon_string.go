// Code generated by "stringer -type=Pokemon"; DO NOT EDIT

package pokemon

import "fmt"

const _Pokemon_name = "BulbasaurIvysaurVenusaurCharmanderCharmeleonCharizardSquirtleWartortleBlastoiseCaterpieMetapodButterfreeWeedleKakunaBeedrillPidgeyPidgeottoPidgeotRattataRaticateSpearowFearowEkansArbokPikachuRaichuSandshrewSandslashNidoran_femaleNidorinaNidoqueenNidoran_maleNidorinoNidokingClefairyClefableVulpixNinetalesJigglypuffWigglytuffZubatGolbatOddishGloomVileplumeParasParasectVenonatVenomothDiglettDugtrioMeowthPersianPsyduckGolduckMankeyPrimeapeGrowlitheArcaninePoliwagPoliwhirlPoliwrathAbraKadabraAlakazamMachopMachokeMachampBellsproutWeepinbellVictreebelTentacoolTentacruelGeodudeGravelerGolemPonytaRapidashSlowpokeSlowbroMagnemiteMagnetonFarfetch_dDoduoDodrioSeelDewgongGrimerMukShellderCloysterGastlyHaunterGengarOnixDrowzeeHypnoKrabbyKinglerVoltorbElectrodeExeggcuteExeggutorCuboneMarowakHitmonleeHitmonchanLickitungKoffingWeezingRhyhornRhydonChanseyTangelaKangaskhanHorseaSeadraGoldeenSeakingStaryuStarmieMr_MimeScytherJynxElectabuzzMagmarPinsirTaurosMagikarpGyaradosLaprasDittoEeveeVaporeonJolteonFlareonPorygonOmanyteOmastarKabutoKabutopsAerodactylSnorlaxArticunoZapdosMoltresDratiniDragonairDragoniteMewtwoMew"

var _Pokemon_index = [...]uint16{0, 9, 16, 24, 34, 44, 53, 61, 70, 79, 87, 94, 104, 110, 116, 124, 130, 139, 146, 153, 161, 168, 174, 179, 184, 191, 197, 206, 215, 229, 237, 246, 258, 266, 274, 282, 290, 296, 305, 315, 325, 330, 336, 342, 347, 356, 361, 369, 376, 384, 391, 398, 404, 411, 418, 425, 431, 439, 448, 456, 463, 472, 481, 485, 492, 500, 506, 513, 520, 530, 540, 550, 559, 569, 576, 584, 589, 595, 603, 611, 618, 627, 635, 645, 650, 656, 660, 667, 673, 676, 684, 692, 698, 705, 711, 715, 722, 727, 733, 740, 747, 756, 765, 774, 780, 787, 796, 806, 815, 822, 829, 836, 842, 849, 856, 866, 872, 878, 885, 892, 898, 905, 912, 919, 923, 933, 939, 945, 951, 959, 967, 973, 978, 983, 991, 998, 1005, 1012, 1019, 1026, 1032, 1040, 1050, 1057, 1065, 1071, 1078, 1085, 1094, 1103, 1109, 1112}

func (i Pokemon) String() string {
	i -= 1
	if i < 0 || i >= Pokemon(len(_Pokemon_index)-1) {
		return fmt.Sprintf("Pokemon(%d)", i+1)
	}
	return _Pokemon_name[_Pokemon_index[i]:_Pokemon_index[i+1]]
}
