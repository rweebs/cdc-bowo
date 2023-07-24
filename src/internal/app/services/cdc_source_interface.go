// \ brief Test interface for different types of service. \ par Function Description Test interface for different types of service
package services

type CDCSourceServicesInterface interface {
	ExecuteDDLChange()
	StartService()
	StopService()
	CheckReplicationCatchUp() (bool, error)
}
