from django.db import models
 #======================================================================
# 
# Encapsulates data for model OfferStatus
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OfferStatus Declaration (enumerated type)
#======================================================================
from enum import Enum 
class OfferStatus(Enum):   # A subclass of Enum
	Draft = 'Draft'
	Sent = 'Sent'
	Accepted = 'Accepted'
	Declined = 'Declined'
	Withdrawn = 'Withdrawn'
	Expired = 'Expired'
