import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

 #======================================================================
# 
# Encapsulates data for model Promotion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PromotionTest Declaration
#======================================================================
class PromotionTest (TestCase) :
	def test_crud(self) :
		promotion = Promotion()
		promotion.name = "default name field value"
		promotion.code = "default code field value"
		promotion.value = "default value field value"
		promotion.startDate = datetime.datetime.now()
		promotion.endDate = datetime.datetime.now()
		promotion.asStackable = False
		promotion.maxRedemptions = 22
		promotion.promotionType = "default promotionType field value"
		promotion.discountType = "default discountType field value"
		
		delegate = PromotionDelegate()
		responseObj = delegate.create(promotion)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


