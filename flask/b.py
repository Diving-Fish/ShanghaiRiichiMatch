f = open('source.txt', 'r', encoding='utf-8')
lines = f.readlines()
for i in range(2, len(lines), 9):
    print("update players set status = 0 where nickname = \"%s\" and sid != 0;" % lines[i].replace('\n', ''))