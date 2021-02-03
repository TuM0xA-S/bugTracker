package types

//BugData is a base type to store data from various sources
type BugData struct {
	CVE string `bson:"CVE"`
	Packages []string `bson:"Packages"`
	Source string `bson:"Source"`
	Data interface{} `bson:"Data"`
}
