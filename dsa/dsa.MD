## Data structures and algorithms	
    
It’s worth noting that, acing coding interviews comprises understanding of data structures and problem solving patterns. Always look at space & time complexity and tradeoff between two.

Order of growth or rate of growth.
			Omega <=   f(x). <=  Big O

Theta  - Characterizes the lower and upper bounds.

Counting rules.
	1.  Each simple operation (+, *, –, =, if, call) takes exactly one time step. 	
    2.  Loops and subroutines are not considered simple operations. 
		These are composition of many single-step operations. It doesn’t makes sense of considering sorting as single-step operation, since sorting 1,000,000 items will certainly take much longer than sorting 10 items. 
	The time it takes to run through a loop or execute a subprogram depends upon the number of loop iterations or the specific nature of the subprogram. 
	3. Each memory access takes exactly one time step.


## Data structures

Contiguous memory blocks (array, matrix) and Linked (Pointer) data structures.

Fundamental abstract types. 

Containers - To organize numerous items independent of content. List, Set etc… Static size data points can be represented by Arrays, Matrix(2D), Tensors(ND), whereas Resizable / Dynamic items
	a. Slice (light weight data structure),  ArrayList
	b. Linked list (singular, double, Circular). XOR List, Skip List, Self organized list.
		
	Non-linear : Binary Tree (Complete, Full, Perfect). Lever order insertion/traversal.
			 BST (AVL tree, RED-Black Tree. Splay Tree)
	
Numeric range queries :  Prefix sum, Segment and Binary Indexed Trees(BIT) (Full binary trees)
Membership queries : Sets, Bitset, Bloom filters.
		
Dictionaries : Can be implemented with Hash Table, BST, SkipLists and  Disk based ( BTree & B+Tree, SSTables, LSM tress)
	    
	Priority queues (Heaps) -  Organizes data as per items priority, which can be minimum or maximum. Heaps are  complete Binary Trees(CBT) and can be represented by arrays or linked structures. Operations are  Inserting(x) , findMin()/findMax(). deleteMini() or deleteMax()	

Treap 

Retrieval based structures. 

Stack / LIFO - Can be put into practice by using both Array or List. Chief operations are push(x), pop() , top(), empty(). Used in Depth first search (DFS).
Queue /  FIFO -  Minimize the maximum time spent waiting. Primarily used in Bread first search (BFS). 
	    Enqueue(x) - Adding new element at the back of the queue.
	    Dequeue()  - Removing element from the front of the queue.
		
Geometry
	        Points, Polygons , KD-Tree.

Text pattern matching or keyword lookup
	Trie or Prefix tree- Useful for key look up and prefix queries.
	Suffix Tree & Suffix array. - Useful for implementing pattern matching queries.
	Effective implementation of text indexing : Finite state automata & Finite state transducer.

## Problem Analysis patterns.

Until problem is understood thoroughly and written down approaches don’t jump on to code.

Read through problem, if it is feasible try to formulate into expression or equation with boundary conditions.
	Ex :
	Triplet sum. x+y+z=sum. 	i>-1 or i<n
	 Largest difference in array =MaxNum-MinNum, smallest.
	Max(s[i-j], S[j:k])
	Sum of the first numbers (n*(n+1))/2

	Problems with count, min or max (optimization problems), shortest, longest words means that there several solutions/options for getting an output. These kind of problems are categorized as optimization problems. The objective here is to minimizes or maximizes the value of some parameter. Primarily involves recursion, so tries to break down the problem into smaller inputs of the same function.

   Min(f(x)) &  Max(f(x))

2.  Picturize the problem and identify problem type. Model the problem mathematically with variables.

	Majority of the algorithms designed to work on rigorously defined abstract structures such as permutations, combinations, graphs, and sets. To exploit the algorithms literature, you must learn to describe your problem abstractly, in terms of procedures on fundamental structures.

	1. Permutations and combinations. (Counting)
	2. Subsets or Subarrays. (Summations). (Interval problem)
	3. Trees
	4. Graphs 
	5. Points
	6. Polygons
	7. Strings.

3.	Jot down combinations of problem inputs with different use cases ( as test cases ).
	Edit distance 
	S1:  “”	,”abac”
	S2:  “abac”,    “” 
	
	What if input is empty?
	What is the starting point and endpoint?
	If there are multiple data points what do you do at each step?

Concrete example : Manually solve concrete instances of the problem and then build general solution.
Case analysis : Split input into multiple cases and solve each in isolation.
Iterative refinement : Most problems can be solved used Brute-force approach. Find that solution and improve upon it.
Reduction :  Use well-known solution to some other problem as a sub routine.
Graph modeling : Describe the problem using graph and solve it using existing algorithm.


Algorithm Design patterns 

1. Brute Force approach.
	Also known as exhaustive search or generate-and-test. It’s straightforward approach to solve a problem with disregards computing resources. Mostly involves iteration. Nevertheless, by identifying pattern in input or  re-ordering data (sorting) or using right data structures, time and space complexity of algorithm can be minimized.

Use below algorithm design patterns sequentially to see can they be applied?

	Iteration starts with value of input at a base / end case, then successively apply the recursive definition to find the values of the function at larger inputs.
	Sum=sum+a{I}
If input is array try to use
	a. Single pointer with beginning either at left or right.
	b. Two pointer approach and indexers movements. (Two ends, Sequential). 
          c. Predominantly for pairs, merging, duplicate removal. non repeating. 
	d. Sliding windows.

2. Sorting.
	Attempt to detect pattern in input data by sorting. For instance sorting of array cuts back time complexity of operations like min(), Max(), Search(key) etc…
	Sorting and Binary Search, 
	Kth smallest or largest use PriorityQueue (min/max).
	
3. Hashing :
	Occurrence (counting), Look up or membership queries can be tackled with Hashing.
	
	1. Hash table (array of size n) - Value is linked list items with same hash value (Chaining ). (Open addressing) 
	2. Hash function.  h(x) => {x /n |  maps to index of hash table}
	
	Hash function should be 1. Efficiently computable. 2. Uniformly distribute the keys. 3. Minimize collisions. 4. Have a low load factor.
4. Recursion
	An algorithm is called recursive if it solves a problem by reducing in to instance smaller input same function. Many useful algorithms are recursive in nature. Code can be written succinctly.  Strings, Tree problems are recursive problems.

	Use recursion if outcome of current step computation is being added to next instance of the problem. If not we can use simple iteration. In recursion final result is achieved by caching the result in method stack. Whereas in iteration temporary / intermediary variables are used.

	1. Head recursion,
	2. Tail recursion (This can be easily replaced with iteration)
	3. Top down
	4. Bottom up
	5 Caching.
	6. Precomputing and using these smaller solutions to build bigger solutions.
	
	f (n) = n   (+,*, -)  f(n-1)
 	Edit distance :     d = 1 + Max(delete, insert, update)

5. Divide-and-conquer:  (Non overlapping recursive problems) 

	Recursively breaks the problem into smaller instances to solve them at smaller inputs, then combine to form final solution.   Divides the problem into disjoint subproblems. 
  Divide -> Conquer -> Combine.
  Merge Sort and Quick sort.

6. Dynamic programming: 
	See for problems of Counting, Minimum, Maximum, Longest or shortest.

	When problem being solved using recursion and if there are overlapping subproblems, then it is time to consider Dynamic programming. Below are the approaches
	1. Top-down with memoization =>  Recursion + Cache
	2.  Bottom-up Built up the sub problems to combine to final solution. Solves the each sub-problem once then save the result into a table (1D, 2D)
	Time - Memory trade off

7. Greedy algorithms: Compute solution in stages, making choices that are locally optimum at step. These choices 				are never undone. Greedy algorithms provide solutions for optimization problems. This algorithm always makes the choice that looks best at the moment. This locally optimal choice leads to a global optimal solution.

Elements of Greedy algorithms
	1. Determine optimal substructure of the problem.
	2. Develop a recursive solution 
	3. Show that if you make greedy choice then only one sub problem remains. 
	4. Prove that it is always safe to make a greedy choice. 
	5. Develop a recursive algorithm that implements greedy strategy. 
	6. Convert the recursive algorithm to an iterative algorithm. 

 Ex. 	1. Finding a shortest path between two cities with small mileage.
	2. Encoding a messing with fewer number of bit. Ex. Huffman coding.	
	3. Choosing denominations - Change making algorithm.
	4. Activity selection problem.
	5. Minimum spanning trees. 
	6. Offline-caching. 

8 Back tracking 
	Combinations or permutations problems can be solved with this approach. It’s a systematic way to iterate through all the possible configurations of a search space exactly once. Use extra memory to record visiting of configuration.

	Exploring the vertices of a graph via depth-first search, also known as backtracking,

Decision, Optimization, Enumeration problems can be solved with this approach.	2. problem - Searching for best solution.