// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/osrg/gobgp/api"
	"github.com/spf13/cobra"
)

var globalOpts struct {
	Host         string `short:"u" long:"url" description:"specifying an url" default:"127.0.0.1"`
	Port         int    `short:"p" long:"port" description:"specifying a port" default:"8080"`
	Debug        bool   `short:"d" long:"debug" description:"use debug"`
	Quiet        bool   `short:"q" long:"quiet" description:"use quiet"`
	Json         bool   `short:"j" long:"json" description:"use json format to output format"`
	GenCmpl      bool   `short:"c" long:"genbashcmpl" description:"use json format to output format"`
	BashCmplFile string
}

var cmds []string
var client api.GrpcClient

func main() {
	rootCmd := &cobra.Command{
		Use: "gobgp",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !globalOpts.GenCmpl {
				conn := connGrpc()
				client = api.NewGrpcClient(conn)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.GenBashCompletionFile(globalOpts.BashCmplFile)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&globalOpts.Host, "host", "u", "127.0.0.1", "host")
	rootCmd.PersistentFlags().IntVarP(&globalOpts.Port, "port", "p", 8080, "port")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.Json, "json", "j", false, "use json format to output format")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.Debug, "debug", "d", false, "use debug")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.Quiet, "quiet", "q", false, "use quiet")
	rootCmd.PersistentFlags().BoolVarP(&globalOpts.GenCmpl, "gen-cmpl", "c", false, "generate completion file")
	rootCmd.PersistentFlags().StringVarP(&globalOpts.BashCmplFile, "bash-cmpl-file", "", "gobgp_completion.bash", "bash cmpl filename")

	globalCmd := NewGlobalCmd()
	neighborCmd := NewNeighborCmd()
	policyCmd := NewPolicyCmd()
	monitorCmd := NewMonitorCmd()
	rootCmd.AddCommand(globalCmd)
	rootCmd.AddCommand(neighborCmd)
	rootCmd.AddCommand(policyCmd)
	rootCmd.AddCommand(monitorCmd)
	rootCmd.Execute()
}
