from django.db import models
 #======================================================================
# 
# Encapsulates data for model WorkCenterType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenterType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WorkCenterType(Enum):   # A subclass of Enum
	Machining = 'Machining'
	Assembly = 'Assembly'
	Painting = 'Painting'
	Packaging = 'Packaging'
	Test = 'Test'
	Warehouse = 'Warehouse'
