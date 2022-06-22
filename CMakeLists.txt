cmake_minimum_required(VERSION 3.22)
project(GIS)

set(CMAKE_CXX_STANDARD 17)

include_directories(.)

add_executable(GIS
        graph.hpp
        json.cpp
        json.hpp
        main.cpp
        router.hpp)
