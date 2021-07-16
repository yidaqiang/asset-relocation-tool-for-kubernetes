package cmd_test

// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/asset-relocation-tool-for-kubernetes/v2/cmd"
)

var _ = Describe("Chart", func() {
	Describe("ParseOutputFlag", func() {
		It("works with default out flag", func() {
			got, err := cmd.ParseOutputFlag(cmd.Output)
			want := "./%s-%s.relocated.tgz"
			Expect(got).To(Equal(want))
			Expect(err).To(BeNil())
		})
		It("rejects out flag without wildcard *", func() {
			_, err := cmd.ParseOutputFlag("nowildcardhere.tgz")
			Expect(err).Should(MatchError(cmd.ErrorMissingOutPlaceHolder))
		})
		It("rejects out flag without proper extension", func() {
			_, err := cmd.ParseOutputFlag("*-wildcardhere")
			Expect(err).Should(MatchError(cmd.ErrorBadExtension))
		})
		It("accepts out flag with wildcard", func() {
			got, err := cmd.ParseOutputFlag("*-wildcardhere.tgz")
			Expect(got).To(Equal("%s-%s-wildcardhere.tgz"))
			Expect(err).To(BeNil())
		})
	})

	Describe("TargetOutput", func() {
		It("works with default out flag", func() {
			outFmt := "./%s-%s.relocated.tgz"
			target := cmd.TargetOutput("path", outFmt, "my-chart", "0.1")
			Expect(target).To(Equal("path/my-chart-0.1.relocated.tgz"))
		})
		It("builds custom out input as expected", func() {
			target := cmd.TargetOutput("path", "%s-%s-wildcardhere.tgz", "my-chart", "0.1")
			Expect(target).To(Equal("path/my-chart-0.1-wildcardhere.tgz"))
		})
	})
})
