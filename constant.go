package main

import (
    "fmt"
    "strings"
)

type Constant struct {
    ResourceOptions
}

func (c *Constant) GetDir() string {
        return "constants"
}

func (c *Constant) GetFileName() string {
    return c.ResourceOptions.Domain + "-constants.js"
}

func (c *Constant) GetFileContent() string {
    domainFirstCharUpper := strings.ToUpper(c.Domain[:1]) + c.Domain[1:]
    return fmt.Sprintf(c.getTemplate(), domainFirstCharUpper, c.Constant)
}

func (c *Constant) getTemplate() string {
    content := "var keyMirror = require('keymirror');\n\n" +
        "var %[1]sConstants = keyMirror({\n";

    if c.IsApi {
        content = content +
        "    %[2]s: null,\n" +
        "    %[2]s_SUCCESS: null\n"
    } else {
        content = content +
        "    %[2]s: null\n"
    }

    content = content +
        "});\n\n" +
        "module.exports = %[1]sConstants;"

    return content
}
