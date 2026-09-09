from django.db import models
 #======================================================================
# 
# Encapsulates data for model SubrogationStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubrogationStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SubrogationStatus(Enum):   # A subclass of Enum
	Open = 'Open'
	Negotiating = 'Negotiating'
	Settled = 'Settled'
	Uncollectible = 'Uncollectible'
	Closed = 'Closed'
