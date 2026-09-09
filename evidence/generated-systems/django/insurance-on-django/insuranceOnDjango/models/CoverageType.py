from django.db import models
 #======================================================================
# 
# Encapsulates data for model CoverageType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CoverageType(Enum):   # A subclass of Enum
	Liability = 'Liability'
	Collision = 'Collision'
	Comprehensive = 'Comprehensive'
	PropertyDamage = 'PropertyDamage'
	BodilyInjury = 'BodilyInjury'
	UninsuredMotorist = 'UninsuredMotorist'
	MedicalPayments = 'MedicalPayments'
	Dwelling = 'Dwelling'
	Contents = 'Contents'
	PersonalLiability = 'PersonalLiability'
	BusinessInterruption = 'BusinessInterruption'
	ProfessionalLiability = 'ProfessionalLiability'
