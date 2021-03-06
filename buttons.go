package naiad

import "github.com/faiface/pixel/pixelgl"

// Key represents a keyboard key.
type Key pixelgl.Button

const (
	KeyUnknown      = Key(pixelgl.KeyUnknown)
	KeySpace        = Key(pixelgl.KeySpace)
	KeyApostrophe   = Key(pixelgl.KeyApostrophe)
	KeyComma        = Key(pixelgl.KeyComma)
	KeyMinus        = Key(pixelgl.KeyMinus)
	KeyPeriod       = Key(pixelgl.KeyPeriod)
	KeySlash        = Key(pixelgl.KeySlash)
	Key0            = Key(pixelgl.Key0)
	Key1            = Key(pixelgl.Key1)
	Key2            = Key(pixelgl.Key2)
	Key3            = Key(pixelgl.Key3)
	Key4            = Key(pixelgl.Key4)
	Key5            = Key(pixelgl.Key5)
	Key6            = Key(pixelgl.Key6)
	Key7            = Key(pixelgl.Key7)
	Key8            = Key(pixelgl.Key8)
	Key9            = Key(pixelgl.Key9)
	KeySemicolon    = Key(pixelgl.KeySemicolon)
	KeyEqual        = Key(pixelgl.KeyEqual)
	KeyA            = Key(pixelgl.KeyA)
	KeyB            = Key(pixelgl.KeyB)
	KeyC            = Key(pixelgl.KeyC)
	KeyD            = Key(pixelgl.KeyD)
	KeyE            = Key(pixelgl.KeyE)
	KeyF            = Key(pixelgl.KeyF)
	KeyG            = Key(pixelgl.KeyG)
	KeyH            = Key(pixelgl.KeyH)
	KeyI            = Key(pixelgl.KeyI)
	KeyJ            = Key(pixelgl.KeyJ)
	KeyK            = Key(pixelgl.KeyK)
	KeyL            = Key(pixelgl.KeyL)
	KeyM            = Key(pixelgl.KeyM)
	KeyN            = Key(pixelgl.KeyN)
	KeyO            = Key(pixelgl.KeyO)
	KeyP            = Key(pixelgl.KeyP)
	KeyQ            = Key(pixelgl.KeyQ)
	KeyR            = Key(pixelgl.KeyR)
	KeyS            = Key(pixelgl.KeyS)
	KeyT            = Key(pixelgl.KeyT)
	KeyU            = Key(pixelgl.KeyU)
	KeyV            = Key(pixelgl.KeyV)
	KeyW            = Key(pixelgl.KeyW)
	KeyX            = Key(pixelgl.KeyX)
	KeyY            = Key(pixelgl.KeyY)
	KeyZ            = Key(pixelgl.KeyZ)
	KeyLeftBracket  = Key(pixelgl.KeyLeftBracket)
	KeyBackslash    = Key(pixelgl.KeyBackslash)
	KeyRightBracket = Key(pixelgl.KeyRightBracket)
	KeyGraveAccent  = Key(pixelgl.KeyGraveAccent)
	KeyWorld1       = Key(pixelgl.KeyWorld1)
	KeyWorld2       = Key(pixelgl.KeyWorld2)
	KeyEscape       = Key(pixelgl.KeyEscape)
	KeyEnter        = Key(pixelgl.KeyEnter)
	KeyTab          = Key(pixelgl.KeyTab)
	KeyBackspace    = Key(pixelgl.KeyBackspace)
	KeyInsert       = Key(pixelgl.KeyInsert)
	KeyDelete       = Key(pixelgl.KeyDelete)
	KeyRight        = Key(pixelgl.KeyRight)
	KeyLeft         = Key(pixelgl.KeyLeft)
	KeyDown         = Key(pixelgl.KeyDown)
	KeyUp           = Key(pixelgl.KeyUp)
	KeyPageUp       = Key(pixelgl.KeyPageUp)
	KeyPageDown     = Key(pixelgl.KeyPageDown)
	KeyHome         = Key(pixelgl.KeyHome)
	KeyEnd          = Key(pixelgl.KeyEnd)
	KeyCapsLock     = Key(pixelgl.KeyCapsLock)
	KeyScrollLock   = Key(pixelgl.KeyScrollLock)
	KeyNumLock      = Key(pixelgl.KeyNumLock)
	KeyPrintScreen  = Key(pixelgl.KeyPrintScreen)
	KeyPause        = Key(pixelgl.KeyPause)
	KeyF1           = Key(pixelgl.KeyF1)
	KeyF2           = Key(pixelgl.KeyF2)
	KeyF3           = Key(pixelgl.KeyF3)
	KeyF4           = Key(pixelgl.KeyF4)
	KeyF5           = Key(pixelgl.KeyF5)
	KeyF6           = Key(pixelgl.KeyF6)
	KeyF7           = Key(pixelgl.KeyF7)
	KeyF8           = Key(pixelgl.KeyF8)
	KeyF9           = Key(pixelgl.KeyF9)
	KeyF10          = Key(pixelgl.KeyF10)
	KeyF11          = Key(pixelgl.KeyF11)
	KeyF12          = Key(pixelgl.KeyF12)
	KeyF13          = Key(pixelgl.KeyF13)
	KeyF14          = Key(pixelgl.KeyF14)
	KeyF15          = Key(pixelgl.KeyF15)
	KeyF16          = Key(pixelgl.KeyF16)
	KeyF17          = Key(pixelgl.KeyF17)
	KeyF18          = Key(pixelgl.KeyF18)
	KeyF19          = Key(pixelgl.KeyF19)
	KeyF20          = Key(pixelgl.KeyF20)
	KeyF21          = Key(pixelgl.KeyF21)
	KeyF22          = Key(pixelgl.KeyF22)
	KeyF23          = Key(pixelgl.KeyF23)
	KeyF24          = Key(pixelgl.KeyF24)
	KeyF25          = Key(pixelgl.KeyF25)
	KeyKP0          = Key(pixelgl.KeyKP0)
	KeyKP1          = Key(pixelgl.KeyKP1)
	KeyKP2          = Key(pixelgl.KeyKP2)
	KeyKP3          = Key(pixelgl.KeyKP3)
	KeyKP4          = Key(pixelgl.KeyKP4)
	KeyKP5          = Key(pixelgl.KeyKP5)
	KeyKP6          = Key(pixelgl.KeyKP6)
	KeyKP7          = Key(pixelgl.KeyKP7)
	KeyKP8          = Key(pixelgl.KeyKP8)
	KeyKP9          = Key(pixelgl.KeyKP9)
	KeyKPDecimal    = Key(pixelgl.KeyKPDecimal)
	KeyKPDivide     = Key(pixelgl.KeyKPDivide)
	KeyKPMultiply   = Key(pixelgl.KeyKPMultiply)
	KeyKPSubtract   = Key(pixelgl.KeyKPSubtract)
	KeyKPAdd        = Key(pixelgl.KeyKPAdd)
	KeyKPEnter      = Key(pixelgl.KeyKPEnter)
	KeyKPEqual      = Key(pixelgl.KeyKPEqual)
	KeyLeftShift    = Key(pixelgl.KeyLeftShift)
	KeyLeftControl  = Key(pixelgl.KeyLeftControl)
	KeyLeftAlt      = Key(pixelgl.KeyLeftAlt)
	KeyLeftSuper    = Key(pixelgl.KeyLeftSuper)
	KeyRightShift   = Key(pixelgl.KeyRightShift)
	KeyRightControl = Key(pixelgl.KeyRightControl)
	KeyRightAlt     = Key(pixelgl.KeyRightAlt)
	KeyRightSuper   = Key(pixelgl.KeyRightSuper)
	KeyMenu         = Key(pixelgl.KeyMenu)
	KeyLast         = Key(pixelgl.KeyLast)
)
