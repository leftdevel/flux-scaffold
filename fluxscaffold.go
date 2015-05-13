package main

import (
    "fmt"
    "flag"
    "os"
)

func main() {
    var domain, constant, action string
    var isApi bool

    flag.StringVar(&domain, "domain", "", "A domain object name e.g user, product");
    flag.StringVar(&constant, "const", "", "A constant name, eg. USERS_FETCH");
    flag.StringVar(&action, "action", "", "An action name, fetchUsers");
    flag.BoolVar(&isApi, "api", false, "If enabled, will create an api file and import it in the action-creator file")

    flag.Parse();

    if (domain == "" || constant == "" || action == "") {
        fmt.Println("domain, constant and action are required \n");
        os.Exit(1);
    }

    resourceOptions := ResourceOptions{Domain: domain, Constant: constant, Action: action, IsApi: isApi}
    constantGenerator := ResourceGenerator{&Constant{ResourceOptions(resourceOptions)}}
    actionGenerator := ResourceGenerator{&Action{ResourceOptions(resourceOptions)}}
    storeGenerator := ResourceGenerator{&Store{ResourceOptions(resourceOptions)}}

    constantGenerator.Execute()
    actionGenerator.Execute()
    storeGenerator.Execute()

    if (isApi) {
        apiGenerator := ResourceGenerator{&Api{ResourceOptions(resourceOptions)}}
        apiGenerator.Execute()
    }
}
