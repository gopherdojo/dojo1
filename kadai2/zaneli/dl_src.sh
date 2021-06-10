#!/bin/sh

REVISION="a7a4f4dbf8a2f574a4c7f071316a6b6b8565ed14"
BASE_URL="https://raw.githubusercontent.com/gopherdojo/dojo1/${REVISION}/kadai1/zaneli/image-converter/converter/"
FILES=("convert.go" "converter.go" "decoder.go" "encoder.go")

cd `dirname $0`

for FILE in "${FILES[@]}"
do
  curl -L "${BASE_URL}/${FILE}" --output "./image-converter/converter/${FILE}"
done
