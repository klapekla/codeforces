import re
hq9plus_string = input()
hq9plus_regexp = re.compile(r'H|Q|9')
if hq9plus_regexp.search(hq9plus_string):
  print('YES')
else:
  print('NO')
