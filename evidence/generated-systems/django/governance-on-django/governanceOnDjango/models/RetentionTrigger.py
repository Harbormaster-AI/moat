from django.db import models
 #======================================================================
# 
# Encapsulates data for model RetentionTrigger
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RetentionTrigger Declaration (enumerated type)
#======================================================================
from enum import Enum 
class RetentionTrigger(Enum):   # A subclass of Enum
	CreationDate = 'CreationDate'
	LastModified = 'LastModified'
	Termination = 'Termination'
	ContractEnd = 'ContractEnd'
	EventCompletion = 'EventCompletion'
	FiscalYearEnd = 'FiscalYearEnd'
