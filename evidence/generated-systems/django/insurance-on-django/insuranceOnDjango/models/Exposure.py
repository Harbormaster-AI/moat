from django.db import models
from insuranceOnDjango.models.ExposureType import ExposureType
from insuranceOnDjango.models.ExposureStatus import ExposureStatus

#======================================================================
# 
# Encapsulates data for model Exposure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Exposure Declaration
#======================================================================
class Exposure (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	claim = models.ForeignKey('Claim', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	policyCoverage = models.ForeignKey('PolicyCoverage', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	insuredObject = models.ForeignKey('InsuredObject', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reserves = models.ManyToManyField('ClaimReserve',  blank=True, related_name='+')
	payments = models.ManyToManyField('ClaimPayment',  blank=True, related_name='+')
	exposureType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExposureType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ExposureStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.exposureType
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Exposure";
    
	def objectType(self):
		return "Exposure";
