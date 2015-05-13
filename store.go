package main

import (
    "fmt"
    "strings"
)

type Store struct {
    ResourceOptions
}

func (s *Store) GetDir() string {
        return "stores"
}

func (s *Store) GetFileName() string {
    return s.ResourceOptions.Domain + "-store.js"
}

func (s *Store) GetFileContent() string {
    domainFirstCharUpper := strings.ToUpper(s.Domain[:1]) + s.Domain[1:]
    return fmt.Sprintf(s.getTemplate(), domainFirstCharUpper, s.Domain, s.Constant)
}

func (s *Store) getTemplate() string {
    content := s.getHeaderTemplate()
    content = content + s.getBodyVariationTemplate()
    content = content + s.getFooterTemplate()

    return content
}

func (s *Store) getHeaderTemplate() string {
    return "var AppDispatcher = require('../dispatcher/app-dispatcher');\n" +
        "var EventEmitter = require('events').EventEmitter;\n" +
        "var assign = require('object-assign');\n" +
        "var %[1]sConstants = require('../constants/%[2]s-constants');\n" +
        "var CHANGE_EVENT = 'change';\n\n"
}

func (s *Store) getBodyVariationTemplate() string {
    if s.IsApi {
        return "" +
            "var _data = {\n" +
            "    %[2]ss: null\n" +
            "};\n\n" +
            "var %[1]sStore = assign({}, EventEmitter.prototype, {\n" +
            "    getData: function() {\n" +
            "        return _data;\n" +
            "    },\n\n" +
            "    get%[1]ss: function() {\n" +
            "        return _data.%[2]ss;\n" +
            "    },\n\n" +
            "    emitChange: function() {\n" +
            "        this.emit(CHANGE_EVENT);\n" +
            "    },\n\n" +
            "    addChangeListener: function(callback) {\n" +
            "        this.on(CHANGE_EVENT, callback);\n" +
            "    },\n\n" +
            "    removeChangeListener: function(callback) {\n" +
            "        this.removeListener(CHANGE_EVENT, callback);\n" +
            "    }\n" +
            "});\n\n" +
            "AppDispatcher.register(function(action) {\n" +
            "    switch (action.actionType) {\n" +
            "        case %[1]sConstants.%[3]s_SUCCESS:\n" +
            "            _data.%[2]ss = action.%[2]ss;\n" +
            "            %[1]sStore.emitChange();\n" +
            "            break;\n\n" +
            "        default:\n" +
            "            // no op\n" +
            "    }\n" +
            "});\n\n"

    } else {
        return "" +
            "var _data = {\n" +
            "    %[2]s: null\n" +
            "};\n\n" +
            "var %[1]sStore = assign({}, EventEmitter.prototype, {\n" +
            "    getData: function() {\n" +
            "        return _data;\n" +
            "    },\n\n" +
            "    get%[1]s: function() {\n" +
            "        return _data.%[2]s;\n" +
            "    },\n\n" +
            "    emitChange: function() {\n" +
            "        this.emit(CHANGE_EVENT);\n" +
            "    },\n\n" +
            "    addChangeListener: function(callback) {\n" +
            "        this.on(CHANGE_EVENT, callback);\n" +
            "    },\n\n" +
            "    removeChangeListener: function(callback) {\n" +
            "        this.removeListener(CHANGE_EVENT, callback);\n" +
            "    }\n" +
            "});\n\n" +
            "AppDispatcher.register(function(action) {\n" +
            "    switch (action.actionType) {\n" +
            "        case %[1]sConstants.%[3]s:\n" +
            "            _data.%[2]s = action.%[2]s;\n" +
            "            %[1]sStore.emitChange();\n" +
            "            break;\n\n" +
            "        default:\n" +
            "            // no op\n" +
            "    }\n" +
            "});\n\n"
    }
}

func (s *Store) getFooterTemplate() string {
    return "module.exports = %[1]sStore;"
}
