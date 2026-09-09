from django.db import models
 #======================================================================
# 
# Encapsulates data for model DashboardStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DashboardStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DashboardStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Live = 'Live'
	Archived = 'Archived'
