from django.db import models
 #======================================================================
# 
# Encapsulates data for model TimesheetStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TimesheetStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TimesheetStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Submitted = 'Submitted'
	Approved = 'Approved'
	Rejected = 'Rejected'
	Processed = 'Processed'
