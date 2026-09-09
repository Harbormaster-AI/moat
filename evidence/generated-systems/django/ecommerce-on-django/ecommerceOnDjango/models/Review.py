from django.db import models
from ecommerceOnDjango.models.ReviewStatus import ReviewStatus

#======================================================================
# 
# Encapsulates data for model Review
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Review Declaration
#======================================================================
class Review (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	rating = models.IntegerField(null=True)
	title = models.CharField(max_length=200, null=True)
	content = models.CharField(max_length=200, null=True)
	createdAt = models.DateField(null=True)
	product = models.ForeignKey('Product', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	customer = models.ForeignKey('Customer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	order = models.ForeignKey('Order', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ReviewStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.rating
		str = str + self.title
		str = str + self.content
		str = str + self.createdAt
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Review";
    
	def objectType(self):
		return "Review";
