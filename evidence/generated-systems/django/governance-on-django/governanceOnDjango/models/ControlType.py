from django.db import models
 #======================================================================
# 
# Encapsulates data for model ControlType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ControlType(Enum):   # A subclass of Enum
	Preventive = 'Preventive'
	Detective = 'Detective'
	Corrective = 'Corrective'
	Directive = 'Directive'
