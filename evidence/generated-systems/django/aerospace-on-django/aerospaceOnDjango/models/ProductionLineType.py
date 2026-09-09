from django.db import models
 #======================================================================
# 
# Encapsulates data for model ProductionLineType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionLineType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ProductionLineType(Enum):   # A subclass of Enum
	FinalAssembly = 'FinalAssembly'
	SubAssembly = 'SubAssembly'
	Integration = 'Integration'
	TestAndDelivery = 'TestAndDelivery'
