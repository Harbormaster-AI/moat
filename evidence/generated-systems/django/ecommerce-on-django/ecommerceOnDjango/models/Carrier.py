from django.db import models
 #======================================================================
# 
# Encapsulates data for model Carrier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Carrier Declaration (enumerated type)
#======================================================================
from enum import Enum 
class Carrier(Enum):   # A subclass of Enum
	UPS = 'UPS'
	FedEx = 'FedEx'
	USPS = 'USPS'
	DHL = 'DHL'
	RoyalMail = 'RoyalMail'
	CanadaPost = 'CanadaPost'
	LocalCourier = 'LocalCourier'
	Other = 'Other'
