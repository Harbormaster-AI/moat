import datetime

from django.test import TestCase
from django.utils import timezone
from hrOnDjango.models.Offer import Offer
from hrOnDjango.delegates.OfferDelegate import OfferDelegate

 #======================================================================
# 
# Encapsulates data for model Offer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OfferTest Declaration
#======================================================================
class OfferTest (TestCase) :
	def test_crud(self) :
		offer = Offer()
		offer.offerNumber = "default offerNumber field value"
		offer.proposedStartDate = datetime.datetime.now()
		offer.baseSalary = "default baseSalary field value"
		offer.signOnBonus = "default signOnBonus field value"
		offer.status = "default status field value"
		
		delegate = OfferDelegate()
		responseObj = delegate.create(offer)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


