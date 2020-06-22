package rules

import(
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// loadRules loads up the JSON files it needs to find in `datadir`
// TODO: error handling
func load(datadir string) Rules {
	return Rules{
		EquipmentList: loadEquipment(datadir + "/" + "5e-equipment.json"),
		SpellList:     loadSpells   (datadir + "/" + "5e-spells.json"),
		MonsterList:   loadMonsters (datadir + "/" + "5e-monsters.json"),
	}
}

// LoadSpells opens a file that should be a ton of JSON objects that parse into spells
func loadSpells(filename string) SpellList {
	sl := map[string]Spell{}

	if jsonF,err := os.Open(filename); err == nil {
		defer jsonF.Close()

		file, _ := ioutil.ReadAll(jsonF)
		spells := []Spell{}
		json.Unmarshal(file, &spells)

		for _,spell := range spells {
			sl[spell.Index] = spell
		}
		log.Printf("%s, loaded %d objects\n", filename, len(sl))
	} else {
		log.Printf("open %s: %v\n", filename, err)
	}
	
	return sl
}

func loadEquipment(filename string) EquipmentList {
	el := map[string]Item{}

	if jsonF,err := os.Open(filename); err == nil {
		defer jsonF.Close()

		file, _ := ioutil.ReadAll(jsonF)
		items := []Item{}
		json.Unmarshal(file, &items)

		for _,item := range items {
			el[item.Index] = item
		}
		log.Printf("%s, loaded %d objects\n", filename, len(el))
	} else {
		log.Printf("open %s: %v\n", filename, err)
	}
	
	return el
}

func loadMonsters(filename string) MonsterList {
	ml := map[string]Monster{}

	if jsonF,err := os.Open(filename); err == nil {
		defer jsonF.Close()

		file, _ := ioutil.ReadAll(jsonF)
		monsters := []Monster{}
		json.Unmarshal(file, &monsters)

		for _,monster := range monsters {
			ml[monster.Index] = monster
		}
		log.Printf("%s, loaded %d objects\n", filename, len(ml))
	} else {
		log.Printf("open %s: %v\n", filename, err)
	}
	
	return ml
}
