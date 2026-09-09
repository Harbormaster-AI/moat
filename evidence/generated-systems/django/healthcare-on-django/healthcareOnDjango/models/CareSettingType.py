from django.db import models
 #======================================================================
# 
# Encapsulates data for model CareSettingType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CareSettingType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CareSettingType(Enum):   # A subclass of Enum
	Inpatient = 'Inpatient'
	Outpatient = 'Outpatient'
	Emergency = 'Emergency'
	HomeHealth = 'HomeHealth'
	Telehealth = 'Telehealth'
