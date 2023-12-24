#ifndef SEGMENTTREE_SEGMENTTREE_H
#define SEGMENTTREE_SEGMENTTREE_H
#include <iostream>
#include <vector>

// todo add comparator || lambda function to determine the comparing strategy by the client
class SegmentTree {
public:
    SegmentTree(std::vector<int>arr){
        this->arrSize = arr.size();
        this->tree.resize(calculateTreeSize(arr));
        buildTree(0, arr.size()-1,0,arr);
    }
    int rangeQuery(int queryStart, int queryEnd);
private:
    int arrSize;
    std::vector<int>tree;
    int calculateTreeSize(std::vector<int>arr);
    void buildTree(int start,int end,int idx,std::vector<int>arr);
    int rangeQuery(int queryStart,int queryEnd,int nodeStart,int nodeEnd,int pos);
};

int SegmentTree::calculateTreeSize(std::vector<int>arr) {
    int allocatedSize = 1;
    while (allocatedSize < this->arrSize){
        allocatedSize <<= 1;
    }
    return allocatedSize * 2 -1;
}
void SegmentTree::buildTree(int start,int end,int idx,std::vector<int>arr) {
    if(start > end){
        return;
    }
    if(start == end){
        this->tree[idx] = arr[start];
        return;
    }
    int mid = (start + end) / 2;
    buildTree(start,mid,idx*2+1,arr);
    buildTree(mid+1,end,idx*2+2,arr);
    this->tree[idx] = std::min(tree[2*idx+1],tree[2*idx+2]);// todo replace with comparator here to decide which is pushed to the vector;
}

int SegmentTree::rangeQuery(int start, int end) {
    return rangeQuery(start,end,0,this->arrSize,0);
}

int SegmentTree::rangeQuery(int queryStart, int queryEnd, int nodeStart, int nodeEnd, int pos) {
    if(queryStart <= nodeStart && queryEnd >= nodeEnd ){// query totally overlap node
        return this->tree[pos];
    }
    if(queryStart > nodeEnd || queryEnd < nodeStart){// no overlap
        return 1e9; // invalid value
    }
    // partial overlap
    int mid = (nodeStart + nodeEnd)/2;
    int left = rangeQuery(queryStart,queryEnd,nodeStart,mid,2*pos+1);
    int right = rangeQuery(queryStart,queryEnd,mid+1,nodeEnd,2*pos+2);
    return std::min(left,right);
}


#endif
