default: all
all: cal
cal: src/main.c
	gcc -g -o calc src/main.c include/addition.h include/subtraction.h include/multiplication.h include/division.h include/modulus.h
	./calc
clean veryclean:
	rm -rf cal