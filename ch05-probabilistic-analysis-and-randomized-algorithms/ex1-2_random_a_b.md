```
def RANDOM(A, B):
    result = A
    FOR i FROM 0 to (B - A + 1)
        result += RANDOM(0, 1)
    RETURN result
```
The runtime of this version of RANDOM(a, b) is `O(f) = B - A`