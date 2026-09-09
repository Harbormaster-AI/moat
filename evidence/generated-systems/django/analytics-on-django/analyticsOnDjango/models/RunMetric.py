from django.db import models

#======================================================================
# 
# Encapsulates data for model RunMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunMetric Declaration
#======================================================================
class RunMetric (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	value = models.CharField(max_length=64, null=True)
	trainingRun = models.ForeignKey('TrainingRun', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	metric = models.ForeignKey('Metric', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.value
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RunMetric";
    
	def objectType(self):
		return "RunMetric";
