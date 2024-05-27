// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: admin/team/team.proto

package team

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTeamAddTeamMember = "/api.admin.team.Team/AddTeamMember"
const OperationTeamCreateTeam = "/api.admin.team.Team/CreateTeam"
const OperationTeamGetTeam = "/api.admin.team.Team/GetTeam"
const OperationTeamListTeam = "/api.admin.team.Team/ListTeam"
const OperationTeamListTeamMember = "/api.admin.team.Team/ListTeamMember"
const OperationTeamMyTeam = "/api.admin.team.Team/MyTeam"
const OperationTeamRemoveTeamAdmin = "/api.admin.team.Team/RemoveTeamAdmin"
const OperationTeamRemoveTeamMember = "/api.admin.team.Team/RemoveTeamMember"
const OperationTeamSetMemberRole = "/api.admin.team.Team/SetMemberRole"
const OperationTeamSetTeamAdmin = "/api.admin.team.Team/SetTeamAdmin"
const OperationTeamTransferTeamLeader = "/api.admin.team.Team/TransferTeamLeader"
const OperationTeamUpdateTeam = "/api.admin.team.Team/UpdateTeam"
const OperationTeamUpdateTeamStatus = "/api.admin.team.Team/UpdateTeamStatus"

type TeamHTTPServer interface {
	// AddTeamMember 添加团队成员
	AddTeamMember(context.Context, *AddTeamMemberRequest) (*AddTeamMemberReply, error)
	// CreateTeam 创建团队
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamReply, error)
	// GetTeam 获取团队详情
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamReply, error)
	// ListTeam 获取团队列表
	ListTeam(context.Context, *ListTeamRequest) (*ListTeamReply, error)
	// ListTeamMember 获取团队成员列表
	ListTeamMember(context.Context, *ListTeamMemberRequest) (*ListTeamMemberReply, error)
	// MyTeam 我的团队， 查看当前用户的团队列表
	MyTeam(context.Context, *MyTeamRequest) (*MyTeamReply, error)
	// RemoveTeamAdmin 移除团队管理员
	RemoveTeamAdmin(context.Context, *RemoveTeamAdminRequest) (*RemoveTeamAdminReply, error)
	// RemoveTeamMember 移除团队成员
	RemoveTeamMember(context.Context, *RemoveTeamMemberRequest) (*RemoveTeamMemberReply, error)
	// SetMemberRole 设置成员角色
	SetMemberRole(context.Context, *SetMemberRoleRequest) (*SetMemberRoleReply, error)
	// SetTeamAdmin 设置成管理员
	SetTeamAdmin(context.Context, *SetTeamAdminRequest) (*SetTeamAdminReply, error)
	// TransferTeamLeader 移交超级管理员
	TransferTeamLeader(context.Context, *TransferTeamLeaderRequest) (*TransferTeamLeaderReply, error)
	// UpdateTeam 更新团队
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamReply, error)
	// UpdateTeamStatus 修改团队状态
	UpdateTeamStatus(context.Context, *UpdateTeamStatusRequest) (*UpdateTeamStatusReply, error)
}

func RegisterTeamHTTPServer(s *http.Server, srv TeamHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/team", _Team_CreateTeam0_HTTP_Handler(srv))
	r.PUT("/v1/team/{id}", _Team_UpdateTeam0_HTTP_Handler(srv))
	r.GET("/v1/team/{id}", _Team_GetTeam0_HTTP_Handler(srv))
	r.POST("/v1/team/list", _Team_ListTeam0_HTTP_Handler(srv))
	r.PUT("/v1/team/{id}/status", _Team_UpdateTeamStatus0_HTTP_Handler(srv))
	r.GET("/v1/my/team", _Team_MyTeam0_HTTP_Handler(srv))
	r.POST("/v1/team/{id}/member/add", _Team_AddTeamMember0_HTTP_Handler(srv))
	r.DELETE("/v1/team/{id}/member/{user_id}", _Team_RemoveTeamMember0_HTTP_Handler(srv))
	r.PUT("/v1/team/{id}/member/{user_id}/admin", _Team_SetTeamAdmin0_HTTP_Handler(srv))
	r.DELETE("/v1/team/{id}/member/{user_id}/admin", _Team_RemoveTeamAdmin0_HTTP_Handler(srv))
	r.PUT("/v1/team/{id}/member/{user_id}/role", _Team_SetMemberRole0_HTTP_Handler(srv))
	r.POST("/v1/team/{id}/member/list", _Team_ListTeamMember0_HTTP_Handler(srv))
	r.PUT("/v1/team/{id}/leader/transfer", _Team_TransferTeamLeader0_HTTP_Handler(srv))
}

func _Team_CreateTeam0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTeamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamCreateTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTeam(ctx, req.(*CreateTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTeamReply)
		return ctx.Result(200, reply)
	}
}

func _Team_UpdateTeam0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTeamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamUpdateTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTeam(ctx, req.(*UpdateTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTeamReply)
		return ctx.Result(200, reply)
	}
}

func _Team_GetTeam0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTeamRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamGetTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTeam(ctx, req.(*GetTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTeamReply)
		return ctx.Result(200, reply)
	}
}

func _Team_ListTeam0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTeamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamListTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTeam(ctx, req.(*ListTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTeamReply)
		return ctx.Result(200, reply)
	}
}

func _Team_UpdateTeamStatus0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTeamStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamUpdateTeamStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTeamStatus(ctx, req.(*UpdateTeamStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTeamStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Team_MyTeam0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MyTeamRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamMyTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MyTeam(ctx, req.(*MyTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MyTeamReply)
		return ctx.Result(200, reply)
	}
}

func _Team_AddTeamMember0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddTeamMemberRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamAddTeamMember)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddTeamMember(ctx, req.(*AddTeamMemberRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddTeamMemberReply)
		return ctx.Result(200, reply)
	}
}

func _Team_RemoveTeamMember0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveTeamMemberRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamRemoveTeamMember)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveTeamMember(ctx, req.(*RemoveTeamMemberRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RemoveTeamMemberReply)
		return ctx.Result(200, reply)
	}
}

func _Team_SetTeamAdmin0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetTeamAdminRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamSetTeamAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetTeamAdmin(ctx, req.(*SetTeamAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetTeamAdminReply)
		return ctx.Result(200, reply)
	}
}

func _Team_RemoveTeamAdmin0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveTeamAdminRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamRemoveTeamAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveTeamAdmin(ctx, req.(*RemoveTeamAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RemoveTeamAdminReply)
		return ctx.Result(200, reply)
	}
}

func _Team_SetMemberRole0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetMemberRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamSetMemberRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetMemberRole(ctx, req.(*SetMemberRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetMemberRoleReply)
		return ctx.Result(200, reply)
	}
}

func _Team_ListTeamMember0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTeamMemberRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamListTeamMember)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTeamMember(ctx, req.(*ListTeamMemberRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTeamMemberReply)
		return ctx.Result(200, reply)
	}
}

func _Team_TransferTeamLeader0_HTTP_Handler(srv TeamHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TransferTeamLeaderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamTransferTeamLeader)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TransferTeamLeader(ctx, req.(*TransferTeamLeaderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TransferTeamLeaderReply)
		return ctx.Result(200, reply)
	}
}

type TeamHTTPClient interface {
	AddTeamMember(ctx context.Context, req *AddTeamMemberRequest, opts ...http.CallOption) (rsp *AddTeamMemberReply, err error)
	CreateTeam(ctx context.Context, req *CreateTeamRequest, opts ...http.CallOption) (rsp *CreateTeamReply, err error)
	GetTeam(ctx context.Context, req *GetTeamRequest, opts ...http.CallOption) (rsp *GetTeamReply, err error)
	ListTeam(ctx context.Context, req *ListTeamRequest, opts ...http.CallOption) (rsp *ListTeamReply, err error)
	ListTeamMember(ctx context.Context, req *ListTeamMemberRequest, opts ...http.CallOption) (rsp *ListTeamMemberReply, err error)
	MyTeam(ctx context.Context, req *MyTeamRequest, opts ...http.CallOption) (rsp *MyTeamReply, err error)
	RemoveTeamAdmin(ctx context.Context, req *RemoveTeamAdminRequest, opts ...http.CallOption) (rsp *RemoveTeamAdminReply, err error)
	RemoveTeamMember(ctx context.Context, req *RemoveTeamMemberRequest, opts ...http.CallOption) (rsp *RemoveTeamMemberReply, err error)
	SetMemberRole(ctx context.Context, req *SetMemberRoleRequest, opts ...http.CallOption) (rsp *SetMemberRoleReply, err error)
	SetTeamAdmin(ctx context.Context, req *SetTeamAdminRequest, opts ...http.CallOption) (rsp *SetTeamAdminReply, err error)
	TransferTeamLeader(ctx context.Context, req *TransferTeamLeaderRequest, opts ...http.CallOption) (rsp *TransferTeamLeaderReply, err error)
	UpdateTeam(ctx context.Context, req *UpdateTeamRequest, opts ...http.CallOption) (rsp *UpdateTeamReply, err error)
	UpdateTeamStatus(ctx context.Context, req *UpdateTeamStatusRequest, opts ...http.CallOption) (rsp *UpdateTeamStatusReply, err error)
}

type TeamHTTPClientImpl struct {
	cc *http.Client
}

func NewTeamHTTPClient(client *http.Client) TeamHTTPClient {
	return &TeamHTTPClientImpl{client}
}

func (c *TeamHTTPClientImpl) AddTeamMember(ctx context.Context, in *AddTeamMemberRequest, opts ...http.CallOption) (*AddTeamMemberReply, error) {
	var out AddTeamMemberReply
	pattern := "/v1/team/{id}/member/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamAddTeamMember))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...http.CallOption) (*CreateTeamReply, error) {
	var out CreateTeamReply
	pattern := "/v1/team"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamCreateTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...http.CallOption) (*GetTeamReply, error) {
	var out GetTeamReply
	pattern := "/v1/team/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTeamGetTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) ListTeam(ctx context.Context, in *ListTeamRequest, opts ...http.CallOption) (*ListTeamReply, error) {
	var out ListTeamReply
	pattern := "/v1/team/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamListTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) ListTeamMember(ctx context.Context, in *ListTeamMemberRequest, opts ...http.CallOption) (*ListTeamMemberReply, error) {
	var out ListTeamMemberReply
	pattern := "/v1/team/{id}/member/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamListTeamMember))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) MyTeam(ctx context.Context, in *MyTeamRequest, opts ...http.CallOption) (*MyTeamReply, error) {
	var out MyTeamReply
	pattern := "/v1/my/team"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTeamMyTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) RemoveTeamAdmin(ctx context.Context, in *RemoveTeamAdminRequest, opts ...http.CallOption) (*RemoveTeamAdminReply, error) {
	var out RemoveTeamAdminReply
	pattern := "/v1/team/{id}/member/{user_id}/admin"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTeamRemoveTeamAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) RemoveTeamMember(ctx context.Context, in *RemoveTeamMemberRequest, opts ...http.CallOption) (*RemoveTeamMemberReply, error) {
	var out RemoveTeamMemberReply
	pattern := "/v1/team/{id}/member/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTeamRemoveTeamMember))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) SetMemberRole(ctx context.Context, in *SetMemberRoleRequest, opts ...http.CallOption) (*SetMemberRoleReply, error) {
	var out SetMemberRoleReply
	pattern := "/v1/team/{id}/member/{user_id}/role"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamSetMemberRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) SetTeamAdmin(ctx context.Context, in *SetTeamAdminRequest, opts ...http.CallOption) (*SetTeamAdminReply, error) {
	var out SetTeamAdminReply
	pattern := "/v1/team/{id}/member/{user_id}/admin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamSetTeamAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) TransferTeamLeader(ctx context.Context, in *TransferTeamLeaderRequest, opts ...http.CallOption) (*TransferTeamLeaderReply, error) {
	var out TransferTeamLeaderReply
	pattern := "/v1/team/{id}/leader/transfer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamTransferTeamLeader))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...http.CallOption) (*UpdateTeamReply, error) {
	var out UpdateTeamReply
	pattern := "/v1/team/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamUpdateTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamHTTPClientImpl) UpdateTeamStatus(ctx context.Context, in *UpdateTeamStatusRequest, opts ...http.CallOption) (*UpdateTeamStatusReply, error) {
	var out UpdateTeamStatusReply
	pattern := "/v1/team/{id}/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamUpdateTeamStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
