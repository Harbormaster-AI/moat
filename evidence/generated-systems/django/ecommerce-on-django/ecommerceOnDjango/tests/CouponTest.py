import datetime

from django.test import TestCase
from django.utils import timezone
from ecommerceOnDjango.models.Coupon import Coupon
from ecommerceOnDjango.delegates.CouponDelegate import CouponDelegate

 #======================================================================
# 
# Encapsulates data for model Coupon
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponTest Declaration
#======================================================================
class CouponTest (TestCase) :
	def test_crud(self) :
		coupon = Coupon()
		coupon.code = "default code field value"
		coupon.usageLimit = 22
		coupon.perCustomerLimit = 22
		coupon.expirationDate = datetime.datetime.now()
		coupon.status = "default status field value"
		
		delegate = CouponDelegate()
		responseObj = delegate.create(coupon)
		
		self.assertEqual(responseObj, delegate.get( responseObj.id ))
	
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 1 )		
		delegate.delete(responseObj.id)
		
		allObj = delegate.getAll()
		self.assertEqual(allObj.count(), 0 )		


