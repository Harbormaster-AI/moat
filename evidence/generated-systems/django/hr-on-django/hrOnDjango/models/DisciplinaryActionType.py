from django.db import models
 #======================================================================
# 
# Encapsulates data for model DisciplinaryActionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DisciplinaryActionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DisciplinaryActionType(Enum):   # A subclass of Enum
	VerbalWarning = 'VerbalWarning'
	WrittenWarning = 'WrittenWarning'
	Suspension = 'Suspension'
	Termination = 'Termination'
