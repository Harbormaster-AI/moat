from django.db import models
 #======================================================================
# 
# Encapsulates data for model OpportunityStage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityStage Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OpportunityStage(Enum):   # A subclass of Enum
	Qualification = 'Qualification'
	Discovery = 'Discovery'
	Proposal = 'Proposal'
	Negotiation = 'Negotiation'
	ClosedWon = 'ClosedWon'
	ClosedLost = 'ClosedLost'
