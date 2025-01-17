(MAIN)
@24576
D=M

@DOWN
D;JGT //key-down (k>0)


//else (key-up)
  @R0
  D=M
  @MAIN
  D;JEQ  //(mode=0): screen empty -> no need to clear

  //(mode=1): clear
  @val
  M=0
  @R0
  M=0
  @LOOP
  0;JMP


//(key-down)
  (DOWN)
  @R0
  M=1
  @val
  M=-1
  @LOOP
  0;JMP


//loop event
(LOOP)
  //set index
  @8159
  D=A
  @i
  M=D
  //loop run
  (RUN)
        @i
        D=M
	@MAIN
	D;JLT  //(if i<0) -> end loop & back to main

	//calculate mem address (base+i)
	@i
	D=M
	@16384
	D=D+A
	@temp
	M=D

	@val
	D=M
	@temp
	A=M
	M=D

	//decrement i
	@i
	M=M-1
	D=M
	@RUN
	D;JGE  //(if i>=0) -> keep loop alive.
	
	
@MAIN
0;JMP
