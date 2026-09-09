from django.db import models
 #======================================================================
# 
# Encapsulates data for model PayrollItemType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PayrollItemType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class PayrollItemType(Enum):   # A subclass of Enum
	Earning = 'Earning'
	Deduction = 'Deduction'
	Tax = 'Tax'
	Benefit = 'Benefit'
