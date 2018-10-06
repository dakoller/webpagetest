# webpagetest
Tiny docker image with go-script to check length of a website


This repository is a project template, which serves me to 
* develop a Go Script,
* get it compiled in a local docker container, then
* actually build a specific container for this script and 
* run it.

The container parts of the script is taken from https://github.com/treeder/tiny-golang-docker and has the simple task of loading a webpage and printing the length of the HTML page in bytes.

The exercise with the local compilation of the script into an executeable serves to minimize the size of the Docker container.
* without this (with Go compilation happening on the fly in a docker container) the Docker container would have 500MB+ of size.
* In the setup used here the Docker container is below 10MB.

## Related resources

* https://blog.iron.io/an-easier-way-to-create-tiny-golang-docker-images/

## Next steps

I'll extend the script in a way taht it can do a slightly more interesting task: grabbing webpage tests from pingdom.com.
In the process I'll also learn how to integrate git-crypt ( https://github.com/AGWA/git-crypt ) to secure credentials, which are needed to for the operation of the container.  

