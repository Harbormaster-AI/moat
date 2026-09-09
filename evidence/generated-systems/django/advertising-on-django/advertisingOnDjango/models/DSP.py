from django.db import models

#======================================================================
# 
# Encapsulates data for model DSP
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DSP Declaration
#======================================================================
class DSP (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	website = models.CharField(max_length=200, null=True)
	region = models.CharField(max_length=200, null=True)
	adAccounts = models.ManyToManyField('AdAccount',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.website
		str = str + self.region
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DSP";
    
	def objectType(self):
		return "DSP";
