from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProcedureStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProcedureStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Aborted = 'Aborted'
