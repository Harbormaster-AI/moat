from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProductionOrderStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionOrderStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProductionOrderStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	Released = 'Released'
	InAssembly = 'InAssembly'
	FlightTest = 'FlightTest'
	Completed = 'Completed'
