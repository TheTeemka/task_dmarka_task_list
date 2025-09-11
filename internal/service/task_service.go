package service

import (
	"github.com/TheTeemka/task_dmarka_task_list/internal/models"
	"github.com/TheTeemka/task_dmarka_task_list/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepo
}

func NewTaskService(taskRepo *repository.TaskRepo) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (ts *TaskService) CreateNewTask(task *models.TaskDTO) (*models.TaskDTO, error) {
	if err := task.Validate(); err != nil {
		return nil, err
	}

	createdTask, err := ts.taskRepo.CreateTask(task.ToModel())
	if err != nil {
		return nil, err
	}

	return createdTask.ToDTO(), nil
}

func (ts *TaskService) GetTaskByID(id int64) (*models.TaskDTO, error) {
	taskModel, err := ts.taskRepo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	return taskModel.ToDTO(), nil
}

func (ts *TaskService) GetListByFilters(filters *models.TaskFilter) ([]*models.TaskDTO, error) {
	taskModels, err := ts.taskRepo.GetListByFilters(filters)
	if err != nil {
		return nil, err
	}

	var tasks []*models.TaskDTO
	for _, tm := range taskModels {
		tasks = append(tasks, tm.ToDTO())
	}
	return tasks, nil
}
