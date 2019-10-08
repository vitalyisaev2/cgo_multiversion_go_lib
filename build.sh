#!/bin/bash

pushd v1
rm -rf build && mkdir build && pushd build
cmake .. && make && make install
popd
popd

pushd v2
rm -rf build && mkdir build && pushd build
cmake .. && make && make install
popd
popd

