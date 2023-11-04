#ifndef TRIES_ITRIE_H
#define TRIES_ITRIE_H
#include <iostream>

class ITrie {
public:
    virtual void insertWord(std::string word) = 0 ;
    virtual void deleteWord(std::string word) = 0 ;
    virtual bool containWord(std::string word) = 0 ;
    virtual bool containPrefix(std::string word) = 0;
};


#endif //TRIES_ITRIE_H
