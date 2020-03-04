CC      = gcc
CFLAGS  = -g
RM      = rm -f


default: all

all: logger

logger: logger.c
	$(CC) $(CFLAGS) -o logger.o logger.c

clean veryclean:
	$(RM) cal
