from django.db import models
from analyticsOnDjango.models.DataFormat import DataFormat

#======================================================================
# 
# Encapsulates data for model DataSet
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataSet Declaration
#======================================================================
class DataSet (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	schemaVersion = models.CharField(max_length=200, null=True)
	refreshSchedule = CronSchedule
	sensitive = models.BooleanField(null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	sources = models.ManyToManyField('DataSource',  blank=True, related_name='+')
	pipelines = models.ManyToManyField('DataPipeline',  blank=True, related_name='+')
	semanticModels = models.ManyToManyField('SemanticModel',  blank=True, related_name='+')
	dimensions = models.ManyToManyField('Dimension',  blank=True, related_name='+')
	measures = models.ManyToManyField('Measure',  blank=True, related_name='+')
	metrics = models.ManyToManyField('Metric',  blank=True, related_name='+')
	qualityRules = models.ManyToManyField('QualityRule',  blank=True, related_name='+')
	lineageNode = models.ForeignKey('LineageNode', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	tags = models.ManyToManyField('Tag',  blank=True, related_name='+')
	dataFormat = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DataFormat])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.schemaVersion
		str = str + self.refreshSchedule
		str = str + self.sensitive
		str = str + self.dataFormat
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DataSet";
    
	def objectType(self):
		return "DataSet";
