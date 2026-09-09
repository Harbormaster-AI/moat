from django.db import models

#======================================================================
# 
# Encapsulates data for model QualitySpecification
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualitySpecification Declaration
#======================================================================
class QualitySpecification (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	specCode = models.CharField(max_length=200, null=True)
	name = models.CharField(max_length=200, null=True)
	version = models.CharField(max_length=200, null=True)
	item = models.ForeignKey('Item', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.specCode
		str = str + self.name
		str = str + self.version
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "QualitySpecification";
    
	def objectType(self):
		return "QualitySpecification";
