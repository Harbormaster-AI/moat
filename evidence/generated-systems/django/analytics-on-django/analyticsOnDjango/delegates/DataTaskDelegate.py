from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.DataTask import DataTask
from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataTask
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataTaskDelegate Declaration
#======================================================================
class DataTaskDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataTaskId ):
		try:	
			dataTask = DataTask.objects.filter(id=dataTaskId)
			return dataTask.first();
		except DataTask.DoesNotExist:
			raise ProcessingError("DataTask with id " + str(dataTaskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataTask):
		for model in serializers.deserialize("json", dataTask):
			model.save()
			return model;

	def create(self, dataTask):
		dataTask.save()
		return dataTask;

	def saveFromJson(self, dataTask):
		for model in serializers.deserialize("json", dataTask):
			model.save()
			return dataTask;
	
	def save(self, dataTask):
		dataTask.save()
		return dataTask;
	
	def delete(self, dataTaskId ):
		errMsg = "Failed to delete DataTask from db using id " + str(dataTaskId)
		
		try:
			dataTask = DataTask.objects.get(id=dataTaskId)
			dataTask.delete()
			return True
		except DataTask.DoesNotExist:
			raise ProcessingError("DataTask with id " + str(dataTaskId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataTask.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataTask from db")
		except Exception:
			return None;
		
	def assignPipeline( self, dataTaskId, pipelineId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to assign element " + str(pipelineId) + " for Pipeline on DataTask"

		try:
			# get the DataTask from db
			dataTask = self.get( dataTaskId ).first()	
			
			# get the DataPipeline from db
			dataPipeline = DataPipelineDelegate().get(pipelineId).first();
			
			# assign the Pipeline		
			dataTask.pipeline = dataPipeline
			
			#save it
			dataTask.save()

			# reload and return the appropriate version					
			return self.get( dataTaskId );
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask with id " + str(dataTaskId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(pipelineId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPipeline( self, dataTaskId ):
		errMsg = "Failed to unassign element " + str(pipelineId) + " for Pipeline on DataTask"

		try:
			# get the DataTask from db
			dataTask = self.get( dataTaskId ).first()	
			
			# assign to None for unassignment
			dataTask.dataPipeline = None			

			#save it
			dataTask.save()

			# reload and return the appropriate version					
			return self.get( dataTaskId );
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask with id " + str(dataTaskId) + " does not exist.")
		except Exception:
			return None;
		
	def addInputDatasets( self, dataTaskId, inputDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(inputDatasetsIds) + " for InputDatasets on DataTask"

		try:
			# get the DataTask
			dataTask = self.get( dataTaskId ).first()
				
			# split on a comma with no spaces
			idList = inputDatasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataTask.inputDatasets.add(dataSet)
				
			# save it		
			dataTask.save()
			
			# reload and return the appropriate version
			return self.get( dataTaskId );
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask with id " + str(dataTaskId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInputDatasets( self, dataTaskId, inputDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(inputDatasetsIds) + " for InputDatasets on DataTask"

		try:
			# get the DataTask
			dataTask = self.get( dataTaskId ).first()
				
			# split on a comma with no spaces
			idList = inputDatasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataTask.inputDatasets.remove(dataSet)
				
			# save it		
			dataTask.save()
			
			# reload and return the appropriate version
			return self.get( dataTaskId );
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask with id " + str(dataTaskId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOutputDatasets( self, dataTaskId, outputDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(outputDatasetsIds) + " for OutputDatasets on DataTask"

		try:
			# get the DataTask
			dataTask = self.get( dataTaskId ).first()
				
			# split on a comma with no spaces
			idList = outputDatasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataTask.outputDatasets.add(dataSet)
				
			# save it		
			dataTask.save()
			
			# reload and return the appropriate version
			return self.get( dataTaskId );
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask with id " + str(dataTaskId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOutputDatasets( self, dataTaskId, outputDatasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(outputDatasetsIds) + " for OutputDatasets on DataTask"

		try:
			# get the DataTask
			dataTask = self.get( dataTaskId ).first()
				
			# split on a comma with no spaces
			idList = outputDatasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataTask.outputDatasets.remove(dataSet)
				
			# save it		
			dataTask.save()
			
			# reload and return the appropriate version
			return self.get( dataTaskId );
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask with id " + str(dataTaskId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
