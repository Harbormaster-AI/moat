from django.db import models
 #======================================================================
# 
# Encapsulates data for model AgreementStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgreementStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AgreementStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Suspended = 'Suspended'
	Terminated = 'Terminated'
