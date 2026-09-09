from django.db import models
 #======================================================================
# 
# Encapsulates data for model TerminationReason
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminationReason Declaration (enumerated type)
#======================================================================
from enum import Enum 
class TerminationReason(Enum):   # A subclass of Enum
	Voluntary = 'Voluntary'
	Involuntary = 'Involuntary'
	Retirement = 'Retirement'
	Redundancy = 'Redundancy'
	EndOfContract = 'EndOfContract'
