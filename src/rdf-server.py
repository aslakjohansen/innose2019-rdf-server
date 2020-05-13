#!/usr/bin/env python3

import sys
import asyncio
from aiohttp import web
import json
from time import time
from rdflib import Graph, Namespace
from os import listdir, mkdir, path

dispatch = {}

###############################################################################
####################################################################### helpers

def valid_python_version ():
    v = sys.version_info
    if v[0] != 3: return False
    if v[1] < 5: return False
    return True

def register_handler (path: str, handler):
    dispatch[path] = handler

def get_latest_model_id (model_dir: str):
    # locate potentials
    potentials = []
    for filename in listdir(model_dir):
        parts = filename.split('.')
        if len(parts)!=2: continue
        if parts[1]!='ttl': continue
        if not parts[0].isdigit(): continue
        potentials.append(int(parts[0]))
    
    # locate latest
    latest_id = sorted(potentials)[-1] if len(potentials)>0 else None
    
    return latest_id

def load_model (model_dir: str, namespace: str):
    global m
    
    latest_id = get_latest_model_id(model_dir)
    latest_model = '%s/%u.ttl' % (model_dir, latest_id) if latest_id else None
    
    m = Graph()
    if latest_model:
        print('STATUS: Loading model "%s".' % latest_model)
        m.parse(latest_model, format='turtle')
    else:
        print('STATUS: No model found in "%s", starting with a blank Brick model.' % model_dir)
        RDF   = Namespace('http://www.w3.org/1999/02/22-rdf-syntax-ns#')
        RDFS  = Namespace('http://www.w3.org/2000/01/rdf-schema#')
        OWL   = Namespace('http://www.w3.org/2002/07/owl#')
        XSD   = Namespace('http://www.w3.org/2001/XMLSchema#')
        BRICK = Namespace('https://brickschema.org/schema/1.1.0/Brick#')
        N     = Namespace(namespace)
        
        m.bind('rdf'  , RDF)
        m.bind('rdfs' , RDFS)
        m.bind('owl'  , OWL)
        m.bind('xsd'  , XSD)
        m.bind('brick', BRICK)
        m.bind('n', N)
        
        m.parse('%s/../Brick_expanded.ttl' % model_dir, format='turtle')

def store_model (model_dir: str):
    latest_id = get_latest_model_id(model_dir)
    if not latest_id: latest_id = 0
    model_filename = '%s/%u.ttl' % (model_dir, latest_id+1)
    print('STATUS: Storing model to "%s".' % model_filename)
    m.serialize(model_filename, 'turtle')
    return model_filename

###############################################################################
###################################################################### handlers

async def handler_time (path: str, payload):
    message = json.dumps({
        'success': True,
        'time': time(),
    }, sort_keys=True, indent=4, separators=(',', ': '))
    return web.Response(status=200, text=message)

async def handler_store (path: str, payload):
    try:
        model_filename = store_model(model_dir)
        message = json.dumps({
            'success': True,
            'filename': model_filename,
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=200, text=message)
    except Exception as e:
        message = json.dumps({
            'success': False,
            'error': str(e),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

async def handler_namespaces (path: str, payload):
    # collect info
    namespaces = {}
    for prefix, namespace in m.namespaces():
        namespaces[prefix] = namespace
    
    message = json.dumps({
        'success': True,
        'namespaces': namespaces,
    }, sort_keys=True, indent=4, separators=(',', ': '))
    return web.Response(status=200, text=message)
    
async def handler (request: web.Request):
    method  =       request.method
    path    =   str(request.rel_url)[1:]
    payload = await request.content.read()
    
    payload = json.loads(payload.decode('utf-8'))
    
    # produce answer
    if path in dispatch:
        return await dispatch[path](path, payload)
    else:
        message = json.dumps({
            'success': False,
            'error': {
                'description': 'No handler registered for path',
                'method':      method,
                'path':        path,
            },
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=404, text=message)

###############################################################################
########################################################################## main

async def main(interface: str, port: int):
    proto  = web.Server(handler)
    server = await loop.create_server(proto, interface, port)
    print('STATUS: Listening on %s:%u' % (interface, port))

# guard: python version
if not valid_python_version():
    print('ERROR: Invalid python version (%s), bust be 3.(5+).' % str(sys.version_info))
    sys.exit(1)

# guard: commandline arguments
if len(sys.argv) != 5:
    print('Syntax: %s INTERFACE PORT MODEL_DIR NAMESPACE' % sys.argv[0])
    print('        %s 0.0.0.0 8001 ../var/model http://ses.sdk.dk/junk/test#' % sys.argv[0])
    sys.exit(2)

# extract parameters
interface =     sys.argv[1]
port      = int(sys.argv[2])
model_dir =     sys.argv[3]
namespace =     sys.argv[4]

# make sure model_dir exists
if not path.exists(model_dir):
    print('STATUS: model_dir does not exist. Creating "%s".' % model_dir)
    mkdir(model_dir)

load_model(model_dir, namespace)

# register handlers
register_handler('time'      , handler_time)
register_handler('store'     , handler_store)
register_handler('namespaces', handler_namespaces)

loop = asyncio.get_event_loop()
asyncio.Task(main(interface, port))

# enter service loop
try:
    loop.run_forever()
except KeyboardInterrupt:
    print('')
    print('STATUS: Exiting ...')
    loop.close()
    exit(0)

########################################################################### EOF
###############################################################################

