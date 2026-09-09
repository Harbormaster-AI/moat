import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.GiftCardRedemption import GiftCardRedemption
from ecommerceOnDjango.delegates.GiftCardRedemptionDelegate import GiftCardRedemptionDelegate

 #======================================================================
# 
# Encapsulates data for model GiftCardRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardRedemptionTest Declaration
#======================================================================
class GiftCardRedemptionTest (TestCase) :
	def test_crud(self) :
		giftCardRedemption = GiftCardRedemption()
		giftCardRedemption.redeemedAt = datetime.datetime.now()
		giftCardRedemption.amount = "default amount field value"
		
		delegate = GiftCardRedemptionDelegate()
		responseObj = delegate.create(giftCardRedemption)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


