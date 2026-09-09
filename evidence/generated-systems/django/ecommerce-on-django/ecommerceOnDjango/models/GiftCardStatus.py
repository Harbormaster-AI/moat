from django.db import models
 #======================================================================
# 
# Encapsulates data for model GiftCardStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class GiftCardStatus(Enum):   # A subclass of Enum
	Active = 'Active'
	Redeemed = 'Redeemed'
	Expired = 'Expired'
	Disabled = 'Disabled'
