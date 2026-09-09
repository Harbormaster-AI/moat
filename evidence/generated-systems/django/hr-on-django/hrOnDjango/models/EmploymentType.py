from django.db import models
 #======================================================================
# 
# Encapsulates data for model EmploymentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class EmploymentType(Enum):   # A subclass of Enum
	FullTime = 'FullTime'
	PartTime = 'PartTime'
	Temporary = 'Temporary'
	Intern = 'Intern'
	Contractor = 'Contractor'
	Seasonal = 'Seasonal'
