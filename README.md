# CINEMA XXI - LAYOUT

CINEMA XXI - LAYOUT

![Flow](cinemaxxi_layout.png)

## Run Project

### Run Go Mod Vendor

```bash
go mod vendor
```

### Run API Gateway

```bash
go run cmd\main.go
```

### Run CINEMA XXI - LAYOUT

```bash
OPEN localhost:8080 on browser
```

## Run Test

```bash
go test -v ./... -cover
```

## Example Request and Response

### Add Seat Configuration

![Flow](cinemaxxi_seat_config.png)

### Booking Seat

![Flow](cinemaxxi_book_seat.png)

### Cancel Seat

![Flow](cinemaxxi_cancel_book.png)

### Seat Status

![Flow](cinemaxxi_plan_report.png)

### Transaction Status

![Flow](cinemaxxi_transaction_report.png)