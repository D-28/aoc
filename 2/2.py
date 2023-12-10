import sys
from collections import defaultdict

with open(sys.argv[1], 'r') as file:
    data = file.read().strip()
    ans1 = 0
    ans2 = 0
    for game in data.split('\n'):
        ok = True
        id_, game = game.split(':')
        V = defaultdict(int)
        for event in game.split(';'):
            event = [e.lstrip() for e in event.split(',')]
            for cube in event:
                n, color = cube.split()
                V[color] = max(V[color], int(n))
                if int(n) > {'red': 12, 'green': 13, 'blue': 14}.get(color, 0):
                    ok = False
        score = 1
        for v in V.values():
            score *= v
        ans2 += score
        if ok:
            # ans1 += int(id_.split()[-1])
            pass
        
    print(ans1)
    print(ans2)


# part two: 
# 1. for each game, get max of each color, save the max values for each color
# 2. for each game, multiply all the max values for each color
# 3. sum all the multiplied values
