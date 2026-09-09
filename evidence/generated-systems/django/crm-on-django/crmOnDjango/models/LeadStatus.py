from django.db import models
 #======================================================================
# 
# Encapsulates data for model LeadStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeadStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LeadStatus(Enum):   # A subclass of Enum
	New = 'New'
	Working = 'Working'
	Nurturing = 'Nurturing'
	Qualified = 'Qualified'
	Unqualified = 'Unqualified'
	Converted = 'Converted'
