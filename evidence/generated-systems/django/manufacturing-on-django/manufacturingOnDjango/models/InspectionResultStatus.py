from django.db import models
 #======================================================================
# 
# Encapsulates data for model InspectionResultStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionResultStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InspectionResultStatus(Enum):   # A subclass of Enum
	Pass = 'Pass'
	Fail = 'Fail'
	Rework = 'Rework'
	Scrap = 'Scrap'
