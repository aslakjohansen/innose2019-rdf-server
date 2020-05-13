#!/usr/bin/env python3

import sys

###############################################################################
####################################################################### helpers

def valid_python_version ():
    v = sys.version_info
    if v[0] != 3: return False
    if v[1] < 5: return False
    return True


###############################################################################
########################################################################## main

# guard: python version
if not valid_python_version():
    print('ERROR: Invalid python version (%s), bust be 3.(5+).' % str(sys.version_info))
    sys.exit(1)


