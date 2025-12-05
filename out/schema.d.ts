// Auto-generated code for schema: schema v1

export class Foo {
	f?: number;
	m?: { [key: number]: string };
	x?: { [key: string]: { [key: number]: boolean } };
	constructor(props?: Omit<Partial<Foo>, 'fromBytes' | 'toBytes'>);
	static readonly TypeID: number;
	toBytes(): Uint8Array;
	fromBytes(bytes: ArrayBuffer | ArrayBufferView): Foo;
}

export function deserialize(bytes: ArrayBuffer | ArrayBufferView): Foo;

