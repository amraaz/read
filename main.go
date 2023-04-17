package main 

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
)

var (
	destination,err = os.Getwd()

	destinationbool = false

)


func ercatch(err error){
	if (err!=nil){
		fmt.Println("Error",err)
	}
}


func main(){

	app:= &cli.App{
		Name:"read",
		Commands: []*cli.Command{
			{
				Name: "dest",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "this",
						Value: destinationbool,
						Action: func(ctx *cli.Context, b bool) error {
							fmt.Println(destination)
							return nil
						},
					},
				},
			},
			{},
		},
	}

	ercatch(app.Run(os.Args))
}
