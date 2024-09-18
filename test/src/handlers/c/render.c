#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void render(char *page, char *args, char *layout) {
  //todo: return html output from exs template
  printf("@PAGE:%s;\n", page);
  printf("@ARGS:%s;\n", args);
  printf("@LAYOUT:%s;\n", layout);
}
