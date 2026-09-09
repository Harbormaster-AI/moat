from django.db import models
 #======================================================================
# 
# Encapsulates data for model PipelineTriggerType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PipelineTriggerType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PipelineTriggerType(Enum):   # A subclass of Enum
	Manual = 'Manual'
	Schedule = 'Schedule'
	Event = 'Event'
