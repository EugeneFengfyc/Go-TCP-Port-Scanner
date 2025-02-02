# TCP Port Scanner

This Go program scans for open TCP ports on a given IP address by attempting to establish a connection to each port within the range of 1-65535. It uses concurrency and atomic operations to efficiently scan multiple ports at once while ensuring thread safety when storing results.

## Features:
- Allows scanning of a user-defined IP address using the `-ip` flag.
- Concurrent scanning of up to 100 ports at a time using goroutines and semaphores.
- Stores open ports in a thread-safe manner using atomic operations.
- Prints the list of open ports once the scan is complete.

## Requirements:
- Go 1.16+ (for `sync/atomic` package)
- A working network connection

## How It Works:
- The program accepts an IP address as input using the `-ip` flag.
- It creates a goroutine for each port in the 1-65535 range to attempt a connection.
- The concurrency is limited to 100 simultaneous goroutines using a semaphore channel.
- It uses the `net.Dial` function to check if the port is open by attempting to establish a TCP connection.
- If the connection is successful, the port is added to the list of open ports.
- The program prints the list of open ports after all the goroutines have finished.

## Example Output:
```bash
Open ports:
192.168.1.1:22
192.168.1.1:80
192.168.1.1:443

Elapsed time: 35 seconds

