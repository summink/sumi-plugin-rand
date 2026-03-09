package sumi_rand

import (
	"fmt"
	"myPlugin/random"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/atotto/clipboard"
)

func RegisterCommand() *command.SCommand {
	i := command.NewCommandFlag[bool]("int").ChangeDescription("set random int")
	f := command.NewCommandFlag[bool]("float").ChangeDescription("set random float")
	s := command.NewCommandFlag[bool]("string").ChangeDescription("set random string").ChangeValue(true)
	r := command.NewCommandFlag[[]string]("ranges").ChangeDescription("set random charsets").ChangeValue([]string{})

	copy := command.NewCommandFlag[bool]("copy").ChangeDescription("copy random value content to clipboard").ChangeValue(false).ChangeShort("")
	max := command.NewCommandFlag[int]("max").ChangeDescription("max number or max string length").ChangeValue(10000).ChangeShort("")
	min := command.NewCommandFlag[int]("min").ChangeDescription("min number or min string length").ChangeValue(1).ChangeShort("")

	cmd := command.
		NewCommand("rand").
		ChangeDescription("Random value").
		AddFlags(i, f, s, max, min, copy, r).
		RegisterHandler(func(cmd *command.SCommand) {
			var value string

			if len(r.Value) > 0 {
				value = random.RandomStringn(min.Value, max.Value, strings.Join(r.Value, ""))
			} else if i.Value {
				value = fmt.Sprint(random.RandomIntn(min.Value, max.Value))
			} else if f.Value {
				value = fmt.Sprint(random.RandomFloatn(min.Value, max.Value))
			} else if s.Value {
				value = random.RandomString(min.Value, max.Value)
			}

			println(value)

			if copy.Value {
				clipboard.WriteAll(string(value))
			}
		})

	return cmd
}
