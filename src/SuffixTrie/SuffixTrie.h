//
// Created by Ismail Magdy on 07/11/2023.
//

#ifndef TRIES_SUFFIXTRIE_H
#define TRIES_SUFFIXTRIE_H
#include <iostream>
#include "SuffixTrieNode.h"

class SuffixTrie {
public:
    SuffixTrie(std::string word);
    bool containSubString(std::string substring);
    bool containSuffix(std::string suffix);
    ~SuffixTrie();
private:
    SuffixTrieNode* root ;
    void insertSuffix(std::string);
};


#endif //TRIES_SUFFIXTRIE_H
