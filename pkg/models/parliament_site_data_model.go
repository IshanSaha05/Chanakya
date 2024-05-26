package models

type ParliamentStateConstituencyName struct {
	ParliamentStateConstituencyMap map[string]string `json:"PARLIAMENT_NAME_DATA,omitempty"`
}

type ParliamentStateConstituencyDataRaw struct {
	Candidate      string  `json:"candidateName"`
	Party          string  `json:"partyName"`
	NumberOfVotes  int     `json:"numberOfVotes"`
	VotePercentage float64 `json:"votePercentage"`
}

type ParliamentListofPartyData struct {
	RawData []ParliamentStateConstituencyDataRaw `json:"listOfPartyData"`
}

type ParliamentElectionData map[string]ParliamentListofPartyData

type ParliamentStateConstituencyResponseData struct {
	ParliamentData struct {
		ParliamentDataRaw ParliamentElectionData `json:"parliamentLevelDataForElections"`
	} `json:"ELECTION_DATA"`
}

type ParliamentStateConstituencyData struct {
	State          string  `bson:"state"`
	Constituency   string  `bson:"constituency"`
	ElectionType   string  `bson:"election_type"`
	Candidate      string  `bson:"candidate"`
	Party          string  `bson:"party"`
	NumberOfVotes  int     `bson:"number_of_votes"`
	VotePercentage float64 `bson:"vote_percentage"`
}
