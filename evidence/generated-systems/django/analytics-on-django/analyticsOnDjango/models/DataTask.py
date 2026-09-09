from django.db import models
from analyticsOnDjango.models.DataTaskType import DataTaskType

#======================================================================
# 
# Encapsulates data for model DataTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataTask Declaration
#======================================================================
class DataTask (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	command = models.CharField(max_length=200, null=True)
	retries = models.IntegerField(null=True)
	pipeline = models.ForeignKey('DataPipeline', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	inputDatasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	outputDatasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	taskType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataTaskType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.command
		str = str + self.retries
		str = str + self.taskType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataTask";
    
	def objectType(self):
		return "DataTask";
