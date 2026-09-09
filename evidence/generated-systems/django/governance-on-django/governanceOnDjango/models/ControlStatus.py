from django.db import models
 #======================================================================
# 
# Encapsulates data for model ControlStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ControlStatus(Enum):   # A subclass of Enum
	Designed = 'Designed'
	Implemented = 'Implemented'
	Operating = 'Operating'
	Retired = 'Retired'
