class AvionicsSuitesController < ApplicationController
  def index
    @avionicsSuites = AvionicsSuite.all
  end
 
  def show
    @avionicsSuite = AvionicsSuite.find(params[:id])
  end
 
  def new
    @avionicsSuite = AvionicsSuite.new
  end
 
  def edit
    @avionicsSuite = AvionicsSuite.find(params[:id])
  end
 
  def create
    @avionicsSuite = AvionicsSuite.new(avionicsSuite_params)
 
    if @avionicsSuite.save
      redirect_to avionicsSuites_path
    else
      render 'new'
    end
  end
 
  def update
    @avionicsSuite = AvionicsSuite.find(params[:id])
 
    if @avionicsSuite.update(avionicsSuite_params)
      redirect_to avionicsSuites_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @avionicsSuite = AvionicsSuite.find(params[:id])
    @avionicsSuite.destroy
    redirect_to avionicsSuites_path
  end

 
  private
    def avionicsSuite_params
      params.require(:avionicsSuite).permit(:suiteName, :softwareBaseline)
    end
end