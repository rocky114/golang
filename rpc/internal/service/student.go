package service

import (
	"context"
	"gitlab.sz.sensetime.com/kubersolver/api/student"
)

type StudentManager struct {
	student.UnimplementedStudentManagerServer
	Addr string
}

func (stu *StudentManager) Echo(ctx context.Context, req *student.StringMessage) (*student.StringMessage, error) {
	return &student.StringMessage{
		Value: req.GetValue() + " from " + stu.Addr,
	}, nil
}
