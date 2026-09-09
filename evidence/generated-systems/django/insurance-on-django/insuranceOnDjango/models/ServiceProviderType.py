from django.db import models
 #======================================================================
# 
# Encapsulates data for model ServiceProviderType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceProviderType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ServiceProviderType(Enum):   # A subclass of Enum
	RepairShop = 'RepairShop'
	Towing = 'Towing'
	MedicalProvider = 'MedicalProvider'
	Attorney = 'Attorney'
	ForensicEngineer = 'ForensicEngineer'
	RentalCar = 'RentalCar'
