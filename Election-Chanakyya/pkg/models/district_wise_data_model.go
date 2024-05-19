package models

// Struct storing the data from district wise site by selecting state and district.
type DistrictWiseStateDistrict struct {
	State                   string `bson:"State"`
	District                string `bson:"District"`
	ElectionType            string `bson:"ElectionType"`
	Year                    string `bson:"Year"`
	Party                   string `bson:"Party"`
	AssemblySegmentsLeading string `bson:"AssemblySegmentsLeading"`
	VoteNumber              string `bson:"VoteNumber"`
	VotePercentage          string `bson:"VotePercentage"`
}
