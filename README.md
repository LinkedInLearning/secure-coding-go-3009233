# Secure Coding in Go
This is the repository for the LinkedIn Learning course Secure Coding in Go. The full course is available from [LinkedIn Learning][lil-course-url].

![Secure Coding in Go][lil-thumbnail-url] 

The Go programming language is growing in popularity. Unfortunately, security issues for Go applications are also on the rise. In this course, learn how to secure your Go application to prevent hackers from stealing data or crashing it. Instructor Miki Tebeka explains why developers should care about security, as well as how to prevent SQL injection attacks, identify places where sensitive data is stored and avoid exposing it to the outside world, write code that protects the integrity of the system, and more. Along the way, he provides challenges that allow you to put your new skills to the test.

## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Go SDK](https://golang.org/dl/)
	- [Git](https://git-scm.com/)
	- [Docker](https://www.docker.com/)
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.
3. Change directory to the root and run `go mod download`


### Instructor

Miki Tebeka 
                            


                            

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/miki-tebeka).

[lil-course-url]: https://www.linkedin.com/learning/secure-coding-in-go
[lil-thumbnail-url]: https://cdn.lynda.com/course/3009233/3009233-1637778695534-16x9.jpg

