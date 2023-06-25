package andMapMemoriesDB

const pinsVisitsTableName = "pins_visits";

type PinVisit struct {
	ID int64
	PinID int64
	VisitDate int64
	photosDir string
}