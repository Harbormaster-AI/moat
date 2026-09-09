from django.db import models

#======================================================================
# 
# Encapsulates data for model CouponRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponRedemption Declaration
#======================================================================
class CouponRedemption (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	redeemedAt = models.DateField(null=True)
	coupon = models.ForeignKey('Coupon', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.redeemedAt
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CouponRedemption";
    
	def objectType(self):
		return "CouponRedemption";
