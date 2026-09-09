from django.db import models
 #======================================================================
# 
# Encapsulates data for model RecordStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecordStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RecordStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Archived = 'Archived'
	PendingDisposition = 'PendingDisposition'
	Disposed = 'Disposed'
	OnHold = 'OnHold'
