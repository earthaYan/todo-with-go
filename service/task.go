package service

import (
	"time"
	"to-do-list/model"
	"to-do-list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"titele" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0 未完成 1已完成
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

type ShowTaskService struct {
}

func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

type ListTaskService struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	var count int64 = 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

type UpdateTaskService struct {
	Title   string `json:"titele" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0 未完成 1已完成
}

func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新成功",
	}
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	var count int64 = 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("title LIKE ? OR content Like ?", "%"+service.Info+"%", "%"+service.Info+"%").Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}
