# Backend Developer Internship Technical Test

**Candidate:** Deren Tanaphan  
**University / Major:** Universitas Gadjah Mada / Information Engineering  

This repository contains my submission for the Backend Developer Internship technical test at PT Fire Rock Indonesia.

## Contents

| File | Description |
|------|-------------|
| [`assessment.md`](./assessment.md) | Full answers — Section 1 (single-choice), Section 2 (multi-select), Section 3 (short essays), and Section 4 (voucher engine) |
| [`main.go`](./main.go) | Runnable Go solution for Section 4 — the "FlashSale Voucher Discount Engine" |

## Running Section 4

Requires Go 1.21+.

```bash
go run main.go
```

The built-in `main()` exercises all business rules: no voucher, discount under the cap, discount capped by `MaxDiscount`, cart below `MinPurchase` (no discount, no error), empty cart, invalid quantity, and negative price.
