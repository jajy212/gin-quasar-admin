package system

import (
	"gin-quasar-admin/global"
)

type SysRole struct {
	global.GqaModel
	RoleCode string    `json:"roleCode" gorm:"comment:角色编码;not null;uniqueIndex;"`
	RoleName string    `json:"roleName" gorm:"comment:角色名称;not null;unique;"`
	User     []SysUser `json:"user" gorm:"many2many:sys_user_role;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Id;joinReferences:SysUserId;"`
	Menu     []SysMenu `json:"menu" gorm:"many2many:sys_role_menu;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Id;joinReferences:SysMenuId;"`
}

type RequestAddRole struct {
	Status   string `json:"status"`
	Sort     int    `json:"sort"`
	Remark     string `json:"remark"`
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}

type RequestRoleCode struct {
	RoleCode string `json:"roleCode"`
}

type RequestRoleUser struct {
	RoleCode string `json:"roleCode"`
	UserId   uint   `json:"userId"`
}

type RequestRoleUserAdd struct {
	RoleCode string `json:"roleCode"`
	UserId   []uint `json:"userId"`
}

type RequestRoleMenuEdit struct {
	RoleCode string            `json:"roleCode"`
	RoleMenu []RequestRoleMenu `json:"roleMenu"`
}

type RequestRoleMenu struct {
	RoleCode string `json:"roleCode"`
	MenuId   uint   `json:"menuId"`
}

type RequestRoleApiEdit struct {
	RoleCode string      `json:"roleCode"`
	Policy   []SysCasbin `json:"policy"`
}

type ResponseRole struct {
	Role []SysRole `json:"role"`
}
