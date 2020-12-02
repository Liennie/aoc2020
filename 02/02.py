#!/usr/bin/python

import re, os

input = "input.txt"

valid = 0

with open(input, 'r') as file:
    for line in file:
        match = re.match(r"^(\d+)-(\d+) (\w): (\w+)$", line)
        if match:
            min = int(match.group(1))
            max = int(match.group(2))
            letter = match.group(3)
            password = match.group(4)

            count = 0
            for char in password:
                if char == letter:
                    count += 1

            if min <= count <= max:
                valid += 1

print(valid)
