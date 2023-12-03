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
ans = 0 
for line in data.split('\n'):
    digits = []
    for char in line:
        if char.isdigit():
            digits.append(char)
    score = int(digits[0] + digits[-1])
    ans += score
print(ans)
