#!/bin/bash

DIST_DIR="runzero-api"
ZIP_NAME="swagger_latest.zip"
DIST_ZIP="runzero-api.zip"
OAS3_FILE="../runzero-api.yml"

## Clean up any directories that may have lingered
if [[ -d "$DIST_DIR" ]]; then
  rm -r "${DIST_DIR:?}"
fi

if [[ -f "$ZIP_NAME" ]]; then
  rm "${ZIP_NAME:?}"
fi

if [[ -f "$DIST_ZIP" ]]; then
  rm "${DIST_ZIP:?}"
fi

find . -name "swagger-api*" -print0 | xargs -r0 rm -r

## Get latest swagger zip from Github
LINK="$(curl -s https://api.github.com/repos/swagger-api/swagger-ui/releases/latest | yq eval '.zipball_url' -)"

## Download an unzip
echo "Downloading latest Swagger UI from: ${LINK}"
curl -s -o swagger_latest.zip -L "${LINK}"
unzip -q swagger_latest.zip
mkdir -p "${DIST_DIR}"
cp -r swagger-api*/dist/* "${DIST_DIR}"

## Copy runZero API spec
cp "$OAS3_FILE" "${DIST_DIR}/runzero-api.yml"
cp README.md "${DIST_DIR}/README.md"

## Convert yaml to json
SPEC_JSON="$(yq eval -o=j -I=-0 "${DIST_DIR}"/runzero-api.yml | sed -e 's/[\/&]/\\&/g' -e 's/console.runzero.com/localhost/g')"

## Edit index to use the runZero API spec
URL_VAL="$(grep url runzero-api/index.html | sed -e 's/[\/&]/\\&/g')"
sed -i "s/${URL_VAL}/spec:${SPEC_JSON},/" runzero-api/index.html

## Zip up the files for distribution
zip -rq runzero-api.zip runzero-api

## Clean up
find . -name "swagger-api*" -print0 | xargs -r0 rm -r
rm swagger_latest.zip
rm -r runzero-api

echo "Finished... ${DIST_ZIP} is ready for distribution."
