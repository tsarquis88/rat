## 🐀 Description

RAT is like TAR but worse 🌠

It is just a work-in-progress project I've created to learn Go, nothing serious (or is it?).

### 🐀 Usage

The following commands are executed in the `cmd/rat` folder.

RAT a file, a set of files, or a folder:
```bash
go run . out.rat file_a
go run . out.rat file_a file_b
go run . out.rat folder_a
```

DERAT a file or a set of files:
```bash
go run . -x out.rat
go run . -x outa.rat outb.rat
```
> TAR files can also be derated with the same command from above.
