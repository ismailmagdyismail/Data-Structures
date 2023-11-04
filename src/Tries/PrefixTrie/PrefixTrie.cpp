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
    if(word.empty() || !containWord(word)){
        return;
    }
    deleteWord(this->root,word,0);
}
void PrefixTrie::deleteWord(PrefixTrieNode *node, std::string word, int index) {
    // todo fix memory leak bug
    if(!node || index > word.size()){
        return;
    }
    deleteWord(node->children[word[index]],word,index+1);
    if(index == word.size() - 1){
        node->isWord = false;
    }
    if(!node->isWord && node->children.empty()){
        delete node;
    }
}
void PrefixTrie::print() {
    printNodes(this->root);
}
void PrefixTrie::printNodes(PrefixTrieNode *node) {
    if(!node){
        return;
    }
    vector<PrefixTrieNode*>children;
    for(auto entry : node->children){
        children.push_back(entry.second);
    }
    for(auto child :children){
        printNodes(child);
    }
}
PrefixTrie::~PrefixTrie() {
    deleteNodes(this->root);
}
void PrefixTrie::deleteNodes(PrefixTrieNode* node){
    if(node == nullptr){
        return;
    }
    for(auto child : node->children){
        deleteNodes(child.second);
    }
    delete node;
}
