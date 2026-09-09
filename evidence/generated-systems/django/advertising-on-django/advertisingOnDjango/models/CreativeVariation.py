from django.db import models

#======================================================================
# 
# Encapsulates data for model CreativeVariation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeVariation Declaration
#======================================================================
class CreativeVariation (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	language = models.CharField(max_length=200, null=True)
	headline = models.CharField(max_length=200, null=True)
	bodyText = models.CharField(max_length=200, null=True)
	callToAction = models.CharField(max_length=200, null=True)
	creativeAsset = models.ForeignKey('CreativeAsset', on_delete=models.CASCADE, null=True, blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.language
		str = str + self.headline
		str = str + self.bodyText
		str = str + self.callToAction
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CreativeVariation";
    
	def objectType(self):
		return "CreativeVariation";
