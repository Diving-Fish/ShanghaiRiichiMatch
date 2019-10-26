f = open('source.txt', 'r', encoding='utf-8')
lines = f.readlines()
for i in range(2, 1536, 12):
    print("update players set status = 1 where nickname = \"%s\" and sid != 0;" % lines[i].replace('\n', ''))