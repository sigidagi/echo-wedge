package cmd

import (
	"os"
	"text/template"

	"echo-wedge/backend/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// when updating this template, don't forget to update config.md!
const configTemplate = `
[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }}

# Log to syslog.
#
# When set to true, log messages are being written to syslog.
log_to_syslog={{ .General.LogToSyslog }}

# Gateway configuration.
#
[gateway]
# hostname:port to connect  to gateway 
url="{{ .Gateway.Url }}"

[rest]
bind="{{ .Rest.Bind }}"
url="{{ .Rest.Url }}"
`

var configCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the Rest service configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
