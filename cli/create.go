package cli

import (
	"fmt"
)

func (instance *Cli) Create() {
	promps := map[string]string{
		"Nom": "",
		"Adresse email": "",
		"Adresse physique": "",
		"Numéro de téléphone": "",
	}

	println("Ajout d'un nouveau contact :")
	for field, _ := range promps {
		println(field, ":")
		line, err := instance.Reader.Readline()
		if err != nil {
			panic(err)
		}
		promps[field] = line
	}

	params := map[string]string{
		"Name": promps["Nom"],
		"Email": promps["Adresse email"],
		"Address": promps["Adresse physique"],
		"Phone": promps["Numéro de téléphone"],
	}
	instance.Book.CreateContact(params)
	fmt.Println(instance.Book.Contacts)

	println("Contact ajouté")
}
