from django.db import models
 #======================================================================
# 
# Encapsulates data for model FeatureStoreType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FeatureStoreType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FeatureStoreType(Enum):   # A subclass of Enum
	Online = 'Online'
	Offline = 'Offline'
	Hybrid = 'Hybrid'
