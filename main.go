// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Blocks")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func incrementScore(colour color.RGBA) {
	switch colour {
	case color.RGBA{R: 0xa3, G: 0x1e, B: 0x0a, A: 0xff}:
		score += 7
	case color.RGBA{R: 0xc2, G: 0x85, B: 0x0a, A: 0xff}:
		score += 5
	case color.RGBA{R: 0x0a, G: 0x85, B: 0x33, A: 0xff}:
		score += 3
	case color.RGBA{R: 0xc2, G: 0xc2, B: 0x29, A: 0xff}:
		score += 1

	}
}
