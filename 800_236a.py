username = input()
num_chars = len(set(username))
if num_chars % 2 == 0:
  print("CHAT WITH HER!")
else:
  print("IGNORE HIM!")