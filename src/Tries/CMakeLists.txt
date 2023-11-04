cmake_minimum_required(VERSION 3.24)
project(Tries)

set(CMAKE_CXX_STANDARD 17)

add_executable(Tries main.cpp ITrie.h PrefixTrie/PrefixTrie.cpp PrefixTrie/PrefixTrie.h PrefixTrie/PrefixTrieNode.h)
