from django.db import models

#======================================================================
# 
# Encapsulates data for model ImagingCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingCenter Declaration
#======================================================================
class ImagingCenter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	imagingOrders = models.ManyToManyField('ImagingOrder',  blank=True, related_name='+')
	imagingReports = models.ManyToManyField('ImagingReport',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ImagingCenter";
    
	def objectType(self):
		return "ImagingCenter";
