f = open('source.txt', 'r', encoding='utf-8')
s = f.readlines()
for line in s:
    arr = line.replace('\n', '').split('\t')
    print('insert into players (school, sid, nickname, password) values ("%s", %d, "%s", "%s");' % (arr[1], 0, arr[2], arr[5]))