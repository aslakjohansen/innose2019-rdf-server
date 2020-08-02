#!/usr/bin/env python3

import sys
import asyncio
from aiohttp import web
import json
from time import time
from os import mkdir, path

import logic

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

###############################################################################
###################################################################### handlers

async def handler_time (path: str, payload):
    s, r = logic.time()
    if s:
        message = json.dumps({
            'success': True,
            'time': r,
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=200, text=message)
    else:
        message = json.dumps({
            'success': False,
            'error': str(r),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

async def handler_store (path: str, payload):
    s, r = logic.store(model_dir)
    if s:
        message = json.dumps({
            'success': True,
            'filename': r,
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=200, text=message)
    else:
        message = json.dumps({
            'success': False,
            'error': str(r),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

async def handler_namespaces (path: str, payload):
    s, r = logic.namespaces()
    if s:
        message = json.dumps({
            'success': True,
            'namespaces': r,
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=200, text=message)
    else:
        message = json.dumps({
            'success': False,
            'error': str(r),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

async def handler_query (path: str, payload):
    s, r = logic.query(payload)
    if s:
        message = json.dumps({
            'success': True,
            'resultset': r,
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=200, text=message)
    else:
        message = json.dumps({
            'success': False,
            'error': str(r),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

async def handler_update (path: str, payload):
    s, r = logic.update(payload)
    if s:
        message = json.dumps({
            'success': True,
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=200, text=message)
    else:
        message = json.dumps({
            'success': False,
            'error': str(r),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

async def handler (request: web.Request):
    method  =       request.method
    path    =   str(request.rel_url)[1:]
    payload = await request.content.read()
    
    try:
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
    except Exception as e:
        message = json.dumps({
            'success': False,
            'exception': str(e),
        }, sort_keys=True, indent=4, separators=(',', ': '))
        return web.Response(status=500, text=message)

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
    print('Syntax: %s INTERFACE PORT MODEL_DIR ONTOLOGY_DIR' % sys.argv[0])
    print('        %s 0.0.0.0 8001 ../var/model ../var/ontologies' % sys.argv[0])
    sys.exit(2)

# extract parameters
interface    =     sys.argv[1]
port         = int(sys.argv[2])
model_dir    =     sys.argv[3]
ontology_dir =     sys.argv[4]

# make sure model_dir exists
if not path.exists(model_dir):
    print('STATUS: model_dir does not exist. Creating "%s".' % model_dir)
    mkdir(model_dir)

logic.load_model(model_dir, ontology_dir)

# register handlers
register_handler('time'      , handler_time)
register_handler('store'     , handler_store)
register_handler('namespaces', handler_namespaces)
register_handler('query'     , handler_query)
register_handler('update'    , handler_update)

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

