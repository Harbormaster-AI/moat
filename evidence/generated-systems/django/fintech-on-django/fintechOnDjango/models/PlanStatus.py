from django.db import models
 #======================================================================
# 
# Encapsulates data for model PlanStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlanStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PlanStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Active = 'Active'
	Suspended = 'Suspended'
	Archived = 'Archived'
