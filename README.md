# Polymarket Low Latency Arbitrage Strategy

This repository contains a Go program that implements a low-latency arbitrage strategy on Polymarket. The program is designed to identify and exploit pricing inefficiencies across different markets in real-time.

## Overview

The goal of this project is to leverage Polymarket's decentralized prediction markets to identify arbitrage opportunities. By monitoring multiple prediction markets, the program can quickly execute trades to take advantage of price differences across markets, generating profit from the discrepancies.

### Key Features:
- **Low-latency performance**: Optimized for real-time market data processing.
- **Automatic arbitrage detection**: The program continuously scans markets for pricing inefficiencies.
- **Execution of trades**: Trades are executed automatically when an arbitrage opportunity is detected.
- **Polymarket API integration**: Uses Polymarket’s API to fetch market data and place orders.

## Prerequisites

To run this program, you'll need:

- **Go 1.18+**: The program is written in Go, and requires Go 1.18 or later.
- **Polymarket API access**: You need API access to Polymarket to retrieve market data and execute trades.
- **A Polymarket wallet**: An active wallet on Polymarket to make transactions.
- **Basic understanding of Arbitrage**: Familiarity with how arbitrage opportunities in financial markets work.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/polymarket-arbitrage.git
   cd polymarket-arbitrage
## Disclaimer
This strategy can be rate limited and is not guaranteed to be profitable. It is heavlity dependent on latency.
