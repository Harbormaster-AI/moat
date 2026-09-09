from django.db import models
from ecommerceOnDjango.models.PromotionType import PromotionType
from ecommerceOnDjango.models.DiscountType import DiscountType

#======================================================================
# 
# Encapsulates data for model Promotion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Promotion Declaration
#======================================================================
class Promotion (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	code = models.CharField(max_length=200, null=True)
	value = models.CharField(max_length=64, null=True)
	startDate = models.DateField(null=True)
	endDate = models.DateField(null=True)
	asStackable = models.BooleanField(null=True)
	maxRedemptions = models.IntegerField(null=True)
	merchant = models.ForeignKey('Merchant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	channels = models.ManyToManyField('Channel',  blank=True, related_name='+')
	applicableProducts = models.ManyToManyField('Product',  blank=True, related_name='+')
	applicableCategories = models.ManyToManyField('Category',  blank=True, related_name='+')
	coupons = models.ManyToManyField('Coupon',  blank=True, related_name='+')
	promotionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PromotionType])
	discountType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DiscountType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.code
		str = str + self.value
		str = str + self.startDate
		str = str + self.endDate
		str = str + self.asStackable
		str = str + self.maxRedemptions
		str = str + self.promotionType
		str = str + self.discountType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Promotion";
    
	def objectType(self):
		return "Promotion";
