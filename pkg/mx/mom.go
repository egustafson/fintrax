package mx

type MOMgr interface {
	Root() MO
}

type BaseMOManager struct {
	root *MO
}

// global Managed Object Manager (MOM)
var mom MOMgr

func MOM() MOMgr {
	return mom
}

func InitMOM() MOMgr {
	root := NewBaseMO()
	root.SetState("type-id", "root-mo")
	mom = &BaseMOManager{
		root: &root,
	}
	return mom
}

func (m *BaseMOManager) Root() MO {
	return *m.root
}
