from django.db import models
 #======================================================================
# 
# Encapsulates data for model EncounterStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EncounterStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EncounterStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	InProgress = 'InProgress'
	OnHold = 'OnHold'
	Discharged = 'Discharged'
	Cancelled = 'Cancelled'
