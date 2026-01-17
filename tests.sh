go run main.go -i schema.lua -o out -l js
go run main.go -i schema.lua -o out -l go
go run main.go -i schema.lua -o out -l c#

go run test/go.go
deno run --allow-write --allow-read test/js.ts
dotnet run test/csharp.cs
go run test/go.go -compare
deno run --allow-write --allow-read test/js.ts --compare
dotnet run test/csharp.cs --compare
