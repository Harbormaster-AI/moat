from django.db import models
from healthcareOnDjango.models.ImagingModality import ImagingModality

#======================================================================
# 
# Encapsulates data for model ImagingOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingOrder Declaration
#======================================================================
class ImagingOrder (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	bodySite = models.CharField(max_length=200, null=True)
	contrast = models.BooleanField(null=True)
	order = models.ForeignKey('ClinicalOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	imagingCenter = models.ForeignKey('ImagingCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reports = models.ManyToManyField('ImagingReport',  blank=True, related_name='+')
	modality = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ImagingModality])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.bodySite
		str = str + self.contrast
		str = str + self.modality
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ImagingOrder";
    
	def objectType(self):
		return "ImagingOrder";
