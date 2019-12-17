package main

type airport struct {
	code string
	name string
}

func loadAirports() map[string]*airport {
	result := make(map[string]*airport, 4)
	result["ATL"] = &airport{code: "ATL", name: "Atlanta"}
	result["ORD"] = &airport{code: "ORD", name: "Chicago"}
	result["LHR"] = &airport{code: "LHR", name: "London"}
	result["HKG"] = &airport{code: "HKG", name: "Hong Kong"}
	return result
}