CC      = gcc
CFLAGS  = -g
RM      = rm -f


default: all

all: cal

cal: src/main.c
	$(CC) $(CFLAGS) -o calc src/main.c include/addition.h include/subtraction.h include/multiplication.h include/division.h include/modulus.h

clean veryclean:
	$(RM) cal