package config

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("Config", func() {

	Describe("Reading external configuration file", func() {

		Context("Reading a valid configuration", func() {

			It("should return a populated viper struct", func() {
				viper, err := readWorkflowDefinition("./testdata/valid.yaml")
				Expect(err).To(BeNil())
				Expect((*viper).GetString("flowit.version")).To(Equal("0.1"))
			})

		})

		Context("Reading an invalid configuration", func() {

			It("should return an informative error", func() {
				viper, err := readWorkflowDefinition("./testdata/incorrect-format.yaml")
				Expect(err).To(Not(BeNil()))
				Expect(errors.Cause(err).Error()).To(MatchRegexp("While parsing config: yaml: line [0-9]+"))
				Expect(viper).To(BeNil())
			})

		})

		Context("Reading a non existent configuration", func() {

			It("should return an informative error", func() {
				viper, err := readWorkflowDefinition("./testdata/non-existent.yaml")
				Expect(err).To(Not(BeNil()))
				Expect(errors.Cause(err).Error()).To(ContainSubstring("Not Found"))
				Expect(viper).To(BeNil())
			})

		})

		Context("Reading a configuration with no read permissions", func() {

			It("should return an informative error", func() {
				fileName := "./testdata/valid.yaml"
				if err := os.Chmod(fileName, 0000); err != nil {
					Fail(fmt.Sprintf("Error changing test file: "+fileName+" permissions: %+v", err))
				}

				defer func() {
					if err := os.Chmod(fileName, 0644); err != nil {
						Fail(fmt.Sprintf("Error restoring test file: "+fileName+" permissions: %+v", err))
					}
				}()
				viper, err := readWorkflowDefinition("./testdata/valid.yaml")
				Expect(err).To(Not(BeNil()))
				Expect(errors.Cause(err).Error()).To(ContainSubstring("permission denied"))
				Expect(viper).To(BeNil())
			})

		})

	})
})
