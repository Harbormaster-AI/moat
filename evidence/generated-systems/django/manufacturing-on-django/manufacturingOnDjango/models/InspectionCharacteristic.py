from django.db import models
from manufacturingOnDjango.models.MeasurementType import MeasurementType

#======================================================================
# 
# Encapsulates data for model InspectionCharacteristic
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionCharacteristic Declaration
#======================================================================
class InspectionCharacteristic (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	characteristicCode = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	lowerSpecLimit = Measurement
	upperSpecLimit = Measurement
	target = Measurement
	inspectionPlan = models.ForeignKey('InspectionPlan', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	measurementType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MeasurementType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.characteristicCode
		str = str + self.name
		str = str + self.lowerSpecLimit
		str = str + self.upperSpecLimit
		str = str + self.target
		str = str + self.measurementType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InspectionCharacteristic";
    
	def objectType(self):
		return "InspectionCharacteristic";
