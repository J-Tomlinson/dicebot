package verbs

import(
	"strings"
)

// Vows is a stateless verb, in that it records the vow in the VerbContext's audit log.
type Vows struct {}
	
func (v *Vows)Help() string { return "to give alms to the poor" }

func (v *Vows)Process(vc VerbContext, args []string) string {
	if len(args) < 1 { return v.Help() }
	
	switch args[0] {
	case "to":
		if vc.User == "" {
			return "no you won't, I don't know you"
		} else if len(args) == 1 {
			return "so you claim, pah"
		}

		vc.LogEvent("vowed to " + strings.Join(args[1:], " "))
		return "thy will be done"
		
	default:
		return v.Help()
	}
}
