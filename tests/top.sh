#!/bin/bash
echo ""
echo ""

echo "Test 1: Top command with -amount and -fiat flags"
echo "Input: top -amount=\"10\" -fiat=\"VND\""
echo "Output:"
go run . top 
echo "--------------------"