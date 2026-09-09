import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.GiftCard import GiftCard
from ecommerceOnDjango.delegates.GiftCardDelegate import GiftCardDelegate

 #======================================================================
# 
# Encapsulates data for model GiftCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardTest Declaration
#======================================================================
class GiftCardTest (TestCase) :
	def test_crud(self) :
		giftCard = GiftCard()
		giftCard.code = "default code field value"
		giftCard.balance = "default balance field value"
		giftCard.expirationDate = datetime.datetime.now()
		giftCard.status = "default status field value"
		
		delegate = GiftCardDelegate()
		responseObj = delegate.create(giftCard)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


