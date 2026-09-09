from django.db import models
from manufacturingOnDjango.models.SamplingPlanType import SamplingPlanType
from manufacturingOnDjango.models.QualityPlanStatus import QualityPlanStatus

#======================================================================
# 
# Encapsulates data for model InspectionPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionPlan Declaration
#======================================================================
class InspectionPlan (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	planNumber = models.CharField(max_length=200, null=True)
	revision = models.CharField(max_length=200, null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	characteristics = models.ManyToManyField('InspectionCharacteristic',  blank=True, related_name='+')
	samplingPlan = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SamplingPlanType])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in QualityPlanStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.planNumber
		str = str + self.revision
		str = str + self.samplingPlan
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InspectionPlan";
    
	def objectType(self):
		return "InspectionPlan";
