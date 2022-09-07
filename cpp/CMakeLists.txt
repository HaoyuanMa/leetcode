cmake_minimum_required(VERSION 3.23)
project(mhy)

set(CMAKE_CXX_STANDARD 14)

include_directories(utils)
add_executable(q-21-merge-two-sorted-lists q-21-merge-two-sorted-lists/q-21-merge-two-sorted-lists.cpp)
add_executable(q-86-partition-list q-86-partition-list/q-86-partition-list.cpp)
add_executable(q-2379-minimum-recolors-to-get-k-consecutive-black-blocks q-2379-minimum-recolors-to-get-k-consecutive-black-blocks/q-2379-minimum-recolors-to-get-k-consecutive-black-blocks.cpp)
add_executable(q-23-merge-k-sorted-lists q-23-merge-k-sorted-lists/q-23-merge-k-sorted-lists.cpp)