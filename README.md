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

[dep](https://github.com/golang/dep)

[python3](https://www.python.org/downloads/release/python-364/)

[go](https://golang.org/)

Optional requirements:

[youtube-dl](https://rg3.github.io/youtube-dl)

----------
Installation:
----------

`go get -u github.com/dreamlyteam/bind `

`cd $GOPATH/src/github.com/dreamlyteam/bind`

`dep ensure` 

`go build`

` ./bind`


------------
Usage
------------

point browser to localhost:8080

(because we need to download the Inception model, the first time running a job python will download a 53 megabyte tensorflow_inception_graph.pb into the bind/models dir)

If you have any problems, `grep` the logrus.log file for error 

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
OSX tensorflow-gpu installation guide...(WIP) 
------------------

The best place to start is to find and install the perfect .whl for your OS version (10.12, 10.11, 10.9, etc...), and install CUDA and Cudnn after that.  That's because every .whl you find out there is going to have different Cuda and Cudnn versions that it wants. Why not install from source? It's slow and unnecessary to be building the .whl yourself from source (which is what most articles focus on). Also, be aware that you'll be wanting to use a .whl that supports your Cuda Compute Capability, in my case I'm running a Geforce 650M which has a 3.0 Cuda Compute Capability

So, here's what worked for me (I personally am running osx 10.12 and used this setup to get GPU working):

I used this wheel:

sudo pip install --upgrade https://github.com/bodak/tensorflow-wheels/releases/download/v1.1.0/tensorflow-1.1.0_GPU-cp36-cp36m-macosx_10_7_x86_64.whl

Now, as we can see from that [github page](https://github.com/bodak/tensorflow-wheels/releases), we need a specific CUDA and Cudnn versions, 8.0 and 5.1 respectively.

also [looking at this guide](https://metakermit.com/2017/compiling-tensorflow-with-gpu-support-on-a-macbook-pro/) you'll see some environmental variables are likely needing to be set, maybe just DYLD_LIBRARY_PATH(I think) 

you shouldn't get any rpath libcudart8.0.0.dylib Image not found errors (or similar errors). If you do [let us know](https://github.com/dreamlyteam/bind/issues/new)


If alls well when you run 

`python3 test.py` (test.py is located within this github projec)

then you are running TF on a mac GPU, impressive! 

***If this doesn't work easily, I'd recommend reverting back to the CPU wheel choices presented near the top of this README.

[Improve this guide](https://github.com/dreamlyteam/bind/issues/new)

------------------
[Share your art on Reddit](http://reddit.com/r/deepdreamvideo) or  [Debug on Slack](https://dreamlycc.slack.com)
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

√ Explore creative approaches to using ML in art
