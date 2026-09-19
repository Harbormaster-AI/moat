class DataPipelinesController < ApplicationController
  def index
    @dataPipelines = DataPipeline.all
  end
 
  def show
    @dataPipeline = DataPipeline.find(params[:id])
  end
 
  def new
    @dataPipeline = DataPipeline.new
  end
 
  def edit
    @dataPipeline = DataPipeline.find(params[:id])
  end
 
  def create
    @dataPipeline = DataPipeline.new(dataPipeline_params)
 
    if @dataPipeline.save
      redirect_to dataPipelines_path
    else
      render 'new'
    end
  end
 
  def update
    @dataPipeline = DataPipeline.find(params[:id])
 
    if @dataPipeline.update(dataPipeline_params)
      redirect_to dataPipelines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataPipeline = DataPipeline.find(params[:id])
    @dataPipeline.destroy
    redirect_to dataPipelines_path
  end

 
  private
    def dataPipeline_params
      params.require(:dataPipeline).permit(:name, :schedule, :TriggerType, :Status)
    end
end