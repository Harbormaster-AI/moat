from django.db import models

#======================================================================
# 
# Encapsulates data for model Wishlist
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Wishlist Declaration
#======================================================================
class Wishlist (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	asPublic = models.BooleanField(null=True)
	createdAt = models.DateField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	items = models.ManyToManyField('WishlistItem',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.asPublic
		str = str + self.createdAt
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Wishlist";
    
	def objectType(self):
		return "Wishlist";
