package models

// Struct storing the data from parliament site by selecting state and parliament constituency.
type ParliamentStateParliamentConstituencyData struct {
	State                  string `bson:"State"`
	ParliamentConstituency string `bson:"ParliamentConstituency"`
	ElectionType           string `bson:"ElectionType"`
	Year                   string `bson:"Year"`
	Candidate              string `bson:"Candidate"`
	Party                  string `bson:"Party"`
	VoteNumber             string `bson:"VoteNumber"`
	VotePercentage         string `bson:"VotePercentage"`
}
