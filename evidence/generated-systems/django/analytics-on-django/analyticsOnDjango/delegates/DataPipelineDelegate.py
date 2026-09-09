from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataTask import DataTask
from analyticsOnDjango.models.DataSource import DataSource
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.LineageNode import LineageNode
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DataPipeline
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DataPipelineDelegate Declaration
#======================================================================
class DataPipelineDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dataPipelineId ):
		try:	
			dataPipeline = DataPipeline.objects.filter(id=dataPipelineId)
			return dataPipeline.first();
		except DataPipeline.DoesNotExist:
			raise ProcessingError("DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dataPipeline):
		for model in serializers.deserialize("json", dataPipeline):
			model.save()
			return model;

	def create(self, dataPipeline):
		dataPipeline.save()
		return dataPipeline;

	def saveFromJson(self, dataPipeline):
		for model in serializers.deserialize("json", dataPipeline):
			model.save()
			return dataPipeline;
	
	def save(self, dataPipeline):
		dataPipeline.save()
		return dataPipeline;
	
	def delete(self, dataPipelineId ):
		errMsg = "Failed to delete DataPipeline from db using id " + str(dataPipelineId)
		
		try:
			dataPipeline = DataPipeline.objects.get(id=dataPipelineId)
			dataPipeline.delete()
			return True
		except DataPipeline.DoesNotExist:
			raise ProcessingError("DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DataPipeline.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DataPipeline from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, dataPipelineId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on DataPipeline"

		try:
			# get the DataPipeline from db
			dataPipeline = self.get( dataPipelineId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			dataPipeline.workspace = analyticsWorkspace
			
			#save it
			dataPipeline.save()

			# reload and return the appropriate version					
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, dataPipelineId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on DataPipeline"

		try:
			# get the DataPipeline from db
			dataPipeline = self.get( dataPipelineId ).first()	
			
			# assign to None for unassignment
			dataPipeline.analyticsWorkspace = None			

			#save it
			dataPipeline.save()

			# reload and return the appropriate version					
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLineageNode( self, dataPipelineId, lineageNodeId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to assign element " + str(lineageNodeId) + " for LineageNode on DataPipeline"

		try:
			# get the DataPipeline from db
			dataPipeline = self.get( dataPipelineId ).first()	
			
			# get the LineageNode from db
			lineageNode = LineageNodeDelegate().get(lineageNodeId).first();
			
			# assign the LineageNode		
			dataPipeline.lineageNode = lineageNode
			
			#save it
			dataPipeline.save()

			# reload and return the appropriate version					
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode with id " + str(lineageNodeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLineageNode( self, dataPipelineId ):
		errMsg = "Failed to unassign element " + str(lineageNodeId) + " for LineageNode on DataPipeline"

		try:
			# get the DataPipeline from db
			dataPipeline = self.get( dataPipelineId ).first()	
			
			# assign to None for unassignment
			dataPipeline.lineageNode = None			

			#save it
			dataPipeline.save()

			# reload and return the appropriate version					
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except Exception:
			return None;
		
	def addTasks( self, dataPipelineId, tasksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataTaskDelegate import DataTaskDelegate

		errMsg = "Failed to add elements " + str(tasksIds) + " for Tasks on DataPipeline"

		try:
			# get the DataPipeline
			dataPipeline = self.get( dataPipelineId ).first()
				
			# split on a comma with no spaces
			idList = tasksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataTask		
				dataTask = DataTaskDelegate().get(id).first();	
				# add the DataTask
				dataPipeline.tasks.add(dataTask)
				
			# save it		
			dataPipeline.save()
			
			# reload and return the appropriate version
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTasks( self, dataPipelineId, tasksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataTaskDelegate import DataTaskDelegate

		errMsg = "Failed to remove elements " + str(tasksIds) + " for Tasks on DataPipeline"

		try:
			# get the DataPipeline
			dataPipeline = self.get( dataPipelineId ).first()
				
			# split on a comma with no spaces
			idList = tasksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataTask		
				dataTask = DataTaskDelegate().get(id).first();	
				# add the DataTask
				dataPipeline.tasks.remove(dataTask)
				
			# save it		
			dataPipeline.save()
			
			# reload and return the appropriate version
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except DataTask.DoesNotExist:
			raise ProcessingError(errMsg + " : DataTask does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSources( self, dataPipelineId, sourcesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

		errMsg = "Failed to add elements " + str(sourcesIds) + " for Sources on DataPipeline"

		try:
			# get the DataPipeline
			dataPipeline = self.get( dataPipelineId ).first()
				
			# split on a comma with no spaces
			idList = sourcesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSource		
				dataSource = DataSourceDelegate().get(id).first();	
				# add the DataSource
				dataPipeline.sources.add(dataSource)
				
			# save it		
			dataPipeline.save()
			
			# reload and return the appropriate version
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSources( self, dataPipelineId, sourcesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

		errMsg = "Failed to remove elements " + str(sourcesIds) + " for Sources on DataPipeline"

		try:
			# get the DataPipeline
			dataPipeline = self.get( dataPipelineId ).first()
				
			# split on a comma with no spaces
			idList = sourcesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSource		
				dataSource = DataSourceDelegate().get(id).first();	
				# add the DataSource
				dataPipeline.sources.remove(dataSource)
				
			# save it		
			dataPipeline.save()
			
			# reload and return the appropriate version
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOutputs( self, dataPipelineId, outputsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(outputsIds) + " for Outputs on DataPipeline"

		try:
			# get the DataPipeline
			dataPipeline = self.get( dataPipelineId ).first()
				
			# split on a comma with no spaces
			idList = outputsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataPipeline.outputs.add(dataSet)
				
			# save it		
			dataPipeline.save()
			
			# reload and return the appropriate version
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOutputs( self, dataPipelineId, outputsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(outputsIds) + " for Outputs on DataPipeline"

		try:
			# get the DataPipeline
			dataPipeline = self.get( dataPipelineId ).first()
				
			# split on a comma with no spaces
			idList = outputsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				dataPipeline.outputs.remove(dataSet)
				
			# save it		
			dataPipeline.save()
			
			# reload and return the appropriate version
			return self.get( dataPipelineId );
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline with id " + str(dataPipelineId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
