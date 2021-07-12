n = int(input())
for i in range(n):
  string = input()
  string_length = len(string)
  if string_length > 10:
    print(string[0:1] + str(string_length-2) + string[(string_length-1):string_length])
  else:
    print(string)
