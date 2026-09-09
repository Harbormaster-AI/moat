from django.db import models

#======================================================================
# 
# Encapsulates data for model WishlistItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistItem Declaration
#======================================================================
class WishlistItem (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	addedDate = models.DateField(null=True)
	wishlist = models.ForeignKey('Wishlist', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.addedDate
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WishlistItem";
    
	def objectType(self):
		return "WishlistItem";
