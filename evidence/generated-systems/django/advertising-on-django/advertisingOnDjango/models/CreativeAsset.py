from django.db import models
from advertisingOnDjango.models.CreativeType import CreativeType
from advertisingOnDjango.models.AdFormat import AdFormat

#======================================================================
# 
# Encapsulates data for model CreativeAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeAsset Declaration
#======================================================================
class CreativeAsset (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	clickUrl = URL
	landingPage = URL
	width = models.IntegerField(null=True)
	height = models.IntegerField(null=True)
	durationSeconds = models.IntegerField(null=True)
	files = models.ManyToManyField('CreativeFile',  blank=True, related_name='+')
	approvals = models.ManyToManyField('CreativeApproval',  blank=True, related_name='+')
	variations = models.ManyToManyField('CreativeVariation',  blank=True, related_name='+')
	lineItems = models.ManyToManyField('LineItem',  blank=True, related_name='+')
	creativeType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CreativeType])
	adFormat = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AdFormat])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.clickUrl
		str = str + self.landingPage
		str = str + self.width
		str = str + self.height
		str = str + self.durationSeconds
		str = str + self.creativeType
		str = str + self.adFormat
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CreativeAsset";
    
	def objectType(self):
		return "CreativeAsset";
