from django.db import models
 #======================================================================
# 
# Encapsulates data for model SpecimenType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SpecimenType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SpecimenType(Enum):   # A subclass of Enum
	Blood = 'Blood'
	Urine = 'Urine'
	Saliva = 'Saliva'
	Sputum = 'Sputum'
	Tissue = 'Tissue'
	CSF = 'CSF'
	Stool = 'Stool'
