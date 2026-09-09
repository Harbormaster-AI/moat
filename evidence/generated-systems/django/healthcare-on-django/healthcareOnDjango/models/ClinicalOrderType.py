from django.db import models
 #======================================================================
# 
# Encapsulates data for model ClinicalOrderType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClinicalOrderType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class ClinicalOrderType(Enum):   # A subclass of Enum
	Medication = 'Medication'
	Laboratory = 'Laboratory'
	Imaging = 'Imaging'
	Procedure = 'Procedure'
	Consultation = 'Consultation'
