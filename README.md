## Hack Assembler
> A simple assembler written in Golang to process `.asm` files and generate corresponding `.hack` files.

### Prerequisites - Install golang

1. Download from [official website](https://go.dev/dl/).
2. Verify installation:
 ```bash
   go version
 ```
   
<br />

### Usage

1. Run without building
```bash
  go run main.go "filename"
```
> Takes "path/filename.asm" --> generate "path/filename.hack"


2. Build the assembler:
```bash
  go build -o assembler
```
run via build:
```bash
  ./assembler "filename"
```

   
