from django.db import models
 #======================================================================
# 
# Encapsulates data for model AcknowledgementStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AcknowledgementStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AcknowledgementStatus(Enum):   # A subclass of Enum
	Pending = 'Pending'
	Acknowledged = 'Acknowledged'
	Declined = 'Declined'
