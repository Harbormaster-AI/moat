from django.db import models
 #======================================================================
# 
# Encapsulates data for model CampaignMemberStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignMemberStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CampaignMemberStatus(Enum):   # A subclass of Enum
	Sent = 'Sent'
	Opened = 'Opened'
	Responded = 'Responded'
	Unsubscribed = 'Unsubscribed'
	Bounced = 'Bounced'
	Registered = 'Registered'
	Attended = 'Attended'
	NoShow = 'NoShow'
