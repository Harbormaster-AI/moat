from django.db import models
 #======================================================================
# 
# Encapsulates data for model DealType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DealType Declaration (enumerated type)
#======================================================================
from enum import Enum 
class DealType(Enum):   # A subclass of Enum
	OpenAuction = 'OpenAuction'
	PrivateAuction = 'PrivateAuction'
	PreferredDeal = 'PreferredDeal'
	ProgrammaticGuaranteed = 'ProgrammaticGuaranteed'
