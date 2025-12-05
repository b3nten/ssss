// Auto-generated code for schema: Schema v1

namespace Schema;

interface ISchema<TSelf> where TSelf : ISchema<TSelf>
{
    public byte[] Serialize();
    public TSelf Deserialize(byte[] data);
}
class Foo : ISchema<Foo>
{
    public float? F;
    public Dictionary<int, string>? M;
    public Dictionary<string, Dictionary<int, bool>>? X;

    public readonly static ushort TypeId = 32471;

    public static Foo CreateFromBytes(byte[] data)
    {
        Foo it = new Foo();
        using (MemoryStream ms = new MemoryStream(data))
        using (BinaryReader r = new BinaryReader(ms))
        {
            _Foo.Deserialize(it, r);
        }
        return it;
    }

    public Foo Deserialize(byte[] data)
    {
        using (MemoryStream ms = new MemoryStream(data))
        using (BinaryReader r = new BinaryReader(ms))
        {
            _Foo.Deserialize(this, r);
        }
        return this;
    }

    public byte[] Serialize()
    {
        using (MemoryStream ms = new MemoryStream())
        using (BinaryWriter w = new BinaryWriter(ms))
        {
            _Foo.Serialize(this, w);
            return ms.ToArray();
        }
    }
}

file class _Foo
{
    static public void Serialize(Foo it, BinaryWriter w)
    {
        w.Write(Foo.TypeId);
        var lengthPos = w.BaseStream.Position;
        w.Write((UInt32)0);
        if (it.F != null)
        {
            w.Write((ushort)1);
            w.Write(it.F.Value);
        }
        if (it.M != null)
        {
            w.Write((ushort)2);
            var length0 = w.BaseStream.Position;
            w.Write((uint)0);
            foreach (var kv0 in it.M)
            {
                w.Write(kv0.Key);
                var bytes = System.Text.Encoding.UTF8.GetBytes(kv0.Value);
                w.Write((uint)bytes.Length);
                w.Write(bytes);
            }
            var end0 = w.BaseStream.Position;
            w.Seek((int)length0, SeekOrigin.Begin);
            w.Write((uint)(end0 - length0 - 4));
            w.Seek(0, SeekOrigin.End);
        }
        if (it.X != null)
        {
            w.Write((ushort)3);
            var length0 = w.BaseStream.Position;
            w.Write((uint)0);
            foreach (var kv0 in it.X)
            {
                var bytes = System.Text.Encoding.UTF8.GetBytes(kv0.Key);
                w.Write((uint)bytes.Length);
                w.Write(bytes);
                var length1 = w.BaseStream.Position;
                w.Write((uint)0);
                foreach (var kv1 in kv0.Value)
                {
                    w.Write(kv1.Key);
                    w.Write(kv1.Value);
                }
                var end1 = w.BaseStream.Position;
                w.Seek((int)length1, SeekOrigin.Begin);
                w.Write((uint)(end1 - length1 - 4));
                w.Seek(0, SeekOrigin.End);
            }
            var end0 = w.BaseStream.Position;
            w.Seek((int)length0, SeekOrigin.Begin);
            w.Write((uint)(end0 - length0 - 4));
            w.Seek(0, SeekOrigin.End);
        }
        var endPos = w.BaseStream.Position;
        w.Seek((int)lengthPos, SeekOrigin.Begin);
        w.Write((UInt32)(endPos - lengthPos - 4));
        w.Seek(0, SeekOrigin.End);
    }
    static public void Deserialize(Foo it, BinaryReader r)
    {
        ushort typeId = r.ReadUInt16();
        if (typeId != Foo.TypeId)
        {
            throw new Exception($"TypeId mismatch: expected Foo.TypeId but got {typeId}");
        }
        uint length = r.ReadUInt32();
        long startPos = r.BaseStream.Position;
        while (r.BaseStream.Position - startPos < length)
        {
            ushort fieldId = r.ReadUInt16();
            switch (fieldId)
            {
                case 1:
                    it.F = r.ReadSingle();
                    break;
                case 2:
                    {
                        uint mapLength0 = r.ReadUInt32();
                        long startPos0 = r.BaseStream.Position;
                        var map0 = new System.Collections.Generic.Dictionary<int, string>();
                        while (r.BaseStream.Position - startPos0 < mapLength0)
                        {
                            int key0;
                            string value0;
                            key0 = r.ReadInt32();
                            {
                                uint strLen = r.ReadUInt32();
                                var strBytes = r.ReadBytes((int)strLen);
                                value0 = System.Text.Encoding.UTF8.GetString(strBytes);
                            }
                            map0.Add(key0, value0);
                        }
                        it.M = map0;
                    }
                    break;
                case 3:
                    {
                        uint mapLength0 = r.ReadUInt32();
                        long startPos0 = r.BaseStream.Position;
                        var map0 = new System.Collections.Generic.Dictionary<string, Dictionary<int, bool>>();
                        while (r.BaseStream.Position - startPos0 < mapLength0)
                        {
                            string key0;
                            Dictionary<int, bool> value0;
                            {
                                uint strLen = r.ReadUInt32();
                                var strBytes = r.ReadBytes((int)strLen);
                                key0 = System.Text.Encoding.UTF8.GetString(strBytes);
                            }
                            {
                                uint mapLength1 = r.ReadUInt32();
                                long startPos1 = r.BaseStream.Position;
                                var map1 = new System.Collections.Generic.Dictionary<int, bool>();
                                while (r.BaseStream.Position - startPos1 < mapLength1)
                                {
                                    int key1;
                                    bool value1;
                                    key1 = r.ReadInt32();
                                    value1 = r.ReadBoolean();
                                    map1.Add(key1, value1);
                                }
                                value0 = map1;
                            }
                            map0.Add(key0, value0);
                        }
                        it.X = map0;
                    }
                    break;
                default:
                    r.BaseStream.Seek(startPos + length, SeekOrigin.Begin);
                    return;
            }
        }
    }
}
