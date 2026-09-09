from django.db import models
 #======================================================================
# 
# Encapsulates data for model KYCStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class KYCStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Verified = 'Verified'
	Rejected = 'Rejected'
	Expired = 'Expired'
