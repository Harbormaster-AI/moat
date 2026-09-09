from django.db import models
 #======================================================================
# 
# Encapsulates data for model SalesCampaignStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesCampaignStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class SalesCampaignStatus(Enum):   # A subclass of Enum
	Prospecting = 'Prospecting'
	Proposal = 'Proposal'
	Negotiation = 'Negotiation'
	Won = 'Won'
	Lost = 'Lost'
