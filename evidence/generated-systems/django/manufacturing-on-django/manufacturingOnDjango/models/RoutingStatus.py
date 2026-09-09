from django.db import models
 #======================================================================
# 
# Encapsulates data for model RoutingStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoutingStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RoutingStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Released = 'Released'
	Obsolete = 'Obsolete'
