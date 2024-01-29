package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugent-library/marc"
)

func main() {
	convertCmd.Flags().StringVarP(&fromFlag, "from", "f", "", "source format")
	convertCmd.Flags().StringVarP(&toFlag, "to", "t", "", "target format")

	rootCmd.AddCommand(convertCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	fromFlag string
	toFlag   string
)

var rootCmd = &cobra.Command{
	Short: "MARC CLI",
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert MARC records",
	Run: func(cmd *cobra.Command, args []string) {
		dec := marc.NewDecoder(fromFlag, os.Stdin)
		enc := marc.NewEncoder(toFlag, os.Stdout)
		if dec == nil {
			log.Fatalf(`No decoder found for format "%s"`, fromFlag)
		}
		if enc == nil {
			log.Fatalf(`No encoder found for format "%s"`, toFlag)
		}

		for {
			rec, err := dec.Decode()
			if err != nil {
				log.Panic(err)
			}
			if rec == nil {
				break
			}
			err = enc.Encode(rec)
			if err != nil {
				log.Panic(err)
			}
		}
	},
}
