version = 1

-- forward declare nested structs
Character = struct {}
Quest = struct {}

Vector3 = struct {
    x = f64[0],
    y = f64[1],
    z = f64[2]
}

Stats = struct {
    health      = int32[0],
    mana        = int32[1],
    stamina     = int32[2],
    critChance  = f64[3],
    critDamage  = f64[4],
    resistances = map(str, int16)[5]
}

Item = struct {
    id          = uint32[0],
    name        = str[1],
    rarity      = uint8[2],
    weight      = f64[3],
    isQuestItem = bool[4],
    tags        = list(str)[5],
    extraData   = map(str, str)[6]
}

EquipmentSlot = struct {
    slotName  = str[0],
    item      = Item[1]
}

Companion = struct {
    name   = str[0],
    level  = uint8[1],
    bond   = f64[2]
}

LootEntry = struct {
    itemId     = uint32[0],
    minQty     = uint8[1],
    maxQty     = uint8[2],
    conditions = map(str, bool)[3]
}

Loot = struct {
    baseChance   = f64[0],
    modifiers    = map(str, f64)[1],
    entries      = list(LootEntry)[2]
}

Character = struct {
    id             = uint64[0],
    name           = str[1],
    position       = Vector3[2],
    stats          = Stats[3],

    inventory      = list(list(Item))[4],

    equipment      = map(str, EquipmentSlot)[5],

    companions     = map(uint16, Companion)[6],

    friends        = list(Character)[7],

    skillProgress  = map(str, list(f64))[8],

    arbitraryData  = map(str, map(str, map(str, int8)))[9]
}

Quest = struct {
    id           = uint32[0],
    title        = str[1],
    difficulty   = uint8[2],
    rewards      = list(Item)[3],
    requiredPos  = list(Vector3)[4],

    objectives   = map(str, list(str))[5],

    prerequisites = list(Quest)[6],
    nextSteps     = map(str, Quest)[7],

    areaLayers    = list(list(list(uint16)))[8],

    description   = str[9]
}

World = struct {
    worldName     = str[0],
    seed          = uint64[1],
    gravity       = f64[2],
    players       = list(Character)[3],
    activeQuests  = list(Quest)[4],

    zoneData      = map(str, map(str, uint32))[5],

    systemFlags   = map(str, bool)[6],

    lootTables    = map(str, Loot)[7]
}
