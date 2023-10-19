# RANDOM DELAY
## A Code Snippet in GoLang

I spent hours tonight, trying to piece together exactly how math/rand's rand.Intn() and time's time.Sleep() worked together, going off of old stack overflow posts that led me on several wild goose chases.

So, to save both my future self, and anyone else who finds themselves in need
of a randomized delay in their app, program, script, etc. I put together
this fully co,mmented code snippet.

I am releasing it under the MIT lisence, except that you DO NOT
need to include a copy of the lisence if you use any or all of the provided code.

I am providing it AS-IS and can't guarantee that the Go team won't break it in the future (as evidenced by my hours of researching such a simple process!)

If in the future it does break, though, I will do my best to update it!  I hope someone besides myself is able to find this useful, but, if not, I have a few APIs planned that I think will benefit nicely.