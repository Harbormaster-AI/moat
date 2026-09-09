from django.db import models
 #======================================================================
# 
# Encapsulates data for model FilingStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FilingStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class FilingStatus(Enum):   # A subclass of Enum
	Single = 'Single'
	MarriedFilingJointly = 'MarriedFilingJointly'
	MarriedFilingSeparately = 'MarriedFilingSeparately'
	HeadOfHousehold = 'HeadOfHousehold'
	QualifyingWidowEr = 'QualifyingWidowEr'
