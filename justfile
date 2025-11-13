VERSION := 'v0.2.0'

version:
  @echo {{VERSION}} 

check-versions:
    #!/usr/bin/env bash
    set -euo pipefail
    expected_go=$(grep '^go = ' mise.toml | cut -d'"' -f2)
    current_go=$(go version | awk '{print $3}' | sed 's/go//')
    if [[ ! "$current_go" =~ ^"$expected_go" ]]; then
        echo "❌ Go version mismatch!"
        echo "   Expected: $expected_go"
        echo "   Current:  $current_go"
        echo "   Run: mise install"
        exit 1
    fi
    echo "✅ All tool versions match mise.toml"

lint:
    golangci-lint run

format:
    golangci-lint run --fix

test:
    go test ./...

doc: 
    go doc -http
