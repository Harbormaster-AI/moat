from django.db import models
 #======================================================================
# 
# Encapsulates data for model NetworkStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NetworkStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class NetworkStatus(Enum):   # A subclass of Enum
	InNetwork = 'InNetwork'
	OutOfNetwork = 'OutOfNetwork'
