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
go run . out.rat.gz folder_a file_a
```
> RAT file can be compressed if the output extension is `.gz`.

DERAT a file or a set of files:
```bash
go run . -x out.rat
go run . -x out.rat another_out.rat.gz
```
