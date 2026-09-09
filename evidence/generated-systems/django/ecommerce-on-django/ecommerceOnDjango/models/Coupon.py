from django.db import models
from ecommerceOnDjango.models.CouponStatus import CouponStatus

#======================================================================
# 
# Encapsulates data for model Coupon
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Coupon Declaration
#======================================================================
class Coupon (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	code = models.CharField(max_length=200, null=True)
	usageLimit = models.IntegerField(null=True)
	perCustomerLimit = models.IntegerField(null=True)
	expirationDate = models.DateField(null=True)
	promotion = models.ForeignKey('Promotion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	redemptions = models.ManyToManyField('CouponRedemption',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CouponStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.code
		str = str + self.usageLimit
		str = str + self.perCustomerLimit
		str = str + self.expirationDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Coupon";
    
	def objectType(self):
		return "Coupon";
