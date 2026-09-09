from django.db import models
 #======================================================================
# 
# Encapsulates data for model ControlEffectiveness
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlEffectiveness Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ControlEffectiveness(Enum):   # A subclass of Enum
	Effective = 'Effective'
	PartiallyEffective = 'PartiallyEffective'
	Ineffective = 'Ineffective'
	NotTested = 'NotTested'
