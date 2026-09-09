from django.db import models
 #======================================================================
# 
# Encapsulates data for model ConnectivityStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConnectivityStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ConnectivityStatus(Enum):   # A subclass of Enum
	Offline = 'Offline'
	Online = 'Online'
	Degraded = 'Degraded'
