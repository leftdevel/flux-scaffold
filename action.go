package main

import (
    "fmt"
    "strings"
)

type Action struct {
    ResourceOptions
}

func (a *Action) GetDir() string {
        return "actions"
}

func (a *Action) GetFileName() string {
    return a.ResourceOptions.Domain + "-actions.js"
}

func (a *Action) GetFileContent() string {
    domainFirstCharUpper := strings.ToUpper(a.Domain[:1]) + a.Domain[1:]
    return fmt.Sprintf(a.getTemplate(), domainFirstCharUpper, a.Domain, a.Action, a.Constant)
}

func (a *Action) getTemplate() string {
    content := a.getHeaderTemplate()
    content = content + a.getBodyTemplateVariation()
    content = content + a.getFooterTemplate()

    return content
}

func (a *Action) getHeaderTemplate() string {
    return "var AppDispatcher = require('../dispatcher/app-dispatcher');\n" +
        "var %[1]sConstants = require('../constants/%[2]s-constants');\n"
}

func (a *Action) getBodyTemplateVariation() string {
    if a.IsApi {
        return "" +
            "var %[1]sApi = require('../webapi/%[2]s-api');\n\n" +
            "var %[1]sActions = {\n" +
            "    %[3]s: function() {\n" +
            "        AppDispatcher.dispatch({\n" +
            "            actionType: %[1]sConstants.%[4]s,\n" +
            "        });\n\n" +
            "        %[1]sApi.%[3]s(%[1]sActions.%[3]sSuccess);\n" +
            "    },\n\n" +
            "    %[3]sSuccess: function(%[2]ss) {\n" +
            "        AppDispatcher.dispatch({\n" +
            "            actionType: %[1]sConstants.%[4]s_SUCCESS,\n" +
            "            %[2]ss: %[2]ss\n" +
            "        });\n" +
            "    }\n"
    } else {
        return "\n" +
            "var %[1]sActions = {\n" +
            "    %[3]s: function(%[2]s) {\n" +
            "        AppDispatcher.dispatch({\n" +
            "            actionType: %[1]sConstants.%[4]s,\n" +
            "            %[2]s: %[2]s\n" +
            "        });\n" +
            "    }\n"
    }
}

func (a *Action) getFooterTemplate() string {
    return "" +
        "};\n\n" +
        "module.exports = %[1]sActions;"
}
