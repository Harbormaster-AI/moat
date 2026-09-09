from django.db import models
 #======================================================================
# 
# Encapsulates data for model LoanStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LoanStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LoanStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Delinquent = 'Delinquent'
	Closed = 'Closed'
	ChargedOff = 'ChargedOff'
