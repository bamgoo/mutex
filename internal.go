package mutex

import "time"

func (m *Module) getInst(conn, key string) (*Instance, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if conn == "" {
		if m.ring == nil {
			return nil, ErrNotReady
		}
		conn = m.ring.Locate(key)
	}

	if conn == "" {
		return nil, ErrInvalidConnection
	}

	if inst, ok := m.instances[conn]; ok {
		return inst, nil
	}

	return nil, ErrInvalidConnection
}

// LockOn locks to a specific connection.
func (m *Module) LockOn(conn string, key string, expires ...time.Duration) error {
	inst, err := m.getInst(conn, key)
	if err != nil {
		return err
	}

	expire := inst.Config.Expire
	if len(expires) > 0 {
		expire = expires[0]
	}

	realKey := inst.Config.Prefix + key
	return inst.conn.Lock(realKey, expire)
}

// Lock locks with auto-selected connection.
func (m *Module) Lock(key string, expires ...time.Duration) error {
	return m.LockOn("", key, expires...)
}

// UnlockOn unlocks on a specific connection.
func (m *Module) UnlockOn(conn, key string) error {
	inst, err := m.getInst(conn, key)
	if err != nil {
		return err
	}

	realKey := inst.Config.Prefix + key
	return inst.conn.Unlock(realKey)
}

// Unlock unlocks with auto-selected connection.
func (m *Module) Unlock(key string) error {
	return m.UnlockOn("", key)
}
