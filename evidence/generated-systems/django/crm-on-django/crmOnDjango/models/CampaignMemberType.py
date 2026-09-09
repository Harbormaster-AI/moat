from django.db import models
 #======================================================================
# 
# Encapsulates data for model CampaignMemberType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignMemberType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class CampaignMemberType(Enum):   # A subclass of Enum
	Lead = 'Lead'
	Contact = 'Contact'
