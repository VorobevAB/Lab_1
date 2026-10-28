#pragma once
#include <iostream>

using namespace std;

template <typename T>
struct Array {
    T* data;
    int size;
    int capacity;

    void Init() {
        capacity = 1;
        size = 0;
        data = new T[capacity];
    }

    void MPush(const T& n) {
        if (size == capacity) {
            T* nd = new T[capacity * 2];
            for (int i = 0; i < size; ++i)
                nd[i] = data[i];
            delete[] data;
            data = nd;
            capacity *= 2;
        }
        data[size++] = n;
    }

    void MPushIdx(int idx, const T& n) {
        if (idx < 0 || idx > size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return;
        }
        MPush(T{});
        for (int i = size - 1; i > idx; --i)
            data[i] = data[i - 1];
        data[idx] = n;
    }

    T MGet(int idx) const {
        if (idx < 0 || idx >= size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return T{};
        }
        return data[idx];
    }

    void MDelInd(int idx) {
        if (idx < 0 || idx >= size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return;
        }
        for (int i = idx; i < size - 1; ++i)
            data[i] = data[i + 1];
        size--;
    }

    void MSetInd(int idx, const T& v) {
        if (idx < 0 || idx >= size) {
            cout << "ОШИБКА: индекс " << idx << " не существует" << endl;
            return;
        }
        data[idx] = v;
    }

    int MLen() const { return size; }

    void MPrint() const {
        cout << "M[";
        for (int i = 0; i < size; i++)
            cout << data[i] << " ";
        cout << "]" << endl;
    }

    void clear() {
        delete[] data;
        data = nullptr;
        size = 0;
        capacity = 0;
    }
};
