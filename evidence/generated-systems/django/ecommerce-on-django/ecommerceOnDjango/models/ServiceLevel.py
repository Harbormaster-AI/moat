from django.db import models
 #======================================================================
# 
# Encapsulates data for model ServiceLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ServiceLevel(Enum):   # A subclass of Enum
	Economy = 'Economy'
	Standard = 'Standard'
	Express = 'Express'
	Priority = 'Priority'
	NextDay = 'NextDay'
