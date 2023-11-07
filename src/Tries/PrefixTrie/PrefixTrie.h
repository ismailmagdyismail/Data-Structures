

#ifndef TRIES_PREFIXTRIE_H
#define TRIES_PREFIXTRIE_H
#include "PrefixTrieNode.h"

class PrefixTrie{
public:
    PrefixTrie();
    void insertWord(std::string word);
    void deleteWord(std::string word);
    bool containWord(std::string word);
    bool containPrefix(std::string word);
    void print();
    ~PrefixTrie();
private:
    PrefixTrieNode* root;
    void deleteNodes(PrefixTrieNode* node);
    bool deleteWord(PrefixTrieNode* node,std::string word,int index);
    void printNodes(PrefixTrieNode* node);
};


#endif //TRIES_PREFIXTRIE_H
