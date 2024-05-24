package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Interface struct {
	Device string `yaml:"device"`
	IP     string `yaml:"ip"`
	Mask   int    `yaml:"mask"`
}

type Config struct {
	Interfaces []struct {
		Interface Interface `yaml:"interface"`
	} `yaml:"-"`
}

/*type RoutingTable struct {

}*/

func main() {
	// Lire le fichier YAML
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier: %v", err)
	}

	// Définir une variable pour stocker la configuration
	var config Config

	// Décoder le fichier YAML en la structure
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Erreur lors de l'analyse du fichier YAML: %v", err)
	}

	// Afficher la configuration
	for _, item := range config.Interfaces {
		fmt.Printf("Device: %s, IP: %s, Mask: %d\n",
			item.Interface.Device, item.Interface.IP, item.Interface.Mask)
	}

	
}
