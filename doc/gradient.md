# Colors and Gradients

The `colorful` and `gradient` packages provide functions for working with gradients.

## Import Path

```go
import (
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
)
```

# Colors

## Defining

Define colors with `colorful.Hex()`.

```go
var (
	Red      = Hex("#FF0000")
	Blue     = Hex("#00F")
	Green50A = Hex("#00FF007F")
)
```

## Boost

Derive out of range colors with `Color.Boost()`.

```go
var (
	Red      = Hex("#FF0000")
	SuperRed = Red.Boost(1.5) // SuperRed has the RGB values 1.5, 0. 0.
)
```

## Sets

Use `colorful.Set`s to work with sets of related colors.

```go
var rgbSet = colorful.NewSet(
    colorful.Red
    colorful.Blue
    colorful.Green
)

func main() {
    rgbSet.Next() // returns colorful.Red
    rgbSet.Next() // returns colorful.Blue
    rgbSet.Next() // returns colorful.Green
    rgbSet.Next() // returns colorful.Red
    
    rgbSet.Rand() // returns a random color from rgbSet
    
    rgbSet.Idx(0)  // returns colorful.Red
    rgbSet.Idx(-1) // returns colorful.Green
    rgbSet.Idx(3)  // returns colorful.Red
}
```

# Gradients

## Defining

Define gradients with equidistant colors using `gradient.New()`.

```go
var (
    rgbGradient = gradient.New(
        colorful.Red,
        colorful.Green,
        colorful.Blue,
    ) // fades from red to green to blue
)
```

Use `gradient.FromSet()` to create a gradient from a `colorful.Set`.

```go
var (
    rbColorSet = colorful.NewSet(
        colorful.Red,
        colorful.Blue,
    )
    
    rbGradient = gradient.FromSet(rbColorSet)
)

```

Define gradients with colors at specific positions like so. The position of each color must lie in the range [0,1] and
the colors must be sorted in ascending order.

```go
var Rainbow = Table{
    {colorful.Hex("#9e0142"), 0.0},
    {colorful.Hex("#d53e4f"), 0.1},
    {colorful.Hex("#f46d43"), 0.2},
    {colorful.Hex("#fdae61"), 0.3},
    {colorful.Hex("#fee090"), 0.4},
    {colorful.Hex("#ffffbf"), 0.5},
    {colorful.Hex("#e6f598"), 0.6},
    {colorful.Hex("#abdda4"), 0.7},
    {colorful.Hex("#66c2a5"), 0.8},
    {colorful.Hex("#3288bd"), 0.9},
    {colorful.Hex("#5e4fa2"), 1.0},
}
```

## Blending

Get the color at position `t` using `Lerp()`. `t` must lie in the range [0,1].

```go
func main() {
    rbGradient := gradient.New(
        colorful.Red,
        colorful.Blue,
    )
    
    rbGradient.Lerp(0.2)
    rbGradient.Lerp(0.8)
}

```

### Blending in Color Spaces

By default, Ilysa blends gradients in the Oklab color space.
See https://raphlinus.github.io/color/2021/01/18/oklab-critique.html for an interactive gradient generator.

Override this default by setting `gradient.DefaultLerpFn`. Override this default on a per-lerp basis using `LerpIn()`.

Ilysa currently supports the following color spaces.

```
colorful.BlendHSV       // HSV
colorful.BlendLAB       // CIE LAB
colorful.BlendLABLCH    // CIE LABLCh
colorful.BlendLUV       // CIE LUV 
colorful.BlendLUVLCh    // CIE LUVLCh
colorful.BlendLinearRGB // Linear RGB
colorful.BlendOklab     // Oklab
colorful.BlendRGB       // sRGB
```

```go
gradient.DefaultLerpFn = colorful.BlendLAB

func main() {
    // all Lerp function calls will blend in the CIE LAB color space.
    
    // except where LerpIn is used
    gradient.Rainbow.LerpIn(0.2, colorful.BlendLinearRGB)
}
```

## Transforming

### Reverse

Reverse the colors of a gradient using `Reverse()`.

```go
var (
    rgbGradient = gradient.New(colorful.Red, colorful.Green, colorful.Blue)
    bgrGradient = rgbGradient.Reverse()
)
```

### Rotate

Rotate the colors of a gradient to the left using `Rorate()`.

```go
var (
    rgbGradient = gradient.New(colorful.Red, colorful.Green, colorful.Blue)
    gbrGradient = rgbGradient.Rotate(1)
    brgGradient = rgbGradient.Rotate(2)
)
```

## Sets

Work with collections of gradients using `gradient.Set`s.

```go
var (
    rbGradient = gradient.New(colorful.Red, colorful.Blue)
    bgGradient = gradient.New(colorful.Blue, colorful.Greeh)
    gradSet    = gradient.NewSet(rbGradient, bgGradient)
)

func main() {
    gradSet.Next() // returns rbGradient
    gradSet.Next() // returns bgGradient 
    gradSet.Next() // returns rbGradient
    
    gradSet.Rand() // returns a random gradient from gradSet
    
    gradSet.Index(0)  // returns rbGradient
    gradSet.Index(-1) // returns bgGradient
    gradSet.Index(2)  // returns rbGradient
}
```

## Gradient Reference Chart

Get a handy chart of all gradients used in your Ilysa project with `GenerateGradientReference()`. Gradients will only
show up in the reference chart if they are declared at the top level. `GenerateGradientReference()` invokes some
black-magic and will only work on the computer you are compiling your project on.

```go
func main(){
    // ... boilerplate snip ...
    
	p := ilysa.New(bsMap)
	
    // - project code goes here -
    
    p.Save() // not necessary
    
    // generate gradient reference and save to gradref.png
	if err := p.GenerateGradientReference("gradref.png"); err != nil {
		fmt.Println(err)
	}
}
```
