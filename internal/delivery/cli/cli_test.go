package cli

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	servicemock "github.com/ahmedshaaban/todolist/internal/delivery/cli/mocks"
)

func TestCli(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Cli Suite")
}

var _ = Describe("Cli Delivery", func() {
	var (
		service *servicemock.TaskService
		cli     *Cli
	)

	BeforeEach(func() {
		service = &servicemock.TaskService{}
		service.On("AddTask", mock.Anything).Return(nil)
		service.On("GetTasks").Return([]string{"1: test", "2: test"}, nil)
		service.On("ChangeTaskStatus", mock.Anything, mock.Anything).Return(nil)
		service.On("CleanDoneTasks").Return(nil)

		cli = New(service)
	})

	Describe("Cli", func() {
		Describe("AddTask", func() {
			It("adds new task", func() {
				cmd := cli.AddTask()
				cmd.SetArgs([]string{"test"})
				err := cmd.Execute()

				Expect(err).To(Not(HaveOccurred()))
			})

			It("returns error if empty args", func() {
				cmd := cli.AddTask()
				cmd.SetArgs([]string{})

				r, w, _ := os.Pipe()
				tmp := os.Stderr
				defer func() {
					os.Stderr = tmp
				}()
				os.Stderr = w

				go func() {
					err := cmd.Execute()
					Expect(err).To(HaveOccurred())

					w.Close()
				}()

				stdout, _ := ioutil.ReadAll(r)
				Expect(string(stdout)).To(Equal("Error: task name is required\nUsage:\n  add [flags]\n\nFlags:\n  -h, --help   help for add\n\n"))
			})
		})

		Describe("GetTasks", func() {
			It("returns all tasks formated", func() {
				cmd := cli.GetTasks()
				cmd.SetArgs([]string{""})

				r, w, _ := os.Pipe()
				tmp := os.Stdout
				defer func() {
					os.Stdout = tmp
				}()
				os.Stdout = w

				go func() {
					err := cmd.Execute()
					Expect(err).To(Not(HaveOccurred()))

					w.Close()
				}()

				stdout, _ := ioutil.ReadAll(r)
				Expect(string(stdout)).To(Equal("1: test\n2: test\n"))
			})
		})

		Describe("DoneTask", func() {
			It("marks task as done", func() {
				cmd := cli.DoneTask()
				cmd.SetArgs([]string{"1"})
				err := cmd.Execute()

				Expect(err).To(Not(HaveOccurred()))
			})

			It("returns error if empty args", func() {
				cmd := cli.DoneTask()
				cmd.SetArgs([]string{})

				r, w, _ := os.Pipe()
				tmp := os.Stderr
				defer func() {
					os.Stderr = tmp
				}()
				os.Stderr = w

				go func() {
					err := cmd.Execute()
					Expect(err).To(HaveOccurred())

					w.Close()
				}()

				stdout, _ := ioutil.ReadAll(r)
				Expect(string(stdout)).To(Equal("Error: task id needed to mark task done\nUsage:\n  done [flags]\n\nFlags:\n  -h, --help   help for done\n\n"))
			})
		})

		Describe("UnDoneTask", func() {
			It("marks task as undone", func() {
				cmd := cli.UnDoneTask()
				cmd.SetArgs([]string{"1"})
				err := cmd.Execute()

				Expect(err).To(Not(HaveOccurred()))
			})

			It("returns error if empty args", func() {
				cmd := cli.UnDoneTask()
				cmd.SetArgs([]string{})

				r, w, _ := os.Pipe()
				tmp := os.Stderr
				defer func() {
					os.Stderr = tmp
				}()
				os.Stderr = w

				go func() {
					err := cmd.Execute()
					Expect(err).To(HaveOccurred())

					w.Close()
				}()

				stdout, _ := ioutil.ReadAll(r)
				Expect(string(stdout)).To(Equal("Error: task id needed to mark task undone\nUsage:\n  undone [flags]\n\nFlags:\n  -h, --help   help for undone\n\n"))
			})
		})

		Describe("Cleanup", func() {
			It("marks task as undone", func() {
				cmd := cli.Cleanup()
				cmd.SetArgs([]string{""})
				err := cmd.Execute()

				Expect(err).To(Not(HaveOccurred()))
			})
		})
	})
})
