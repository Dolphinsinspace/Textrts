package main

import ("fmt"
		"time")

var cityDescriptions = map[string]string{		
	"Billen":"Distant outpost near a river. Closest cities are Volomia and Fospacia across river.",
	"Coulder":"This town brought to you by the letter C!",
	"Dalebba":"Surrounded by mountains, plains, and rivers. Dalebba has a road to Owlder.",
	"Enver":"Central town. Wasteland of empty buildings. Road north to Coulder, East to Lochlo, South to Zillson, and it's possible to cross a river southwest to Dalebba.",
	"Fospacia":"In a secluded plain surounded by hills, mountains, and rivers. North to Lochlo and south to Volomia.",
	"Highlands Ranch":"Though not near any noteworthy landforms Highlands Ranch conects to: Volomia, Billen, Fospacia, Wess, Owld, and Dalebba.",
	"Jade":"High in the Jade Mountains. Roads to Enver and Lochlo, and a path to Fospacia.",
	"Lochlo":"Directly east of Enver in the foothills of the Jade Mountains. Lots of good hiding spaces around here.",
	"Owlder":"Across the river from Wess, and southwest of Dalebba. One of the first cities to spring back after the 8th extinction.",
	"Samburg":"In the western region south of Tareville.",
	"Tareville":"On the road between Coulder to the east and Samburg to the south.",
	"Volomia":"Between Highlands Ranch and Fospacia. Legend has it that terrible things happened here.",
	"Wess":"By a river in the southwest reaches. Owlder is across the river, and west is Highlands Ranch.",
	"Yetistone":"Directly east of Coulder. There were never any Yetis spotted around here. Sounds like somthing a Yeti would say!",
	"Zillson":"South of Enver between a river and hills. Beautiful views used to attract vacationers, and there are giant abandoned resorts around the area. Further south is Highlands Ranch.",
}

func main() {
	var input string
	location := "Enver"
	var description string
	for input !="quit" {
		fmt.Println("Send Dummy to mountain, forest, or river:")
		fmt.Scanln(&input)
		var ok bool
		if input == "look"{
			fmt.Printf("%v\n",description)
		}
		if description, ok = cityDescriptions[input]; ok {
			location = input
		}
		ptimer := time.NewTimer(1 * time.Second)
		<-ptimer.C
		fmt.Printf("Dummy has made it to %v.",location)
	}
}

