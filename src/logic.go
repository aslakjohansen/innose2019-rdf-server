package main

import (
    "runtime"
    "fmt"
    python "github.com/sbinet/go-python"
)

var (
    module_name string = "logic"
    time *python.PyObject
)

func Init () {
    runtime.LockOSThread() // stick go routine to thread
    
    err := python.Initialize()
    if err != nil {
        fmt.Println("Unable to initialize python", err)
    }
    
    module := python.PyImport_ImportModule(module_name)
    if module == nil {
        fmt.Println("Unable to import module '"+module_name+"'")
    }
    
    time = module.GetAttrString("time")
    if time == nil {
        fmt.Println("Unable to name function 'time' in module '"+module_name+"'")
    }
}

func Finalize () {
    err := python.Finalize()
    if err != nil {
        fmt.Println("Unable to finalize python", err)
    }
}

func Time () (float64, bool) {
    state, gstate := enter()
    
    resPython := time.Call(python.PyTuple_New(0), python.PyDict_New())
    success, result := unpack_float64(resPython)
    
    leave(state, gstate)
    return result, success;
}

func enter () (*python.PyThreadState, python.PyGILState) {
    var state  *python.PyThreadState = python.PyEval_SaveThread()
    var gstate  python.PyGILState    = python.PyGILState_Ensure()
    return state, gstate
}

func leave (state *python.PyThreadState, gstate python.PyGILState) {
    python.PyGILState_Release(gstate)
    python.PyEval_RestoreThread(state)
}

func unpack (tuple *python.PyObject) (bool, *python.PyObject) {
    // guard: is a tuple
    if !python.PyTuple_Check(tuple) {
        return false, nil
    }
    
    // guard: has two elements
    if python.PyTuple_GET_SIZE(tuple)!=2 {
        return false, nil
    }
    
    var success *python.PyObject = python.PyTuple_GET_ITEM(tuple, 0)
    var result  *python.PyObject = python.PyTuple_GET_ITEM(tuple, 1)
    
    // guard: 
    if !python.PyBool_Check(success) {
        return false, nil
    }
    
    success_long := python.PyInt_AsLong(success)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
        return false, nil
    }
    
    return success_long!=0, result
}

func unpack_float64 (tuple *python.PyObject) (bool, float64) {
    var success  bool
    var result  *python.PyObject
    
    success, result = unpack(tuple)
    
    // guard: not success
    if !success {
        return false , 1.0
    }
    
    // guard: result is a float
    if !python.PyFloat_Check(result) {
        return false, 2.0
    }
    
    res := python.PyFloat_AsDouble(result)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
        return false, 3.0
    }
    
    return true, res
}

