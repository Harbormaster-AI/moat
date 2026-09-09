from django.db import models
 #======================================================================
# 
# Encapsulates data for model DisputeReason
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DisputeReason Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DisputeReason(Enum):   # A subclass of Enum
	Fraud = 'Fraud'
	Duplicate = 'Duplicate'
	NotAsDescribed = 'NotAsDescribed'
	NotReceived = 'NotReceived'
	ProcessingError = 'ProcessingError'
