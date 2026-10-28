#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct CBNode {
    T value;
    CBNode<T>* left = nullptr;
    CBNode<T>* right = nullptr;
};

template <typename T>
struct CompleteBinaryTree {
    CBNode<T>* root = nullptr;
    int size = 0;

    void TInsert(const T& v) {
        root = insertAt(root, size, 0, v);
        ++size;
    }

    CBNode<T>* insertAt(CBNode<T>* r, int target, int idx, const T& v) {
        if (!r) {
            if (idx == target)
                return new CBNode<T>{ v, nullptr, nullptr };
            return nullptr;
        }
        if (2 * idx + 1 <= target)
            r->left = insertAt(r->left, target, 2 * idx + 1, v);
        if (2 * idx + 2 <= target)
            r->right = insertAt(r->right, target, 2 * idx + 2, v);
        return r;
    }

    void TFind(const T& v) const {
        int depth = 0;
        if (findRec(root, v, depth))
            cout << "Элемент был найден на глубине " << depth << endl;
        else
            cout << "Элемент не был найден в дереве" << endl;
    }

    bool findRec(CBNode<T>* r, const T& v, int& depth) const {
        if (!r)
            return false;
        else if (r->value == v)
            return true;
        else if (findRec(r->left, v, depth)) {
            ++depth;
            return true;
        }
        else if (findRec(r->right, v, depth)) {
            ++depth;
            return true;
        }
        return false;
    }

    void TPrint() const {
        cout << "T[";
        inorder(root);
        cout << "]" << endl;
    }

    void inorder(CBNode<T>* r) const {
        if (!r)
            return;
        inorder(r->left);
        cout << r->value << " ";
        inorder(r->right);
    }

    void TPrintTree() const {
        printTree(root, 0);
    }

    void printTree(CBNode<T>* r, int level) const {
        if (!r)
            return;
        printTree(r->right, level + 1);
        for (int i = 0; i < level; ++i)
            cout << "\t";
        cout << r->value << endl;
        printTree(r->left, level + 1);
    }

    bool TIsComplete() const {
        return isCompleteRec(root, 0, size);
    }

    bool isCompleteRec(CBNode<T>* r, int i, int n) const {
        if (!r)
            return true;
        if (i >= n)
            return false;
        return isCompleteRec(r->left, 2 * i + 1, n)
            && isCompleteRec(r->right, 2 * i + 2, n);
    }

    bool TIsFull() const {
        return isFullRec(root);
    }

    bool isFullRec(CBNode<T>* r) const {
        if (!r)
            return true;
        if (!r->left && !r->right)
            return true;
        if (r->left && r->right)
            return isFullRec(r->left) && isFullRec(r->right);
        return false;
    }

    void clear() {
        destroy(root);
        root = nullptr;
        size = 0;
    }

    void destroy(CBNode<T>* r) {
        if (!r)
            return;
        destroy(r->left);
        destroy(r->right);
        delete r;
    }
};
