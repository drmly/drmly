Bind is in early development! There are many rough edges. File issues for any bugs you find.

![Alt Text](bg.png?raw=true "will")
----------
Gifs made with bind: (to see vids go to http://reddit.com/r/deepdreamvideo)
----------

![Shapes](https://media.giphy.com/media/xULW8qKMNmfa4RZIPe/giphy.gif)
![Human](https://media.giphy.com/media/3oFzmnlg0UXEgkNGh2/giphy.gif)
![Celebrity](https://media.giphy.com/media/xULW8CulD7x86n4Hdu/giphy.gif)
![Wrestling](https://media.giphy.com/media/3oFzmf2YjR0CskBB1m/giphy.gif)




-----------
Requirements:
-----------

[Tensorflow](https://www.tensorflow.org/install/) (pip install tensorflow, or for an optimized version, use a .whl for your OS)
possible shortcuts to get an optimized tensorflow installation: 
some whl's for tensorflow here:
https://github.com/lakshayg/tensorflow-build
and more to be found at http://ci.tensorflow.org

[ffmpeg](https://ffmpeg.org/download.html)

ffprobe (probably installed along with ffmpeg)

[wget](https://www.gnu.org/software/wget/)

[python3](https://www.python.org/downloads/release/python-364/)

[go](https://golang.org/)


----------
Installation:
----------

go get -u github.com/dreamlyteam/bind 

cd $GOPATH/src/github.com/dreamlyteam/bind
dep ensure
go build

 ./bind


------------
Usage
------------

point browser to localhost:8080

(because we need to download the Inception model, the first time running a job python will download a 53 megabyte tensorflow_inception_graph.pb into the bind/models dir)

------------
Filetypes supported:
------------

Tries to support anything with pixels automatically, videos of any extension, gifs, and images of course. [Suggest a new filetype](https://github.com/dreamlyteam/bind/issues/new)

------------
Want to take video on your cell phone from exotic locations and send the job to your computer at home? 
------------

download https://ngrok.com/download  -> start it with:

./ngrok http 8080

ånd then your terminal will display the ngrok url to use on your cell phone, for example: http://c55d5584.ngrok.io     Type that into your phone. 

------------------
OSX tensorflow-gpu installation walkthroughs guide...(WIP) 
------------------

// Here be dragons. Thou art forewarned (difficulty level: Shao Kahn from MK2 tough.)


[maybe try this one](https://metakermit.com/2017/compiling-tensorflow-with-gpu-support-on-a-macbook-pro/)

[or maybe this ones better](https://medium.com/@mattias.arro/installing-tensorflow-1-2-from-sources-with-gpu-support-on-macos-4f2c5cab8186)

Or maybe something like a .whl would make it faster figuring this out?  [try this](https://github.com/qinyuhang/tensorflow-mac-build)

[Improve this guide](https://github.com/dreamlyteam/bind/issues/new)

------------------
[Share your results on Reddit](http://reddit.com/r/deepdreamvideo) [Debug on Slack](https://dreamlycc.slack.com)
------------------




------
We use
-------
https://github.com/git-up/GitUp/wiki/Using-GitUp-Advanced-Commit-View

https://github.com/Microsoft/vscode-go


------
Roadmap
------
√ Optical flow (see code at: https://github.com/ksk-S/DeepDreamVideoOpticalFlow/blob/master/dreamer.py)

√ More parameters and more parameter automations implemented

√ Clear steps to reproducible builds for any computer.

√ Make deep dreaming accessible as art

√ Explore creative approaches to using ML in art


-------------
Wishlist
-------------
• A tool that helps people install tensorflow-gpu for macs


------------------
Donations accepted, send me an email, just decrypt it first: ∂®´åµ¬¥†´åµ™©µåˆ¬≥çøµ 
------------------

LTC:
 Lfoa64kkZS9gDihVA9PEQY46NjUvhZBs9p

BTC:
16bfdgzL5st8bVPGV2JCerkgSAMXNuciEE

BCH:
18tvu2tf6ZcFKEFDP26Q5gBmNa5d62q9jM

ETH:
0x876b28d1aB248A9E05D7a2ef904095987c83E516



Venmo:
https://venmo.com/Tim-Bone
