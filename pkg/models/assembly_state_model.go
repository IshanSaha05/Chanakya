package models

type AssemblyStateDataRaw struct {
	Party                   string  `json:"partyName"`
	AssemblySegmentsLeading int     `json:"numberOfSeatLeading"`
	VotePercentage          float64 `json:"votePercentage"`
}

type AssemblyStateDataMap map[string][]AssemblyStateDataRaw

type AssemblyStateResponseData struct {
	StateLevelData struct {
		StateDataRaw AssemblyStateDataMap `json:"stateLevelData"`
	} `json:"ELECTION_DATA"`
}

type AssemblyStateData struct {
	State                   string  `bson:"state"`
	ElectionType            string  `bson:"election_type"`
	Party                   string  `bson:"party"`
	AssemblySegmentsLeading int     `bson:"assembly_segments_leading"`
	VotePercentage          float64 `bson:"vote_percentage"`
}

type AssemblyStateConstituencyNames struct {
	ConstituencyName map[string]string `json:"ASSEMBLY_NAME_DATA"`
}

type AssemblyStateConstituencyDataRaw struct {
	Party          string  `json:"partyName"`
	Candidate      string  `json:"candidateName"`
	NumberOfVotes  int     `json:"numberOfVotes"`
	VotePercentage float64 `json:"votePercentage"`
}

type AssemblyListofPartyData struct {
	RawData []AssemblyStateConstituencyDataRaw `json:"listOfPartyData"`
}

type AssemblyElectionData map[string]AssemblyListofPartyData

type AssemblyStateConstituencyResponseData struct {
	AssemblyData struct {
		AssemblyDataRaw AssemblyElectionData `json:"assemblyLevelDataForElections"`
	} `json:"ELECTION_DATA"`
}

type AssemblyStateConstituencyData struct {
	State          string  `bson:"state"`
	Constituency   string  `bson:"constituency"`
	ElectionType   string  `bson:"election_type"`
	Candidate      string  `bson:"candidate"`
	Party          string  `bson:"party"`
	NumberOfVotes  int     `bson:"number_of_votes"`
	VotePercentage float64 `bson:"vote_percentage"`
}
