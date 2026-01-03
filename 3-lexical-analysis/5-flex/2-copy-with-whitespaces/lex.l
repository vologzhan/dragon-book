%%
[ \t\n]+ { printf(" "); }
%%

int yywrap() {
    return 1;
}

int main() {
    printf("\"");
    yylex();
    printf("\"\n");
    return 0;
}
