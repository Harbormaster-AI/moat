from django.db import models
 #======================================================================
# 
# Encapsulates data for model BenefitEnrollmentStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BenefitEnrollmentStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class BenefitEnrollmentStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Active = 'Active'
	Waived = 'Waived'
	Cancelled = 'Cancelled'
	Terminated = 'Terminated'
