// This module implements a Proxy of the ScheduableService defined and
// implemented inside the service package. The Proxy forwards the requests in the
// form of command.Command instances to the server.
package proxy_service

import (
	"brevity/client"
	"brevity/command"
	"brevity/task"
	"time"
)

type ProxyScheduableService struct {
	client *client.Client
	commandFactory command.AbstractCommandFactory
}

// Create a new instance of ProxyScheduableService. Gets pointer to client.Client
// as argument.
func NewProxyScheduableService(client *client.Client) *ProxyScheduableService {
	return &ProxyScheduableService{client: client,
		commandFactory: command.Factory{}}
}

// Add the container.
func (p ProxyScheduableService) AddContainer(startYear int, startMonth time.Month, startDay, startHour, startMin int,
	endYear int, endMonth time.Month, endDay, endHour, endMin int) {
	startDate := time.Date(startYear, startMonth, startDay, startHour, startMin, 0, 0, time.Local)
	endDate := time.Date(endYear, endMonth, endDay, endHour, endMin,0 , 0, time.Local)

	com := p.commandFactory.AddContainer(*task.NewContainer(0, startDate, endDate))
	_, _ = p.client.SendCommand(com)
}

// Remove a container by id.
func (p ProxyScheduableService) RemoveContainer(id uint64) bool {
	com := p.commandFactory.RemoveContainer(id)
	resp, _ := p.client.SendCommand(com)
	return resp.Data.(bool)
}

// Get a list of all the containers.
func (p ProxyScheduableService) GetAllContainers() []task.Container {
	com := p.commandFactory.GetAllContainers()
	resp, _ := p.client.SendCommand(com)
	return resp.Data.([]task.Container)
}
