package main

import (
	"github.com/KlemenGod/VPSA-vaja6-moduli/redovalnica"
	"fmt"
	"github.com/urfave/cli/v3"
	"context"
	"os"
	"log"
)


func main(){

	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Vpis ocen in izracun statistike",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Min stevilo ocen za pozitivno",
				Value: 3,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjsa mozna ocena",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Najvecja mozna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	studenti := make(map[string]redovalnica.Student)
	studenti["63230081"] = redovalnica.NewStudent("kl", "go", []int{6, 7, 8})
	studenti["63230082"] = redovalnica.NewStudent("kl2", "go2", []int{7, 7, 8})
	studenti["63230083"] = redovalnica.NewStudent("kl3", "go3", []int{6, 5, 9})
	studenti["63230084"] = redovalnica.NewStudent("kl3", "go3", []int{9, 10, 9})
	studenti["63230085"] = redovalnica.NewStudent("kl3", "go3", []int{6, 5, 5})
	studenti["63230086"] = redovalnica.NewStudent("kl3", "go3", []int{7, 8, 9})


	redovalnica.DodajOceno(studenti,"63230081",6)


	redovalnica.IzpisRedovalnice(studenti)

	fmt.Println()

	redovalnica.IzpisiKoncniUspeh(studenti)
}


func nekej(st int){
	fmt.Println("st ocen")
}