package main

import (
    "runtime"
    "fmt"
    python "github.com/sbinet/go-python"
)

var (
    module_name string = "logic"
    python_time       *python.PyObject
    python_store      *python.PyObject
    python_load_model *python.PyObject
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// lifecycle management

func Init (model_dir string, ontology_dir string) {
    runtime.LockOSThread() // stick go routine to thread
    
    err := python.Initialize()
    if err != nil {
        fmt.Println("Unable to initialize python", err)
    }
    
    module := python.PyImport_ImportModule(module_name)
    if module == nil {
        fmt.Println("Unable to import module '"+module_name+"'")
    }
    
    python_time = module.GetAttrString("time")
    if python_time == nil {
        fmt.Println("Unable to name function 'time' in module '"+module_name+"'")
    }
    
    python_store = module.GetAttrString("store")
    if python_store == nil {
        fmt.Println("Unable to name function 'store' in module '"+module_name+"'")
    }
    
    python_load_model = module.GetAttrString("load_model")
    if python_load_model == nil {
        fmt.Println("Unable to name function 'load_model' in module '"+module_name+"'")
    }
    
    load_model(model_dir, ontology_dir)
}

func Finalize () {
    err := python.Finalize()
    if err != nil {
        fmt.Println("Unable to finalize python", err)
    }
}

func load_model (model_dir string, ontology_dir string) {
    state, gstate := enter()
    
    // construct arguments
    args := python.PyTuple_New(2)
    python.PyTuple_SET_ITEM(args, 0, python.PyString_FromString(model_dir))
    python.PyTuple_SET_ITEM(args, 1, python.PyString_FromString(ontology_dir))
    
    python_load_model.Call(args, python.PyDict_New())
    
    leave(state, gstate)
}


///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// handlers

func Time () (float64, bool) {
    state, gstate := enter()
    
    resPython := python_time.Call(python.PyTuple_New(0), python.PyDict_New())
    success, result := unpack_float64(resPython)
    
    leave(state, gstate)
    return result, success;
}

func Store (model_dir string) (string, bool) {
    state, gstate := enter()
    
    // construct arguments
    args := python.PyTuple_New(1)
    python.PyTuple_SET_ITEM(args, 0, python.PyString_FromString(model_dir))
    
    resPython := python_store.Call(args, python.PyDict_New())
    success, result := unpack_string(resPython)
    
    leave(state, gstate)
    return result, success;
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////// helpers

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
    
    // guard: type of success
    if !python.PyBool_Check(success) {
        return false, nil
    }
    
    // parse success
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
    
    // convert
    res := python.PyFloat_AsDouble(result)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
        return false, 3.0
    }
    
    return true, res
}

func unpack_string (tuple *python.PyObject) (bool, string) {
    var success  bool
    var result  *python.PyObject
    
    success, result = unpack(tuple)
    
    // guard: not success
    if !success {
        return false , "no success"
    }
    
    // guard: result is a float
    if !python.PyString_Check(result) {
        return false, "result not a float"
    }
    
    // convert
    res := python.PyString_AsString(result)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
        return false, "does not decode as string"
    }
    
    return true, res
}

