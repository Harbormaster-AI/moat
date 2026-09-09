from django.db import models
 #======================================================================
# 
# Encapsulates data for model PipelineStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PipelineStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PipelineStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Active = 'Active'
	Paused = 'Paused'
	Failed = 'Failed'
	Succeeded = 'Succeeded'
