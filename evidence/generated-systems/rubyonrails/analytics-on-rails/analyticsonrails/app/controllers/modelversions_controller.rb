class ModelVersionsController < ApplicationController
  def index
    @modelVersions = ModelVersion.all
  end
 
  def show
    @modelVersion = ModelVersion.find(params[:id])
  end
 
  def new
    @modelVersion = ModelVersion.new
  end
 
  def edit
    @modelVersion = ModelVersion.find(params[:id])
  end
 
  def create
    @modelVersion = ModelVersion.new(modelVersion_params)
 
    if @modelVersion.save
      redirect_to modelVersions_path
    else
      render 'new'
    end
  end
 
  def update
    @modelVersion = ModelVersion.find(params[:id])
 
    if @modelVersion.update(modelVersion_params)
      redirect_to modelVersions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @modelVersion = ModelVersion.find(params[:id])
    @modelVersion.destroy
    redirect_to modelVersions_path
  end

 
  private
    def modelVersion_params
      params.require(:modelVersion).permit(:version, :Lifecycle, :TrainingStatus)
    end
end