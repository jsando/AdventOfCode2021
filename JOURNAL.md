# Day 1

Here we go again with Advent of Code!

Decided to use Go again this year, but to try to do strict Test-Driven-Design in the implementation.  I've always been a fan of unit tests and any other automated testing approaches ... but somehow I skipped over seriously doing strict TDD in the last 20 years.  As I hear many devs say "I tried it a few times, it has its place but its not how I think."  Re-watching some TDD youtube videos and seeing a few tweets this year got me thinking that maybe I just haven't tried hard enough, so here goes.

I should say the very first challenge is that I have a lifetime of programming NOT using this approach!  It is challenging to not lay out the overall solution in my mind and then start working towards it, but instead think "what's a simple test I can do in like 60 seconds or less and get it to fail and then fix it and then iterate?"  I'm so used to sketching an approach to a puzzle in my mind, and then start implementing, and then figure out how to test and refactor to make that easier.  Run code coverage and see what I missed and add more tests for that.  But TDD flips the process around.

Another challenge I'm seeing with TDD on day1 is that I still want to test with the input file and not just from a unit test.  But I don't want a unit test readig from a file?  I'm sure I'll figure out a better solution in the coming days.

# Day 2

Since I did strict test-driven, introducing a struct and member functions felt natural to make the assertions easier, versus passing in and returning a tuple or passing in pointers to depth/horizontal. If I hadn't done test-driven, I would have ended up with probably a single function ... which I think would be simpler to read and maintain for such a simple problem. 

Also, I still have to test calling it from the CLI with some test input, so even though I have that same test in unit test, its not testing the file parsing code or the hook through main.go and all that ... if I'm going to do that test anyway then why bother duplicate it in a unit test? 

BUT ... part2 changes the rules.  If I'd done a simple function for part1, I can have another simple function for part2 without breaking part1.  But now I have a whole type and member functions that will have to be duplicated.  Over-complicated.  Ugh.

But hey ... I've got unit tests so I can always refactor ;) I won't do much though just so I can read this journal and see the state of the code later on.

Done with day 2 ... feels like I'm writing a lot more code than is warranted for simple problems with the unit tests in there, but I will stick with it.  

# Day 3 

Uh ... gosh this strict TDD thing is really hard.  When I see other people do it, they seem to magically find a simple, incremental test without thinking through the entire problem and then through the series of little increments at some point magically announce "and that's the final solution!".  I can't even IMAGINE what to test without thinking of the overall solution?  Maybe its BS?  Maybe they have just practiced the same "fizz buzz" code dojo so many times they already know the "simple incremental tests" that lead to the final solution?

Side note ... I somehow had Go amd64 installed on my arm64 mac???  No idea how I did that and failed to notice it.  Had to wipe it out and reinstall Go and all my ~/go/bin stuff.

My solution is really bugging me ... there's gotta be a bitwise op to do this versus using an array of counters?

I went back and rewrote into a simple to read function, and ended up making lots of smaller tests ... and the overall coverage is 85%.  Just the error handling paths isn't being tested.  

# Day 4 / day 5

Its kind of natural to take the example straight from the puzzle and make that the unit test ... which is what I lazily did in day 5.  But that's not really using test driven design right?  That's the high-level test, but there should be a series of smaller tests along the way.  Like the first test should just be can I represent the grid, and then plot points, and then check counts, or whatever.  Little steps like that.  I find myself coding the big picture test and then writing code for 15-20 minutes to make it pass ... TDD is supposed to be very short red/green/refactor cycles.  So, I guess i'm still doing it wrong.  Still trying.

# Day 6

Still struggling to figure out incremental tests to derive the solution without a plan up-front.  Prob because my brain is thinking of the solution up front.  I am doing better with tests though, last several days have 90% coverage and I'm not even trying for coverage.  

# Day 7

Same.

Although with the unit tests I'm writing way more code than I would normally for this type of thing, its not really taking me that much longer to complete the puzzles.

# Day 8

OMG I got so stuck on part 2 and TDD was somewhat part of my problem.  Also I was waffling about how to represent segments in the display.  Got focused finally and re-read the problem and finally realized I was hella overthinking it, and approching it all wrong.  Once I had the right approach, the TDD flowed naturally.  I ended up having to write out the algorithm in pseudocode.  TDD was supposed to make it so I didn't have to think it through up front.

# Day 9

Just realized that in doing TDD in Go, I wish I could put the TestXXX function in the same source file as the function itself (ie, solution.go instead of solution_test.go).  In doing strict TDD I write the test function first, then create the functions its calling and that is easiest if I let VSCode create it in the same source file (solution_test.go).  I make the test pass, and then have to manualy move the function over to solution.go.  Would be nice if I could just put the Test functions in the solution.go file :) 

I realize I am doing more test driven design as I go along, and I mean that in the sense of "discovering" the solution along the way.  I haven't been going back to clean this up but realize now that I'm using some "pathfinder" tests to help discover the solution ... in prod code I would def go back and clean that up.  As well as looking at overall coverage, and fuzzing, and ... and ... and.

# Day 12-13 ... a month later!

It's now January 17th!  I got sidetracked by the holidays, then struggled to jump back into day 12.  The struggle was in large part TDD ... the problem was obviously easily solveable with a simple recursive function, but I really struggled to derive it via TDD.  I got there finally though but dang ... it was beating me up.

Day 12 part 2, and day 13 went pretty quick.

I'm mostly been using GoConvey, so all the tests run on every save and it gives me a toast popup with the results within a second.  It sometimes gets lost and I have to restart GoConvey but it mostly works.

I'm writing more code with TDD than I would otherwise, and I don't just mean the unit tests, I mean the overall solution seems like it has more code.  I'm more likely to introduce a struct for example to make it easier to test some things.  If I was doing an actual competition I'd be hammering stright into the solution as quick as possible and not wasting time with TDD.  

Of course, actual *production* code needs maintainability a gagillion times more than it needs some random algorithm implemented in as big a rush as you can.

My code so far is NOT production code, maybe I'm being a bit split-minded here ... though I'm doing TDD, I'm also taking shortcuts for AOC such as hard-coding limits of arrays, or panicking if an i/o returns an error.  And I'm not refactoring my solution once I get it working.  But my goal this year was to practice the TDD part, and I am for sure achieving that.

# Day 14

The unit tests I created were not the ones I ended with, as the solution evolved.  But that's ok, it got me there.

