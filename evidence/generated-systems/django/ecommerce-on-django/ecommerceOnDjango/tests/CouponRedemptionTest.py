import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.CouponRedemption import CouponRedemption
from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

 #======================================================================
# 
# Encapsulates data for model CouponRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponRedemptionTest Declaration
#======================================================================
class CouponRedemptionTest (TestCase) :
	def test_crud(self) :
		couponRedemption = CouponRedemption()
		couponRedemption.redeemedAt = datetime.datetime.now()
		
		delegate = CouponRedemptionDelegate()
		responseObj = delegate.create(couponRedemption)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


