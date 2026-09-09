from django.db import models
 #======================================================================
# 
# Encapsulates data for model IncidentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IncidentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class IncidentStatus(Enum):   # A subclass of Enum
	Identified = 'Identified'
	Contained = 'Contained'
	Notified = 'Notified'
	Resolved = 'Resolved'
	Closed = 'Closed'
