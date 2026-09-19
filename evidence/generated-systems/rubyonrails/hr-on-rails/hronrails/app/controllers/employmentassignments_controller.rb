class EmploymentAssignmentsController < ApplicationController
  def index
    @employmentAssignments = EmploymentAssignment.all
  end
 
  def show
    @employmentAssignment = EmploymentAssignment.find(params[:id])
  end
 
  def new
    @employmentAssignment = EmploymentAssignment.new
  end
 
  def edit
    @employmentAssignment = EmploymentAssignment.find(params[:id])
  end
 
  def create
    @employmentAssignment = EmploymentAssignment.new(employmentAssignment_params)
 
    if @employmentAssignment.save
      redirect_to employmentAssignments_path
    else
      render 'new'
    end
  end
 
  def update
    @employmentAssignment = EmploymentAssignment.find(params[:id])
 
    if @employmentAssignment.update(employmentAssignment_params)
      redirect_to employmentAssignments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @employmentAssignment = EmploymentAssignment.find(params[:id])
    @employmentAssignment.destroy
    redirect_to employmentAssignments_path
  end

 
  private
    def employmentAssignment_params
      params.require(:employmentAssignment).permit(:startDate, :endDate, :primary, :AssignmentType, :Status)
    end
end