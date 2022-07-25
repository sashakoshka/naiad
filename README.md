# naiad

![Naiad logo](assets/logo.png)

Naiad lets you create windows with shapes in them. It is designed for creating
graphical user interfaces. It wraps [pixel](https://github.com/faiface/pixel)
and provides a simple way to create and manipulate persistent shapes on screen.
It is as cross platform as pixel is, which is to say, very. This library is not
designed to be used in conjunction with pixel, it abstracts it away. Eventually
it will just use raw OpenGL as a backend.

The reason this exists is because the browser DOM is very useful for building
user interfaces because you can just make things and the browser abstracts away
all the rendering logic. However, it is inherently resource intensive, has too
much vestigial nonsense built into it, and it wasn't even designed for
constructing user interfaces (being more oriented towards text documents).

Naiad is designed to let you have your cake and eat it too. You have full
freedom to create whatever ridiculous shapes you want, with the ease of working
with a DOM, and the assurance that your program will not end up being a 127
megabyte, cpu eating, ram eating piece of garbage.

## Disclaimer

It is not done yet dont use it.

## Diagram that looks more complicated than it actually is

On-screen shapes in Naiad are structured using groups. These groups contain
paths (and soon text as well). At the top there is a root group which you are
not allowed to mess with, but you can add things to it and remove things from
it through methods defined on the window object.

```
          ┌────────┐
        ┌─┤ Window │
        │ └────────┘
        │
        │ ┌────────────┐
        └►│ Root Group │
          └─────┬───┬──┘
                │   │
                │   └──────┐
                ▼          ▼
            ┌───────┐  ┌───────┐
┌─────┬─────┤ Group │  │ Group ├──────┐
│     │     └───┬───┘  └───┬───┘      ▼
│     ▼         │          │      ┌───────┐
│ ┌──────┐      ▼          ▼      │ Path  │
│ │ Path │  ┌───────┐  ┌──────┐   └───────┘
│ └──────┘  │ Group │  │ Path │
│           └───┬───┘  └──────┘
│ ┌──────┐      │
└►│ Path │      ▼
  └──────┘  ┌──────┐
            │ Path │
            └──────┘
```

Paths, in turn, are just lists of points. Each point contains information like
position, color, and line cap type. The entire path contains stuff like line
thickness, and if it should be filled or stroked.

```
Point     Point
  ┌─┬─────┬─┐
  ├─┘     └─┤
  │  Path   │
  ├─┐     ┌─┤
  └─┴─────┴─┘
Point     Point
```

Groups are very cool because they each have their own internal buffer, so
objects inside of them don't need to redraw themselves all the dang time.


## Epic checklist

- [X] Paths
- [X] Shape Groups
- [ ] Calculate mouse click, drag, and hover information in relation to shapes
- [ ] Pass keyboard input as well
- [ ] Text
- [ ] Stop using pixel as a backend
- [ ] Texture paths, text, etc with images
- [ ] Make these images work as mini raster graphics contexts
