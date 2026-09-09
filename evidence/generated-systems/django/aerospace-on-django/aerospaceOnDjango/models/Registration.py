from django.db import models

#======================================================================
# 
# Encapsulates data for model Registration
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Registration Declaration
#======================================================================
class Registration (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	tailNumber = TailNumber
	registryCountry = models.CharField(max_length=200, null=True)
	aircraft = models.OneToOneField('Aircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.tailNumber
		str = str + self.registryCountry
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Registration";
    
	def objectType(self):
		return "Registration";
