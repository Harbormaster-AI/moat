from django.db import models
 #======================================================================
# 
# Encapsulates data for model TagCategory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TagCategory Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TagCategory(Enum):   # A subclass of Enum
	Domain = 'Domain'
	Sensitivity = 'Sensitivity'
	Priority = 'Priority'
	Lifecycle = 'Lifecycle'
