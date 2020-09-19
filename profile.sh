set -ex

go test -bench='.*' -cpuprofile=cpu.profile $@