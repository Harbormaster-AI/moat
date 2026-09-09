from django.db import models
 #======================================================================
# 
# Encapsulates data for model SalaryComponentType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalaryComponentType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SalaryComponentType(Enum):   # A subclass of Enum
	BaseSalary = 'BaseSalary'
	Allowance = 'Allowance'
	OvertimeRate = 'OvertimeRate'
	Commission = 'Commission'
	ShiftDifferential = 'ShiftDifferential'
