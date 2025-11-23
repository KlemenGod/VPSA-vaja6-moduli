package redovalnica

import "fmt"

type Student struct {
    ime     string
    priimek string
    ocene   []int
}

var studenti map[string]Student


//NewStudent vrne novo narejenega študenta
func NewStudent(ime string, priimek string, ocene []int) Student{
	return Student{ime, priimek, ocene}
}

//DodajOceno doda oceno studentu z vpisno stevilko
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int){
	stud, ok := studenti[vpisnaStevilka]

	if ok{
		if ocena <= 10 && ocena >= 0{
			stud.ocene = append(stud.ocene, ocena)
			studenti[vpisnaStevilka] = stud
		}
	}else{
		fmt.Println("Študenta ni na seznamu")
	}

}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64{
	stud, ok := studenti[vpisnaStevilka]

	if ok{
		pov := 0.0

		for _,v := range(stud.ocene){
			pov += float64(v)
		}
		
		pov /= float64(len(stud.ocene))

		if pov < 6{
			return 0
		}else{
			return pov
		}
	}else{
		return -1
	}
}

//IzpisiRedovalnico izpiše celotno redovalnico študentov
func IzpisRedovalnice(studenti map[string]Student){
	fmt.Println("REDOVALNICA: ")
	for vpisna ,stud := range(studenti){
		fmt.Printf("%s - %s %s: %v\n",vpisna,stud.ime, stud.priimek,stud.ocene)
	}
}

//IzpisiKoncniUspeh izpiše končni uspeh študenta
func IzpisiKoncniUspeh(studenti map[string]Student){
	for vpisna ,stud := range(studenti){
		uspesnost := ""
		pov := povprecje(studenti,vpisna)

		if pov >= 9{
			uspesnost = "Odličen študent!"
		}else if pov >= 6 && pov <= 9{
			uspesnost = "Povprečen študent"
		}else{
			uspesnost = "Neuspešen študent"
		}

		fmt.Printf("%s %s: povprečna ocena %.1f -> %s\n", stud.ime,stud.priimek,pov,uspesnost)
	}
}