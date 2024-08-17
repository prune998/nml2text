# NML to Text

This project takes a Playlist file in NML format, as exported from Traktor, and generate a text list output that you can use
to share along your mixes.

## Usage 

```bash
./nml2text -nmlFile djset-highnrj-house.nml

1) Alright - Micky More & Andy Tee
2) Love Sweet Sound (Friscia & Lamboy House Mix) - Groove Armada
3) Wanna Hum Hum (Original Mix) - Andruss
```

## Building

This is go...

```
go get -u ./...
go build
```