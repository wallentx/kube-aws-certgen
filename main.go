package main

import (
    "github.com/wallentx/kube-aws-certgen/cmd"
    "os"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        os.Exit(2)
    }
}