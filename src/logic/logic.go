package logic

import (
    "runtime"
    "fmt"
    "sync"
    "encoding/json"
    
    python "github.com/sbinet/go-python"
    
    "innose2019-rdf-server/config"
)

var (
    module_name string = "logic"
    python_load_model *python.PyObject
    python_time       *python.PyObject
    python_store      *python.PyObject
    python_namespaces *python.PyObject
    python_query      *python.PyObject
    python_update     *python.PyObject
    mutex sync.Mutex
)

type LogicModuleConfig struct {
    config.ModuleConfig
    ModelDir    string `json:"modeldir"`
    OntologyDir string `json:"ontologydir"`
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// lifecycle management

func Init (configraw *json.RawMessage) {
    var config LogicModuleConfig
    
    // parse config
    err := json.Unmarshal(*configraw, &config)
    if err!=nil {
        fmt.Println("Unable to unmarshal config for module 'logic':", err)
    }
    
    // python.PyEval_InitThreads() // https://stackoverflow.com/questions/8451334/why-is-pygilstate-release-throwing-fatal-python-errors
    
    runtime.LockOSThread() // stick go routine to thread
    
    err = python.Initialize()
    if err != nil {
        fmt.Println("Unable to initialize python", err)
    }
    
    // ensure mutual exclusion
    state, gstate := enter()
    defer leave(state, gstate)
    
    module := python.PyImport_ImportModule(module_name)
    if module == nil {
        fmt.Println("Unable to import module '"+module_name+"'")
    }
    
    python_load_model = module.GetAttrString("load_model")
    if python_load_model == nil {
        fmt.Println("Unable to name function 'load_model' in module '"+module_name+"'")
    }
    
    python_time = module.GetAttrString("time")
    if python_time == nil {
        fmt.Println("Unable to name function 'time' in module '"+module_name+"'")
    }
    
    python_store = module.GetAttrString("store")
    if python_store == nil {
        fmt.Println("Unable to name function 'store' in module '"+module_name+"'")
    }
    
    python_namespaces = module.GetAttrString("namespaces")
    if python_namespaces == nil {
        fmt.Println("Unable to name function 'namespaces' in module '"+module_name+"'")
    }
    
    python_query = module.GetAttrString("query")
    if python_query == nil {
        fmt.Println("Unable to name function 'query' in module '"+module_name+"'")
    }
    
    python_update = module.GetAttrString("update")
    if python_update == nil {
        fmt.Println("Unable to name function 'update' in module '"+module_name+"'")
    }
    
    load_model(config.ModelDir, config.OntologyDir)
}

func Finalize () {
    err := python.Finalize()
    if err != nil {
        fmt.Println("Unable to finalize python", err)
    }
}

func load_model (model_dir string, ontology_dir string) {
    arg0 := python.PyString_FromString(model_dir)
    arg1 := python.PyString_FromString(ontology_dir)
    
    // construct arguments
    args := python.PyTuple_New(2)
    python.PyTuple_SET_ITEM(args, 0, arg0)
    python.PyTuple_SET_ITEM(args, 1, arg1)
    
    python_load_model.Call(args, python.PyDict_New())
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// handlers

func Time () (float64, bool) {
    state, gstate := enter()
    defer leave(state, gstate)
    
    resPython := python_time.Call(python.PyTuple_New(0), python.PyDict_New())
    success, result := unpack_float64(resPython)
    
    return result, success;
}

func Store (model_dir string) (string, bool) {
    state, gstate := enter()
    defer leave(state, gstate)
    
    // construct arguments
    args := python.PyTuple_New(1)
    python.PyTuple_SET_ITEM(args, 0, python.PyString_FromString(model_dir))
    
    resPython := python_store.Call(args, python.PyDict_New())
    success, result := unpack_string(resPython)
    
    return result, success;
}

func Namespaces () (map[string]string, bool) {
    state, gstate := enter()
    defer leave(state, gstate)
    
    resPython := python_namespaces.Call(python.PyTuple_New(0), python.PyDict_New())
    success, result := unpack_string2string(resPython)
    
    return result, success;
}

func Query (q string) ([][]string, bool) {
    state, gstate := enter()
    defer leave(state, gstate)
    
    // construct arguments
    args := python.PyTuple_New(1)
    python.PyTuple_SET_ITEM(args, 0, python.PyString_FromString(q))
    
    resPython := python_query.Call(args, python.PyDict_New())
    success, result := unpack_string2d(resPython)
    
    return result, success;
}

func Update (q string) (bool, bool) {
    state, gstate := enter()
    defer leave(state, gstate)
    
    // construct arguments
    args := python.PyTuple_New(1)
    python.PyTuple_SET_ITEM(args, 0, python.PyString_FromString(q))
    
    resPython := python_update.Call(args, python.PyDict_New())
    success, _ := unpack(resPython)
    
    return success, success;
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////// helpers

func enter () (*python.PyThreadState, python.PyGILState) {
    mutex.Lock()
    var state  *python.PyThreadState = python.PyEval_SaveThread()
    var gstate  python.PyGILState    = python.PyGILState_Ensure()
    return state, gstate
}

func leave (state *python.PyThreadState, gstate python.PyGILState) {
    python.PyGILState_Release(gstate)
    python.PyEval_RestoreThread(state)
    mutex.Unlock()
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
    
    // guard: result is a string
    if !python.PyString_Check(result) {
        return false, "result not a string"
    }
    
    // convert
    res := python.PyString_AsString(result)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
        return false, "does not decode as string"
    }
    
    return true, res
}

func unpack_string2string (tuple *python.PyObject) (bool, map[string]string) {
    var success  bool
    var result  *python.PyObject
    
    success, result = unpack(tuple)
    
    // guard: not success
    if !success {
        fmt.Println("Unpacking of python return value was not a success")
        return false , nil
    }
    
    // guard: result is a dict
    if !python.PyDict_Check(result) {
        fmt.Println("Python return value is not a dict")
        return false, nil
    }
    
    // convert to items
    items := python.PyDict_Items(result)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Extraction of items from python return value failed")
        return false, nil
    }
    
    // guard: items is a list
    if !python.PyList_Check(items) {
        fmt.Println("Items from dict is not a list")
        return false, nil
    }
    
    size := python.PyList_GET_SIZE(items)
    
    // construct map
    var res map[string]string = make(map[string]string)
    for i := 0 ; i < size ; i++ {
        var tuple *python.PyObject = python.PyList_GET_ITEM(items, i)
        
        // guard: is a tuple
        if !python.PyTuple_Check(tuple) {
            fmt.Println("Element",i, "in list not a tuple")
            return false, nil
        }
        
        // guard: has two elements
        if python.PyTuple_GET_SIZE(tuple)!=2 {
            fmt.Println("Tuple does not have exactly 2 elements")
            return false, nil
        }
        
        var key   *python.PyObject = python.PyTuple_GET_ITEM(tuple, 0)
        var value *python.PyObject = python.PyTuple_GET_ITEM(tuple, 1)
        
        // guard: key is a string
        if !python.PyString_Check(key) {
            fmt.Println("Key is not a string")
            return false, nil
        }
        
        // guard: value is a string
        if !python.PyString_Check(value) {
            fmt.Println("Value is not a string")
            return false, nil
        }
        
        // extract string representation of key
        var key_str string = python.PyString_AsString(key)
        if python.PyErr_Occurred()!=nil {
            fmt.Println("Decoding of python key failed")
            return false, nil
        }
        
        // extract string representation of value
        var value_str string = python.PyString_AsString(value)
        if python.PyErr_Occurred()!=nil {
            fmt.Println("Decoding of python value failed")
            return false, nil
        }
        
        res[key_str] = value_str
    }
    
    return true, res
}

func unpack_string2d (tuple *python.PyObject) (bool, [][]string) {
    var success  bool
    var result  *python.PyObject
    
    success, result = unpack(tuple)
    
    // guard: not success
    if !success {
        fmt.Println("Unpacking of python return value was not a success")
        return false , nil
    }
    
    // guard: result is a dict
    if !python.PyList_Check(result) {
        fmt.Println("Python return value is not a list")
        return false, nil
    }
    
    row_count := python.PyList_GET_SIZE(result)
    
    var col_count int
    var res       [][]string = make([][]string, row_count)
    for r := 0 ; r<row_count ; r++ {
        var row *python.PyObject = python.PyList_GET_ITEM(result, r)
        
        // guard: is a list
        if !python.PyList_Check(row) {
            fmt.Println("Row", r, "in result not a list")
            return false, nil
        }
        
        // allocate and construct the 2d array
        if r==0 {
            col_count = python.PyList_GET_SIZE(row)
        }
        res[r] = make([]string, col_count)
        
        for c:=0 ; c<col_count ; c++ {
            var cell *python.PyObject = python.PyList_GET_ITEM(row, c)
            
            // guard: cell is nil
            if cell==nil {
                fmt.Println("Cell is nil")
                return false, nil
            }
            
//            // guard: cell is a string
//            if !python.PyString_Check(cell) {
//                fmt.Println("Cell is not a string")
//                return false, nil
//            }
            
            res[r][c] = python.PyString_AsString(cell)
            if python.PyErr_Occurred()!=nil {
                fmt.Println("Decoding of cell failed")
                return false, nil
            }
        }
    }
    
    return true, res
}

