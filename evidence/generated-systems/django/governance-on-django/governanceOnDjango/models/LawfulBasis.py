from django.db import models
 #======================================================================
# 
# Encapsulates data for model LawfulBasis
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LawfulBasis Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LawfulBasis(Enum):   # A subclass of Enum
	Consent = 'Consent'
	Contract = 'Contract'
	LegalObligation = 'LegalObligation'
	VitalInterests = 'VitalInterests'
	PublicTask = 'PublicTask'
	LegitimateInterests = 'LegitimateInterests'
