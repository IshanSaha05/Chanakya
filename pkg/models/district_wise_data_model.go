package models

type DistrictStateDistrictName struct {
	DistrictStateDistrictMap map[string]string `json:"DISTRICT_NAME_DATA,omitempty"`
}

type DistrictStateDistrictDataRaw struct {
	PartyName               string  `json:"partyName"`
	NumberOfVotes           int     `json:"numberOfVotes"`
	VotePercentage          float64 `json:"votePercentage"`
	AssemblySegmentsLeading int     `json:"numberOfAssemblySegmentWin"`
}

type DistrictListofPartyData struct {
	RawData []DistrictStateDistrictDataRaw `json:"listOfPartyData"`
}

type DistrictElectionData map[string]DistrictListofPartyData

type DistrictStateDistrictResponseData struct {
	DistrictData struct {
		DistrictDataRaw DistrictElectionData `json:"districtLevelDataForElections"`
	} `json:"ELECTION_DATA"`
}

type DistrictStateDistrictData struct {
	State                   string  `bson:"state"`
	District                string  `bson:"district"`
	ElectionType            string  `bson:"election_type"`
	Party                   string  `bson:"party"`
	AssemblySegmentsLeading int     `bson:"assembly_segments_leading"`
	NumberOfVotes           int     `bson:"number_of_votes"`
	VotePercentage          float64 `bson:"vote_percentage"`
}
