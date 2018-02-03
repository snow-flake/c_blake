cmake_minimum_required(VERSION 3.10)
project(c_blake)

set(CMAKE_BUILD_TYPE Release)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_FLAGS "-march=native -mtune=native ${CMAKE_CXX_FLAGS}")
set(CMAKE_C_FLAGS "-march=native -mtune=native ${CMAKE_C_FLAGS}")

#Bring the headers, such as library.h into the project
include_directories(include)

#However, the file(GLOB...) allows for wildcard additions:
file(GLOB SOURCES "src/*.cpp" "src/*.c")
file(GLOB TEST_SOURCES "test/test_library.cpp")

#Generate the shared library from the sources
#add_library(c_blake STATIC ${SOURCES})

set(EXECUTABLE_OUTPUT_PATH "bin")
add_executable(test_c_blake ${SOURCES} ${TEST_SOURCES})
