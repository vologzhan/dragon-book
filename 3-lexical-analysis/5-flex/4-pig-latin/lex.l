%{
int pigLatin();
%}

%%
[A-Za-z]+ {pigLatin();}
%%

#include <ctype.h>

int yywrap() {
    return 1;
}

int main() {
    yylex();
    return 0;
}

int pigLatin() {
    char *src = yytext;

    char first = src[0];
    char lower = tolower(first);
    int isLower = first == lower;

    switch (lower) {
        case 'b': case 'c': case 'd': case 'f': case 'g': case 'h': case 'j': case 'k': case 'l': case 'm': case 'n':
        case 'p': case 'q': case 'r': case 's': case 't': case 'v': case 'w': case 'x': case 'y': case 'z':
            size_t len = strlen(src);
            memmove(src, src + 1, len - 1);
            src[len-1] = lower;

            if (!isLower) {
                *src = toupper(*src);
            }
    }

    printf("%say", yytext);

    return 0;
}
