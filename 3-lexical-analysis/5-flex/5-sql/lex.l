%{
#define SELECT 1
#define FROM   2
#define WHERE  3
#define ID     4

int yylval;

int installId();
%}

%option caseless

%%
[ \n\t]+       {}
select         { yylval = SELECT;       return(SELECT); }
from           { yylval = FROM;         return(FROM);   }
where          { yylval = WHERE;        return(WHERE);  }
[A-Za-z0-9_]+  { yylval = installId();  return(ID);     }
%%

int yywrap() {
    return 1;
}

int main() {
    int token;
    while((token = yylex())) {
        printf("Token code: %d, yylval: %d\n", token, yylval);
    }
    return 0;
}

int installId() {
    printf("[ID %s] ", yytext);
    return 0;
}
