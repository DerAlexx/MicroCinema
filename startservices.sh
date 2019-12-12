#!/bin/bash

echo "[+] Starting Services"
go run ./cinemahall/main.go &
echo "[+] Started Cinemahall-Service"
go run ./movies/main.go &
echo "[+] Started Movie-Service"
go run ./reservation/main.go &
echo "[+] Started Reservation-Service"
go run ./show/main.go &
echo "[+] Started Show-Service"
go run ./users/main.go &
echo "[+] Started User-Service"