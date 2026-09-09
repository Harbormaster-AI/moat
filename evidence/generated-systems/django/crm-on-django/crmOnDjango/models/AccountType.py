from django.db import models
 #======================================================================
# 
# Encapsulates data for model AccountType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AccountType(Enum):   # A subclass of Enum
	Prospect = 'Prospect'
	Customer = 'Customer'
	Partner = 'Partner'
	Vendor = 'Vendor'
	Competitor = 'Competitor'
