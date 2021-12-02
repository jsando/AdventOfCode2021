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