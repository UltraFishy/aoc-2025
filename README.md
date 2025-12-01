# Advent of Code

## Overview
This repository contains my Go solutions for [Advent of Code 2025](https://adventofcode.com/2025)challenges.

## Running Solutions
```bash
go run ./days/day** -part *
```
Where you choose the day and the part you want.

## Scripts
Using a makefile we have some automation when setting up an environment for a day.
The three different targets are the following `make <target>`:

1. `get-aoc-cookie`: ensures $AOC_SESSION_COOKIE is set
2. `input`: gets the input text from the website
3. `skeleton`: gets the template ready for the day
  
