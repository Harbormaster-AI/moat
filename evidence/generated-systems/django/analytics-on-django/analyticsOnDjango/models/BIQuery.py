from django.db import models
from analyticsOnDjango.models.SQLDialect import SQLDialect

#======================================================================
# 
# Encapsulates data for model BIQuery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BIQuery Declaration
#======================================================================
class BIQuery (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	text = models.CharField(max_length=200, null=True)
	workspace = models.ForeignKey('AnalyticsWorkspace', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	reports = models.ManyToManyField('Report',  blank=True, related_name='+')
	dashboards = models.ManyToManyField('Dashboard',  blank=True, related_name='+')
	notebooks = models.ManyToManyField('Notebook',  blank=True, related_name='+')
	dialect = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in SQLDialect])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.text
		str = str + self.dialect
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BIQuery";
    
	def objectType(self):
		return "BIQuery";
