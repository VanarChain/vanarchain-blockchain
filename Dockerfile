# Build Geth in a stock Go builder container
FROM golang:1.21.6 as builder


# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/


ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM golang:1.21.6

COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth", "--testnet", "--http.addr=0.0.0.0", "--http.api", "eth,net,web3,txpool", "--ws.addr=0.0.0.0", "--ws.api", "eth,net,web3,txpool", "--http.vhosts", "*", "--http.corsdomain", "*", "--ws.origins", "*", "--ethstats", "rpc-cloud:WdF5dXPRJBR8QxF0A7B@stats-testnet.vanarchain.com" ]
