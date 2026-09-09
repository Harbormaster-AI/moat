from django.db import models
 #======================================================================
# 
# Encapsulates data for model DischargeDisposition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DischargeDisposition Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DischargeDisposition(Enum):   # A subclass of Enum
	Home = 'Home'
	HomeWithHomeCare = 'HomeWithHomeCare'
	SkilledNursingFacility = 'SkilledNursingFacility'
	AcuteCareFacility = 'AcuteCareFacility'
	Expired = 'Expired'
	AgainstMedicalAdvice = 'AgainstMedicalAdvice'
