from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.EvaluationMetric import EvaluationMetric
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model EvaluationMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvaluationMetricDelegate Declaration
#======================================================================
class EvaluationMetricDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, evaluationMetricId ):
		try:	
			evaluationMetric = EvaluationMetric.objects.filter(id=evaluationMetricId)
			return evaluationMetric.first();
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError("EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, evaluationMetric):
		for model in serializers.deserialize("json", evaluationMetric):
			model.save()
			return model;

	def create(self, evaluationMetric):
		evaluationMetric.save()
		return evaluationMetric;

	def saveFromJson(self, evaluationMetric):
		for model in serializers.deserialize("json", evaluationMetric):
			model.save()
			return evaluationMetric;
	
	def save(self, evaluationMetric):
		evaluationMetric.save()
		return evaluationMetric;
	
	def delete(self, evaluationMetricId ):
		errMsg = "Failed to delete EvaluationMetric from db using id " + str(evaluationMetricId)
		
		try:
			evaluationMetric = EvaluationMetric.objects.get(id=evaluationMetricId)
			evaluationMetric.delete()
			return True
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError("EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = EvaluationMetric.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all EvaluationMetric from db")
		except Exception:
			return None;
		
	def assignModelVersion( self, evaluationMetricId, modelVersionId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to assign element " + str(modelVersionId) + " for ModelVersion on EvaluationMetric"

		try:
			# get the EvaluationMetric from db
			evaluationMetric = self.get( evaluationMetricId ).first()	
			
			# get the ModelVersion from db
			modelVersion = ModelVersionDelegate().get(modelVersionId).first();
			
			# assign the ModelVersion		
			evaluationMetric.modelVersion = modelVersion
			
			#save it
			evaluationMetric.save()

			# reload and return the appropriate version					
			return self.get( evaluationMetricId );
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModelVersion( self, evaluationMetricId ):
		errMsg = "Failed to unassign element " + str(modelVersionId) + " for ModelVersion on EvaluationMetric"

		try:
			# get the EvaluationMetric from db
			evaluationMetric = self.get( evaluationMetricId ).first()	
			
			# assign to None for unassignment
			evaluationMetric.modelVersion = None			

			#save it
			evaluationMetric.save()

			# reload and return the appropriate version					
			return self.get( evaluationMetricId );
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMetric( self, evaluationMetricId, metricId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to assign element " + str(metricId) + " for Metric on EvaluationMetric"

		try:
			# get the EvaluationMetric from db
			evaluationMetric = self.get( evaluationMetricId ).first()	
			
			# get the Metric from db
			metric = MetricDelegate().get(metricId).first();
			
			# assign the Metric		
			evaluationMetric.metric = metric
			
			#save it
			evaluationMetric.save()

			# reload and return the appropriate version					
			return self.get( evaluationMetricId );
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMetric( self, evaluationMetricId ):
		errMsg = "Failed to unassign element " + str(metricId) + " for Metric on EvaluationMetric"

		try:
			# get the EvaluationMetric from db
			evaluationMetric = self.get( evaluationMetricId ).first()	
			
			# assign to None for unassignment
			evaluationMetric.metric = None			

			#save it
			evaluationMetric.save()

			# reload and return the appropriate version					
			return self.get( evaluationMetricId );
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, evaluationMetricId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on EvaluationMetric"

		try:
			# get the EvaluationMetric from db
			evaluationMetric = self.get( evaluationMetricId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			evaluationMetric.dataset = dataSet
			
			#save it
			evaluationMetric.save()

			# reload and return the appropriate version					
			return self.get( evaluationMetricId );
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, evaluationMetricId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on EvaluationMetric"

		try:
			# get the EvaluationMetric from db
			evaluationMetric = self.get( evaluationMetricId ).first()	
			
			# assign to None for unassignment
			evaluationMetric.dataSet = None			

			#save it
			evaluationMetric.save()

			# reload and return the appropriate version					
			return self.get( evaluationMetricId );
		except EvaluationMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : EvaluationMetric with id " + str(evaluationMetricId) + " does not exist.")
		except Exception:
			return None;
		
