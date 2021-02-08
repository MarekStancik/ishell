// +build windows

package ishell

import (
	"github.com/MarekStancik/readline"
)

func clearScreen(s *Shell) error {
	return readline.ClearScreen(s.writer)
}
