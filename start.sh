CONTAINER_NAME=mecab-api
if [ $MODE = DEVELOPMENT ]; then
    docker build -t $CONTAINER_NAME .
elif [ $MODE = RUN ]; then
    docker run -it -v $PWD/src/:/go/src/ $CONTAINER_NAME /bin/bash
else
    echo "Not DEVELOPMENT"
fi
docker run -it -v $PWD/src/:/go/src/ $CONTAINER_NAME /bin/bash
