package cmd

import (
    "github.com/spf13/cobra"
)

var (
    RootCmd = &cobra.Command{
        Use:   "kube-aws-certgen",
        Short: "Tool for creating private TLS assets for kube-aws production cluster",
        Long:  ``,
    }

    configPath = "cluster.yaml"
)