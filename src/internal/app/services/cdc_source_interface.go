package services

type CDCSourceServicesInterface interface {
	ExecuteDDLChange()
	StartService()
	StopService()
	CheckReplicationCatchUp() (bool, error)
}
