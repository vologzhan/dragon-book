OUT_FILE=/tmp/a.out

compile-and-run:
	flex -t lex.l | gcc -x c - -o $(OUT_FILE) && $(OUT_FILE) < input.txt
