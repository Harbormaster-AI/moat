from django.db import models
 #======================================================================
# 
# Encapsulates data for model MRPRunStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MRPRunStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class MRPRunStatus(Enum):   # A subclass of Enum
	Started = 'Started'
	Completed = 'Completed'
	Failed = 'Failed'
	Cancelled = 'Cancelled'
