# Advent of Go 2021

These are my solutions (in Go) to Advent of Code 2021 (https://adventofcode.com/2020), a daily programming challenge.

Most days require an input file:

    go run main.go -d 2 -i day02/input.txt

# Calling It

I used this year's puzzles to practice Test Driven Design (TDD), and kept an informal journal in JOURNAL.md.  I worked on it when I had time, no rush ... after taking a break for the holidays I jumped back in and did a few more days but ran out of motivation on Day 17.

I did improve my Go TDD skills, which was the goal, and I mostly enjoyed it.  I will continue to prefer TDD over after-the-fact unit tests.

One key thing I learned is that the tests you come up with *during* TDD are not necessarily the tests you want *forever*.  In TDD you use the tests to organically "discover" a working solution ... once you've arrived, you need to clean up the tests, look at coverage, fuzzing, benchmarking, etc.  TDD gives you a way to cut through analysis paralysis ... but you still need to edit your final solution once its working, if you want it to be polished.  Imagine writing an essary for school, and you manage to get a minimum string of paragraphs together that qualify as "ok" ... if you care about your grade, you go back and edit it before turning it in.

I still encountered scenarios where the effort to come up with a test seemed to block forward progress rather than enable it.  TDD is easiest when the problem you're solving has a similar pattern to ones you've done before.  For example, I took a break from this to go goof off on another project (https://github.com/jsando/mpu).  It had been a few years since I wrote a lexer and parser from scratch, and trying to do TDD on those really frustrated me.  Now that they are done, I believe I could go back and do it once more using TDD ... maybe I'll make up another toy language and do that, just to see if that's true.

