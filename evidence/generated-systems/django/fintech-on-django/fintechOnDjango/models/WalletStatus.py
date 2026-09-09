from django.db import models
 #======================================================================
# 
# Encapsulates data for model WalletStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WalletStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class WalletStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Suspended = 'Suspended'
	Closed = 'Closed'
