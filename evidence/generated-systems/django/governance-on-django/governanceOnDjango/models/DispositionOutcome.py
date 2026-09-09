from django.db import models
 #======================================================================
# 
# Encapsulates data for model DispositionOutcome
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispositionOutcome Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DispositionOutcome(Enum):   # A subclass of Enum
	Approved = 'Approved'
	Deferred = 'Deferred'
	Rejected = 'Rejected'
	Executed = 'Executed'
