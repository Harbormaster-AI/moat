from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Dimension import Dimension
from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.BusinessGlossaryTerm import BusinessGlossaryTerm
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Dimension
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DimensionDelegate Declaration
#======================================================================
class DimensionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dimensionId ):
		try:	
			dimension = Dimension.objects.filter(id=dimensionId)
			return dimension.first();
		except Dimension.DoesNotExist:
			raise ProcessingError("Dimension with id " + str(dimensionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dimension):
		for model in serializers.deserialize("json", dimension):
			model.save()
			return model;

	def create(self, dimension):
		dimension.save()
		return dimension;

	def saveFromJson(self, dimension):
		for model in serializers.deserialize("json", dimension):
			model.save()
			return dimension;
	
	def save(self, dimension):
		dimension.save()
		return dimension;
	
	def delete(self, dimensionId ):
		errMsg = "Failed to delete Dimension from db using id " + str(dimensionId)
		
		try:
			dimension = Dimension.objects.get(id=dimensionId)
			dimension.delete()
			return True
		except Dimension.DoesNotExist:
			raise ProcessingError("Dimension with id " + str(dimensionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Dimension.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Dimension from db")
		except Exception:
			return None;
		
	def assignSemanticModel( self, dimensionId, semanticModelId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to assign element " + str(semanticModelId) + " for SemanticModel on Dimension"

		try:
			# get the Dimension from db
			dimension = self.get( dimensionId ).first()	
			
			# get the SemanticModel from db
			semanticModel = SemanticModelDelegate().get(semanticModelId).first();
			
			# assign the SemanticModel		
			dimension.semanticModel = semanticModel
			
			#save it
			dimension.save()

			# reload and return the appropriate version					
			return self.get( dimensionId );
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension with id " + str(dimensionId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSemanticModel( self, dimensionId ):
		errMsg = "Failed to unassign element " + str(semanticModelId) + " for SemanticModel on Dimension"

		try:
			# get the Dimension from db
			dimension = self.get( dimensionId ).first()	
			
			# assign to None for unassignment
			dimension.semanticModel = None			

			#save it
			dimension.save()

			# reload and return the appropriate version					
			return self.get( dimensionId );
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension with id " + str(dimensionId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, dimensionId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Dimension"

		try:
			# get the Dimension
			dimension = self.get( dimensionId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dimension.datasets.add(dataSet)
				
			# save it		
			dimension.save()
			
			# reload and return the appropriate version
			return self.get( dimensionId );
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension with id " + str(dimensionId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, dimensionId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Dimension"

		try:
			# get the Dimension
			dimension = self.get( dimensionId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dimension.datasets.remove(dataSet)
				
			# save it		
			dimension.save()
			
			# reload and return the appropriate version
			return self.get( dimensionId );
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension with id " + str(dimensionId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGlossaryTerms( self, dimensionId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to add elements " + str(glossaryTermsIds) + " for GlossaryTerms on Dimension"

		try:
			# get the Dimension
			dimension = self.get( dimensionId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				dimension.glossaryTerms.add(businessGlossaryTerm)
				
			# save it		
			dimension.save()
			
			# reload and return the appropriate version
			return self.get( dimensionId );
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension with id " + str(dimensionId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGlossaryTerms( self, dimensionId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to remove elements " + str(glossaryTermsIds) + " for GlossaryTerms on Dimension"

		try:
			# get the Dimension
			dimension = self.get( dimensionId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				dimension.glossaryTerms.remove(businessGlossaryTerm)
				
			# save it		
			dimension.save()
			
			# reload and return the appropriate version
			return self.get( dimensionId );
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension with id " + str(dimensionId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
