f = open('a.txt', encoding='utf-8')
lines = f.readlines()
j = 0
k = 0
bigarr = [[], [], [], []]
for l in lines:
    l = l.replace('\n', '')
    arr = l.split('\t')
    for i in [1, 2, 3, 4]:
        bigarr[i - 1].append(arr[i].split(' ')[0])
for r in bigarr:
    for m in r:
        if k >= 4:
            j = j + 1
            k = 0
        k = k + 1
        print('insert into groups (game_name, round, group_id, process) values (\"%s\", 6, %d, 1);' % (m, j))
