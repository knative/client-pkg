/*
 Copyright 2024 The Knative Authors

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"knative.dev/client-pkg/pkg/output"
	"knative.dev/client-pkg/pkg/output/term"
)

const spinnerColor = lipgloss.Color("205")

type Spinner interface {
	Runnable[Spinner]
}

func (w *widgets) NewSpinner(message string) Spinner {
	return &BubbleSpinner{
		InputOutput: output.PrinterFrom(w.ctx),
		Message:     Message{Text: message},
	}
}

type BubbleSpinner struct {
	output.InputOutput
	Message

	spin spinner.Model
	tea  *tea.Program
}

func (b *BubbleSpinner) With(fn func(Spinner) error) error {
	b.start()
	defer b.stop()
	return fn(b)
}

func (b *BubbleSpinner) Init() tea.Cmd {
	return b.spin.Tick
}

func (b *BubbleSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, c := b.spin.Update(msg)
	b.spin = m
	return b, c
}

func (b *BubbleSpinner) View() string {
	return fmt.Sprintf("%s %s", b.Message.Text, b.spin.View())
}

func (b *BubbleSpinner) start() {
	b.spin = spinner.New(
		spinner.WithSpinner(spinner.Meter),
		spinner.WithStyle(spinnerStyle()),
	)
	out := b.OutOrStdout()
	b.tea = tea.NewProgram(b,
		tea.WithInput(b.InOrStdin()),
		tea.WithOutput(out),
	)
	go func() {
		t := b.tea
		_, _ = t.Run()
		if term.IsTerminal(out) {
			_ = t.ReleaseTerminal()
		}
	}()
}

func (b *BubbleSpinner) stop() {
	if b.tea == nil {
		return
	}
	b.tea.Quit()
	b.tea.Wait()
	b.tea = nil
	endMsg := fmt.Sprintf("%s %s\n",
		b.Message.Text, spinnerStyle().Render("Done"))
	_, _ = b.OutOrStdout().Write([]byte(endMsg))
}

func spinnerStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(spinnerColor)
}
