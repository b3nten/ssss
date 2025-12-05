# [S]tupidly [S]imple [S]erialization [S]ystem

SSSS is a cross-language data serialization system designed to be simple to implement and use.

*SSSS is in pre-release and is not yet suitable for production use. Critical features like maps are missing, along with a test suite.*

## Overview

There are many great serialization systems available, such as [Proto3](https://protobuf.dev), [Avro](https://avro.apache.org), [Flatbuffers](https://flatbuffers.dev), and many more.

The goals of SSSS are as follows:

1. Generate single-file artifacts with zero dependencies.
	- There should no reliance on a runtime or external libraries. The tool should generate all required code.
	
2. Schema compatibility is crucial.
	- New fields may be added to structs without breaking compatibility with older versions.
	- If a field does not exist in the schema it is simply ignored. This enables asynchronous deployment of schema changes.
	
3. All values can be nullable.
	- Every field is optional and may be omitted in a payload. This enables schema evolution and compatibility.
	
4. Prefer ease of use over performace.
	- Working with the generated data types should be idiomatic to each language, as if they were hand-written.
	
5. Prefer simplicity over features.
	- SSSS relies on simple data types: numbers strings and bools, structs, lists, and maps.
	
6. Generated code should be human-readable.
	- The generated code should be easy to read and understand, making debugging and maintenance easier.
	
7. Additional languages or modifications to existing codegen should be possible with minimal effort.
	- Codegen for most languages is implemented in Lua, and the CLI can take an arbitrary Lua script as input.
	
## Usage

Schema files are written in Lua. Each field needs a unique id, 
which is used in the serialized format to identify fields:

```lua
version = 1

item = struct {
	id = int32 [1],
	name = str [2],
	price = f64 [3],
	tags = list(string) [4],
}

vec3 = struct {
	x = f32 [1],
	y = f32 [2],
	z = f32 [3],
}

player = struct {
	id = int32 [1],
	name = str [2],
	position = vec3 [3],
	inventory = list(item) [4],
}
```

You can generate code for a specific language using the CLI tool:

```bash
ssss -input <input_file> -lang <language> -output <output_file>
```

The supported languages currently are:
- js (With d.ts type definitions)
- go
- c#

You can also pass the path to a Lua script as the lang flag to implement custom code generation for other languages.
See examples of Lua templating in `/writer/js.lua` and `/writer/csharp.lua`.

### Data types

SSSS supports the following data types:
- bool
- (u)int(8|16|32|64)
- float(32|64)
- string
- list<T>
- map<K,V> where K is a primitive type (excluding string & bool) and V is any type
- struct

### Serialization format

Each struct can be serialized and deserialized to/from a compact binary format.
Structs are serialized as follows:

`[type_id(uint16), length(uint32), ...fields]`

Where `type_id` is a unique identifier for the struct type, `length` is the length of the serialized struct in bytes, and `...fields` are the serialized fields of the struct.

Each field begins with it's id (uint16) followed by the serialized value. Because all fields are optional, missing fields are simply omitted from the serialized data. Since the type of each field is known from the schema, the format can omit type information for each field (structs always contain type info).

All non-fixed types (string, list, struct) are prefixed with their length in bytes (uint32).

## Custom Code Generation

You can pass a custom Lua script to the `-lang` flag to implement code generation for other languages.
The script will receive a global `Schema` table containing the following data:

```lua
Schema = {
	version = number,
	name = string,
	structs = { struct_name : struct }
}

struct = {
	name = string,
	id = number, -- the generated unique id for the struct (hashed from the name)
	fields = { field_name : field }
}

field = {
	name = string,
	id = number, -- the user-provided id for the field
	type = "primitive" | "struct" | "list" | "map",
}

primitive_field = {
	type = "primitive",
	name = "bool" | "int8" | "uint8" | "int16" | "uint16" | "int32" | "uint32" | "int64" | "uint64" | "f32" | "f64" | "string",
}

list_field = {
	type = "list",
	of = field, -- the type of the list elements (nested lists are supported)
}

struct_field = {
	type = "struct",
	name = struct,
}

map_field = {
	type = "map",
	from = "primitive", -- key type (must be a primitive type except string/bool)
	to = field,      -- value type (can be any type)
}

-- The script should populate the Output table with file names as keys and file contents as values.
Output["file.ext"] = "file contents as a string"
```

See the existing codegen scripts in `/writer/` for examples of how to use the Schema table to generate code.

### Upcoming Features

- Custom field metadata.
- Default values (under consideration).

### Limitations

- Complex features like Enums, Unions, or service generation / RPC are not planned. This is to keep
	the system simple to implement across new languages.
- No built-in schema registry or versioning system. This is left to the user to implement as needed.
- No built-in compression or encryption. This is left to the user to implement as needed.
- No built-in validation or constraints on field values.
- No built-in support for cyclic references in structs.
- No built-in support for inheritance or polymorphism in structs.

### License
SSSS is licensed under the MIT License.
