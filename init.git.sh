git clone  git clone git@github.com:Leon-OnlineLearning/leon-all.git
git submodule init
git submodule update --init --recursive
mkdir nginx/logs
touch nginx/cert/fullchain.pem
touch nginx/cert/privkey.pem