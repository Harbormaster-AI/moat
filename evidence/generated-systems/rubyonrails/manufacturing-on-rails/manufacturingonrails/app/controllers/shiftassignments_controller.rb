class ShiftAssignmentsController < ApplicationController
  def index
    @shiftAssignments = ShiftAssignment.all
  end
 
  def show
    @shiftAssignment = ShiftAssignment.find(params[:id])
  end
 
  def new
    @shiftAssignment = ShiftAssignment.new
  end
 
  def edit
    @shiftAssignment = ShiftAssignment.find(params[:id])
  end
 
  def create
    @shiftAssignment = ShiftAssignment.new(shiftAssignment_params)
 
    if @shiftAssignment.save
      redirect_to shiftAssignments_path
    else
      render 'new'
    end
  end
 
  def update
    @shiftAssignment = ShiftAssignment.find(params[:id])
 
    if @shiftAssignment.update(shiftAssignment_params)
      redirect_to shiftAssignments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @shiftAssignment = ShiftAssignment.find(params[:id])
    @shiftAssignment.destroy
    redirect_to shiftAssignments_path
  end

 
  private
    def shiftAssignment_params
      params.require(:shiftAssignment).permit(:assignmentDate)
    end
end