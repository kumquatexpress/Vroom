package helpers

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vroom Options", func() {
	var (
		filename string
	)

	Describe("Parse Options", func() {
		It("returns an error when given nil", func() {
			opts, err := parseOpts(nil)
			Expect(opts).To(BeNil())
			Expect(err).To(BeAssignableToTypeOf(&json.SyntaxError{}))
		})

		It("returns the correct options when given json data", func() {
			data := map[string]string{"TemplateDirectory": "test"}
			bytes, err := json.Marshal(data)
			Expect(err).To(BeNil())

			opts, err := parseOpts(bytes)
			Expect(opts).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
	})

	Describe("Default Options", func() {
		It("returns the correct type as a default option set", func() {
			Expect(NewVroomOpts()).To(BeAssignableToTypeOf(&VroomOpts{}))
		})
	})

	Describe("NewVroomOptsFromFile", func() {
		It("returns the default when given a non-existant filename", func() {
			filename = "YouCan'tFindMe"
			Expect(NewVroomOptsFromFile(filename)).To(Equal(NewVroomOpts()))
		})

		It("returns the default when given malformed data", func() {
			filename = "../test/test_opts_malformed.json"
			Expect(NewVroomOptsFromFile(filename)).To(Equal(NewVroomOpts()))
		})

		It("returns a correct opt struct, ignoring irrelevant data", func() {
			filename = "../test/test_opts.json"
			opts := NewVroomOptsFromFile(filename)

			Expect(opts.TemplateDirectory).To(Equal("testdir/templates"))
			Expect(opts.BuildDirectory).To(Equal("testdir/compiled"))
			Expect(opts.Metadata).To(Equal(
				map[string]string{"title": "test", "body": "hello world"}))
		})
	})
})
