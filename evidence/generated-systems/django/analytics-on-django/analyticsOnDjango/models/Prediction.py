from django.db import models

#======================================================================
# 
# Encapsulates data for model Prediction
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Prediction Declaration
#======================================================================
class Prediction (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	referenceKey = models.CharField(max_length=200, null=True)
	predictedAt = models.DateField(null=True)
	score = models.CharField(max_length=64, null=True)
	endpoint = models.ForeignKey('InferenceEndpoint', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	modelVersion = models.ForeignKey('ModelVersion', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.referenceKey
		str = str + self.predictedAt
		str = str + self.score
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Prediction";
    
	def objectType(self):
		return "Prediction";
