from django.db import models
from analyticsOnDjango.models.DimensionType import DimensionType

#======================================================================
# 
# Encapsulates data for model Dimension
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Dimension Declaration
#======================================================================
class Dimension (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	typeTime = models.BooleanField(null=True)
	semanticModel = models.ForeignKey('SemanticModel', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	datasets = models.ManyToManyField('DataSet',  blank=True, related_name='+')
	glossaryTerms = models.ManyToManyField('BusinessGlossaryTerm',  blank=True, related_name='+')
	dimensionType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DimensionType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.typeTime
		str = str + self.dimensionType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Dimension";
    
	def objectType(self):
		return "Dimension";
