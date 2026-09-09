from django.db import models
 #======================================================================
# 
# Encapsulates data for model LoanPurpose
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanPurpose Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LoanPurpose(Enum):   # A subclass of Enum
	HomeImprovement = 'HomeImprovement'
	Education = 'Education'
	DebtConsolidation = 'DebtConsolidation'
	Business = 'Business'
	Other = 'Other'
