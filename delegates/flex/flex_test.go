package flex_test

import (
	"fmt"
	"testing"
	"time"

	tfl "github.com/iAmPlus/go-tflite"
	"github.com/iAmPlus/go-tflite/delegates/flex"
)

func printDims(interpreter *tfl.Interpreter) {
	for i := 0; i < interpreter.GetInputTensorCount(); i++ {
		fmt.Printf("Input Dim %d : ", i)
		tensor := interpreter.GetInputTensor(i)
		for d := 0; d < tensor.NumDims(); d++ {
			fmt.Printf("%d ", tensor.Dim(d))
		}
		fmt.Println("")
	}

	for i := 0; i < interpreter.GetOutputTensorCount(); i++ {
		fmt.Printf("Output Dim %d : ", i)
		tensor := interpreter.GetOutputTensor(i)
		for d := 0; d < tensor.NumDims(); d++ {
			fmt.Printf("%d ", tensor.Dim(d))
		}
		fmt.Println("")
	}
}


func TestFlexDelegateLoading(t *testing.T) {
	modelPath := "test_data/model.tflite"

	loadStartTime := time.Now()
	model := tfl.NewModelFromFile(modelPath)
	if model == nil {
		t.Errorf("Cannot load model : %s", modelPath)
		return
	}
	fmt.Println("time taken to load from file : ", time.Since(loadStartTime))

	fmt.Println("Loading new interpreter")
	options := tfl.NewInterpreterOptions()
	options.SetNumThread(4)
	flexDelegate := flex.NewFlex()
	if flexDelegate != nil {
		options.AddDelegate(flexDelegate)
	}

	interpreterStartTime := time.Now()
	interpreter := tfl.NewInterpreter(model, options)
	if interpreter == nil {
		t.Errorf("Unable to load model : %s", modelPath)
		return
	}

	fmt.Println("time to to create interpreter :", time.Since(interpreterStartTime))

	fmt.Println("allocating tensors")
	allocateStartTime := time.Now()
	status := interpreter.AllocateTensors()
	if status != 0 {
		t.Errorf("Unable to allocate tensors : %s", modelPath)
		return		
	}
	fmt.Println("time to allocate tensors :", time.Since(allocateStartTime))
	fmt.Println("total time to load :", time.Since(loadStartTime))
	fmt.Println("printing tensors")
	printDims(interpreter)

	fmt.Println("invoking ..")
	invlokeStartTime := time.Now()
	status = interpreter.Invoke()
	if status != 0 {
		t.Errorf("Exception : Invoke Failed : %d", status)
		return
	}
	fmt.Println("time to invoke :", time.Since(invlokeStartTime))
}
