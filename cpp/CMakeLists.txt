cmake_minimum_required(VERSION 3.10)

project(Duncode)

set(CMAKE_CXX_STANDARD 11)


add_executable(duncoder  Duncoder.cpp  Duncode.cpp Tool.cpp Duncode.h)
add_library(Duncode  Duncoder.cpp Duncode.cpp Tool.cpp Duncode.h)