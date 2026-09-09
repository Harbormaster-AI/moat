from django.db import models
 #======================================================================
# 
# Encapsulates data for model RoutingType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoutingType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RoutingType(Enum):   # A subclass of Enum
	Standard = 'Standard'
	Alternate = 'Alternate'
	Rework = 'Rework'
