// A series of utilities so I can build simple CI/CD tools for Azure DevOps
package azdevops

import (
	"fmt"
	"strings"
)

func SetVariable(arg SetVariableArgs) error {
	if strings.TrimSpace(arg.Variable) == "" {
		return &ArgumentNilOrEmptyError{ArgumentName: "arg.Variable"}
	}

	extra := ""
	if arg.IsSecret {
		extra += "issecret=true;"
	}
	if arg.IsOutput {
		extra += "isoutput=true;"
	}
	if arg.IsReadonly {
		extra += "isreadonly=true;"
	}

	fmt.Printf("##vso[task.setvariable variable=%s;%s]%s\n", arg.Variable, extra, arg.Content)
	return nil
}
func LogFmtGroup(group string) {
	fmt.Printf("##[group]%s\n", group)
}
func LogFmtGroupEnd() {
	fmt.Printf("##[endgroup]\n")
}
func LogFmtCommand(command string) {
	fmt.Printf("##[command]%s\n", command)
}
func LogFmtWarning(message string) {
	fmt.Printf("##[warning]%s\n", message)
}
func LogFmtError(message string) {
	fmt.Printf("##[error]%s\n", message)
}
func LogFmtDebug(message string) {
	fmt.Printf("##[debug]%s\n", message)
}
func LogFmtSection(message string) {
	fmt.Printf("##[section]%s\n", message)
}

func LogIssueError(message string) {
	fmt.Printf("##vso[task.logissue type=error]%s\n", message)
}
func LogIssueWarning(message string) {
	fmt.Printf("##vso[task.logissue type=warning]%s\n", message)
}
func LogIssueErrorSource(message string, linenumber int, columnnumber int, code int) {
	fmt.Printf("##vso[task.logissue type=error;linenumber=%d;columnnumber=%d;code=%d;]%s\n", linenumber, columnnumber, code, message)
}
func LogIssueWarningSource(message string, linenumber int, columnnumber int, code int) {
	fmt.Printf("##vso[task.logissue type=warning;linenumber=%d;columnnumber=%d;code=%d;]%s\n", linenumber, columnnumber, code, message)
}
