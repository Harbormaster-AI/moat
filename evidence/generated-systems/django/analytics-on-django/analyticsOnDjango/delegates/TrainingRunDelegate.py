from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.TrainingRun import TrainingRun
from analyticsOnDjango.models.Experiment import Experiment
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Feature import Feature
from analyticsOnDjango.models.RunMetric import RunMetric
from analyticsOnDjango.models.RunParameter import RunParameter
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TrainingRun
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TrainingRunDelegate Declaration
#======================================================================
class TrainingRunDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, trainingRunId ):
		try:	
			trainingRun = TrainingRun.objects.filter(id=trainingRunId)
			return trainingRun.first();
		except TrainingRun.DoesNotExist:
			raise ProcessingError("TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, trainingRun):
		for model in serializers.deserialize("json", trainingRun):
			model.save()
			return model;

	def create(self, trainingRun):
		trainingRun.save()
		return trainingRun;

	def saveFromJson(self, trainingRun):
		for model in serializers.deserialize("json", trainingRun):
			model.save()
			return trainingRun;
	
	def save(self, trainingRun):
		trainingRun.save()
		return trainingRun;
	
	def delete(self, trainingRunId ):
		errMsg = "Failed to delete TrainingRun from db using id " + str(trainingRunId)
		
		try:
			trainingRun = TrainingRun.objects.get(id=trainingRunId)
			trainingRun.delete()
			return True
		except TrainingRun.DoesNotExist:
			raise ProcessingError("TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TrainingRun.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TrainingRun from db")
		except Exception:
			return None;
		
	def assignExperiment( self, trainingRunId, experimentId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to assign element " + str(experimentId) + " for Experiment on TrainingRun"

		try:
			# get the TrainingRun from db
			trainingRun = self.get( trainingRunId ).first()	
			
			# get the Experiment from db
			experiment = ExperimentDelegate().get(experimentId).first();
			
			# assign the Experiment		
			trainingRun.experiment = experiment
			
			#save it
			trainingRun.save()

			# reload and return the appropriate version					
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment with id " + str(experimentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignExperiment( self, trainingRunId ):
		errMsg = "Failed to unassign element " + str(experimentId) + " for Experiment on TrainingRun"

		try:
			# get the TrainingRun from db
			trainingRun = self.get( trainingRunId ).first()	
			
			# assign to None for unassignment
			trainingRun.experiment = None			

			#save it
			trainingRun.save()

			# reload and return the appropriate version					
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Exception:
			return None;
		
	def assignModelVersion( self, trainingRunId, modelVersionId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to assign element " + str(modelVersionId) + " for ModelVersion on TrainingRun"

		try:
			# get the TrainingRun from db
			trainingRun = self.get( trainingRunId ).first()	
			
			# get the ModelVersion from db
			modelVersion = ModelVersionDelegate().get(modelVersionId).first();
			
			# assign the ModelVersion		
			trainingRun.modelVersion = modelVersion
			
			#save it
			trainingRun.save()

			# reload and return the appropriate version					
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion with id " + str(modelVersionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignModelVersion( self, trainingRunId ):
		errMsg = "Failed to unassign element " + str(modelVersionId) + " for ModelVersion on TrainingRun"

		try:
			# get the TrainingRun from db
			trainingRun = self.get( trainingRunId ).first()	
			
			# assign to None for unassignment
			trainingRun.modelVersion = None			

			#save it
			trainingRun.save()

			# reload and return the appropriate version					
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Exception:
			return None;
		
	def addInputDatasets( self, trainingRunId, inputDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(inputDatasetsIds) + " for InputDatasets on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = inputDatasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				trainingRun.inputDatasets.add(dataSet)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInputDatasets( self, trainingRunId, inputDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(inputDatasetsIds) + " for InputDatasets on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = inputDatasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				trainingRun.inputDatasets.remove(dataSet)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFeatures( self, trainingRunId, featuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureDelegate import FeatureDelegate

		errMsg = "Failed to add elements " + str(featuresIds) + " for Features on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = featuresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Feature		
				feature = FeatureDelegate().get(id).first();	
				# add the Feature
				trainingRun.features.add(feature)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatures( self, trainingRunId, featuresIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureDelegate import FeatureDelegate

		errMsg = "Failed to remove elements " + str(featuresIds) + " for Features on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = featuresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Feature		
				feature = FeatureDelegate().get(id).first();	
				# add the Feature
				trainingRun.features.remove(feature)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except Feature.DoesNotExist:
			raise ProcessingError(errMsg + " : Feature does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRunMetrics( self, trainingRunId, runMetricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.RunMetricDelegate import RunMetricDelegate

		errMsg = "Failed to add elements " + str(runMetricsIds) + " for RunMetrics on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = runMetricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RunMetric		
				runMetric = RunMetricDelegate().get(id).first();	
				# add the RunMetric
				trainingRun.runMetrics.add(runMetric)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRunMetrics( self, trainingRunId, runMetricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.RunMetricDelegate import RunMetricDelegate

		errMsg = "Failed to remove elements " + str(runMetricsIds) + " for RunMetrics on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = runMetricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RunMetric		
				runMetric = RunMetricDelegate().get(id).first();	
				# add the RunMetric
				trainingRun.runMetrics.remove(runMetric)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except RunMetric.DoesNotExist:
			raise ProcessingError(errMsg + " : RunMetric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRunParameters( self, trainingRunId, runParametersIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.RunParameterDelegate import RunParameterDelegate

		errMsg = "Failed to add elements " + str(runParametersIds) + " for RunParameters on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = runParametersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RunParameter		
				runParameter = RunParameterDelegate().get(id).first();	
				# add the RunParameter
				trainingRun.runParameters.add(runParameter)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except RunParameter.DoesNotExist:
			raise ProcessingError(errMsg + " : RunParameter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRunParameters( self, trainingRunId, runParametersIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.RunParameterDelegate import RunParameterDelegate

		errMsg = "Failed to remove elements " + str(runParametersIds) + " for RunParameters on TrainingRun"

		try:
			# get the TrainingRun
			trainingRun = self.get( trainingRunId ).first()
				
			# split on a comma with no spaces
			idList = runParametersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RunParameter		
				runParameter = RunParameterDelegate().get(id).first();	
				# add the RunParameter
				trainingRun.runParameters.remove(runParameter)
				
			# save it		
			trainingRun.save()
			
			# reload and return the appropriate version
			return self.get( trainingRunId );
		except TrainingRun.DoesNotExist:
			raise ProcessingError(errMsg + " : TrainingRun with id " + str(trainingRunId) + " does not exist.")
		except RunParameter.DoesNotExist:
			raise ProcessingError(errMsg + " : RunParameter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
