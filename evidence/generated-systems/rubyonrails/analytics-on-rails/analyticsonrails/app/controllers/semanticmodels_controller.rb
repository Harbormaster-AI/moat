class SemanticModelsController < ApplicationController
  def index
    @semanticModels = SemanticModel.all
  end
 
  def show
    @semanticModel = SemanticModel.find(params[:id])
  end
 
  def new
    @semanticModel = SemanticModel.new
  end
 
  def edit
    @semanticModel = SemanticModel.find(params[:id])
  end
 
  def create
    @semanticModel = SemanticModel.new(semanticModel_params)
 
    if @semanticModel.save
      redirect_to semanticModels_path
    else
      render 'new'
    end
  end
 
  def update
    @semanticModel = SemanticModel.find(params[:id])
 
    if @semanticModel.update(semanticModel_params)
      redirect_to semanticModels_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @semanticModel = SemanticModel.find(params[:id])
    @semanticModel.destroy
    redirect_to semanticModels_path
  end

 
  private
    def semanticModel_params
      params.require(:semanticModel).permit(:name, :version, :grain)
    end
end