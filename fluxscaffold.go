package main

import (
    "fmt"
    "flag"
    "os"
)

var domain, constant, action string
var isApi bool
var resourceOptions ResourceOptions
var generators []ResourceGenerator

func main() {
    parseInputOrExit()
    createResourceOptions()
    createGenerators()
    executeGenerators()
}

func parseInputOrExit() {
    flag.StringVar(&domain, "domain", "", "A domain object name e.g user, product");
    flag.StringVar(&constant, "const", "", "A constant name, eg. USERS_FETCH");
    flag.StringVar(&action, "action", "", "An action name, fetchUsers");
    flag.BoolVar(&isApi, "api", false, "If enabled, will create an api file and import it in the action-creator file")

    flag.Parse();

    if (domain == "" || constant == "" || action == "") {
        fmt.Println("domain, constant and action are required \n");
        os.Exit(1);
    }
}

func createResourceOptions() {
    resourceOptions = ResourceOptions{Domain: domain, Constant: constant, Action: action, IsApi: isApi}
}

func createGenerators() {
    constantGenerator := ResourceGenerator{&Constant{ResourceOptions(resourceOptions)}}
    actionGenerator := ResourceGenerator{&Action{ResourceOptions(resourceOptions)}}
    storeGenerator := ResourceGenerator{&Store{ResourceOptions(resourceOptions)}}

    generators = append(generators, constantGenerator)
    generators = append(generators, actionGenerator)
    generators = append(generators, storeGenerator)

    if (isApi) {
        apiGenerator := ResourceGenerator{&Api{ResourceOptions(resourceOptions)}}
        generators = append(generators, apiGenerator)
    }
}

func executeGenerators() {
    race := make(chan bool)

    for _, generator := range generators {
        go generator.Execute(race)
        <- race
    }
}
