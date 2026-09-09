from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.RunMetric import RunMetric
from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RunMetric
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RunMetricDelegate Declaration
#======================================================================
class RunMetricDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, runMetricId ):
		try:	
			runMetric = RunMetric.objects.filter(id=runMetricId)
			return runMetric.first();
		except RunMetric.DoesNotExist:
			raise ProcessingError("RunMetric with id " + str(runMetricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, runMetric):
		for model in serializers.deserialize("json", runMetric):
			model.save()
			return model;

	def create(self, runMetric):
		runMetric.save()
		return runMetric;

	def saveFromJson(self, runMetric):
		for model in serializers.deserialize("json", runMetric):
			model.save()
			return runMetric;
	
	def save(self, runMetric):
		runMetric.save()
		return runMetric;
	
	def delete(self, runMetricId ):
		errMsg = "Failed to delete RunMetric from db using id " + str(runMetricId)
		
		try:
			runMetric = RunMetric.objects.get(id=runMetricId)
			runMetric.delete()
			return True
		except RunMetric.DoesNotExist:
			raise ProcessingError("RunMetric with id " + str(runMetricId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RunMetric.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RunMetric from db")
		except Exception:
			return None;
		
	def assignTrainingRun( self, runMetricId, trainingRunId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.TrainingRunDelegate import TrainingRunDelegate

		errMsg = "Failed to assign element " + str(trainingRunId) + " for TrainingRun on RunMetric"

		try:
			# get the RunMetric from db
			runMetric = self.get( runMetricId ).first()	
			
			# get the TrainingRun from db
			trainingRun = TrainingRunDelegate().get(trainingRunId).first();
			
			# assign the TrainingRun		
			runMetric.trainingRun = trainingRun
			
			#save it
			runMetric.save()

			# reload and return the appropriate version					
			return self.get( runMetricId );
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric with id " + str(runMetricId) + " does not exist.")
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTrainingRun( self, runMetricId ):
		errMsg = "Failed to unassign element " + str(trainingRunId) + " for TrainingRun on RunMetric"

		try:
			# get the RunMetric from db
			runMetric = self.get( runMetricId ).first()	
			
			# assign to None for unassignment
			runMetric.trainingRun = None			

			#save it
			runMetric.save()

			# reload and return the appropriate version					
			return self.get( runMetricId );
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric with id " + str(runMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMetric( self, runMetricId, metricId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to assign element " + str(metricId) + " for Metric on RunMetric"

		try:
			# get the RunMetric from db
			runMetric = self.get( runMetricId ).first()	
			
			# get the Metric from db
			metric = MetricDelegate().get(metricId).first();
			
			# assign the Metric		
			runMetric.metric = metric
			
			#save it
			runMetric.save()

			# reload and return the appropriate version					
			return self.get( runMetricId );
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric with id " + str(runMetricId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric with id " + str(metricId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMetric( self, runMetricId ):
		errMsg = "Failed to unassign element " + str(metricId) + " for Metric on RunMetric"

		try:
			# get the RunMetric from db
			runMetric = self.get( runMetricId ).first()	
			
			# assign to None for unassignment
			runMetric.metric = None			

			#save it
			runMetric.save()

			# reload and return the appropriate version					
			return self.get( runMetricId );
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric with id " + str(runMetricId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, runMetricId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on RunMetric"

		try:
			# get the RunMetric from db
			runMetric = self.get( runMetricId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			runMetric.dataset = dataSet
			
			#save it
			runMetric.save()

			# reload and return the appropriate version					
			return self.get( runMetricId );
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric with id " + str(runMetricId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, runMetricId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on RunMetric"

		try:
			# get the RunMetric from db
			runMetric = self.get( runMetricId ).first()	
			
			# assign to None for unassignment
			runMetric.dataSet = None			

			#save it
			runMetric.save()

			# reload and return the appropriate version					
			return self.get( runMetricId );
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric with id " + str(runMetricId) + " does not exist.")
		except Exception:
			return None;
		
