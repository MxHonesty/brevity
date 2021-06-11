// This module implements a Proxy of the ScheduableService defined and
// implemented inside the service package. The Proxy forwards the requests in the
// form of command.Command instances to the server.
package proxy_service

import (
	"brevity/task"
	"time"
)

// TODO: implement
type ProxyScheduableService struct {
	
}

func (p ProxyScheduableService) AddContainer(startYear int, startMonth time.Month, startDay, startHour, startMin int,
	endYear int, endMonth time.Month, endDay, endHour, endMin int) {
	panic("implement me")
}

func (p ProxyScheduableService) RemoveContainer(id uint64) bool {
	panic("implement me")
}

func (p ProxyScheduableService) GetAllContainers() []task.Container {
	panic("implement me")
}
