Bind is in early development! There are many rough edges. Please file issues for any bugs you find.

----------
What it looks like:
----------
![Alt Text](bg.png?raw=true "will")
----------
Gifs made with bind: (to see vids go to http://reddit.com/r/deepdreamvideo)
----------

![Alt Text](https://media.giphy.com/media/3oFzmnlg0UXEgkNGh2/giphy.gif)
![Alt Text](https://media.giphy.com/media/xULW8CulD7x86n4Hdu/giphy.gif)
![Alt Text](https://media.giphy.com/media/3oFzmf2YjR0CskBB1m/giphy.gif)


-----------
Requirements:
-----------

Tensorflow (pip install tensorflow, or for an optimized version, use a .whl for your OS)
possible shortcuts to get an optimized tensorflow installation: 
some whl's for tensorflow here:
https://github.com/lakshayg/tensorflow-build
and more to be found at http://ci.tensorflow.org

ffmpeg (brew install ffmpeg, or apt-get install ffmpeg)

ffprobe

wget

python3 

go


----------
Installation:
----------

either go get OR dep ensure 

go build

 ./bind

point browser to localhost:8080


------------
Want to run jobs from your cell phone?
------------

download https://ngrok.com/download  -> start it with:

./ngrok http 8080

ånd †hen your terminal will display the ngrok url to use on your cell phone, for example: http://c55d5584.ngrok.io     Type that into your phone. 


------------------
donations accepted:  dreamlyteam@gmail.com
------------------

LTC:
 Lfoa64kkZS9gDihVA9PEQY46NjUvhZBs9p

BCH:
18tvu2tf6ZcFKEFDP26Q5gBmNa5d62q9jM

ETH:
0x876b28d1aB248A9E05D7a2ef904095987c83E516

BTC:
16bfdgzL5st8bVPGV2JCerkgSAMXNuciEE


Venmo:
https://venmo.com/Tim-Bone

------------------
Join, Debug, Follow, Apply, Like,:
------------------

https://www.facebook.com/dreamly.cc/

https://dreamlycc.slack.com

https://join.slack.com/t/dreamlycc/shared_invite/enQtMjg3NzMxOTQ3OTg5LTNjNjU1ZWYzMGFmMTFkZWY5ZTRhMWY0MDM5NTRiMzI3NmI1MGE2NGMyOWI5MTU1YTQ4ZjUwN2YxNWU5ODMyYTc

https://docs.google.com/forms/d/1KvOZyPu3QamuF_YZtyrCbpS7JbC1oxoyUB3GZVrfVLY/edit

------
We use
-------
https://github.com/git-up/GitUp/wiki/Using-GitUp-Advanced-Commit-View

https://github.com/Microsoft/vscode-go


------
Roadmap
------
√ Optical flow (see code at: https://github.com/ksk-S/DeepDreamVideoOpticalFlow/blob/master/dreamer.py)

√ Clear steps to reproducible builds for any computer.

√ Make deep dreaming accessible as art

√ Explore creative approaches to using ML in art


-------------
Wishlist
-------------
• A tool that helps people install tensorflow-gpu for macs


License
-------------

MIT 2017  Tim Bone
