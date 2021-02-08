// +build darwin dragonfly freebsd linux,!appengine netbsd openbsd solaris

package ishell

import (
	"github.com/MarekStancik/readline"
)

func clearScreen(s *Shell) error {
	_, err := readline.ClearScreen(s.writer)
	return err
}
