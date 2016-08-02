package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dmmcquay/sqrl"
	"github.com/spf13/cobra"
)

func main() {
	var verbose bool

	var cpu = &cobra.Command{
		Use:   "cpu",
		Short: "displays only cpu info",
		Long:  `Will only display the information outputted by the cpu function`,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := sqrl.GetCPUInfo()
			if err != nil {
				log.Fatal(err)
			}
			c.LogWriter()
			if verbose {
				fmt.Println(c)
			}
		},
	}

	var ros = &cobra.Command{
		Use:   "ros",
		Short: "displays ram, os, and swap memory info",
		Long: `Will only display the information outputted by the ram, os, and swap
			memory function`,
		Run: func(cmd *cobra.Command, args []string) {
			r, o, s, err := sqrl.RamOSInfo()
			if err != nil {
				log.Fatal(err)
			}
			r.LogWriter()
			o.LogWriter()
			s.LogWriter()
			if verbose {
				fmt.Println(r)
				fmt.Println(s)
				fmt.Println(o)
			}
		},
	}

	var net = &cobra.Command{
		Use:   "net",
		Short: "displays network info",
		Long:  `Will only display the information outputted by the network function`,
		Run: func(cmd *cobra.Command, args []string) {
			i, err := sqrl.GetInterfaces()
			if err != nil {
				log.Fatal(err)
			}
			if verbose {
				fmt.Println(i)
			}
		},
	}

	var all = &cobra.Command{
		Use:   "all",
		Short: "displays all info",
		Long: `Will display all the information from the cpu, ram, os, swap memory,
			and network functions`,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := sqrl.GetCPUInfo()
			if err != nil {
				log.Fatal(err)
			}
			c.LogWriter()
			r, o, s, err := sqrl.RamOSInfo()
			if err != nil {
				log.Fatal(err)
			}
			r.LogWriter()
			o.LogWriter()
			s.LogWriter()
			i, err := sqrl.GetInterfaces()
			if err != nil {
				log.Fatal(err)
			}
			if verbose {
				fmt.Println(c)
				fmt.Println(r)
				fmt.Println(o)
				fmt.Println(s)
				fmt.Println(i)
			}
		},
	}

	var las = &cobra.Command{
		Use:   "las",
		Short: "starts listen and serve functionality",
		Long:  `will start a listen and serve function on localhost so the user can curl specific commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			host := "localhost"
			port := 8888

			sm := http.NewServeMux()

			sqrl.AddRoutes(sm)
			log.Printf("serving at: http://%s:%d/", host, port)

			addr := fmt.Sprintf("%s:%d", host, port)
			err := http.ListenAndServe(addr, sm)
			if err != nil {
				log.Printf("%+v", err)
				os.Exit(1)
			}
		},
	}

	cpu.Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"print to stdout the information that would be logged",
	)

	ros.Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"print to stdout the information that would be logged",
	)
	net.Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"print to stdout the information that would be logged",
	)

	all.Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"print to stdout the information that would be logged",
	)

	las.Flags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"print to stdout the information that would be logged",
	)
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cpu, ros, net, all, las)
	rootCmd.Execute()
}
