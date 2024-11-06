#!/bin/bash

#* C
gcc -o index index.c && ./index

#* Elixir
elixir index.exs
iex index.exs

#* Scala
# build
scalac main.scala
# run
scala app

#* Python
python app.py

#* Haskell
# build
ghc --make main.hs -o app
# run
./app

#* C#
dotnet run
# build
dotnet publish -o app
# run
./app/c#
