set -e

if [ -z "$1" ] || [ -z "$2" ]; then
    echo "usage: $0 REGISTRY_ID VERSION_TAG";
    exit 1;
fi

REGISTRY_ID="$1"
VERSION_TAG="$2"
YA_TOKEN=`cat ~/.ya_token`

docker login -u oauth -p "${YA_TOKEN}" cr.yandex
docker build . -t "cr.yandex/${REGISTRY_ID}/my_app:${VERSION_TAG}"
docker push "cr.yandex/${REGISTRY_ID}/my_app:${VERSION_TAG}"
