package task

import (
	"testing"

	repomock "github.com/ahmedshaaban/todolist/internal/usecase/task/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

func TestTask(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Task Suite")
}

var _ = Describe("Task Service", func() {
	var (
		taskService *Service
		repo        *repomock.Repo
	)

	BeforeEach(func() {
		repo = &repomock.Repo{}
		repo.On("Add", mock.Anything).Return(nil)
		repo.On("GetAll").Return([][]string{{"1", "test description", "true", "false"}, {"2", "test description", "true", "false"}}, nil)
		repo.On("UpdateStatus", mock.Anything, mock.Anything).Return(nil)
		repo.On("DeleteDisabled").Return(nil)

		taskService = New(repo)
	})

	Describe("Task", func() {
		Describe("Add", func() {
			It("should propagate input to repo", func() {
				err := taskService.AddTask([]string{"test"})
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Describe("ChangeTaskStatus", func() {
			It("should propagate input to repo", func() {
				err := taskService.ChangeTaskStatus(1, false)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Describe("CleanDoneTasks", func() {
			It("should propagate input to repo", func() {
				err := taskService.CleanDoneTasks()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Describe("GetTasks", func() {
			It("should propagate input to repo and return correct format", func() {
				result, err := taskService.GetTasks()

				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal([]string{"1: test description", "2: test description"}))
			})
		})
	})
})
