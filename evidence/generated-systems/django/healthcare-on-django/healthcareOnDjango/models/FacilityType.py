from django.db import models
 #======================================================================
# 
# Encapsulates data for model FacilityType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FacilityType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FacilityType(Enum):   # A subclass of Enum
	Hospital = 'Hospital'
	Clinic = 'Clinic'
	AmbulatorySurgeryCenter = 'AmbulatorySurgeryCenter'
	UrgentCare = 'UrgentCare'
	Laboratory = 'Laboratory'
	ImagingCenter = 'ImagingCenter'
	Pharmacy = 'Pharmacy'
