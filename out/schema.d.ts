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
		
			export class Loot {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Loot, "serialize" | "deserialize">>);

									baseChance?: number;
									modifiers?: Record<string, number>;
									entries?: LootEntry[];
				
				deserialize(bytes: Uint8Array): Loot;
				serialize(): Uint8Array;
			}
		
			export class World {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<World, "serialize" | "deserialize">>);

									lootTables?: Record<string, Loot>;
									worldName?: string;
									seed?: bigint;
									gravity?: number;
									players?: Character[];
									activeQuests?: Quest[];
									zoneData?: Record<string, Record<string, number>>;
									systemFlags?: Record<string, boolean>;
				
				deserialize(bytes: Uint8Array): World;
				serialize(): Uint8Array;
			}
		
			export class Character {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Character, "serialize" | "deserialize">>);

									companions?: Record<number, Companion>;
									name?: string;
									position?: Vector3;
									equipment?: Record<string, EquipmentSlot>;
									friends?: Character[];
									arbitraryData?: Record<string, Record<string, Record<string, number>>>;
									inventory?: Item[][];
									skillProgress?: Record<string, number[]>;
									id?: bigint;
									stats?: Stats;
				
				deserialize(bytes: Uint8Array): Character;
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
		
			export class Companion {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Companion, "serialize" | "deserialize">>);

									bond?: number;
									name?: string;
									level?: number;
				
				deserialize(bytes: Uint8Array): Companion;
				serialize(): Uint8Array;
			}
		
			export class EquipmentSlot {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<EquipmentSlot, "serialize" | "deserialize">>);

									item?: Item;
									slotName?: string;
				
				deserialize(bytes: Uint8Array): EquipmentSlot;
				serialize(): Uint8Array;
			}
		
			export class Quest {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Quest, "serialize" | "deserialize">>);

									title?: string;
									rewards?: Item[];
									areaLayers?: number[][][];
									description?: string;
									objectives?: Record<string, string[]>;
									prerequisites?: Quest[];
									id?: number;
									nextSteps?: Record<string, Quest>;
									difficulty?: number;
									requiredPos?: Vector3[];
				
				deserialize(bytes: Uint8Array): Quest;
				serialize(): Uint8Array;
			}
		
			export class Item {
				static readonly TypeID: number;

				constructor(fields?: Partial<Omit<Item, "serialize" | "deserialize">>);

									tags?: string[];
									extraData?: Record<string, string>;
									id?: number;
									name?: string;
									rarity?: number;
									weight?: number;
									isQuestItem?: boolean;
				
				deserialize(bytes: Uint8Array): Item;
				serialize(): Uint8Array;
			}
		
