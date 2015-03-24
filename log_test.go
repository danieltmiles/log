package log

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/monsooncommerce/mockwriter"
	. "github.com/onsi/gomega"
)

func TestLogging(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Logging", func() {
		var m *mockwriter.MockWriter
		var logger *Logger

		g.BeforeEach(func() {
			m = &mockwriter.MockWriter{}
			logger = New(m, Debug)
		})

		g.It("should write a debug message", func() {
			logger.Debug("debug message")
			Expect(m.Written).To(ContainSubstring("debug message"))
			Expect(m.Written).To(ContainSubstring("DEBUG"))
		})

		g.It("should write an info message", func() {
			logger.Info("info message")
			Expect(m.Written).To(ContainSubstring("info message"))
			Expect(m.Written).To(ContainSubstring("INFO"))
		})

		g.It("should write a notice message", func() {
			logger.Notice("notice message")
			Expect(m.Written).To(ContainSubstring("notice message"))
			Expect(m.Written).To(ContainSubstring("NOTICE"))
		})

		g.It("should write a warning message", func() {
			logger.Warning("warning message")
			Expect(m.Written).To(ContainSubstring("warning message"))
			Expect(m.Written).To(ContainSubstring("WARNING"))
		})

		g.It("should write a error message", func() {
			logger.Error("error message")
			Expect(m.Written).To(ContainSubstring("error message"))
			Expect(m.Written).To(ContainSubstring("ERROR"))
		})
	})

	g.Describe("Thresholds", func() {
		g.It("should write debug and higher severity", func() {
			m := &mockwriter.MockWriter{}
			logger := New(m, Debug)
			logger.Debug("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("DEBUG"))

			logger.Info("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("INFO"))

			logger.Notice("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("NOTICE"))

			logger.Warning("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("WARNING"))

			logger.Error("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("ERROR"))
		})

		g.It("should write info and higher severity", func() {
			m := &mockwriter.MockWriter{}
			logger := New(m, Info)
			logger.Debug("test message")
			Expect(m.Written).To(BeNil())

			logger.Info("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("INFO"))

			logger.Notice("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("NOTICE"))

			logger.Warning("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("WARNING"))

			logger.Error("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("ERROR"))
		})

		g.It("should write notice and higher severity", func() {
			m := &mockwriter.MockWriter{}
			logger := New(m, Notice)
			logger.Debug("test message")
			Expect(m.Written).To(BeNil())

			logger.Info("test message")
			Expect(m.Written).To(BeNil())

			logger.Notice("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("NOTICE"))

			logger.Warning("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("WARNING"))

			logger.Error("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("ERROR"))
		})

		g.It("should write warning and higher severity", func() {
			m := &mockwriter.MockWriter{}
			logger := New(m, Warning)
			logger.Debug("test message")
			Expect(m.Written).To(BeNil())

			logger.Info("test message")
			Expect(m.Written).To(BeNil())

			logger.Notice("test message")
			Expect(m.Written).To(BeNil())

			logger.Warning("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("WARNING"))

			logger.Error("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("ERROR"))
		})

		g.It("should write error and higher severity", func() {
			m := &mockwriter.MockWriter{}
			logger := New(m, Error)
			logger.Debug("test message")
			Expect(m.Written).To(BeNil())

			logger.Info("test message")
			Expect(m.Written).To(BeNil())

			logger.Notice("test message")
			Expect(m.Written).To(BeNil())

			logger.Warning("test message")
			Expect(m.Written).To(BeNil())

			logger.Error("test message")
			Expect(m.Written).To(ContainSubstring("test message"))
			Expect(m.Written).To(ContainSubstring("ERROR"))
		})
	})
}
