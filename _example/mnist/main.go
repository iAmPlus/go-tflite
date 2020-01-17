package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/iAmPlus/go-tflite"
	"github.com/nfnt/resize"
)

func top(a []float32) int {
	t := 0
	m := float32(0)
	for i, e := range a {
		if i == 0 || e > m {
			m = e
			t = i
		}
	}
	return t
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	model := tflite.NewModelFromFile("mnist_model.tflite")
	if model == nil {
		log.Fatal("cannot load model")
	}
	defer model.Delete()

	interpreter := tflite.NewInterpreter(model, nil)
	defer interpreter.Delete()

	status := interpreter.AllocateTensors()
	if status != tflite.OK {
		log.Fatal("allocate failed")
	}

	input := interpreter.GetInputTensor(0)
	resized := resize.Resize(28, 28, img, resize.NearestNeighbor)
	in := input.Float32s()
	for y := 0; y < 28; y++ {
		for x := 0; x < 28; x++ {
			r, g, b, _ := resized.At(x, y).RGBA()
			in[y*28+x] = (float32(b) + float32(g) + float32(r)) / 3.0 / 65535.0
		}
	}
	status = interpreter.Invoke()
	if status != tflite.OK {
		log.Fatal("invoke failed")
	}

	output := interpreter.GetOutputTensor(0)
	out := output.Float32s()
	fmt.Println(top(out))
}
