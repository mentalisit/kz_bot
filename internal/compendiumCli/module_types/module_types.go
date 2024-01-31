package module_types

var ModuleMap = map[string]int{
	"RedStarScanner":       701,
	"ShipmentRelay":        702,
	"AllianceLevel":        801,
	"Transport":            103,
	"Miner":                102,
	"Battleship":           101,
	"TransportCapacity":    401,
	"ShipmentComputer":     402,
	"RemoteRepair":         413,
	"Rush":                 404,
	"Stealth":              608,
	"TradeBurst":           405,
	"ShipmentDrone":        406,
	"RedStarExtender":      603,
	"RelicDrone":           412,
	"Dispatch":             411,
	"CargoRocket":          414,
	"MiningBoost":          501,
	"HydroReplicator":      511,
	"ArtifactBoost":        512,
	"MassMining":           504,
	"Genesis":              508,
	"Enrich":               503,
	"Crunch":               507,
	"HydrogenUpload":       505,
	"HydroRocket":          510,
	"BlastDrone":           513,
	"Laser":                203,
	"MassBattery":          204,
	"Battery":              202,
	"DualLaser":            205,
	"Barrage":              206,
	"DartLauncher":         207,
	"ChainRay":             208,
	"PlayerRocketLauncher": 209,
	"Pulse":                210,
	"AlphaShield":          301,
	"ImpulseShield":        302,
	"PassiveShield":        303,
	"OmegaShield":          304,
	"BlastShield":          306,
	"MirrorShield":         305,
	"AreaShield":           307,
	"MotionShield":         308,
	"EMP":                  601,
	"Solitude":             625,
	"Fortify":              609,
	"Teleport":             602,
	"DamageAmplifier":      626,
	"Destiny":              614,
	"Barrier":              615,
	"Vengeance":            616,
	"DeltaRocket":          617,
	"Leap":                 618,
	"Bond":                 619,
	"Suspend":              622,
	"OmegaRocket":          621,
	"RemoteBomb":           623,
	"DecoyDrone":           901,
	"RepairDrone":          902,
	"RocketDrone":          904,
	"LaserTurret":          624,
	"ChainRayDrone":        905,
	"DeltaDrones":          906,
	"DroneSquad":           907,
	"repair":               604, // Deprecated = устарело
	"warp":                 605, // Deprecated
	"unity":                606, // Deprecated
	"sanctuary":            607, // Deprecated
	"impulse":              610, // Deprecated
	"rocket":               611, // deprecated
	"salvage":              612, // deprecated
	"suppress":             613, // Deprecated
	"alphadrone":           620, // Deprecated
	"hydrobay":             502, // Deprecated
	"miningunity":          506, // Deprecated
	"minedrone":            509, // Deprecated
	"tradeboost":           403, // Deprecated
	"offload":              407, // deprecated
	"beam":                 408, // deprecated
	"entrust":              409, // deprecated
	"recall":               410, // deprecated

}

var invertedData map[int]string
var invertedDataInitialized = false

func getTechIndex(tech string) int {
	if value, ok := ModuleMap[tech]; ok {
		return value
	}
	return 0
}

func checkInvert() {
	if !invertedDataInitialized {
		invertedData = make(map[int]string)
		for key, value := range ModuleMap {
			invertedData[value] = key
		}
		invertedDataInitialized = true
	}
}

func GetTechFromIndex(index int) string {
	checkInvert()
	if value, ok := invertedData[index]; ok {
		return value
	}
	return ""
}

func getTypes() map[string]int {
	return ModuleMap
}

func getInvertedTypes() map[int]string {
	checkInvert()
	return invertedData
}
