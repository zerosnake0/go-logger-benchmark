FILE=$1
if [ "$FILE" == "" ]; then
  FILE=cpu.profile
fi

go tool pprof -web $FILE