from django.db import models

#======================================================================
# 
# Encapsulates data for model RunParameter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunParameter Declaration
#======================================================================
class RunParameter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	value = models.CharField(max_length=200, null=True)
	trainingRun = models.ForeignKey('TrainingRun', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

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
		return "RunParameter";
    
	def objectType(self):
		return "RunParameter";
