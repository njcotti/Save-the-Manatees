# PLC-HW-1
Manatee sorting by gender, age, and size. Written in GoLang. Assignment creation credits go to Dr. R Stansifer of Flroida Tech. Found at: https://cs.fit.edu/~ryan/cse4250/projects/evacuation/

Professor Smith of the FIT department of marine biology is preparing to evacuate a large number manattees from the IRL before they starve death. Professor Smith has ordered them to line up by age in two rows with the males in the first row and the females in the second row. Conveniently there are the same number in each row. The youngest ones are on the left. If two manatees have the same age, then the their order in the row is immaterial.

Male manatees are generally smaller than the females of the same age. In fact, while males tend to grow only up to around 9 feet in length and 1,200-1,800 pounds in weight, female manatees are about 10-13 feet in length and can weigh up to 3,500 pounds.

A team of graduate student scuba divers are to be sent into the water in front of the manatees to insure all the manatees are in good order and ready to move out. A problem has arisen when a diver pointed out that they might not be able to see a manattee in the second row if the manttee directly in front is bigger and hides the smaller manatee directly.

Professor Smith comes to you with the data and asks if it possible to arrange the two rows such that each manatee in the back row is bigger than the manatee in front of it. The data includes the age of the manatee and the size of the manatee.

Input and Output All the input comes from the standard input stream. The first line of input contains an integer *n* (1 ≤ *n* ≤ 5 ≤ 109), the number of manatees in each row. The next four lines contain 
 integer each. The first pair of lines represents the female manatees (the back row) and the second pair of lines repreesents the male manatees (the front row). Manatees in each row are tattooed with a number 1 to 
 according to their order in the input. The first line in each pair of lines contains 
 integer *n* integers *s**i* (1 ≤ *s**i* ≤ 109 for each *i*), where 
is the age of the manatee with the tattoo 
. The second line in each pair of lines contains 
 integers 
 ($1\le s_i\le 10^9$ for each 
), where 
 is the size of the manatee with the tatoo 
. If there is a valid ordering, output it as two lines of integers, each consisting of a permutation of the tile numbers from to 
 to 
. The first line represents the ordering of the manatees in the back row and the second represents the ordering of the manatees in the front row. If more than one pair of permutations satisfies the constraints, any such pair will be accepted. If no ordering exists, output “impossible”.

Sample Input 1 
4 
3 2 1 2 
2 3 4 3 
2 1 2 1 
2 2 1 3 

Sample Output 1 
3 2 4 1 
4 2 1 3 

Sample Input 2 
2 
1 2 
2 3 
2 8 
2 1 

Sample Output 2 
impossible 

Submission It is not necessary to write the most efficient solution possible, It takes knowledge of data structures to write the
n (1 ≤ n ≤ 5 ≤ 109)

ai (1 ≤ ai ≤ 109 for each i)

n integers si (1 ≤ si ≤ 109 for each i)

 ($1\le n \le 5 \le 10^9$)

 ($1\le a_i\le 10^9$ for each 
)

 integers 
 ($1\le s_i\le 10^9$ for each 
)
