#!/bin/bash
echo "开始执行程序！！！"


echo "git clone loading ..."

git clone git@github.com:hejin0929/root_student.git

echo "git clone end and start build ..."

mkdir dist

# shellcheck disable=SC2164
cd ./dist

dir=$(pwd)

cd ../

# shellcheck disable=SC2164
cd "./root_student"

echo "this is a dir value ? $(pwd)"

yarn build

# shellcheck disable=SC2164
cd "./dist"

 
dis=$(pwd)/

cp -af "${dis}" "${dir}"

cd ../../


rm -rf ./index

rm -rf ./root_student

echo "$(pwd) ls"


# cp dis res

# ls


echo "程序执行结束。。。"