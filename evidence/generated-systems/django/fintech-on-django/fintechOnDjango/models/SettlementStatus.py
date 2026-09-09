from django.db import models
 #======================================================================
# 
# Encapsulates data for model SettlementStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SettlementStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SettlementStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Processing = 'Processing'
	Closed = 'Closed'
	Reconciled = 'Reconciled'
