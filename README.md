# 9tools

Tools I use with [p9p](https://github.com/9fans/plan9port).

## Dependencies

- editinacme: 9fans.net/go/acme/editinacme

## Tips

### Acme

- Delete all text in window: `Edit ,d`
- Search and replace:
	- All: `Edit ,x/old/c/new/`
	- Selection: `Edit x/old/c/new/`
	- Next: `Edit /old/c/new/`
- Big font:
	- Proportional width: `Font /mnt/font/LiberationSans/18a/font`
	- Fixed width: `Font /mnt/font/LiberationMono/18a/font`

## Resources

Interesting links about using Acme, Sam and other plan9 programs:

- [A Tour of Acme](https://research.swtch.com/acme) & [HN thread](https://news.ycombinator.com/item?id=10957576)
- [rsc on 9fans: Using the Acme Editor](https://groups.google.com/d/msg/comp.os.plan9/_YUEVbTFuME/tJHB8y8-0vYJ)
- [jinyangustc/acme-editor](https://github.com/jinyangustc/acme-editor)
- [evbogdanov/acme](https://github.com/evbogdanov/acme)
