using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Schema;
using KellermanSoftware.CompareNetObjects;

class Program
{
    static void Main(string[] args)
    {
        var v1 = new Vector3
        {
            X = 1,
            Y = 2,
            Z = 3
        };

        var stats1 = new Stats
        {
            Health = 100,
            Mana = 50,
            Stamina = 200,
            CritChance = 0.25,
            CritDamage = 1.5,
            Resistances = new Dictionary<string, short>
            {
                { "air", 1 },
                { "fire", 2 }
            }
        };

        var item1 = new Item
        {
            Id = 9487329,
            Name = "Item One",
            Rarity = 255,
            Weight = 12.5,
            IsQuestItem = false,
            Tags = new List<string> { "sharp", "shiny" },
            ExtraData = new Dictionary<string, string>
            {
                { "origin", "dungeon" }
            }
        };

        var item2 = new Item
        {
            Id = 239487,
            Name = "Item Two",
            Rarity = 5,
            Weight = 3.2,
            IsQuestItem = true,
            Tags = new List<string> { "fragile" },
            ExtraData = new Dictionary<string, string>
            {
                { "origin", "castle" }
            }
        };

        var equipSlot1 = new EquipmentSlot
        {
            SlotName = "head",
            Item = item1
        };

        var companion1 = new Companion
        {
            Name = "Fido",
            Level = 5,
            Bond = 75.0
        };

        var companion2 = new Companion
        {
            Name = "Whiskers",
            Level = 3,
            Bond = 50.0
        };

        var lootEntry1 = new LootEntry
        {
            ItemId = 348384,
            MinQty = 0,
            MaxQty = 3,
            Conditions = new Dictionary<string, bool>
            {
                { "damaged", false }
            }
        };

        var lootEntry2 = new LootEntry
        {
            ItemId = 3748,
            MinQty = 129,
            MaxQty = 190,
            Conditions = new Dictionary<string, bool>
            {
                { "wet", true }
            }
        };

        var loot1 = new Loot
        {
            BaseChance = 0.5,
            Modifiers = new Dictionary<string, double>
            {
                { "luck", 0.1 }
            },
            Entries = new List<LootEntry> { lootEntry1, lootEntry2 }
        };

        var char1 = new Character
        {
            Id = 1234567890,
            Name = "Hero",
            Position = v1,
            Stats = stats1,
            Equipment = new Dictionary<string, EquipmentSlot>
            {
                { "backpack", equipSlot1 }
            },
            Companions = new Dictionary<ushort, Companion>
            {
                { 1, companion1 },
                { 5, companion2 }
            },
            Friends = new List<Character>(),
            SkillProgress = new Dictionary<string, List<double>>
            {
                { "archery", new List<double> { 0.3, 0.5, 0.6 } },
                { "sneaking", new List<double> { 0.1, 0.2 } }
            },
            ArbitraryData = new Dictionary<string, Dictionary<string, Dictionary<string, sbyte>>>
            {
                {
                    "foo", new Dictionary<string, Dictionary<string, sbyte>>
                    {
                        {
                            "bar", new Dictionary<string, sbyte>
                            {
                                { "baz", 1 }
                            }
                        }
                    }
                },
                {
                    "a", new Dictionary<string, Dictionary<string, sbyte>>
                    {
                        {
                            "b", new Dictionary<string, sbyte>
                            {
                                { "c", 127 }
                            }
                        }
                    }
                }
            }
        };

        var char2 = new Character
        {
            Id = 9876543210,
            Name = "Sidekick",
            Position = new Vector3 { X = -5, Y = 0, Z = 10 },
            Stats = new Stats
            {
                Health = 80,
                Mana = 30,
                Stamina = 150,
                CritChance = 0.15,
                CritDamage = 1.2,
                Resistances = new Dictionary<string, short>
                {
                    { "earth", 3 },
                    { "water", 1 }
                }
            },
            Equipment = new Dictionary<string, EquipmentSlot>(),
            Companions = new Dictionary<ushort, Companion>(),
            Friends = new List<Character> { char1 },
            SkillProgress = new Dictionary<string, List<double>>
            {
                { "magic", new List<double> { 0.4, 0.6 } }
            },
            ArbitraryData = new Dictionary<string, Dictionary<string, Dictionary<string, sbyte>>>
            {
                { "b", new Dictionary<string, Dictionary<string, sbyte>>() }
            }
        };

        var prereq1 = new Quest
        {
            Id = 1234567,
            Title = "Gather Information",
            Description = "Talk to the village elder to learn more about the goblin threat.",
            Difficulty = 2,
            RequiredPos = new List<Vector3>
            {
                new Vector3 { X = 5, Y = 0, Z = 0 }
            },
            Objectives = new Dictionary<string, List<string>>
            {
                { "main", new List<string> { "Speak with the village elder" } },
                { "side", new List<string>() }
            }
        };

        var prereq2 = new Quest
        {
            Id = 2345678,
            Title = "Prepare Defenses",
            Description = "Help the villagers set up defenses against the goblin attack.",
            Difficulty = 3,
            RequiredPos = new List<Vector3>
            {
                new Vector3 { X = 8, Y = 0, Z = -3 }
            },
            Objectives = new Dictionary<string, List<string>>
            {
                { "main", new List<string> { "Build barricades", "Train the villagers" } },
                { "side", new List<string>() }
            }
        };

        var nextQuest = new Quest
        {
            Id = 8765432,
            Title = "Report Back",
            Description = "Go back to the village elder and inform him of your success.",
            Difficulty = 1,
            RequiredPos = new List<Vector3>
            {
                new Vector3 { X = 0, Y = 0, Z = 0 }
            },
            Objectives = new Dictionary<string, List<string>>
            {
                { "main", new List<string> { "Speak with the village elder" } },
                { "side", new List<string>() }
            }
        };

        var quest = new Quest
        {
            Id = 7748392,
            Title = "Save the Village",
            Description = "Help the villagers fend off the goblin attack. This quest is dangerious, mainly because the goblins are not very nice and will do mean things to you if you dont fight them off. However they are also a bit misunderstood and lonely so physical force isn't the only way you could fight them offf. Good luck my friend.",
            Difficulty = 4,
            Rewards = new List<Item> { item1, item2 },
            RequiredPos = new List<Vector3>
            {
                new Vector3 { X = 10, Y = 0, Z = -5 },
                new Vector3 { X = 15, Y = 0, Z = -10 }
            },
            Objectives = new Dictionary<string, List<string>>
            {
                { "main", new List<string> { "Defeat the goblin leader", "Protect the villagers" } },
                { "side", new List<string> { "Find the lost amulet", "Collect 10 healing herbs" } }
            },
            Prerequisites = new List<Quest> { prereq1, prereq2 },
            NextSteps = new Dictionary<string, Quest>
            {
                { "Return to the village elder for your reward", nextQuest }
            },
            AreaLayers = new List<List<List<ushort>>>
            {
                new List<List<ushort>>
                {
                    new List<ushort> { 1, 2, 3 },
                    new List<ushort> { 4, 5, 6 }
                },
                new List<List<ushort>>
                {
                    new List<ushort> { 7, 8, 9 }
                }
            }
        };

        var dragonLoot = new Loot
        {
            BaseChance = 0.2,
            Modifiers = new Dictionary<string, double>
            {
                { "luck", 0.05 }
            },
            Entries = new List<LootEntry>
            {
                new LootEntry
                {
                    ItemId = 999999,
                    MinQty = 1,
                    MaxQty = 1,
                    Conditions = new Dictionary<string, bool>
                    {
                        { "legendary", true }
                    }
                }
            }
        };

        var world = new World
        {
            WorldName = "Fantasy Land",
            Seed = 12345678987654,
            Gravity = -9.89,
            Players = new List<Character> { char1, char2 },
            ActiveQuests = new List<Quest> { quest },
            ZoneData = new Dictionary<string, Dictionary<string, uint>>
            {
                {
                    "forest", new Dictionary<string, uint>
                    {
                        { "hitRates", 2 }
                    }
                },
                {
                    "desert", new Dictionary<string, uint>
                    {
                        { "hitRates", 1 }
                    }
                }
            },
            SystemFlags = new Dictionary<string, bool>
            {
                { "isPVPEnabled", true },
                { "isHardcoreMode", false }
            },
            LootTables = new Dictionary<string, Loot>
            {
                { "goblinCamp", loot1 },
                { "dragonLair", dragonLoot }
            }
        };

        byte[] bytes = world.Serialize();
        var newWorld = new World().Deserialize(bytes);

        File.WriteAllBytes("test/csharp.bin", bytes);

        if (args.Contains("--compare"))
        {
        		CompareLogic compareLogic = new CompareLogic();
            try
            {
                var jsBytes = File.ReadAllBytes("test/js.bin");
                var jsWorld = new World().Deserialize(jsBytes);
                var result = compareLogic.Compare(newWorld, jsWorld);
								if (result.AreEqual) Console.WriteLine("JS -> C# ✅");
								else Console.WriteLine($"JS -> C# 💀 {result.DifferencesString}");
            }
            catch (Exception e)
            {
                Console.WriteLine($"JS -> C# 💀 {e.Message}");
            }

            try
						{
								var goBytes = File.ReadAllBytes("test/go.bin");
								var goWorld = new World().Deserialize(goBytes);
								var result = compareLogic.Compare(newWorld, goWorld);
								if (result.AreEqual) Console.WriteLine("Go -> C# ✅");
								else Console.WriteLine($"Go -> C# 💀 {result.DifferencesString}");
						}
						catch (Exception e)
						{
								Console.WriteLine($"Go -> C# 💀 {e.Message}");
						}
        }
    }
}
