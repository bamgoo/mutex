package mutex

import (
	"github.com/infrago/base"
	"github.com/infrago/infra"
)

func (m *Module) Ready() bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.started && len(m.instances) > 0
}

func (m *Module) Health() infra.ModuleHealth {
	m.mutex.RLock()
	started := m.started
	connections := len(m.instances)
	m.mutex.RUnlock()
	return infra.NewModuleHealth("mutex", started && connections > 0, nil, base.Map{"connections": connections})
}

func (m *Module) MonitorStats() infra.ModuleStats {
	m.mutex.RLock()
	started := m.started
	connections := len(m.instances)
	m.mutex.RUnlock()
	return infra.NewModuleStats("mutex", started && connections > 0, base.Map{"connections": connections})
}
