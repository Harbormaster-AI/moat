from django.db import models
from advertisingOnDjango.models.ReportType import ReportType

#======================================================================
# 
# Encapsulates data for model Report
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Report Declaration
#======================================================================
class Report (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reportName = models.CharField(max_length=200, null=True)
	generatedAt = models.CharField(max_length=64, null=True)
	fileUrl = URL
	adAccount = models.ForeignKey('AdAccount', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	campaign = models.ForeignKey('Campaign', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	lineItem = models.ForeignKey('LineItem', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	reportType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReportType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reportName
		str = str + self.generatedAt
		str = str + self.fileUrl
		str = str + self.reportType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Report";
    
	def objectType(self):
		return "Report";
