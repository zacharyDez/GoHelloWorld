package main

import (
    "fmt"
    "os"

    "example.com/greetings"
)

// flatten string array to spaced string
func flattenStringArray(strArray []string) string {
    var name string = ""
    for i := 0; i < len(strArray); i++ {
        name +=  strArray[i]
        if i != len(strArray) - 1 { name+= " " }
    }

    return name
}

// get a meeting message and print it
func main() {
    // can't have name for greeting be longer than 5
    var nameArgs []string
    
    if len(os.Args) == 1 {
        nameArgs = append(nameArgs, "World")
    } else if len(os.Args)<=6 {
        nameArgs = os.Args[1:]
    } else {
        fmt.Println("You've entered too many arguments.")
        return
    }

    var name string  = flattenStringArray(nameArgs)
    
    message := greetings.Hello(name)
    fmt.Println(message)
}