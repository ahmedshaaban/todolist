package db

import (
	"encoding/csv"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDB(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "DB Suite")
}

var _ = BeforeSuite(func() {
	file, err := os.Create("test.csv")
	Expect(err).To(BeNil())
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.WriteAll([][]string{{"1", "test description", "true", "false"}})
	Expect(err).To(BeNil())
})

var _ = AfterSuite(func() {
	err := os.Remove("test.csv")
	Expect(err).To(BeNil())
})

var _ = Describe("DB", func() {
	var (
		db *DB
	)

	BeforeEach(func() {
		db = New("test.csv")
	})

	AfterEach(func() {
		file, err := os.Create("test.csv")
		Expect(err).To(BeNil())
		defer file.Close()

		w := csv.NewWriter(file)
		defer w.Flush()

		err = w.WriteAll([][]string{{"1", "test description", "true", "false"}})
		Expect(err).To(BeNil())
	})

	Describe("GetAll", func() {
		It("retrives the correct data", func() {
			result, err := db.GetAll()

			Expect(err).To(BeNil())
			Expect(result).To(Equal([][]string{{"1", "test description", "true", "false"}}))
		})

		It("should return file not found error", func() {
			db = New("invalid")
			_, err := db.GetAll()

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("open invalid: no such file or directory"))
		})
	})

	Describe("Add", func() {
		It("adds a new input to the csv", func() {
			err := db.Add([]string{"new test description", "true", "false"})
			Expect(err).To(BeNil())

			result, err := db.GetAll()

			Expect(err).To(BeNil())
			Expect(result).To(Equal([][]string{{"1", "test description", "true", "false"}, {"2", "new test description", "true", "false"}}))
		})

		Context("empty file", func() {
			BeforeEach(func() {
				err := os.Remove("test.csv")
				Expect(err).To(BeNil())
				file, err := os.Create("test.csv")
				Expect(err).To(BeNil())
				defer file.Close()
			})

			It("adds to empty file", func() {
				err := db.Add([]string{"new test description", "true", "false"})
				Expect(err).To(BeNil())

				result, err := db.GetAll()

				Expect(err).To(BeNil())
				Expect(result).To(Equal([][]string{{"1", "new test description", "true", "false"}}))
			})
		})

		It("should return file not found error", func() {
			db = New("invalid")
			err := db.Add([]string{})

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("open invalid: no such file or directory"))
		})
	})

	Describe("UpdateStatus", func() {
		It("should update status of an input", func() {
			err := db.UpdateStatus(1, false)
			Expect(err).To(BeNil())

			result, err := db.GetAll()

			Expect(err).To(BeNil())
			Expect(result).To(Equal([][]string{}))
		})

		It("should retrun index not found error", func() {
			err := db.UpdateStatus(10, false)
			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("index out of range"))
		})

		It("should not update deleted record", func() {
			err := db.UpdateStatus(1, false)
			Expect(err).To(BeNil())

			err = db.DeleteDisabled()
			Expect(err).To(BeNil())

			err = db.UpdateStatus(1, false)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("index not found"))
		})
	})

	Describe("DeleteDisabled", func() {
		It("should update status of an input", func() {
			err := db.UpdateStatus(1, false)
			Expect(err).To(BeNil())

			err = db.DeleteDisabled()
			Expect(err).To(BeNil())

			result, err := db.GetAll()

			Expect(err).To(BeNil())
			Expect(result).To(Equal([][]string{}))
		})

		It("should return file not found error", func() {
			db = New("invalid")
			err := db.UpdateStatus(2, false)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(Equal("open invalid: no such file or directory"))
		})
	})
})
