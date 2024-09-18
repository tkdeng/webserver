#include <stdio.h>
#include "../handlers/c/render.c"
#include "../handlers/c/json.c"

int main(){
  printf("%s\n", "Hello, World!");

  char *json = json_map();

  json_arg(&json, JSON_STR, "key", "value");

  char *list = json_arr();
  json_arg(&list, JSON_INT, 1);
  json_arg(&list, JSON_FLOAT, 2.5);
  json_arg(&list, JSON_INT, 3);

  json_arg(&json, JSON_OBJ, "list", list);

  json_end(&json);

  //todo: improve render method
  render("index", json, "layout");

  return 0;
}
