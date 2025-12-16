import * as Schema from "../out/schema"
// @ts-ignore
import { deepStrictEqual } from "node:assert"
// @ts-ignore
import * as fs from "node:fs"
// @ts-ignore
import * as process from "node:process"

const v1 = new Schema.Vector3({ x: 1, y: 2, z: 3 })

const stats1 = new Schema.Stats({
	health: 100,
	mana: 50,
	stamina: 200,
	critChance: .25,
	critDamage: 1.5,
	resistances: {
		air: 1,
		fire: 2,
	}
})

const item1 = new Schema.Item({
	id: 9487329,
	name: "Item One",
	rarity: 255,
	weight: 12.5,
	isQuestItem: false,
	tags: ["sharp", "shiny"],
	extraData: {
		origin: "dungeon",
	}
})

const item2 = new Schema.Item({
	id: 239487,
	name: "Item Two",
	rarity: 5,
	weight: 3.2,
	isQuestItem: true,
	tags: ["fragile"],
	extraData: {
		origin: "castle",
	}
})

const equiptSlot1 = new Schema.EquipmentSlot({
	slotName: "head",
	item: item1,
})

const companion1 = new Schema.Companion({
	name: "Fido",
	level: 5,
	bond: 75,
})

const companion2 = new Schema.Companion({
	name: "Whiskers",
	level: 3,
	bond: 50,
})

const lootentry1 = new Schema.LootEntry({
	itemId: 348384,
	minQty: 0,
	maxQty: 3,
	conditions: {
		damaged: false
	}
})

const lootentry2 = new Schema.LootEntry({
	itemId: 3748,
	minQty: 129,
	maxQty: 190,
	conditions: {
		wet: true
	}
})

const loot1 = new Schema.Loot({
	baseChance: 0.5,
	modifiers: {
		luck: 0.1,
	},
	entries: [lootentry1, lootentry2],
})

const char1 = new Schema.Character({
	id: BigInt(1234567890),
	name: "Hero",
	position: v1,
	stats: stats1,
	equipment: {
		backpack: equiptSlot1,
	},
	companions: {
		1: companion1,
		5: companion2,
	},
	friends: [],
	skillProgress: {
		"archery": [.3, .5, .6],
		"sneaking": [.1, .2],
	},
	arbitraryData: {
		foo: {
			bar: {
				baz: 1
			}
		},
		a: {
			b: {
				c: 127
			}
		}
	}
})

const char2 = new Schema.Character({
	id: BigInt(9876543210),
	name: "Sidekick",
	position: new Schema.Vector3({ x: -5, y: 0, z: 10 }),
	stats: new Schema.Stats({
		health: 80,
		mana: 30,
		stamina: 150,
		critChance: .15,
		critDamage: 1.2,
		resistances: {
			earth: 3,
			water: 1,
		}
	}),
	equipment: {},
	companions: {},
	friends: [char1],
	skillProgress: {
		"magic": [.4, .6],
	},
	arbitraryData: {
		b: {}
	},
})

const quest = new Schema.Quest({
	id: 7748392,
	title: "Save the Village",
	description: "Help the villagers fend off the goblin attack. This quest is dangerious, mainly because the goblins are not very nice and will do mean things to you if you dont fight them off. However they are also a bit misunderstood and lonely so physical force isn't the only way you could fight them offf. Good luck my friend.",
	difficulty: 4,
	rewards: [item1, item2],
	requiredPos: [
		new Schema.Vector3({ x: 10, y: 0, z: -5 }),
		new Schema.Vector3({ x: 15, y: 0, z: -10 })
	],
	objectives: {
		main: ["Defeat the goblin leader", "Protect the villagers"],
		side: ["Find the lost amulet", "Collect 10 healing herbs"]
	},
	prerequisites: [
		new Schema.Quest({
			id: 1234567,
			title: "Gather Information",
			description: "Talk to the village elder to learn more about the goblin threat.",
			difficulty: 2,
			requiredPos: [
				new Schema.Vector3({ x: 5, y: 0, z: 0 })
			],
			objectives: {
				main: ["Speak with the village elder"],
				side: []
			},
		}),
		new Schema.Quest({
			id: 2345678,
			title: "Prepare Defenses",
			description: "Help the villagers set up defenses against the goblin attack.",
			difficulty: 3,
			requiredPos: [
				new Schema.Vector3({ x: 8, y: 0, z: -3 })
			],
			objectives: {
				main: ["Build barricades", "Train the villagers"],
				side: []
			},
		})
	],
	nextSteps: {
		"Return to the village elder for your reward": new Schema.Quest({
			id: 8765432,
			title: "Report Back",
			description: "Go back to the village elder and inform him of your success.",
			difficulty: 1,
			requiredPos: [
				new Schema.Vector3({ x: 0, y: 0, z: 0 })
			],
			objectives: {
				main: ["Speak with the village elder"],
				side: []
			},
		})
	},
	areaLayers: [[[1, 2, 3], [4, 5, 6]], [[7, 8, 9]]],
})

const world = new Schema.World({
	worldName: "Fantasy Land",
	seed: BigInt(12345678987654),
	gravity: -9.89,
	players: [
		char1,
		char2
	],
	activeQuests: [
		quest
	],
	zoneData: {
		"forest": {
			hitRates: 2,
		},
		"desert": {
			hitRates: 1,
		}
	},
	systemFlags: {
		isPVPEnabled: true,
		isHardcoreMode: false,
	},
	lootTables: {
		"goblinCamp": loot1,
		"dragonLair": new Schema.Loot({
			baseChance: 0.2,
			modifiers: {
				luck: 0.05,
			},
			entries: [
				new Schema.LootEntry({
					itemId: 999999,
					minQty: 1,
					maxQty: 1,
					conditions: {
						legendary: true
					}
				})
			]
		})
	}
})

const bytes = world.serialize()
const newWorld = new Schema.World().deserialize(bytes)

try {
	deepStrictEqual(world, newWorld)
}
catch (e) {
	console.log(e)
}

fs.writeFileSync("test/js.bin", bytes)

if (process.argv.includes("--compare")) {
	try {
		const goBytes = fs.readFileSync("test/go.bin")
		const goWorld = new Schema.World().deserialize(goBytes)
		deepStrictEqual(world, goWorld)
		console.log("Go -> JS ✅")
	}
	catch (e) {
		console.log("Go -> JS 💀", e)
	}
	try {
		const cSharpBytes = fs.readFileSync("test/csharp.bin")
		const cSharpWorld = new Schema.World().deserialize(cSharpBytes)
		deepStrictEqual(world, cSharpWorld)
		console.log("C# -> JS ✅")
	}
	catch (e) {
		console.log("C# -> JS 💀", e)
	}
}
