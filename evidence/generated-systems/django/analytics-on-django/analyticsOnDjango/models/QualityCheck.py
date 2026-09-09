from django.db import models
from analyticsOnDjango.models.QualityStatus import QualityStatus

#======================================================================
# 
# Encapsulates data for model QualityCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityCheck Declaration
#======================================================================
class QualityCheck (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	checkedAt = models.DateField(null=True)
	observedValue = models.CharField(max_length=64, null=True)
	sampleSize = models.IntegerField(null=True)
	rule = models.ForeignKey('QualityRule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in QualityStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.checkedAt
		str = str + self.observedValue
		str = str + self.sampleSize
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "QualityCheck";
    
	def objectType(self):
		return "QualityCheck";
