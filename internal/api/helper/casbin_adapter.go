package helper

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"go.uber.org/zap"
	"subflow-core-go/pkg/ent"
	entrole "subflow-core-go/pkg/ent/role"
	entuser "subflow-core-go/pkg/ent/user"
)

type CasbinAdapter struct {
	client   *ent.Client
	ctx      context.Context
	filtered bool
}

type Option func(a *CasbinAdapter) error

func NewAdapter(client *ent.Client, options ...Option) (*CasbinAdapter, error) {
	a := &CasbinAdapter{
		client: client,
		ctx:    context.Background(),
	}
	for _, option := range options {
		if err := option(a); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (a *CasbinAdapter) LoadPolicy(model model.Model) error {
	users, err := a.client.User.
		Query().
		WithRoles().
		All(a.ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		for _, role := range user.Edges.Roles {
			policyLine := fmt.Sprintf("g,%d,%s", user.ID, role.Name)
			zap.S().Debugw(
				"LoadPolicy",
				"policy", policyLine,
			)
			err := persist.LoadPolicyLine(policyLine, model)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *CasbinAdapter) SavePolicy(model model.Model) error {
	return nil
}

func (a *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	if ptype != "g" || len(rule) < 2 {
		return errors.New("invalid parameters")
	}
	userID, _ := strconv.Atoi(rule[0])
	roleName := rule[1]

	user, err := a.client.User.
		Query().
		Where(entuser.IDEQ(userID)).
		Only(a.ctx)
	if err != nil {
		return fmt.Errorf("the uid %d does not exist", userID)
	}

	role, err := a.client.Role.
		Query().
		Where(entrole.NameEQ(roleName)).
		Only(a.ctx)
	if err != nil {
		return fmt.Errorf("the role %s does not exist", roleName)
	}

	_, err = a.client.User.
		UpdateOne(user).
		AddRoles(role).
		Save(a.ctx)
	if err != nil {
		return fmt.Errorf("failed to add role %s to uid %d", roleName, userID)
	}
	return nil
}

// RemovePolicy 从持久层删除单条policy规则
func (a *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	if ptype != "g" || len(rule) < 2 {
		return errors.New("invalid parameters")
	}
	userID, _ := strconv.Atoi(rule[0])
	roleName := rule[1]

	// 获取用户和角色实体
	user, err := a.client.User.
		Query().
		Where(entuser.IDEQ(userID)).
		Only(a.ctx)
	if err != nil {
		return fmt.Errorf("the uid %d does not exist", userID)
	}
	role, err := a.client.Role.
		Query().
		Where(entrole.NameEQ(roleName)).
		Only(a.ctx)
	if err != nil {
		return fmt.Errorf("the role %s does not exist", roleName)
	}

	// 从用户中移除角色
	_, err = a.client.User.
		UpdateOne(user).
		RemoveRoles(role).
		Save(a.ctx)
	if err != nil {
		return fmt.Errorf("failed to remove role %s from uid %d", roleName, userID)
	}

	return nil
}

// RemoveFilteredPolicy 从持久层删除符合筛选条件的policy规则
func (a *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
