			export class Character {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Character, "serialize" | "deserialize">>);

									name?: string;
									position?: Vector3;
									id?: bigint;
									equipment?: Record<string, EquipmentSlot>;
									skillProgress?: Record<string, number[]>;
									stats?: Stats;
									arbitraryData?: Record<string, Record<string, Record<string, number>>>;
									inventory?: Item[][];
									companions?: Record<number, Companion>;
									friends?: Character[];
				
				deserialize(bytes: Uint8Array): Character;
				serialize(): Uint8Array;
			}
		
			export class World {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<World, "serialize" | "deserialize">>);

									activeQuests?: Quest[];
									zoneData?: Record<string, Record<string, number>>;
									systemFlags?: Record<string, boolean>;
									lootTables?: Record<string, Loot>;
									worldName?: string;
									seed?: bigint;
									gravity?: number;
									players?: Character[];
				
				deserialize(bytes: Uint8Array): World;
				serialize(): Uint8Array;
			}
		
			export class Companion {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Companion, "serialize" | "deserialize">>);

									name?: string;
									level?: number;
									bond?: number;
				
				deserialize(bytes: Uint8Array): Companion;
				serialize(): Uint8Array;
			}
		
			export class Vector3 {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Vector3, "serialize" | "deserialize">>);

									x?: number;
									y?: number;
									z?: number;
				
				deserialize(bytes: Uint8Array): Vector3;
				serialize(): Uint8Array;
			}
		
			export class LootEntry {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<LootEntry, "serialize" | "deserialize">>);

									itemId?: number;
									minQty?: number;
									maxQty?: number;
									conditions?: Record<string, boolean>;
				
				deserialize(bytes: Uint8Array): LootEntry;
				serialize(): Uint8Array;
			}
		
			export class Stats {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Stats, "serialize" | "deserialize">>);

									health?: number;
									mana?: number;
									stamina?: number;
									critChance?: number;
									critDamage?: number;
									resistances?: Record<string, number>;
				
				deserialize(bytes: Uint8Array): Stats;
				serialize(): Uint8Array;
			}
		
			export class Quest {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Quest, "serialize" | "deserialize">>);

									id?: number;
									difficulty?: number;
									requiredPos?: Vector3[];
									description?: string;
									title?: string;
									objectives?: Record<string, string[]>;
									nextSteps?: Record<string, Quest>;
									areaLayers?: number[][][];
									rewards?: Item[];
									prerequisites?: Quest[];
				
				deserialize(bytes: Uint8Array): Quest;
				serialize(): Uint8Array;
			}
		
			export class Loot {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Loot, "serialize" | "deserialize">>);

									baseChance?: number;
									modifiers?: Record<string, number>;
									entries?: LootEntry[];
				
				deserialize(bytes: Uint8Array): Loot;
				serialize(): Uint8Array;
			}
		
			export class EquipmentSlot {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<EquipmentSlot, "serialize" | "deserialize">>);

									slotName?: string;
									item?: Item;
				
				deserialize(bytes: Uint8Array): EquipmentSlot;
				serialize(): Uint8Array;
			}
		
			export class Item {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Item, "serialize" | "deserialize">>);

									extraData?: Record<string, string>;
									id?: number;
									name?: string;
									rarity?: number;
									weight?: number;
									isQuestItem?: boolean;
									tags?: string[];
				
				deserialize(bytes: Uint8Array): Item;
				serialize(): Uint8Array;
			}
		
