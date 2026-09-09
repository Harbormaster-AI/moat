from django.db import models

#======================================================================
# 
# Encapsulates data for model Catalog
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Catalog Declaration
#======================================================================
class Catalog (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	catalogCode = models.CharField(max_length=200, null=True)
	asActive = models.BooleanField(null=True)
	channel = models.ForeignKey('Channel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	categories = models.ManyToManyField('Category',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.catalogCode
		str = str + self.asActive
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Catalog";
    
	def objectType(self):
		return "Catalog";
