from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.BusinessGlossaryTerm import BusinessGlossaryTerm
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Dimension import Dimension
from analyticsOnDjango.models.Measure import Measure
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BusinessGlossaryTerm
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BusinessGlossaryTermDelegate Declaration
#======================================================================
class BusinessGlossaryTermDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, businessGlossaryTermId ):
		try:	
			businessGlossaryTerm = BusinessGlossaryTerm.objects.filter(id=businessGlossaryTermId)
			return businessGlossaryTerm.first();
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError("BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, businessGlossaryTerm):
		for model in serializers.deserialize("json", businessGlossaryTerm):
			model.save()
			return model;

	def create(self, businessGlossaryTerm):
		businessGlossaryTerm.save()
		return businessGlossaryTerm;

	def saveFromJson(self, businessGlossaryTerm):
		for model in serializers.deserialize("json", businessGlossaryTerm):
			model.save()
			return businessGlossaryTerm;
	
	def save(self, businessGlossaryTerm):
		businessGlossaryTerm.save()
		return businessGlossaryTerm;
	
	def delete(self, businessGlossaryTermId ):
		errMsg = "Failed to delete BusinessGlossaryTerm from db using id " + str(businessGlossaryTermId)
		
		try:
			businessGlossaryTerm = BusinessGlossaryTerm.objects.get(id=businessGlossaryTermId)
			businessGlossaryTerm.delete()
			return True
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError("BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BusinessGlossaryTerm.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BusinessGlossaryTerm from db")
		except Exception:
			return None;
		
	def addRelatedTerms( self, businessGlossaryTermId, relatedTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to add elements " + str(relatedTermsIds) + " for RelatedTerms on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = relatedTermsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				businessGlossaryTerm.relatedTerms.add(businessGlossaryTerm)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRelatedTerms( self, businessGlossaryTermId, relatedTermsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.BusinessGlossaryTermDelegate import BusinessGlossaryTermDelegate

		errMsg = "Failed to remove elements " + str(relatedTermsIds) + " for RelatedTerms on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = relatedTermsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BusinessGlossaryTerm		
				businessGlossaryTerm = BusinessGlossaryTermDelegate().get(id).first();	
				# add the BusinessGlossaryTerm
				businessGlossaryTerm.relatedTerms.remove(businessGlossaryTerm)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMetrics( self, businessGlossaryTermId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to add elements " + str(metricsIds) + " for Metrics on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				businessGlossaryTerm.metrics.add(metric)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMetrics( self, businessGlossaryTermId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to remove elements " + str(metricsIds) + " for Metrics on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				businessGlossaryTerm.metrics.remove(metric)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, businessGlossaryTermId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				businessGlossaryTerm.datasets.add(dataSet)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, businessGlossaryTermId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				businessGlossaryTerm.datasets.remove(dataSet)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDimensions( self, businessGlossaryTermId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to add elements " + str(dimensionsIds) + " for Dimensions on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				businessGlossaryTerm.dimensions.add(dimension)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDimensions( self, businessGlossaryTermId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to remove elements " + str(dimensionsIds) + " for Dimensions on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				businessGlossaryTerm.dimensions.remove(dimension)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMeasures( self, businessGlossaryTermId, measuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

		errMsg = "Failed to add elements " + str(measuresIds) + " for Measures on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = measuresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Measure		
				measure = MeasureDelegate().get(id).first();	
				# add the Measure
				businessGlossaryTerm.measures.add(measure)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMeasures( self, businessGlossaryTermId, measuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MeasureDelegate import MeasureDelegate

		errMsg = "Failed to remove elements " + str(measuresIds) + " for Measures on BusinessGlossaryTerm"

		try:
			# get the BusinessGlossaryTerm
			businessGlossaryTerm = self.get( businessGlossaryTermId ).first()
				
			# split on a comma with no spaces
			idList = measuresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Measure		
				measure = MeasureDelegate().get(id).first();	
				# add the Measure
				businessGlossaryTerm.measures.remove(measure)
				
			# save it		
			businessGlossaryTerm.save()
			
			# reload and return the appropriate version
			return self.get( businessGlossaryTermId );
		except BusinessGlossaryTerm.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessGlossaryTerm with id " + str(businessGlossaryTermId) + " does not exist.")
		except Measure.DoesNotExist:
			raise ProcessingError(errMsg + " : Measure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
