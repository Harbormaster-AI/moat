from django.db import models
 #======================================================================
# 
# Encapsulates data for model VerificationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VerificationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class VerificationStatus(Enum):   # A subclass of Enum
	Unverified = 'Unverified'
	Verified = 'Verified'
	Failed = 'Failed'
