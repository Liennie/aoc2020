#!/usr/bin/python

import re, os

input = "input.txt"

passwords = []

with open(input, 'r') as file:
    for line in file:
        match = re.match(r"^(\d+)-(\d+) (\w): (\w+)$", line)
        if match:
            min = int(match.group(1))
            max = int(match.group(2))
            letter = match.group(3)
            password = match.group(4)

            passwords.append({
                "min": min,
                "max": max,
                "letter": letter,
                "password": password,
            })

# Part 1

valid = 0
for password in passwords:
    min = password["min"]
    max = password["max"]
    letter = password["letter"]
    word = password["password"]

    count = 0
    for char in word:
        if char == letter:
            count += 1

    if min <= count <= max:
        valid += 1

print("Part 1:", valid)

# Part 2

valid = 0
for password in passwords:
    min = password["min"]
    max = password["max"]
    letter = password["letter"]
    word = password["password"]

    if len(word) >= max:
        if (word[min-1] == letter) != (word[max-1] == letter):
            valid += 1

print("Part 2:", valid)
