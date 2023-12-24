#ifndef INTERVAL_BST_TREE_INTERVALTREE_H
#define INTERVAL_BST_TREE_INTERVALTREE_H
#include <iostream>
using namespace std;

struct Interval{
    int low;
    int high;
};
struct Node{
    Interval interval;
    Node* left;
    Node* right;
    int highestSubTreeEndPoint = -1;
    Node(int low,int high,int highestSubTreeEndPoint){
        this->interval.low = low;
        this->interval.high = high;
        this->highestSubTreeEndPoint = highestSubTreeEndPoint;
    }
};

class IntervalTree {
public:
    void addInterval(int low ,int high);
    bool existIntersection(int low, int high);
    IntervalTree findIntersectingInterval(int low,int high);
    ~IntervalTree();

private:
    Node* root = nullptr;
    void addInterval(Node* node,int low ,int high);
    bool existIntersection(Node* node,int low,int high);
};

void IntervalTree::addInterval(int low, int high) {
    if(this->root == nullptr){
        this->root = new Node(low,high,high);
        return;
    }
    addInterval(this->root,low,high);
}

void IntervalTree::addInterval(Node* node,int low,int high){
    node->highestSubTreeEndPoint = max(high,node->highestSubTreeEndPoint);
    if(low < node->interval.low){
        if(node->left != nullptr){
            addInterval(node->left,low,high);
            return;
        }
        Node* newNode = new Node(low,high,high);
        node->left = newNode;
    }else{
        if(node->right != nullptr){
            addInterval(node->right,low,high);
            return;
        }
        Node* newNode = new Node(low,high,high);
        node->right = newNode;
    }
}

bool IntervalTree::existIntersection(int low, int high) {
    return existIntersection(this->root,low,high);
}
bool IntervalTree::existIntersection(Node* node,int low, int high) {
    if(node == nullptr){
        return false;
    }
    if(low <= node->interval.high && high >= node->interval.low){
        return true;
    }
    if(node->left == nullptr){
        return existIntersection(node->right,low, high);
    }
    if(node->left->highestSubTreeEndPoint < low){
        return existIntersection(node->right,low,high);
    }
    return existIntersection(node->left,low,high);
}

IntervalTree::~IntervalTree() {
    /// todo Free Allocated Nodes
}


#endif //INTERVAL_BST_TREE_INTERVALTREE_H
