import os
import time
import sys

# start_time = time.time()
# with open('day1-file.txt', 'r') as f:
#     result = 0
#     line = f.readline()
#     while line != '':
#         first, last = '', ''
#         foundFirst = False
#         for i, char in enumerate(line):
#             if char.isdigit():
#                 if not foundFirst:
#                     first = char
#                     foundFirst = True
#                 last = char
        
#         if first and last:
#             strResult = first + last
#             result += int(strResult)
#         line = f.readline()
    
# end_time = time.time()
# print(f'Answer: {result}')
# print(f'Esxecution time: {(end_time - start_time) * 1_000_000}µs')

# more pythonic
# def process(line):
#     digits = [char for char in line if char.isdigit()]
#     return digits[0] + digits[-1] if digits else None
#
# start_time = time.time()
#
# with open('day1-file.txt', 'r') as f:
#     result = sum(int(process(line)) for line in f if process(line))
#
# end_time = time.time()
#
# print(f'Answer: {result}')
# print(f'Execution time: {(end_time - start_time) * 1_000_000}µs')

data = open(sys.argv[1]).read().strip()
ans1 = 0 
ans2 = 0 
for line in data.split('\n'):
    digits1 = []
    digits2 = []
    for i,c in enumerate(line):
        if c.isdigit():
            digits1.append(c)
            digits2.append(c)
        for d, val in enumerate(['one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine']):
            if line[i:].startswith(val):
                digits2.append(str(d+1))
    ans1 += int(digits1[0] + digits1[-1])
    ans2 += int(digits2[0] + digits2[-1])
print(ans1)
print(ans2)
