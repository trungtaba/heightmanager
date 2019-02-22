# heightmanager
based on binary search
in my scenario, I don't know about height and time value so I will check it value and (height, time) can be anything
## Validate:
-check conflict first
-check height or time is presented in list or not?
  +Based on hashtable to store value: O(1) for checking
-finding heightIndex and timeIndex: take O(logn) time
-return index

##AddNewPair
-Check validate
-Add pair to list
-Add data to hashtable

#nearestHeightBefore
-binary search
-if dont have value return -1;
