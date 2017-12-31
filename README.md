drmly is a cross platform web interface for creatively running deep dream on anything with pixels.

Are you an artist looking to use dd to do art? This project's goal is make that easy for you! Pull requests and issues are welcome here. 

![Alt Text](bg.png?raw=true "will")
----------
Gifs made with drmly: (to see vids go to http://reddit.com/r/deepdreamvideo)
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

[cv2](https://duckduckgo.com/?q=install+cv2&ia=qa)

[ffmpeg](https://ffmpeg.org/download.html)

ffprobe (probably installed along with ffmpeg)

[wget](https://www.gnu.org/software/wget/)

[python3](https://www.python.org/downloads/release/python-364/)

[go](https://golang.org/)

Optional requirements:

[youtube-dl](https://rg3.github.io/youtube-dl)

[dep](https://github.com/golang/dep)

----------
Installation:
----------

`go get -u github.com/drmly/drmly `

`cd $GOPATH/src/github.com/drmly/drmly`

`go install`

`drmly`


------------
Usage
------------

point browser to localhost:8080

(because we need to download the Inception model, the first time running a job python will download a 53 megabyte tensorflow_inception_graph.pb into the drmly/models dir)

If you have any problems, watch the terminal output from drmly, if there's errors they will be listed as error



------------
Filetypes supported:
------------

Tries to support anything with pixels automatically, videos of any extension, gifs, and images of course. [Suggest a new filetype](https://github.com/drmly/drmly/issues/new)

------------
Want to make pixels on your cell phone from exotic locations and send the job to your computer at home? 
------------

download https://ngrok.com/download  -> start it with:

./ngrok http 8080

ånd then your terminal will display the ngrok url to use on your cell phone, for example: http://c55d5584.ngrok.io     Type that into your phone and then you should see the bind localhost:8080 web portal running on your local computer.

-----------
Clipboard Support
------------

Bind also works with youtube-dl, allowing you to run jobs from the clipboard on youtube-dl compliant URL's. Just copy any semi popular website and customize the job at localhost:8080. Make sure the checkbox yt-dl clipboard is checked.

-----------
Screenshot Support
------------

Any image files added to the Desktop will automatically be randomly deep dreamed. Turn this feature off by commenting out or deleting screenshot.go

Want to do batch jobs? You can also download or move many images at once to the desktop and drmly will finish one at a time. 

------------------
OSX tensorflow-gpu installation guide (this is an optional optimization, only suggested if you're running many jobs) 
------------------

Tf doesn't support GPU on Macs out of the box (after v1.1). There are several solutions to this problem, and lot's of blog posts out there to consider. Imo the best place to start figuring this out is to find and install the perfect .whl for your OS version (10.12, 10.11, 10.9, etc...), and then install CUDA and Cudnn after that.  That's because every .whl you find out there is going to have different Cuda and Cudnn versions that it wants. Why not install from source? It's slow and unnecessary to be building the .whl yourself from source (which is what most articles focus on). Also, be aware that you'll be wanting to use a .whl that supports your Cuda Compute Capability, in my case I'm running a Geforce 650M which has a 3.0 Cuda Compute Capability. This .whl works for my system:

sudo pip install --upgrade https://github.com/bodak/tensorflow-wheels/releases/download/v1.1.0/tensorflow-1.1.0_GPU-cp36-cp36m-macosx_10_7_x86_64.whl

Now, assuming you didn't get an error that this .whl is not for your OSX version, move on to the next step.

We can see from this [github page](https://github.com/bodak/tensorflow-wheels/releases) that we need to install specific CUDA and Cudnn versions, 8.0 and 5.1 respectively.

[Next, Looking at this guide](https://metakermit.com/2017/compiling-tensorflow-with-gpu-support-on-a-macbook-pro/) you'll see some environmental variables should be set.

If you don't get errors when you run:

`python3 test.py` (test.py is located within this github projec)

then you are running TF on a mac GPU, impressive! 

At this point you shouldn't get any rpath libcudart8.0.0.dylib Image not found errors (or similar errors). If you do [let us know](https://github.com/drmly/drmly/issues/new)

***If this doesn't work easily, I'd recommend reverting back to the CPU .whl presented near the top of this README. Or, just run pip install tensorflow (not optimized, but acceptable for exploration or low numbers of jobs)

[Improve this guide](https://github.com/drmly/drmly/issues/new)

------------------
[Share your art on Reddit](http://reddit.com/r/deepdreamvideo)
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


------------------
Developing
------------------
Development mode started with d argument

drmly d 