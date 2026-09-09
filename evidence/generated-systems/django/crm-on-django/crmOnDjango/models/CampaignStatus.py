from django.db import models
 #======================================================================
# 
# Encapsulates data for model CampaignStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CampaignStatus(Enum):   # A subclass of Enum
	Planned = 'Planned'
	InProgress = 'InProgress'
	Completed = 'Completed'
	OnHold = 'OnHold'
	Cancelled = 'Cancelled'
