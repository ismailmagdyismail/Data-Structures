//
// Created by Ismail Magdy on 04/11/2023.
//

#include <iostream>
#include "PrefixTrie.h"
using namespace std;

PrefixTrie::PrefixTrie() {
    this->root = new PrefixTrieNode();
    this->root->isWord = false;
}
void PrefixTrie::insertWord(std::string word) {
    PrefixTrieNode* currentNode = this->root;
    for(int i = 0 ;i<word.size();i++){
        if(currentNode->children[word[i]] == nullptr){
            currentNode->children[word[i]] = new PrefixTrieNode();
        }
        currentNode = currentNode->children[word[i]];
        if(i == word.size() - 1){
            currentNode->isWord = true;
        }
    }
}
bool PrefixTrie::containWord(std::string word) {
    PrefixTrieNode* currentNode = this->root;
    for(int i = 0 ;i<word.size();i++){
        if(currentNode->children[word[i]] != nullptr){
            currentNode = currentNode->children[word[i]];
        }else{
            return false;
        }
        if(i == word.size()-1 && !currentNode->isWord){
            return false;
        }
    }
    return true;
}
bool PrefixTrie::containPrefix(std::string word){
    PrefixTrieNode* currentNode = this->root;
    for(int i = 0 ;i<word.size();i++){
        if(currentNode->children[word[i]] != nullptr){
            currentNode = currentNode->children[word[i]];
        }else{
            return false;
        }
    }
    return true;
}
void PrefixTrie::deleteWord(std::string word) {

}
void deleteNodes(PrefixTrieNode* node){
    if(node == nullptr){
        return;
    }
    for(auto child : node->children){
        cout<<"going to delete "<<child.first<<" and its children\n";
        deleteNodes(child.second);
    }
    cout<<"delete\n";
    delete node;
}
PrefixTrie::~PrefixTrie() {
    deleteNodes(this->root);
}
