# Simple-diff
This is a simple function for finding difference between two strings in go language.


Sample/Usage is given in the file

Function returns an array of int containing starting position and ending position of sub-substrings which are different.One after the other. If ending position is less than starting position, it means that the other string has a substring inserted at this position.

Output looks like

--------------------------------------------------
     old word marker-: [4 4 6 6]
     new word marker-: [4 5 7 6]

    testing my program
        - - 
    test on my program
        -- ^
       _ is a change 
       ^ is a insertion 
-------------------------------------------------
