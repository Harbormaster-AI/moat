from django.db import models
 #======================================================================
# 
# Encapsulates data for model InspectionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class InspectionType(Enum):   # A subclass of Enum
	Incoming = 'Incoming'
	InProcess = 'InProcess'
	Final = 'Final'
	Audit = 'Audit'
