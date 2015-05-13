package main

import (
    "fmt"
    "strings"
)

type Api struct {
    ResourceOptions
}

func (s *Api) GetDir() string {
    return "webapi"
}

func (s *Api) GetFileName() string {
    return s.ResourceOptions.Domain + "-api.js"
}

func (s *Api) GetFileContent() string {
    domainFirstCharUpper := strings.ToUpper(s.Domain[:1]) + s.Domain[1:]
    return fmt.Sprintf(s.getTemplate(), domainFirstCharUpper, s.Action, s.Domain)
}

func (s *Api) getTemplate() string {
    return "var request = require('superagent');\n" +
        "var ApiResponseHandler = require('./api-response-handler');\n\n" +
        "var %[1]sApi = {\n" +
        "    %[2]s: function(successCallback) {\n" +
        "        var url = Routing.generate('%[3]s_list');\n\n" +
        "        request\n" +
        "            .get(url)\n" +
        "            .set('Accept', 'application/json')\n" +
        "            .end(function(err, res) {\n" +
        "                ApiResponseHandler.handle(err, res, successCallback);\n" +
        "            })\n" +
        "        ;\n" +
        "    }\n" +
        "};\n\n" +

        "module.exports = %[1]sApi;"
}
