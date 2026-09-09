from django.db import models
 #======================================================================
# 
# Encapsulates data for model CampaignType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CampaignType(Enum):   # A subclass of Enum
	Email = 'Email'
	Social = 'Social'
	Event = 'Event'
	Webinar = 'Webinar'
	Advertising = 'Advertising'
	ContentMarketing = 'ContentMarketing'
	Referral = 'Referral'
