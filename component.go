package main

import (
    "fmt"
    "strings"
)

type Component struct {
    ResourceOptions
}

func (c *Component) GetDir() string {
    return "components"
}

func (c *Component) GetFileName() string {
    if c.IsApi {
        return c.ResourceOptions.Domain + "s.js"
    } else {
        return c.ResourceOptions.Domain + ".js"
    }
}

func (c *Component) GetFileContent() string {
    domainFirstCharUpper := strings.ToUpper(c.Domain[:1]) + c.Domain[1:]
    return fmt.Sprintf(c.getTemplate(), domainFirstCharUpper, c.Domain)
}

func (c *Component) getTemplate() string {
    content := c.getHeaderTemplate()
    content = content + c.getBodyTemplateVariation()

    return content
}

func (c *Component) getHeaderTemplate() string {
    return "" +
        "var React = require('react');\n" +
        "var %[1]sStore = require('../stores/%[2]s-store');\n\n"
}

func (c *Component) getBodyTemplateVariation() string {
    common := c.getCommonBodyTemplate();

    if c.IsApi {
        return "" +
            "function getState() {\n" +
            "    return {\n" +
            "        %[2]ss: ToastStore.get%[1]ss()\n" +
            "    };\n" +
            "}\n\n" +
            "var %[1]ss = React.createClass({\n" +
            common +
            "    render: function() {\n" +
            "        return (\n" +
            "            <h1>%[1]ss Component</h1>\n" +
            "        );\n" +
            "    }\n" +
            "});\n\n" +
            "module.exports = %[1]ss;\n"
    } else {
        return "" +
            "function getState() {\n" +
            "    return {\n" +
            "        %[2]s: ToastStore.get%[1]s()\n" +
            "    };\n" +
            "}\n\n" +
            "var %[1]s = React.createClass({\n" +
            common +
            "    render: function() {\n" +
            "        return (\n" +
            "            <h1>%[1]s Component</h1>\n" +
            "        );\n" +
            "    }\n" +
            "});\n\n" +
            "module.exports = %[1]s;\n"
    }

}

func (c *Component) getCommonBodyTemplate() string {
    return "" +
    "    getInitialState: function() {\n" +
    "        return getState();\n" +
    "    },\n\n" +
    "    componentDidMount: function() {\n" +
    "        %[1]sStore.addChangeListener(this._onChange);\n" +
    "    },\n\n" +
    "    componentWillUnmount: function() {\n" +
    "        %[1]sStore.removeChangeListener(this._onChange);\n" +
    "    },\n\n" +
    "    _onChange: function() {\n" +
    "        this.setState(getState());\n" +
    "    },\n\n"
}
