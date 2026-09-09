from django.db import models
from analyticsOnDjango.models.DataSourceType import DataSourceType
from analyticsOnDjango.models.DataFormat import DataFormat

#======================================================================
# 
# Encapsulates data for model DataSource
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSource Declaration
#======================================================================
class DataSource (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	connection = ConnectionInfo
	streaming = models.BooleanField(null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	producedDatasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	pipelines = models.ManyToManyField('DataPipeline',  blank=True, related_name='+')
	sourceType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataSourceType])
	format = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataFormat])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.connection
		str = str + self.streaming
		str = str + self.sourceType
		str = str + self.format
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataSource";
    
	def objectType(self):
		return "DataSource";
