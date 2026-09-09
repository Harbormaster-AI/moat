from django.db import models
from analyticsOnDjango.models.QualityDimension import QualityDimension
from analyticsOnDjango.models.ComparisonOperator import ComparisonOperator

#======================================================================
# 
# Encapsulates data for model QualityRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityRule Declaration
#======================================================================
class QualityRule (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	threshold = Threshold
	targetField = models.CharField(max_length=200, null=True)
	dataset = models.ForeignKey('DataSet', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	checks = models.ManyToManyField('QualityCheck',  blank=True, related_name='+')
	dimension = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in QualityDimension])
	operator = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ComparisonOperator])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.threshold
		str = str + self.targetField
		str = str + self.dimension
		str = str + self.operator
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "QualityRule";
    
	def objectType(self):
		return "QualityRule";
