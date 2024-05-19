package models

// Struct storing the data from assembly site by selecting state only.
type AssemblyStateData struct {
	State                   string `bson:"State"`
	ElectionType            string `bson:"ElectionType"`
	Year                    string `bson:"Year"`
	Party                   string `bson:"Party"`
	AssemblySegmentsLeading string `bson:"AssemblySegmentsLeading"`
	VoteSharePercentage     string `bson:"VoteSharePercentage"`
}

// Struct storing the data from assembly site by selecting state and assembly constituency.
type AssemblyStateAssemblyConstituencyData struct {
	State                string `bson:"State"`
	AssemblyConstituency string `bson:"AssemblyConstituency"`
	ElectionType         string `bson:"ElectionType"`
	Year                 string `bson:"Year"`
	Candidate            string `bson:"Candidate"`
	Party                string `bson:"Party"`
	VoteNumber           string `bson:"VoteNumber"`
	VotePercentage       string `bson:"VotePercentage"`
}

// Struct storing the data from assembly site by selecting state, assembly constituency and pooling booth.
type AssemblyStateAssemblyConstituencyPoolingBoothData struct {
	State                string `bson:"State"`
	AssemblyConstituency string `bson:"AssemblyConstituency"`
	PoolingBooth         string `bson:"PoolingBooth"`
	ElectionType         string `bson:"ElectionType"`
	Year                 string `bson:"Year"`
	Candidate            string `bson:"Candidate"`
	Party                string `bson:"Party"`
	VotePolled           string `bson:"VotePolled"`
	VotePercentage       string `bson:"VotePercentage"`
}
