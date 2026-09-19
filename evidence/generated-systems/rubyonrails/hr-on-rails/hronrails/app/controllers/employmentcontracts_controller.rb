class EmploymentContractsController < ApplicationController
  def index
    @employmentContracts = EmploymentContract.all
  end
 
  def show
    @employmentContract = EmploymentContract.find(params[:id])
  end
 
  def new
    @employmentContract = EmploymentContract.new
  end
 
  def edit
    @employmentContract = EmploymentContract.find(params[:id])
  end
 
  def create
    @employmentContract = EmploymentContract.new(employmentContract_params)
 
    if @employmentContract.save
      redirect_to employmentContracts_path
    else
      render 'new'
    end
  end
 
  def update
    @employmentContract = EmploymentContract.find(params[:id])
 
    if @employmentContract.update(employmentContract_params)
      redirect_to employmentContracts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @employmentContract = EmploymentContract.find(params[:id])
    @employmentContract.destroy
    redirect_to employmentContracts_path
  end

 
  private
    def employmentContract_params
      params.require(:employmentContract).permit(:contractNumber, :startDate, :endDate, :workHoursPerWeek, :EmploymentType, :Status, :PayFrequency)
    end
end