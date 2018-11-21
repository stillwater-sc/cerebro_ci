package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start a Cerebro Service",
	Run: serverRun,
}

var portNr string

func init() {
	serverCmd.Flags().StringVarP(&portNr, "portNr", "p", "9090", "port number the Cerebro server will listen to")
	rootCmd.AddCommand(serverCmd)
}

func work(w http.ResponseWriter, r *http.Request) {
	log.Println("Work function")
	w.Write([]byte("OK"))
}

func serverRun(cmd *cobra.Command, args []string) {
	workHandler := http.HandlerFunc(work)

	http.Handle("/", workHandler)
	if err := http.ListenAndServe(":"+portNr, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
