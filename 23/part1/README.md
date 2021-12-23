The code here does not build at all and it probably doesn't make much sense. That's because I didn't use any code for this task, I just used my reasonably wrinkly brain and some paper. So instead of code you get an explanation of my thought process, enjoy!

My input looks like this:
#############
#...........#
###B#B#D#D###
  #D#C#B#A#
  #D#B#A#C#
  #C#C#A#A#
  #########

First a general observation:
1. The B column consists of only B and C, so if the C column can be emptied then the B column kan be emptied very easily.

Then I realized two things:
2. While there is a D in the A column I can't put an A in the middle of the hallway, because it will block the D's in the A column forever (unless they can escape to the left, as we will see later ;))
3. Same thing as above but for D while there is A's in the D column.

These realizations also lead to the realization that A and D must be packed into the edges of the hallway to not block everything else. I tried this for a long time and finally realized another thing:
4. A or D must be _inside_ D or A to not block eachother.

Because the hallway is sort of like a stack you have to put the letters in the correct order so you can retrieve them in the correct order later. This lead to a state where I could empty the C column like this:

#############
#B......A.AD#
###B#B#.#D###
  #D#C#.#A#
  #D#B#.#C#
  #C#C#.#A#
  #########

Notice that the A's in the hallway are _inside_ the D to the right. So the A-column needs to be cleared right after the B and C are cleared. After clearing B and C I got this state:

#############
#DD.....A.AD#
###.#B#.#D###
  #.#B#C#A#
  #.#B#C#C#
  #.#B#C#A#
  #########

Notice the two D's to the left that we're put there from the A column. This allows the correctly stacked A's in the hallway to slide into the A column. After this the solution is trivial.