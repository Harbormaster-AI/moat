from django.db import models
from healthcareOnDjango.models.ResultStatus import ResultStatus

#======================================================================
# 
# Encapsulates data for model ImagingReport
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ImagingReport Declaration
#======================================================================
class ImagingReport (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reportNumber = models.CharField(max_length=200, null=True)
	impression = models.CharField(max_length=200, null=True)
	reportedDate = models.CharField(max_length=64, null=True)
	imagingOrder = models.ForeignKey('ImagingOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	clinician = models.ForeignKey('Clinician', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	encounter = models.ForeignKey('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	imagingCenter = models.ForeignKey('ImagingCenter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ResultStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reportNumber
		str = str + self.impression
		str = str + self.reportedDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ImagingReport";
    
	def objectType(self):
		return "ImagingReport";
