class EngineTypesController < ApplicationController
  def index
    @engineTypes = EngineType.all
  end
 
  def show
    @engineType = EngineType.find(params[:id])
  end
 
  def new
    @engineType = EngineType.new
  end
 
  def edit
    @engineType = EngineType.find(params[:id])
  end
 
  def create
    @engineType = EngineType.new(engineType_params)
 
    if @engineType.save
      redirect_to engineTypes_path
    else
      render 'new'
    end
  end
 
  def update
    @engineType = EngineType.find(params[:id])
 
    if @engineType.update(engineType_params)
      redirect_to engineTypes_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @engineType = EngineType.find(params[:id])
    @engineType.destroy
    redirect_to engineTypes_path
  end

 
  private
    def engineType_params
      params.require(:engineType).permit(:engineModelCode, :maxThrustKn, :Category)
    end
end