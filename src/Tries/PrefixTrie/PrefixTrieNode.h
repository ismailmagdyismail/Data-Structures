//
// Created by Ismail Magdy on 04/11/2023.
//

#ifndef TRIES_PREFIXTRIENODE_H
#define TRIES_PREFIXTRIENODE_H
#include <map>

struct PrefixTrieNode {
    std::unordered_map<char,PrefixTrieNode*>children;
    bool isWord;
};


#endif //TRIES_PREFIXTRIENODE_H
