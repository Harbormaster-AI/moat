from django.db import models
 #======================================================================
# 
# Encapsulates data for model AccountLifecycleStage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountLifecycleStage Declaration (enumerated type)
#======================================================================
from enum import Enum 
class AccountLifecycleStage(Enum):   # A subclass of Enum
	Subscriber = 'Subscriber'
	Lead = 'Lead'
	MarketingQualified = 'MarketingQualified'
	SalesQualified = 'SalesQualified'
	Customer = 'Customer'
	Evangelist = 'Evangelist'
	Churned = 'Churned'
