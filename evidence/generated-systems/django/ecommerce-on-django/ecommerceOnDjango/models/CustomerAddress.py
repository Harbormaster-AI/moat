from django.db import models

#======================================================================
# 
# Encapsulates data for model CustomerAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerAddress Declaration
#======================================================================
class CustomerAddress (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	label = models.CharField(max_length=200, null=True)
	address = Address
	asDefaultShipping = models.BooleanField(null=True)
	asDefaultBilling = models.BooleanField(null=True)
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.label
		str = str + self.address
		str = str + self.asDefaultShipping
		str = str + self.asDefaultBilling
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CustomerAddress";
    
	def objectType(self):
		return "CustomerAddress";
