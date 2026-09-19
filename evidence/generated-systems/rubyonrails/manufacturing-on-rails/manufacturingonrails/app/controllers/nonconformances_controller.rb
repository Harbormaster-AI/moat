class NonconformancesController < ApplicationController
  def index
    @nonconformances = Nonconformance.all
  end
 
  def show
    @nonconformance = Nonconformance.find(params[:id])
  end
 
  def new
    @nonconformance = Nonconformance.new
  end
 
  def edit
    @nonconformance = Nonconformance.find(params[:id])
  end
 
  def create
    @nonconformance = Nonconformance.new(nonconformance_params)
 
    if @nonconformance.save
      redirect_to nonconformances_path
    else
      render 'new'
    end
  end
 
  def update
    @nonconformance = Nonconformance.find(params[:id])
 
    if @nonconformance.update(nonconformance_params)
      redirect_to nonconformances_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @nonconformance = Nonconformance.find(params[:id])
    @nonconformance.destroy
    redirect_to nonconformances_path
  end

 
  private
    def nonconformance_params
      params.require(:nonconformance).permit(:ncNumber, :description, :containmentAction, :NcType, :Severity, :Status)
    end
end