package constants

type Role struct {
	ID   int
	Name string
	Desc string
}

const (
	RoleNameSuperuser           = "superuser"
	RoleNameAdmin               = "admin"
	RoleNameLeader              = "leader"
	RoleNameViceLeader          = "vice_leader"
	RoleNameTranslator          = "translator"
	RoleNameTimelineMaker       = "timeline_maker"
	RoleNameProofreader         = "proofreader"
	RoleNamePostProductionStaff = "post_production_staff"
	RoleNameCompressor          = "compressor"
	RoleNameSupervisor          = "supervisor"
)

var (
	roleSuperuser = Role{
		ID:   1000,
		Name: RoleNameSuperuser,
		Desc: "超级管理员",
	}
	roleAdmin = Role{
		ID:   1001,
		Name: RoleNameAdmin,
		Desc: "管理员",
	}
	roleLeader = Role{
		ID:   2000,
		Name: RoleNameLeader,
		Desc: "组长",
	}
	roleViceLeader = Role{
		ID:   2001,
		Name: RoleNameViceLeader,
		Desc: "副组长",
	}
	roleTranslator = Role{
		ID:   3000,
		Name: RoleNameTranslator,
		Desc: "翻译",
	}
	roleTimelineMaker = Role{
		ID:   3001,
		Name: RoleNameTimelineMaker,
		Desc: "时轴",
	}
	roleProofreader = Role{
		ID:   3002,
		Name: RoleNameProofreader,
		Desc: "校对",
	}
	rolePostProductionStaff = Role{
		ID:   3003,
		Name: RoleNamePostProductionStaff,
		Desc: "后期",
	}
	roleCompressor = Role{
		ID:   3004,
		Name: RoleNameCompressor,
		Desc: "压制",
	}
	roleSupervisor = Role{
		ID:   3005,
		Name: RoleNameSupervisor,
		Desc: "监制",
	}
)

func GetAllRoles() []Role {
	return []Role{
		roleAdmin,
		roleSuperuser,
		roleLeader,
		roleViceLeader,
		roleTranslator,
		roleTimelineMaker,
		roleProofreader,
		rolePostProductionStaff,
		roleCompressor,
		roleSupervisor,
	}
}
