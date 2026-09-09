from django.db import models
 #======================================================================
# 
# Encapsulates data for model AdmissionType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AdmissionType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AdmissionType(Enum):   # A subclass of Enum
	Elective = 'Elective'
	Emergency = 'Emergency'
	Urgent = 'Urgent'
	Newborn = 'Newborn'
	Trauma = 'Trauma'
