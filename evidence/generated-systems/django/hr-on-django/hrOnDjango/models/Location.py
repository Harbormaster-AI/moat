from django.db import models

#======================================================================
# 
# Encapsulates data for model Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Location Declaration
#======================================================================
class Location (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	address = Address
	timezone = models.CharField(max_length=200, null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	departments = models.ManyToManyField('Department',  blank=True, related_name='+')
	positions = models.ManyToManyField('Position',  blank=True, related_name='+')
	employees = models.ManyToManyField('Employee',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.address
		str = str + self.timezone
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Location";
    
	def objectType(self):
		return "Location";
