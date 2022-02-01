Various implementations of the [merge sort](https://en.wikipedia.org/wiki/Merge_sort).

                  |7 3 4 1 6 8 2 5|
                 /                 \
           |7 3 4 1|               |6 8 2 5|
         /          \             /           \
      |7 3|        |4 1|       |6 8|         |2 5|
     /     \      /     \     /     \       /     \
    |7|    |3|   |4|    |1|  |6|    |8|   |2|    |5|
     \     /      \     /     \     /       \     /
      |3 7|        |1 4|       |6 8|         |2 5|
         \          /             \           /
           |1 3 4 7|               |2 5 6 8|
                 \                 /
                  |1 2 3 4 5 6 7 8|
