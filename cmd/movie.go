package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
	"reflect"
)


// movieCmd represents the movie command
var movieCmd = &cobra.Command{
	Use:   "movie",
	Short: "Shows Movie Information",
	Long: `Shows various information of movies like cast, budget, ratings...
	Usage :
	moviegenie movie [content-title]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting Movie Information...")
		content := args[0]
		var finalurl string
		finalurl = "http://www.omdbapi.com/?apikey=*****" + content

		response, err := http.Get(finalurl)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var data map[string]interface{}
		err = json.Unmarshal([]byte(responseData), &data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		//ratings := data["Ratings"].([]interface{})
		for key, value := range data {

			if reflect.TypeOf(data[key]).Kind() == reflect.Slice {
				fmt.Println(key)
				for _, a := range value.([]interface{}) {
					value2 := a.(map[string]interface{})
					fmt.Printf("\t")
					for _, value3 := range value2{
						fmt.Printf("%v\t",value3)
					}
					fmt.Println()
				}
			} else {
				fmt.Printf("%s :%v\n",key,value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(movieCmd)
}
