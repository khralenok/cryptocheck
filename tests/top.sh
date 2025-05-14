#!/bin/bash
echo ""
echo ""

echo "Test 1: Top command with no flags at all"
echo "Input: top"
echo "Output:"
go run . top
echo "--------------------"

echo "Test 2: Top command with -amount and -fiat flags"
echo "Input: top -amount=\"10\" -fiat=\"VND\""
echo "Output:"
go run . top -amount "10" -fiat="VND"
echo "--------------------"

echo "Test 3: Top command with -amount fewer than 10(Minimal coindesk query)"
echo "Input: top -amount=\"3\""
echo "Output:"
go run . top -amount "3"
echo "--------------------"

echo "Test 4: Top command with too large -amount value"
echo "Input: top -amount=\"15000\""
echo "Output:"
go run . top -amount "15000"
echo "--------------------"

echo "Test 5: Top command with wrong -fiat flag"
echo "Input: top -fiat=\"jjlkfjkdjf\""
echo "Output:"
go run . top -amount "10" -fiat="jjlkfjkdjf"
echo "--------------------"

echo "Test 6: Top command with multiple values in -fiat flag"
echo "Input: top -fiat=\"VND, USD\""
echo "Output:"
go run . top -amount "10" -fiat="VND, USD"
echo "--------------------"