## 1-1 What are the minimum and maximum numbers of elements in a heap of height h?

The minimum nubmer of elements is $2^{h-1}$.
The maximum nubmer of elements is $2^{h} - 1$

## 1-2 Show that an n-element heap has height $\lfloor lg n \rfloor$
pass

## 1-3 Show that in any subtree of a max-heap, the root of the subtree contains the largest value occurring anywhere in that subtree.
pass

## 1-4 Where in a max-heap might the smallest element reside, assuming that all elements are distinct?
Any leaf nodes.

## 1-5 At which levels in a max-heap might the k th largest element reside, for $2 <= k <= \lfloor n / 2 \rfloor$, assuming that all elements are distinct?

## 1-6 Is array that is in sorted order a min-heap?
Yes

## 1-7 Is the array with values (33, 19, 20, 15, 13, 10, 2, 13, 16, 12) i a max-heap?
No

33
|\
19 20
| \  \ \
15 13 10 2
| \  \
13 16 12

16 is bigger than its parents 15

## 1-8 Show that, with the array for storing an n-element heap, the leaves are the nodes indexed by $\lfloor n/2 \rfloor + 1$, $\lfloor n/2 \rfloor + 2$, ..., $n$.