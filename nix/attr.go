package nix

import (
	"os"
	"os/exec"
	"strings"
)

// An abstract Nix attribute, possibly nested.
type Attr interface {
	Show(string, *Options) error
}


// A Nix scalar value.
type Value struct {
	Documentation string
}

func (v Value) Show(attr string, options *Options) error {
	less := exec.Command("less")
	less.Stdin = strings.NewReader(v.Documentation)
	less.Stderr = os.Stderr
	less.Stdout = os.Stdout
	return less.Run()
}


// A Nix attribute set.
type Attrset struct {
	Attrs []string
}

func (a Attrset) Show(attr string, options *Options) error {
	var prompt string
	if len(attr) == 0 {
		prompt = "> "
	} else {
		prompt = attr + "."
	}

	for {
		fzf := exec.Command("fzf", "--prompt", prompt, "--expect=.")
		fzf.Stdin = strings.NewReader(strings.Join(a.Attrs, "\n"))
		fzf.Stderr = os.Stderr
		output, err := fzf.Output()
		if err != nil {
			switch v := err.(type) {
			case *exec.ExitError:
				switch v.ExitCode() {
				case 1:
				// No match, do nothing and repeat the prompt.
				default:
					return err
				}
			default:
				return err
			}
		}

		newAttr := strings.Split(string(output), "\n")[1]
		var attrPath string
		if len(attr) > 0 {
			attrPath = attr + "." + newAttr
		} else {
			attrPath = newAttr
		}
		options.Show(attrPath)
	}
}
