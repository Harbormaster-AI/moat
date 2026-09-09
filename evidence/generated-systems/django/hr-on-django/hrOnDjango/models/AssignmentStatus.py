from django.db import models
 #======================================================================
# 
# Encapsulates data for model AssignmentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AssignmentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AssignmentStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Active = 'Active'
	Completed = 'Completed'
	Cancelled = 'Cancelled'
