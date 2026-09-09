from django.db import models

#======================================================================
# 
# Encapsulates data for model Category
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Category Declaration
#======================================================================
class Category (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	slug = models.CharField(max_length=200, null=True)
	position = models.IntegerField(null=True)
	asActive = models.BooleanField(null=True)
	catalog = models.ForeignKey('Catalog', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	parentCategory = models.ForeignKey('self', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	subcategories = models.ManyToManyField('Category',  blank=True, related_name='+')
	products = models.ManyToManyField('Product',  blank=True, related_name='+')

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.slug
		str = str + self.position
		str = str + self.asActive
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Category";
    
	def objectType(self):
		return "Category";
