from django.db import models
 #======================================================================
# 
# Encapsulates data for model BillingStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BillingStatus(Enum):   # A subclass of Enum
	Current = 'Current'
	Delinquent = 'Delinquent'
	Collections = 'Collections'
	Closed = 'Closed'
