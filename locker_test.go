package filelock_test

import (
	"fmt"
	"sync"

	"code.cloudfoundry.org/filelock"
	"code.cloudfoundry.org/filelock/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Locker", func() {
	var (
		locker *filelock.Locker
		flock  *fakes.FileLocker
	)
	BeforeEach(func() {
		flock = &fakes.FileLocker{}
		locker = &filelock.Locker{
			FileLocker: flock,
			Mutex:      &sync.Mutex{},
		}
	})
	Describe("Lifecycle", func() {
		It("locks and unlocks", func() {
			err := locker.Lock()
			Expect(err).NotTo(HaveOccurred())

			Expect(flock.OpenCallCount()).To(Equal(1))
		})
		Context("when fileLocker fails to open", func() {
			BeforeEach(func() {
				flock.OpenReturns(nil, fmt.Errorf("banana"))

			})
			It("should return the error", func() {
				err := locker.Lock()
				Expect(err).To(MatchError("open lock file: banana"))
			})
		})
	})
})
