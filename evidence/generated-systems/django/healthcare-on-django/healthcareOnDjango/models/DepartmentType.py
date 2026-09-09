from django.db import models
 #======================================================================
# 
# Encapsulates data for model DepartmentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DepartmentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DepartmentType(Enum):   # A subclass of Enum
	Emergency = 'Emergency'
	Cardiology = 'Cardiology'
	Oncology = 'Oncology'
	Orthopedics = 'Orthopedics'
	Pediatrics = 'Pediatrics'
	Radiology = 'Radiology'
	Pathology = 'Pathology'
	Pharmacy = 'Pharmacy'
	IntensiveCare = 'IntensiveCare'
