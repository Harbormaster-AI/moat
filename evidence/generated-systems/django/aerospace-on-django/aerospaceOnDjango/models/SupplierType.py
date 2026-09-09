from django.db import models
 #======================================================================
# 
# Encapsulates data for model SupplierType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SupplierType(Enum):   # A subclass of Enum
	Airframe = 'Airframe'
	Engine = 'Engine'
	Avionics = 'Avionics'
	Systems = 'Systems'
	Materials = 'Materials'
	MRO = 'MRO'
	Testing = 'Testing'
