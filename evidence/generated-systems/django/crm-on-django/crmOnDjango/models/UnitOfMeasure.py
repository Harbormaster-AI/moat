from django.db import models
 #======================================================================
# 
# Encapsulates data for model UnitOfMeasure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnitOfMeasure Declaration (enumerated type)
#======================================================================
from enum import Enum 
class UnitOfMeasure(Enum):   # A subclass of Enum
	Each = 'Each'
	Hour = 'Hour'
	Day = 'Day'
	Month = 'Month'
	User = 'User'
	Package = 'Package'
