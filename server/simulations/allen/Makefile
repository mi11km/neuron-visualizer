CC = gcc
CFLAGS = -O3 -std=c11 -Wall

all: main

main: main.o cell.o ion.o solve.o csparse.o timer.o
	$(CC) ${CFLAGS} -o $@ $^

main.o: main.c param.h cell.h ion.h solve.h
	$(CC) ${CFLAGS} -c $<

cell.o: cell.c param.h cell.h ion.h solve.h
	$(CC) ${CFLAGS} -c $<

ion.o: ion.c param.h cell.h ion.h solve.h
	$(CC) ${CFLAGS} -c $<

solve.o: solve.c param.h cell.h ion.h solve.h
	$(CC) ${CFLAGS} -c $<

csparse.o: csparse.c csparse.h
	$(CC) ${CFLAGS} -c $<

timer.o: timer.c
	$(CC) ${CFLAGS} -c $<

clean:
	rm -f *.o *~

distclean: clean
	rm -f main *.csv *.dat
