from django.db import models
 #======================================================================
# 
# Encapsulates data for model LeadSource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeadSource Declaration (enumerated type)
#======================================================================
from enum import Enum 
class LeadSource(Enum):   # A subclass of Enum
	Web = 'Web'
	Referral = 'Referral'
	Event = 'Event'
	Partner = 'Partner'
	Advertisement = 'Advertisement'
	Outbound = 'Outbound'
	Inbound = 'Inbound'
	Social = 'Social'
	Other = 'Other'
