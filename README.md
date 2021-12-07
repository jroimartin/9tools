# 9tools

Tools I use with [p9p](https://github.com/9fans/plan9port).

## Dependencies

- editinacme: `go install 9fans.net/go/acme/editinacme@latest`
- [Go fonts](https://go.dev/blog/go-fonts)

## Tips

### Acme

- Delete all text in window: `Edit ,d`
- Search and replace:
	- All: `Edit ,x/old/c/new/`
	- Selection: `Edit x/old/c/new/`
	- Next: `Edit /old/c/new/`

### Common

- Format text to 80 columns: `| 9 fmt -j -w 80`

## Resources

Interesting links about using Acme, Sam and other plan9 programs:

- [A Tour of Acme](https://research.swtch.com/acme) & [HN thread](https://news.ycombinator.com/item?id=10957576)
- [Acme: A User Interface for Programmers](https://research.swtch.com/acme.pdf)
- [The Text Editor sam](https://research.swtch.com/sam.pdf)
- [A tutorial for the sam command language](http://doc.cat-v.org/bell_labs/sam_lang_tutorial/sam_tut.pdf)
