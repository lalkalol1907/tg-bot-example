brew install gperf cmake openssl
git clone https://github.com/tdlib/td.git
cd td
rm -rf build
mkdir build
cd build
cmake -DCMAKE_BUILD_TYPE=Release -DOPENSSL_ROOT_DIR=/opt/homebrew/opt/openssl/ -DOPENSSL_LIBRARIES=/opt/homebrew/opt/openssl/lib -DCMAKE_INSTALL_PREFIX:PATH=/usr/local ..
cmake --build . --target install --parallel 8
cd ..
cd ..
ls -l /usr/local