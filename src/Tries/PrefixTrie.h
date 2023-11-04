//
// Created by Ismail Magdy on 04/11/2023.
//

#ifndef TRIES_PREFIXTRIE_H
#define TRIES_PREFIXTRIE_H

#include "../ITrie.h"
#include "PrefixTrieNode.h"

class PrefixTrie : public ITrie{
private:
    PrefixTrieNode* root;
public:
    PrefixTrie();
    virtual void insertWord(std::string word);
    virtual void deleteWord(std::string word);
    virtual bool containWord(std::string word);
    virtual bool containPrefix(std::string word);
    virtual ~PrefixTrie();

};


#endif //TRIES_PREFIXTRIE_H
