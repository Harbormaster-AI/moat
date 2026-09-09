from django.db import models
 #======================================================================
# 
# Encapsulates data for model CoverageLevel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageLevel Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CoverageLevel(Enum):   # A subclass of Enum
	EmployeeOnly = 'EmployeeOnly'
	EmployeeSpouse = 'EmployeeSpouse'
	EmployeeChildren = 'EmployeeChildren'
	Family = 'Family'
