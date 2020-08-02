from os import listdir, path
from time import time as time_time
from rdflib import Graph, Namespace
import json

###############################################################################
################################################################ input / output

def get_latest_model_id (model_dir):
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

def load_model (model_dir, ontology_dir):
    global m
    
    latest_id = get_latest_model_id(model_dir)
    latest_model = '%s/%u.ttl' % (model_dir, latest_id) if latest_id else None
    
    m = Graph()
    if latest_model:
        print('STATUS: Loading model "%s".' % latest_model)
        m.parse(latest_model, format='turtle')
    else:
        print('STATUS: No model found in "%s", starting with a blank model.' % model_dir)
        
        # fixed namespaces
        RDF   = Namespace('http://www.w3.org/1999/02/22-rdf-syntax-ns#')
        RDFS  = Namespace('http://www.w3.org/2000/01/rdf-schema#')
        OWL   = Namespace('http://www.w3.org/2002/07/owl#')
        XSD   = Namespace('http://www.w3.org/2001/XMLSchema#')
        BRICK = Namespace('https://brickschema.org/schema/1.1.0/Brick#')
        
        # fixed namespace mapping
        m.bind('rdf'  , RDF)
        m.bind('rdfs' , RDFS)
        m.bind('owl'  , OWL)
        m.bind('xsd'  , XSD)
        
        # load extra namespaces
        namespace_map_filename = '%s/namespaces.json' % ontology_dir
        if path.exists(namespace_map_filename):
            print('NOTICE: Loading namespace mapping from "%s":' % namespace_map_filename)
            with open(namespace_map_filename) as fo:
                namespace_map = json.loads(''.join(fo.readlines()))
                for key in namespace_map:
                    print('NOTICE: - %s: %s' % (key, namespace_map[key]))
                    m.bind(key, Namespace(namespace_map[key]))
        else:
            print('NOTICE: No extra namespaces mapped')
        
        # load ontologies
        if path.exists(ontology_dir):
            print('NOTICE: Loading ontologies from "%s":' % ontology_dir)
            for filename in listdir(ontology_dir):
                if not filename.endswith('.ttl'): continue
                true_filename = '%s/%s' % (ontology_dir, filename)
                print('NOTICE: - %s' % true_filename)
                m.parse(true_filename, format='turtle')
        else:
            print('NOTICE: No directory for loading ontologies')
        

def store_model (model_dir):
    latest_id = get_latest_model_id(model_dir)
    if not latest_id: latest_id = 0
    model_filename = '%s/%u.ttl' % (model_dir, latest_id+1)
    print('STATUS: Storing model to "%s".' % model_filename)
    m.serialize(model_filename, 'turtle')
    return model_filename

###############################################################################
###################################################################### handlers

def time ():
    try:
        return True, float(time_time())
    except Exception as e:
        return False, str(e)

def store (model_dir):
    try:
        model_filename = store_model(model_dir)
        return True, model_filename
    except Exception as e:
        return False, str(e)

def namespaces ():
    try:
        namespaces = {}
        for prefix, namespace in m.namespaces():
            namespaces[str(prefix)] = str(namespace)
        return True, namespaces
    except Exception as e:
        return False, str(e)

def query (q):
    try:
        resultset = []
        for row in m.query(q):
            resultset.append(list(map(lambda element: '%s' % element, row)))
        return True, resultset
    except Exception as e:
        return False, str(e)

def update (q):
    try:
        m.update(q)
        return True, None
    except Exception as e:
        return False, str(e)

