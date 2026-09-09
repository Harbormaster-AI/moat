from django.db import models
from ecommerceOnDjango.models.MediaType import MediaType

#======================================================================
# 
# Encapsulates data for model MediaAsset
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MediaAsset Declaration
#======================================================================
class MediaAsset (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	url = models.CharField(max_length=200, null=True)
	altText = models.CharField(max_length=200, null=True)
	position = models.IntegerField(null=True)
	product = models.ForeignKey('Product', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	variant = models.ForeignKey('ProductVariant', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	mediaType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in MediaType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.url
		str = str + self.altText
		str = str + self.position
		str = str + self.mediaType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MediaAsset";
    
	def objectType(self):
		return "MediaAsset";
