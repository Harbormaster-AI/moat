from django.db import models
from manufacturingOnDjango.models.BOMStatus import BOMStatus

#======================================================================
# 
# Encapsulates data for model BOM
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BOM Declaration
#======================================================================
class BOM (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	bomNumber = models.CharField(max_length=200, null=True)
	revision = models.CharField(max_length=200, null=True)
	effectivityStart = models.DateField(null=True)
	effectivityEnd = models.DateField(null=True)
	parentItem = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	bomItems = models.ManyToManyField('BOMItem',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in BOMStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.bomNumber
		str = str + self.revision
		str = str + self.effectivityStart
		str = str + self.effectivityEnd
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "BOM";
    
	def objectType(self):
		return "BOM";
