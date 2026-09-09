from django.db import models
 #======================================================================
# 
# Encapsulates data for model DirectDebitScheme
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DirectDebitScheme Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DirectDebitScheme(Enum):   # A subclass of Enum
	SEPA = 'SEPA'
	ACH = 'ACH'
	BACS = 'BACS'
	BECS = 'BECS'
