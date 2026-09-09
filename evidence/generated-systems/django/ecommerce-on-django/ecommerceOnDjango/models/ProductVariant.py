from django.db import models
from ecommerceOnDjango.models.WeightUnit import WeightUnit

#======================================================================
# 
# Encapsulates data for model ProductVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductVariant Declaration
#======================================================================
class ProductVariant (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	sku = SKU
	barcode = models.CharField(max_length=200, null=True)
	title = models.CharField(max_length=200, null=True)
	weight = models.CharField(max_length=64, null=True)
	requiresShipping = models.BooleanField(null=True)
	product = models.ForeignKey('Product', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	pricing = models.ManyToManyField('ProductPricing',  blank=True, related_name='+')
	inventoryItems = models.ManyToManyField('InventoryItem',  blank=True, related_name='+')
	mediaAssets = models.ManyToManyField('MediaAsset',  blank=True, related_name='+')
	subscriptions = models.ManyToManyField('Subscription',  blank=True, related_name='+')
	cartItems = models.ManyToManyField('CartItem',  blank=True, related_name='+')
	orderLines = models.ManyToManyField('OrderLine',  blank=True, related_name='+')
	wishlistItems = models.ManyToManyField('WishlistItem',  blank=True, related_name='+')
	weightUnit = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in WeightUnit])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.sku
		str = str + self.barcode
		str = str + self.title
		str = str + self.weight
		str = str + self.requiresShipping
		str = str + self.weightUnit
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ProductVariant";
    
	def objectType(self):
		return "ProductVariant";
