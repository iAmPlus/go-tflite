#ifndef GO_FLEX_TFLITE_H
#define GO_FLEX_TFLITE_H

#define _GNU_SOURCE

#include <stdint.h>
#include "tensorflow/lite/c/common.h"

// Creates a new flex delegate instance that need to be destroyed with
// TfLiteFlexDelegateDelete when delegate is no longer used by TFLite.
//
TfLiteDelegate* TfLiteFlexDelegateCreate();

// Destroys a delegate created with `TfLiteFlexDelegateCreate` call.
void TfLiteFlexDelegateDelete(TfLiteDelegate* delegate);

#endif  // GO_FLEX_TFLITE_H