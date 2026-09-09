from django.db import models
from manufacturingOnDjango.models.InspectionResultStatus import InspectionResultStatus

#======================================================================
# 
# Encapsulates data for model InspectionResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionResult Declaration
#======================================================================
class InspectionResult (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	resultValue = Measurement
	recordedOn = models.CharField(max_length=64, null=True)
	notes = models.CharField(max_length=200, null=True)
	inspectionLot = models.ForeignKey('InspectionLot', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	characteristic = models.ForeignKey('InspectionCharacteristic', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	resultStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in InspectionResultStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.resultValue
		str = str + self.recordedOn
		str = str + self.notes
		str = str + self.resultStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "InspectionResult";
    
	def objectType(self):
		return "InspectionResult";
