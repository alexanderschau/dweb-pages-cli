#!/bin/sh

ipfs init
ipfs daemon &
waitForIpfs

PROJECT_CID=$(ipfs cid base32 $(dweb-pages add$(for tag in $(echo $TAGS | tr ", " "\n"); do echo -n " -t $tag"; done)))

if [[ $ENDPOINT ]]; then
  ipfs pin remote service add origin $ENDPOINT $ACCESS_TOKEN

  OLD_CID=$(ipfs pin remote ls --service=origin --name=$NAME | awk '{print $1}')
  if [[ $OLD_CID ]]; then
  ipfs pin remote rm --service=origin --cid=$OLD_CID
    ipfs pin remote add --service=origin --name=$NAME $PROJECT_CID
  else
    ipfs pin remote add --service=origin --name=$NAME $PROJECT_CID
  fi
fi

echo "::cid::$PROJECT_CID"
