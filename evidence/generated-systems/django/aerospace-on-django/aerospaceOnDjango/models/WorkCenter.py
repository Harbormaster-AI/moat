from django.db import models

#======================================================================
# 
# Encapsulates data for model WorkCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkCenter Declaration
#======================================================================
class WorkCenter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	capability = models.CharField(max_length=200, null=True)
	productionLine = models.ForeignKey('ProductionLine', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.capability
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "WorkCenter";
    
	def objectType(self):
		return "WorkCenter";
