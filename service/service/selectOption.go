package service

import (
	"code.shomes.cn/tfblue/smartgy/protobuf"
	"code.shomes.cn/tfblue/smartgy/tool"
	"context"
)

type SelectOption struct {
}

func (s *SelectOption) StructureType(ctx context.Context, in *protobuf.SelectOptionStructureTypeIn) (*protobuf.SelectOptionStructureTypeOut, error) {
	db := tool.Smart(ctx)
	_ = db.StartTrans()
	out := protobuf.SelectOptionStructureTypeOut{}
	_ = db.TableName(&out.Data, "structure_type").Filed("id,name").Select()
	return &out, nil
}

// 进度类型
func (s *SelectOption) ScheduleType(ctx context.Context, in *protobuf.SelectOptionScheduleTypeIn) (*protobuf.SelectOptionScheduleTypeOut, error) {
	db := tool.Smart(ctx)
	_ = db.StartTrans()
	out := protobuf.SelectOptionScheduleTypeOut{}
	_ = db.TableName(&out.Data, "schedule_type").Filed("id,name").Select()
	return &out, nil
}

// 工程用途
func (s *SelectOption) ProjectUse(ctx context.Context, in *protobuf.SelectOptionProjectUseIn) (*protobuf.SelectOptionProjectUseOut, error) {
	db := tool.Smart(ctx)
	_ = db.StartTrans()
	out := protobuf.SelectOptionProjectUseOut{}
	_ = db.TableName(&out.Data, "project_use").Filed("id,name").Select()
	return &out, nil
}

// 工程类别
func (s *SelectOption) ProjectType(ctx context.Context, in *protobuf.SelectOptionProjectTypeIn) (*protobuf.SelectOptionProjectTypeOut, error) {
	db := tool.Smart(ctx)
	_ = db.StartTrans()
	out := protobuf.SelectOptionProjectTypeOut{}
	_ = db.TableName(&out.Data, "project_type").Filed("id,name").Select()
	return &out, nil
}
