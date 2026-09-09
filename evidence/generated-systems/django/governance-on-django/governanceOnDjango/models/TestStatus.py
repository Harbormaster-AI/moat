from django.db import models
 #======================================================================
# 
# Encapsulates data for model TestStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TestStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TestStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	InProgress = 'InProgress'
	Completed = 'Completed'
	Blocked = 'Blocked'
	Cancelled = 'Cancelled'
