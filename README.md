# go-play-ascii
A project to take a video and render it as ascii in a terminal. When I'm done it will be available as a command line tool and I probably won't make a library version of it because of a limited use case but also how would that work? pipes maybe? lmao

## why go? wasnt C fast enough?
I wasn't happy with the state of cmd-play so I decided to start rewriting it in go because of the concurrency features and the ease of using ffmpeg with go. It seems like it will be just as challenging but go is a much more modern language and not only will I be able to prototype faster in this language, but I will be able to take advantage of modern features for a (hopefully) more smooth experience. Not to mention the modularity potential with go to support more advanced playback features like buffering.

## todo
- create `ffmpeg` bindings because I want the challenge of not using an existing library and I want it to be more extensive than the mock-functions I used in the C version.
- create render function for single frame, send back through channel
    - pipe frame object into manager to determine playback order and pacing 
    - possible speed sampling implementaion available?
- divide frame buffer to multiple goroutines 
- print on recieve or queue into output buffer

## late night thoughts
maybe ill make this a file format because that would be so cool
