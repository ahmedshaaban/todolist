// +build e2e

package e2e

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

var todoBinary string

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Suite")
}

var _ = BeforeSuite(func() {
	var err error

	todoBinary, err = gexec.Build("github.com/ahmedshaaban/todolist")
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	err := os.Remove("db.csv")
	Expect(err).ShouldNot(HaveOccurred())

	gexec.CleanupBuildArtifacts()
})

var _ = Describe("E2E", func() {
	Describe("add command", func() {
		It("should add new task", func() {
			cmd := exec.Command(todoBinary, "add", "e2e description")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal(""))

		})

		It("should list the newly added task", func() {
			cmd := exec.Command(todoBinary, "list")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal("1: e2e description\n"))
		})

	})

	Describe("done command", func() {
		It("should mark item as done", func() {
			cmd := exec.Command(todoBinary, "done", "1")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal(""))

		})

		It("should not be shown when listing", func() {
			cmd := exec.Command(todoBinary, "list")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal("\n"))
		})
	})

	Describe("undone command", func() {
		It("should mark item as undone", func() {
			cmd := exec.Command(todoBinary, "undone", "1")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal(""))

		})

		It("should be shown when listing", func() {
			cmd := exec.Command(todoBinary, "list")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal("1: e2e description\n"))
		})
	})

	Describe("cleanup command", func() {
		It("should mark item as done", func() {
			cmd := exec.Command(todoBinary, "done", "1")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal(""))

		})

		It("should cleanup done items", func() {
			cmd := exec.Command(todoBinary, "cleanup")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal(""))

		})

		It("should be shown when listing", func() {
			cmd := exec.Command(todoBinary, "list")

			resultBytes, err := cmd.Output()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(resultBytes)).Should(Equal("\n"))
		})
	})
})
