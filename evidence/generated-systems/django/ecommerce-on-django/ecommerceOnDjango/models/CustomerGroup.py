from django.db import models
 #======================================================================
# 
# Encapsulates data for model CustomerGroup
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerGroup Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CustomerGroup(Enum):   # A subclass of Enum
	Retail = 'Retail'
	Wholesale = 'Wholesale'
	VIP = 'VIP'
	Employee = 'Employee'
