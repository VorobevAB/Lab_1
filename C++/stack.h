#pragma once
#include <iostream>
#include "forward_list.h"

using namespace std;

template <typename T>
struct Stack {
    Node<T>* head = nullptr;

    bool SEmpty() const {
        return head == nullptr;
    }

    void SPush(const T& v) {
        head = new Node<T>{ v, head };
    }

    void SPop() {
        if (SEmpty()) {
            cout << "ОШИБКА: Стек пуст, удалять нечего" << endl;
            return;
        }
        Node<T>* d = head;
        head = d->next;
        delete d;
    }

    T STop() const {
        if (SEmpty()) {
            cout << "ОШИБКА: Стек пуст" << endl;
            return T{};
        }
        return head->value;
    }

    void SPrint() const {
        cout << "S[";
        for (Node<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void clear() {
        while (head) {
            Node<T>* n = head->next;
            delete head;
            head = n;
        }
    }
};
