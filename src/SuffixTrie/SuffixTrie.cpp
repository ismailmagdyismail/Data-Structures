#include "SuffixTrie.h"

SuffixTrie::SuffixTrie(std::string word) {
    this->root =  new SuffixTrieNode();
    for(int i =0 ;i<word.size();i++){
        std::string suffix = word.substr(i);
        this->insertSuffix(suffix);
    }
}
bool SuffixTrie::containSubString(std::string substring) {
    SuffixTrieNode* currentNode = this->root;
    for(char i : substring){
        if(currentNode->children[i] == nullptr){
            return false;
        }
        currentNode = currentNode->children[i];
    }
    return true;
}
bool SuffixTrie::containSuffix(std::string suffix) {
    SuffixTrieNode* currentNode = this->root;
    for(char i : suffix){
        if(currentNode->children[i] == nullptr){
            return false;
        }
        currentNode = currentNode->children[i];
    }
    return currentNode->children['$'] != nullptr;
}
void SuffixTrie::insertSuffix(std::string suffix) {
    SuffixTrieNode* currentNode = this->root;
    for(int i =0 ;i<suffix.size();i++){
        if(currentNode->children[suffix[i]] == nullptr){
            currentNode->children[suffix[i]] = new SuffixTrieNode();
        }
        currentNode = currentNode->children[suffix[i]];
    }
    currentNode->children['$'] = new SuffixTrieNode();
}

SuffixTrie::~SuffixTrie() {
    // todo add de-allocation
}


