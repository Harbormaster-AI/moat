from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.SemanticModel import SemanticModel
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.Dimension import Dimension
from analyticsOnDjango.models.Measure import Measure
from analyticsOnDjango.models.BusinessGlossaryTerm import BusinessGlossaryTerm
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SemanticModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SemanticModelDelegate Declaration
#======================================================================
class SemanticModelDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, semanticModelId ):
		try:	
			semanticModel = SemanticModel.objects.filter(id=semanticModelId)
			return semanticModel.first();
		except SemanticModel.DoesNotExist:
			raise ProcessingError("SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, semanticModel):
		for model in serializers.deserialize("json", semanticModel):
			model.save()
			return model;

	def create(self, semanticModel):
		semanticModel.save()
		return semanticModel;

	def saveFromJson(self, semanticModel):
		for model in serializers.deserialize("json", semanticModel):
			model.save()
			return semanticModel;
	
	def save(self, semanticModel):
		semanticModel.save()
		return semanticModel;
	
	def delete(self, semanticModelId ):
		errMsg = "Failed to delete SemanticModel from db using id " + str(semanticModelId)
		
		try:
			semanticModel = SemanticModel.objects.get(id=semanticModelId)
			semanticModel.delete()
			return True
		except SemanticModel.DoesNotExist:
			raise ProcessingError("SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SemanticModel.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SemanticModel from db")
		except Exception:
			return None;
		
	def addDatasets( self, semanticModelId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				semanticModel.datasets.add(dataSet)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, semanticModelId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				semanticModel.datasets.remove(dataSet)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMetrics( self, semanticModelId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to add elements " + str(metricsIds) + " for Metrics on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				semanticModel.metrics.add(metric)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMetrics( self, semanticModelId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to remove elements " + str(metricsIds) + " for Metrics on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				semanticModel.metrics.remove(metric)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDimensions( self, semanticModelId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to add elements " + str(dimensionsIds) + " for Dimensions on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				semanticModel.dimensions.add(dimension)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDimensions( self, semanticModelId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to remove elements " + str(dimensionsIds) + " for Dimensions on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				semanticModel.dimensions.remove(dimension)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMeasures( self, semanticModelId, measuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

		errMsg = "Failed to add elements " + str(measuresIds) + " for Measures on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = measuresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Measure		
				measure = MeasureDelegate().get(id).first();	
				# add the Measure
				semanticModel.measures.add(measure)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMeasures( self, semanticModelId, measuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

		errMsg = "Failed to remove elements " + str(measuresIds) + " for Measures on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = measuresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Measure		
				measure = MeasureDelegate().get(id).first();	
				# add the Measure
				semanticModel.measures.remove(measure)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGlossaryTerms( self, semanticModelId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to add elements " + str(glossaryTermsIds) + " for GlossaryTerms on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				semanticModel.glossaryTerms.add(businessGlossaryTerm)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGlossaryTerms( self, semanticModelId, glossaryTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to remove elements " + str(glossaryTermsIds) + " for GlossaryTerms on SemanticModel"

		try:
			# get the SemanticModel
			semanticModel = self.get( semanticModelId ).first()
				
			# split on a comma with no spaces
			idList = glossaryTermsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				semanticModel.glossaryTerms.remove(businessGlossaryTerm)
				
			# save it		
			semanticModel.save()
			
			# reload and return the appropriate version
			return self.get( semanticModelId );
		except SemanticModel.DoesNotExist:
			raise ProcessingError(errMsg + " : SemanticModel with id " + str(semanticModelId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
