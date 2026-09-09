from django.db import models
 #======================================================================
# 
# Encapsulates data for model ExperimentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExperimentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ExperimentStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Running = 'Running'
	Completed = 'Completed'
	Failed = 'Failed'
	Stopped = 'Stopped'
