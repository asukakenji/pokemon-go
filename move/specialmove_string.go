// Code generated by "stringer -type=SpecialMove"; DO NOT EDIT

package move

import "fmt"

const _SpecialMove_name = "AerialAceAirCutterAncientPowerAquaJetAquaTailBlizzardBodySlamBoneClubBrickBreakBrineBubbleBeamBugBuzzBulldozeCrossChopCrossPoisonDarkPulseDazzlingGleamDigDisarmingVoiceDischargeDragonClawDragonPulseDrainingKissDrillPeckDrillRunEarthquakeFireBlastFirePunchFlameBurstFlameChargeFlameWheelFlamethrowerFlashCannonGunkShotHeatWaveHornAttackHurricaneHydroPumpHyperBeamHyperFangIceBeamIcePunchIcyWindIronHeadLeafBladeLowSweepMagnetBombMegahornMoonblastMudBombNightSlashOminousWindPetalBlizzardPlayRoughPoisonFangPowerGemPowerWhipPsybeamPsychicPsyshockRockSlideRockTombScaldSeedBombShadowBallSignalBeamSludgeSludgeBombSludgeWaveSolarBeamStompStoneEdgeStruggleSubmissionSwiftThunderThunderPunchThunderboltTwisterViceGripWaterPulseWrapX_Scissor"

var _SpecialMove_index = [...]uint16{0, 9, 18, 30, 37, 45, 53, 61, 69, 79, 84, 94, 101, 109, 118, 129, 138, 151, 154, 168, 177, 187, 198, 210, 219, 227, 237, 246, 255, 265, 276, 286, 298, 309, 317, 325, 335, 344, 353, 362, 371, 378, 386, 393, 401, 410, 418, 428, 436, 445, 452, 462, 473, 486, 495, 505, 513, 522, 529, 536, 544, 553, 561, 566, 574, 584, 594, 600, 610, 620, 629, 634, 643, 651, 661, 666, 673, 685, 696, 703, 711, 721, 725, 734}

func (i SpecialMove) String() string {
	i -= 41
	if i < 0 || i >= SpecialMove(len(_SpecialMove_index)-1) {
		return fmt.Sprintf("SpecialMove(%d)", i+41)
	}
	return _SpecialMove_name[_SpecialMove_index[i]:_SpecialMove_index[i+1]]
}
