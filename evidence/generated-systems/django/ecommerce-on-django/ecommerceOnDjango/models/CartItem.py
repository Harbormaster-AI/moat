from django.db import models

#======================================================================
# 
# Encapsulates data for model CartItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartItem Declaration
#======================================================================
class CartItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	quantity = models.IntegerField(null=True)
	unitPrice = Money
	totalPrice = Money
	cart = models.ForeignKey('Cart', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	appliedPromotions = models.ManyToManyField('Promotion',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.quantity
		str = str + self.unitPrice
		str = str + self.totalPrice
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CartItem";
    
	def objectType(self):
		return "CartItem";
