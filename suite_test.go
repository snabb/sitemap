package sitemap_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSitemap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sitemap Suite")
}
