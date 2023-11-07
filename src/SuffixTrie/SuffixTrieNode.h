//
// Created by Ismail Magdy on 07/11/2023.
//

#ifndef TRIES_SUFFIXTRIENODE_H
#define TRIES_SUFFIXTRIENODE_H

struct SuffixTrieNode {
    std::unordered_map<char,SuffixTrieNode*>children;
};


#endif //TRIES_SUFFIXTRIENODE_H
