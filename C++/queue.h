#pragma once
#include <iostream>
#include "forward_list.h"

using namespace std;

template <typename T>
struct Queue {
    Node<T>* head = nullptr;
    Node<T>* tail = nullptr;

    bool QEmpty() const {
        return head == nullptr;
    }

    void QPush(const T& v) {
        Node<T>* n = new Node<T>{ v, nullptr };
        if (!tail) {
            head = n;
            tail = n;
        }
        else {
            tail->next = n;
            tail = n;
        }
    }

    void QPop() {
        if (QEmpty()) {
            cout << "ОШИБКА: Очередь пуста, удалять нечего" << endl;
            return;
        }
        Node<T>* d = head;
        head = d->next;
        delete d;
        if (!head)
            tail = nullptr;
    }

    T QFront() const {
        if (QEmpty()) {
            cout << "ОШИБКА: Очередь пуста" << endl;
            return T{};
        }
        return head->value;
    }

    void QPrint() const {
        cout << "Q[";
        for (Node<T>* c = head; c; c = c->next)
            cout << c->value << " ";
        cout << "]" << endl;
    }

    void clear() {
        Node<T>* c = head;
        while (c) {
            Node<T>* n = c->next;
            delete c;
            c = n;
        }
        head = nullptr;
        tail = nullptr;
    }
};
