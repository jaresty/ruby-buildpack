package supply_test

import (
	"io/ioutil"
	"os"
	"ruby/supply"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -source=supply.go -destination=mocks_supply_test.go --package=supply_test

var _ = Describe("Supply", func() {
	var (
		buildDir string
		depDir   string
		err      error

		mockCtrl     *gomock.Controller
		mockManifest *MockManifest
		mockLogger   *MockLogger
		mockRunner   *MockRunner
		subject      *supply.Supply
	)

	BeforeEach(func() {
		buildDir, err = ioutil.TempDir("", "nodejs-buildpack.build.")
		Expect(err).To(BeNil())

		depDir, err = ioutil.TempDir("", "nodejs-buildpack.dep.")
		Expect(err).To(BeNil())

		mockCtrl = gomock.NewController(GinkgoT())
		mockManifest = NewMockManifest(mockCtrl)
		mockLogger = NewMockLogger(mockCtrl)
		mockRunner = NewMockRunner(mockCtrl)
	})

	JustBeforeEach(func() {
		subject = &supply.Supply{BuildDir: buildDir, DepDir: depDir, Manifest: mockManifest, Log: mockLogger, Runner: mockRunner}
	})

	AfterEach(func() {
		mockCtrl.Finish()

		err = os.RemoveAll(buildDir)
		Expect(err).To(BeNil())

		err = os.RemoveAll(depDir)
		Expect(err).To(BeNil())
	})

})
