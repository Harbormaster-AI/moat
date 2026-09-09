from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Measure import Measure
from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.BusinessGlossaryTerm import BusinessGlossaryTerm
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Measure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MeasureDelegate Declaration
#======================================================================
class MeasureDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, measureId ):
		try:	
			measure = Measure.objects.filter(id=measureId)
			return measure.first();
		except Measure.DoesNotExist:
			raise ProcessingError("Measure with id " + str(measureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, measure):
		for model in serializers.deserialize("json", measure):
			model.save()
			return model;

	def create(self, measure):
		measure.save()
		return measure;

	def saveFromJson(self, measure):
		for model in serializers.deserialize("json", measure):
			model.save()
			return measure;
	
	def save(self, measure):
		measure.save()
		return measure;
	
	def delete(self, measureId ):
		errMsg = "Failed to delete Measure from db using id " + str(measureId)
		
		try:
			measure = Measure.objects.get(id=measureId)
			measure.delete()
			return True
		except Measure.DoesNotExist:
			raise ProcessingError("Measure with id " + str(measureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Measure.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Measure from db")
		except Exception:
			return None;
		
	def assignSemanticModel( self, measureId, semanticModelId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.SemanticModelDelegate import SemanticModelDelegate

		errMsg = "Failed to assign element " + str(semanticModelId) + " for SemanticModel on Measure"

		try:
			# get the Measure from db
			measure = self.get( measureId ).first()	
			
			# get the SemanticModel from db
			semanticModel = SemanticModelDelegate().get(semanticModelId).first();
			
			# assign the SemanticModel		
			measure.semanticModel = semanticModel
			
			#save it
			measure.save()

			# reload and return the appropriate version					
			return self.get( measureId );
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure with id " + str(measureId) + " does not exist.")
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSemanticModel( self, measureId ):
		errMsg = "Failed to unassign element " + str(semanticModelId) + " for SemanticModel on Measure"

		try:
			# get the Measure from db
			measure = self.get( measureId ).first()	
			
			# assign to None for unassignment
			measure.semanticModel = None			

			#save it
			measure.save()

			# reload and return the appropriate version					
			return self.get( measureId );
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure with id " + str(measureId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, measureId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Measure"

		try:
			# get the Measure
			measure = self.get( measureId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				measure.datasets.add(dataSet)
				
			# save it		
			measure.save()
			
			# reload and return the appropriate version
			return self.get( measureId );
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure with id " + str(measureId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, measureId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Measure"

		try:
			# get the Measure
			measure = self.get( measureId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				measure.datasets.remove(dataSet)
				
			# save it		
			measure.save()
			
			# reload and return the appropriate version
			return self.get( measureId );
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure with id " + str(measureId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGlossaryTerms( self, measureId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to add elements " + str(glossaryTermsIds) + " for GlossaryTerms on Measure"

		try:
			# get the Measure
			measure = self.get( measureId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				measure.glossaryTerms.add(businessGlossaryTerm)
				
			# save it		
			measure.save()
			
			# reload and return the appropriate version
			return self.get( measureId );
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure with id " + str(measureId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGlossaryTerms( self, measureId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to remove elements " + str(glossaryTermsIds) + " for GlossaryTerms on Measure"

		try:
			# get the Measure
			measure = self.get( measureId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				measure.glossaryTerms.remove(businessGlossaryTerm)
				
			# save it		
			measure.save()
			
			# reload and return the appropriate version
			return self.get( measureId );
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure with id " + str(measureId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
