from django.db import models
 #======================================================================
# 
# Encapsulates data for model DistributionChannelType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DistributionChannelType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DistributionChannelType(Enum):   # A subclass of Enum
	Agency = 'Agency'
	Broker = 'Broker'
	Direct = 'Direct'
	Bancassurance = 'Bancassurance'
	AffinityPartner = 'AffinityPartner'
	OnlineAggregator = 'OnlineAggregator'
