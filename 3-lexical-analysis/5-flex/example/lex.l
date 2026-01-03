%{
/* Определения именованных констант */
#define LT      1
#define LE      2
#define EQ      3
#define NE      4
#define GT      5
#define GE      6
#define IF      7
#define THEN    8
#define ELSE    9
#define WHILE   10
#define ID      11
#define NUMBER  12
#define RELOP   13

int yylval;

/* Прототипы функций */
int installID();
int installNum();
%}

/* Регулярные определения */
delim   [ \t\n]
ws      {delim}+
letter  [A-Za-z]
digit   [0-9]
id      {letter}({letter}|{digit}|_)*
number  {digit}+(\.{digit}+)?(E[+-]?{digit}+)?

%%

{ws}       {/* Нет действий и выхода из функции */}
if         {return(IF);}
then       {return(THEN);}
else       {return(ELSE);}
while      {return(WHILE);}
{id}       {yylval = (int) installID(); return(ID);}
{number}   {yylval = (int) installNum(); return(NUMBER);}
"<"        {yylval = LT; return(RELOP);}
"<="       {yylval = LE; return(RELOP);}
"="        {yylval = EQ; return(RELOP);}
"!="       {yylval = NE; return(RELOP);}
">"        {yylval = GT; return(RELOP);}
">="       {yylval = GE; return(RELOP);}

%%

/* Реализация функций */
int installID() {
    printf("[ID: %s] ", yytext);
    return 0;
}

int installNum() {
    printf("[NUM: %s] ", yytext);
    return 0;
}

/* Функция завершения (необходима для Flex) */
int yywrap() {
    return 1;
}

/* Точка входа для проверки */
int main() {
    int token;
    while((token = yylex())) {
        printf("Token code: %d, yylval: %d\n", token, yylval);
    }
    return 0;
}
