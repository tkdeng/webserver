#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdarg.h>

const int JSON_OBJ = 0;
const int JSON_STR = 1;
const int JSON_INT = 2;
const int JSON_FLOAT = 3;
const int JSON_BOOL = 4;
const int JSON_NULL = 5;

int endsWith(char *string, char *suffix){
  string = strrchr(string, suffix[0]);

  if (string != NULL){
    return strcmp(string, suffix);
  }

  return -1;
}

char* json_map(){
  return "{%s}";
}

char* json_arr(){
  return "[%s]";
}

void json_end(char **json){
  if(strlen(*json) < 4 || (endsWith(*json, ",%s}") == -1 && endsWith(*json, ",%s]") == -1)){
    return;
  }

  char b[1024*1024];
  sprintf(b, *json, "");

  b[strlen(b)-1] = '\0';

  if(*json[0] == '{'){
    b[strlen(b)-1] = '}';
  }else if(*json[0] == '['){
    b[strlen(b)-1] = ']';
  }

  // *json = b;
  *json = malloc(strlen(b));
  strcpy(*json, b);
}

void json_arg(char **json, int type, ...){
  va_list args;
  va_start(args, type);

  char *key, *val;
  int val_i;
  double val_f;

  char b[1024*1024];

  if(*json[0] == '{'){
    key = va_arg(args, char*);

    if(type == JSON_OBJ){
      val = va_arg(args, char*);
      json_end(&val);
      sprintf(b, *json, "\"%s\": %s,%s");
    }else if(type == JSON_STR){
      val = va_arg(args, char*);
      sprintf(b, *json, "\"%s\": \"%s\",%s");
    }else if(type == JSON_INT){
      val_i = va_arg(args, int);
      sprintf(b, *json, "\"%s\": %d,%s");
    }else if(type == JSON_FLOAT){
      val_f = va_arg(args, double);
      sprintf(b, *json, "\"%s\": %f,%s");
    }else if(type == JSON_BOOL){
      val_i = va_arg(args, int);
      sprintf(b, *json, "\"%s\": %s,%s");
    }else{
      sprintf(b, *json, "\"%s\": %s,%s");
    }
  }else if(*json[0] == '['){
    if(type == JSON_OBJ){
      val = va_arg(args, char*);
      json_end(&val);
      sprintf(b, *json, "%s,%s");
    }else if(type == JSON_STR){
      val = va_arg(args, char*);
      sprintf(b, *json, "\"%s\",%s");
    }else if(type == JSON_INT){
      val_i = va_arg(args, int);
      sprintf(b, *json, "%d,%s");
    }else if(type == JSON_FLOAT){
      val_f = va_arg(args, double);
      sprintf(b, *json, "%f,%s");
    }else if(type == JSON_BOOL){
      val_i = va_arg(args, int);
      sprintf(b, *json, "%s,%s");
    }else{
      sprintf(b, *json, "%s,%s");
    }
  }else{
    va_end(args);
    return;
  }

  va_end(args);

  char buf[1024*1024];

  if(*json[0] == '{'){
    if(type == JSON_OBJ || type == JSON_STR){
      sprintf(buf, b, key, val, "%s");
    }else if(type == JSON_INT){
      sprintf(buf, b, key, val_i, "%s");
    }else if(type == JSON_FLOAT){
      sprintf(buf, b, key, val_f, "%s");
    }else if(type == JSON_BOOL){
      if(val_i >= 0){
        sprintf(buf, b, key, "false", "%s");
      }else{
        sprintf(buf, b, key, "true", "%s");
      }
    }else{
      sprintf(buf, b, key, "null", "%s");
    }
  }else if(*json[0] == '['){
    if(type == JSON_OBJ || type == JSON_STR){
      sprintf(buf, b, val, "%s");
    }else if(type == JSON_INT){
      sprintf(buf, b, val_i, "%s");
    }else if(type == JSON_FLOAT){
      sprintf(buf, b, val_f, "%s");
    }else if(type == JSON_BOOL){
      if(val_i >= 0){
        sprintf(buf, b, "false", "%s");
      }else{
        sprintf(buf, b, "true", "%s");
      }
    }else{
      sprintf(buf, b, "null", "%s");
    }
  }

  // *json = buf;
  *json = malloc(strlen(buf));
  strcpy(*json, buf);
}
