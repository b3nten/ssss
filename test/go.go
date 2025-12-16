package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	schema "github.com/b3nten/ssss/out"
)

func ptr[T any](val T) *T {
	return &val
}

func main() {
	v1 := &schema.Vector3{
		X: ptr(float64(1)),
		Y: ptr(float64(2)),
		Z: ptr(float64(3)),
	}

	stats1 := &schema.Stats{
		Health:     ptr(int32(100)),
		Mana:       ptr(int32(50)),
		Stamina:    ptr(int32(200)),
		CritChance: ptr(0.25),
		CritDamage: ptr(1.5),
		Resistances: &map[string]int16{
			"air":  1,
			"fire": 2,
		},
	}

	item1 := schema.Item{
		Id:          ptr(uint32(9487329)),
		Name:        ptr("Item One"),
		Rarity:      ptr(uint8(255)),
		Weight:      ptr(12.5),
		IsQuestItem: ptr(false),
		Tags:        &[]string{"sharp", "shiny"},
		ExtraData: &map[string]string{
			"origin": "dungeon",
		},
	}

	item2 := schema.Item{
		Id:          ptr(uint32(239487)),
		Name:        ptr("Item Two"),
		Rarity:      ptr(uint8(5)),
		Weight:      ptr(3.2),
		IsQuestItem: ptr(true),
		Tags:        &[]string{"fragile"},
		ExtraData: &map[string]string{
			"origin": "castle",
		},
	}

	equipSlot1 := schema.EquipmentSlot{
		SlotName: ptr("head"),
		Item:     &item1,
	}

	companion1 := schema.Companion{
		Name:  ptr("Fido"),
		Level: ptr(uint8(5)),
		Bond:  ptr(75.0),
	}

	companion2 := schema.Companion{
		Name:  ptr("Whiskers"),
		Level: ptr(uint8(3)),
		Bond:  ptr(50.0),
	}

	lootEntry1 := schema.LootEntry{
		ItemId: ptr(uint32(348384)),
		MinQty: ptr(uint8(0)),
		MaxQty: ptr(uint8(3)),
		Conditions: &map[string]bool{
			"damaged": false,
		},
	}

	lootEntry2 := schema.LootEntry{
		ItemId: ptr(uint32(3748)),
		MinQty: ptr(uint8(129)),
		MaxQty: ptr(uint8(190)),
		Conditions: &map[string]bool{
			"wet": true,
		},
	}

	loot1 := schema.Loot{
		BaseChance: ptr(0.5),
		Modifiers: &map[string]float64{
			"luck": 0.1,
		},
		Entries: &[]schema.LootEntry{lootEntry1, lootEntry2},
	}

	char1 := schema.Character{
		Id:       ptr(uint64(1234567890)),
		Name:     ptr("Hero"),
		Position: v1,
		Stats:    stats1,
		Equipment: &map[string]schema.EquipmentSlot{
			"backpack": equipSlot1,
		},
		Companions: &map[uint16]schema.Companion{
			1: companion1,
			5: companion2,
		},
		Friends: &[]schema.Character{},
		SkillProgress: &map[string][]float64{
			"archery":  {.3, .5, .6},
			"sneaking": {.1, .2},
		},
		ArbitraryData: &map[string]map[string]map[string]int8{
			"foo": {
				"bar": {
					"baz": 1,
				},
			},
			"a": {
				"b": {
					"c": 127,
				},
			},
		},
	}

	char2 := schema.Character{
		Id:   ptr(uint64(9876543210)),
		Name: ptr("Sidekick"),
		Position: &schema.Vector3{
			X: ptr(float64(-5)),
			Y: ptr(float64(0)),
			Z: ptr(float64(10)),
		},
		Stats: &schema.Stats{
			Health:     ptr(int32(80)),
			Mana:       ptr(int32(30)),
			Stamina:    ptr(int32(150)),
			CritChance: ptr(.15),
			CritDamage: ptr(1.2),
			Resistances: &map[string]int16{
				"earth": 3,
				"water": 1,
			},
		},
		Equipment:  &map[string]schema.EquipmentSlot{},
		Companions: &map[uint16]schema.Companion{},
		Friends:    &[]schema.Character{char1},
		SkillProgress: &map[string][]float64{
			"magic": {.4, .6},
		},
		ArbitraryData: &map[string]map[string]map[string]int8{
			"b": {},
		},
	}

	prereq1 := schema.Quest{
		Id:          ptr(uint32(1234567)),
		Title:       ptr("Gather Information"),
		Description: ptr("Talk to the village elder to learn more about the goblin threat."),
		Difficulty:  ptr(uint8(2)),
		RequiredPos: &[]schema.Vector3{
			{
				X: ptr(float64(5)),
				Y: ptr(float64(0)),
				Z: ptr(float64(0)),
			},
		},
		Objectives: &map[string][]string{
			"main": {"Speak with the village elder"},
			"side": {},
		},
	}

	prereq2 := schema.Quest{
		Id:          ptr(uint32(2345678)),
		Title:       ptr("Prepare Defenses"),
		Description: ptr("Help the villagers set up defenses against the goblin attack."),
		Difficulty:  ptr(uint8(3)),
		RequiredPos: &[]schema.Vector3{
			{
				X: ptr(float64(8)),
				Y: ptr(float64(0)),
				Z: ptr(float64(-3)),
			},
		},
		Objectives: &map[string][]string{
			"main": {"Build barricades", "Train the villagers"},
			"side": {},
		},
	}

	nextQuest := schema.Quest{
		Id:          ptr(uint32(8765432)),
		Title:       ptr("Report Back"),
		Description: ptr("Go back to the village elder and inform him of your success."),
		Difficulty:  ptr(uint8(1)),
		RequiredPos: &[]schema.Vector3{
			{
				X: ptr(float64(0)),
				Y: ptr(float64(0)),
				Z: ptr(float64(0)),
			},
		},
		Objectives: &map[string][]string{
			"main": {"Speak with the village elder"},
			"side": {},
		},
	}

	quest := schema.Quest{
		Id:          ptr(uint32(7748392)),
		Title:       ptr("Save the Village"),
		Description: ptr("Help the villagers fend off the goblin attack. This quest is dangerious, mainly because the goblins are not very nice and will do mean things to you if you dont fight them off. However they are also a bit misunderstood and lonely so physical force isn't the only way you could fight them offf. Good luck my friend."),
		Difficulty:  ptr(uint8(4)),
		Rewards:     &[]schema.Item{item1, item2},
		RequiredPos: &[]schema.Vector3{
			{
				X: ptr(float64(10)),
				Y: ptr(float64(0)),
				Z: ptr(float64(-5)),
			},
			{
				X: ptr(float64(15)),
				Y: ptr(float64(0)),
				Z: ptr(float64(-10)),
			},
		},
		Objectives: &map[string][]string{
			"main": {"Defeat the goblin leader", "Protect the villagers"},
			"side": {"Find the lost amulet", "Collect 10 healing herbs"},
		},
		Prerequisites: &[]schema.Quest{prereq1, prereq2},
		NextSteps: &map[string]schema.Quest{
			"Return to the village elder for your reward": nextQuest,
		},
		AreaLayers: &[][][]uint16{
			{{1, 2, 3}, {4, 5, 6}},
			{{7, 8, 9}},
		},
	}

	dragonLoot := schema.Loot{
		BaseChance: ptr(0.2),
		Modifiers: &map[string]float64{
			"luck": 0.05,
		},
		Entries: &[]schema.LootEntry{
			{
				ItemId: ptr(uint32(999999)),
				MinQty: ptr(uint8(1)),
				MaxQty: ptr(uint8(1)),
				Conditions: &map[string]bool{
					"legendary": true,
				},
			},
		},
	}

	world := &schema.World{
		WorldName:    ptr("Fantasy Land"),
		Seed:         ptr(uint64(12345678987654)),
		Gravity:      ptr(-9.89),
		Players:      &[]schema.Character{char1, char2},
		ActiveQuests: &[]schema.Quest{quest},
		ZoneData: &map[string]map[string]uint32{
			"forest": {
				"hitRates": 2,
			},
			"desert": {
				"hitRates": 1,
			},
		},
		SystemFlags: &map[string]bool{
			"isPVPEnabled":   true,
			"isHardcoreMode": false,
		},
		LootTables: &map[string]schema.Loot{
			"goblinCamp": loot1,
			"dragonLair": dragonLoot,
		},
	}

	serialized, err := schema.MarshalBytes(world)
	if err != nil {
		panic(err)
	}
	newWorld := &schema.World{}
	err = schema.UnmarshalBytes(serialized, newWorld)
	if err != nil {
		panic(err)
	}

	result := cmp.Diff(world, newWorld)
	fmt.Println("Diff:", result)
}
