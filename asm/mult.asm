@R0
D=M
@i
M=D

(LOOP)
  @i
  D=M
  @END
  D;JLE

  @R1
  D=M
  @R2
  M=M+D

  @i
  M=M-1

  D=M
  @LOOP
  D;JGT

(END)
  @END
  0;JMP



